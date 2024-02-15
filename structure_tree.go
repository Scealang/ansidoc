package main

import (
	"os"
	"log"
)


func structureBlock(rolePath string)(string) {
	var structureBlockText,indent, indentFigure string

	structureBlockText = "## Структура проекта\n```"

	
	indent = ""
	files, err := os.ReadDir(rolePath)

	if err != nil {
		log.Fatalf("Error reading directory: %s\n", err)
	} 

	for _, file := range files {
		indent = ""
		indentFigure = ""
		structureBlockText = createTree(file, rolePath, indent, indentFigure, structureBlockText)
	}
	structureBlockText += "\n```"
	return structureBlockText
}

func createTree( file os.DirEntry, rolePath string, indent string, indentFigure string, structureBlockText string) string{
	structureBlockText += "\n" + indent + indentFigure  + file.Name()
	if file.IsDir() {
		indent +="   "
		indentFigure ="├──"
		subrolePath := rolePath + "\\" + file.Name()
		subfiles, suberror := os.ReadDir(subrolePath)

		if suberror != nil {
			log.Fatalf("Error reading subdirectory: %s\n", suberror)
		}
		for index, subfile := range subfiles {
			if (len(subfiles) == index + 1) {
				indentFigure ="└──"
			}
			structureBlockText = createTree(subfile, subrolePath, indent, indentFigure, structureBlockText)
		}
	}
	return structureBlockText;
}