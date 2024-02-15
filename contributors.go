package main

import (
	"log"
	"os"
	"strings"
	"gopkg.in/yaml.v3"
)

type Contributors struct {
	Author string
	Company string
	Block string
}

func contributorsBlock(rolePath string)(string){
	rolePath += "\\meta"
	contributorsBlockText := "## Авторы\n|Автор|Компания|Блок|\n|:-|:-:|:-:|\n"
	
	files, err := os.ReadDir(rolePath)
	if err != nil {
		log.Fatalf("Error reading directory: %s\n", err)
	} 
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yml") || (strings.HasSuffix(file.Name(), ".yaml")){
			data, _ := readYamlFile(rolePath, file.Name())
			parsedYaml, _ := parseMetaFile(data)
			contributorsList :=  createList(parsedYaml)
			contributorsBlockText = createBlock(contributorsList, contributorsBlockText)
		}
	}
	if contributorsBlockText == "|Автор|Компания|Блок|\n|:-|:-:|:-:|\n" {
		contributorsBlockText = ""
	}
	return contributorsBlockText
}

func parseMetaFile(data []byte) (map[string]interface{}, error) {
	var parsedYaml map[string]interface{}
	err := yaml.Unmarshal(data, &parsedYaml)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
		return nil, err
	}
	return parsedYaml, err
}

func createList(parsedYaml map[string]interface{})([]Contributors){
	var contributors []Contributors
	for block, value := range parsedYaml {
		subMap, ok := value.(map[string]interface{})	
		author, ok := subMap["author"].(string)
		if !ok {
			author = ""
		}
		company, ok := subMap["company"].(string)
		if !ok {
			company =""
		}
		if author + company == "" {
			continue
		}
		contributors = append(contributors, Contributors{Author: author, Company: company, Block: block})
	}
	return contributors
}

func createBlock(contributorsList []Contributors, contributorsBlockText string)(string){
	for _, contributor := range contributorsList {
		contributorsBlockText += "|" + contributor.Author + "|" + contributor.Company + "|" + contributor.Block + "|\n"
	} 
	return contributorsBlockText
}