package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {

	var args arguments = initArgs()
	if args.targetMode.IsDir() {

		filename := args.target + "/describeFiles.json"
		createOutputFile(filename)
		outFile, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		freezeDirectoryFileData(args.target, args.recursive, outFile)
	} else if args.targetMode.IsRegular() {
		log.Fatal("Regular files are not supported at this time.")
	} else {
		log.Fatal("Not a directory or regular file")
	}
}

func freezeDirectoryFileData(directory string, recursive bool, out *os.File) {

	files := make(chan fileInfo)
	var wg sync.WaitGroup
	go getFiles(directory, recursive, files)
	wg.Add(1)
	go writeFileInfo(files, out, &wg)
	wg.Wait()
}

func createOutputFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

type fileInfo struct {
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	Mode         string `json:"mode"`
	LastModified string `json:"modTime"`
	Dir          string `json:"directory"`
}

func writeFileInfo(files <-chan fileInfo, out *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	out.WriteString("[")
	for file := range files {

		data, err := json.Marshal(file)

		if err != nil {
			log.Fatal(err)
		}
		out.WriteString(string(data))
		if files != nil {
			out.WriteString(",")
		}
	}

	// Overwrite trailing comma
	_, err := out.Seek(-1, 1)
	if err != nil {
		log.Fatal(err)
	}
	out.WriteString("]")
}

func getFiles(dir string, recursive bool, infoChan chan<- fileInfo) {
	log.Printf("Scanning '%s'...\n", dir)
	if recursive {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				infoChan <- fileInfo{
					Name:         info.Name(),
					Size:         info.Size(),
					Mode:         info.Mode().String(),
					LastModified: info.ModTime().String(),
					Dir:          filepath.Dir(path),
				}
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fileInfos, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for f := range fileInfos {
			if !fileInfos[f].IsDir() {
				infoChan <- fileInfo{
					fileInfos[f].Name(),
					fileInfos[f].Size(),
					fileInfos[f].Mode().String(),
					fileInfos[f].ModTime().String(),
					dir,
				}
			}
		}
	}
	close(infoChan)
}
