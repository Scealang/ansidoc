package main

import (
	"os"
)

func main () {
	var rolePath string

	rolePath = os.Getenv("ANSIBLE_ROLE_DIR")
	createFile(rolePath)
	appendToFile(rolePath, headerBlock(rolePath) + "\n")
	appendToFile(rolePath, structureBlock(rolePath) + "\n")
	appendToFile(rolePath, contributorsBlock(rolePath) + "\n")
	appendToFile(rolePath, tagBlock(rolePath) +"\n")
	appendToFile(rolePath, variableBlock(rolePath) +"\n")
}

