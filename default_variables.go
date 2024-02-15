package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Variable struct {
	Name  string
	Value string
	Comment string
}

func variableBlock(rolePath string)(string) {
	variablesBlockText := "## Переменные по умолчанию\n|Переменная|Значение|Комментарий|\n|:-|:-:|:-:|\n"
	// Replace with the path to your Ansible variable file
	file, err := os.Open("C:\\Users\\Lenovo\\Documents\\Работа\\ansible-etcd\\defaults\\main.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	varList := []Variable{}

	// Regular expression pattern to match variable assignments and comments
	pattern := regexp.MustCompile(`^(\w+)\s*:\s*(.*?)(#.*)?$`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindStringSubmatch(line)
		if len(matches) >  0 {
			var varEntry Variable
			varEntry.Name = matches[1]
			varEntry.Value = strings.TrimSpace(matches[2])
			if len(matches) ==  4 {
				varEntry.Comment = strings.TrimPrefix(matches[3], "#")
			}
			varList = append(varList, varEntry)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Output the parsed variables
	
	for _, varEntry := range varList {
		varEntry.Value = strings.ReplaceAll(varEntry.Value, "|", "\\|")
		variableLine := fmt.Sprintf("|%s|%s|%s|\n", varEntry.Name, varEntry.Value, varEntry.Comment)
		variablesBlockText = variablesBlockText + variableLine
	}

	return variablesBlockText
}
