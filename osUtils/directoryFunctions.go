package osUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type FileEntry struct {
	Path string `json:"path"`
	IsDirectory bool `json:"is_directory"`
	Size int64 `json:"size"`
	LastModified time.Time `json:"last_modified"`
	FileEntries []FileEntry `json:"file_entries"`
}

func (f FileEntry) Print() {
	json, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}

func ReadFileFromPath(filePath string) []byte {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func ListDirectoryContents(directoryPath string, skipHidden bool) FileEntry {
	dirInfo, err := os.Stat(directoryPath)
	if err != nil {
		log.Fatal(err)
	}
	
	root := FileEntry{
		Path: directoryPath,
		IsDirectory: true,
		Size: 0,
		LastModified: dirInfo.ModTime(),
		FileEntries: []FileEntry{},
	}
	entries, err := os.ReadDir(directoryPath) 
	if err != nil {
		fmt.Println("Error reading directory. Finished prematurely:", err)
	}
	
	for _, entry := range entries {
		// Skip hidden files
		if skipHidden && entry.Name() != "." && entry.Name()[0:1] == "." {
			continue
		}
		if entry.IsDir() {
			child := ListDirectoryContents(directoryPath + "/" + entry.Name(), skipHidden)
			root.FileEntries = append(root.FileEntries, child)
		} else {
			entryInfo, err := entry.Info()
			if err != nil {
				log.Fatal(err)
			}
			
			child := FileEntry{
				Path: directoryPath + "/" + entry.Name(),
				IsDirectory: false,
				Size: entryInfo.Size(),
				LastModified: entryInfo.ModTime(),
				FileEntries: []FileEntry{},
			}
			root.FileEntries = append(root.FileEntries, child)
		}
	}

	return root
}