package main

import (
	"log"
	"os"
	"gopkg.in/yaml.v3"
	"golang.org/x/exp/slices"
	"fmt"
)

type tagInfo struct {
	name string
	tags []string
}


func tagBlock(rolePath string)(string){
	var allTagInfoList []tagInfo
	var allTags []string
	var tagBlockText string

	rolePath += "\\tasks"
	tagBlockText = "## Использование тегов\n Данный сценарий может использовать следущие теги:\n"

	files, err := os.ReadDir(rolePath)
	if err != nil {
		log.Fatalf("Error reading directory: %s\n", err)
	} 
	for _, file := range files {
		data, _ := readYamlFile(rolePath, file.Name())
		parsedYaml, _ := parsePlaybook(data)
		tagInfoList,tags := createTagInfoList(parsedYaml)
		allTags,tagBlockText = allUniqueTags(allTags,tags,tagBlockText)
		allTagInfoList = combineTagInfoList(allTagInfoList, tagInfoList)		
	}
	tagBlockText = getTagsTasks(allTagInfoList, allTags, tagBlockText)
	return tagBlockText
}

func parsePlaybook(data []byte) ([]map[string]interface{}, error) {
	var parsedYaml []map[string]interface{}
	err := yaml.Unmarshal(data, &parsedYaml)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
		return nil, err
	}
	return parsedYaml, err
}

func createTagInfoList(parsedYaml []map[string]interface{})([]tagInfo, []string) {
	var tagsInfoList []tagInfo
	var allTagsList []string
	for _, task := range parsedYaml {
		if block,ok := task["block"].([]interface{}); ok {
			var blockmap []map[string]interface{}
			for _, item := range block {
				blockmap = append(blockmap, item.(map[string]interface{}))
			}
			fmt.Println(blockmap)
			tagsInfoList,allTagsList = createTagInfoList(blockmap)
		}
		name,ok := task["name"].(string)
		if !ok {
			name ="Без имени"
		}
		var tags []string
		if taskTags, ok := task["tags"].([]interface{}); ok {
			for _, tag := range taskTags {
				tags = append(tags, tag.(string))
				allTagsList = append(allTagsList, tag.(string))
			}
		} else if tag, ok := task["tags"].(string); ok {
			tags = append(tags, tag)
			allTagsList = append(allTagsList, tag)
		} else {
			tags = append(tags, "")
		}
		tagsInfoList = append(tagsInfoList, tagInfo{name: name, tags: tags})
	}
	return tagsInfoList,allTagsList
}

func allUniqueTags(allTags []string,tags []string, tagBlockText string)([]string,string){
	for _,tag := range tags{
		if !slices.Contains(allTags, tag) {
			allTags = append(allTags, tag)
			tagBlockText += "- " + tag +"\n"
		}	
	}

	return allTags, tagBlockText
}

func combineTagInfoList(allTagInfoList []tagInfo,tagInfoList []tagInfo)([]tagInfo){
	for _,tagInfo := range tagInfoList{
		allTagInfoList = append(allTagInfoList, tagInfo)
	}
	return allTagInfoList
}


func getTagsTasks(allTagInfoList []tagInfo, allTags []string, tagBlockText string) (string) {
	for _, tagname := range allTags {
		tagBlockText += "\nШаги сценария для тега **" + tagname + "**:\n"
		for _, tagInfo := range allTagInfoList {
			for _, tag := range tagInfo.tags {
				if tag == tagname {
					tagBlockText += "- " + tagInfo.name + "\n"
					break
				}
			}
		}
	}
	return tagBlockText
}
