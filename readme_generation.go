package main

import (
	"os"
	"log"
	"io/fs"
)

func createFile(path string)(){
	err := os.Remove(path + "\\README.md")
	file, err := os.Create(path + "\\README.md")
	if err != nil { 
        log.Fatal("Failed to recreate file: %s\n", err ) 
    }
	file.Close()
}

func appendToFile(path string, content string)(){
	f, err := os.OpenFile(path + "\\README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open in file: %s\n", err)
	}
	if _, err := f.Write([]byte(content)); err != nil {
		f.Close() 
		log.Fatal("Failed to write in file: %s\n", err)
	}
	if err := f.Close(); err != nil {
		log.Fatal("Failed to finish writing in file: %s\n", err)
	}
}

func readYamlFile(rolePath string, filename string) ([]byte, error) {
	yamlFile, err := fs.ReadFile(os.DirFS(rolePath), filename)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
		return nil, err
	}
	return yamlFile, err
}