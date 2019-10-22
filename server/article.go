package main

import (
	"io/ioutil"
	"strings"
)

const ARTICLE_DIR = "./articles"

type ArticleListing struct{
	Name string
	File string
	Author string
	Date string
}

func getArticleList() []ArticleListing {
	files, err := ioutil.ReadDir(ARTICLE_DIR)
	if err != nil {
		myLog("Failed to list the articles directory")
		return nil
	}

	ret := make([]ArticleListing, 0)
	for _, file := range files {
		data, err := ioutil.ReadFile(ARTICLE_DIR + "/" + file.Name())
		if err != nil{
			myLog("Could not load '"+file.Name()+"'... skipping")
			continue
		}
		var article ArticleListing
		article.File = file.Name()
		spl := strings.Split(string(data), "\n");
		for _, line := range spl{
			if line == "~"{
				break;
			}
			lsplit := strings.Split(line, ":")
			if lsplit[0] == "name"{
				article.Name = lsplit[1]
			}else if lsplit[0] == "author"{
				article.Author = lsplit[1]
			}else if lsplit[0] == "date"{
				article.Date = lsplit[1]
			}else{
				myLog(lsplit[0] + " is not a header field we know about")
			}
		}
		ret = append(ret, article)
	}

	return ret;
}

func encodeListing(listing ArticleListing) string{
	var sb strings.Builder
	sb.WriteString(listing.Name)
	sb.WriteString("#")
	sb.WriteString(listing.File)
	sb.WriteString("#")
	sb.WriteString(listing.Author)
	sb.WriteString("#")
	sb.WriteString(listing.Date)

	return sb.String()
}

func getArticleBody(fileName string) string{
	data, err := ioutil.ReadFile(ARTICLE_DIR + "/" + fileName)
	if err != nil{
		myLog("Cannot get body at file '"+fileName+"'")
		return ""
	}
	//Very bad
	return strings.Split(string(data), "~")[1]
}
