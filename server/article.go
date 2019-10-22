package main

import (
	"io/ioutil"
)

const ARTICLE_DIR = "./articles"

type ArticleListing struct{
	Name string
	File string
}

func getArticleList() []ArticleListing {
	files, err := ioutil.ReadDir(ARTICLE_DIR)
	if err != nil {
		myLog("Failed to list the articles directory")
		return nil
	}

	ret := make([]ArticleListing, 0)
	for _, file := range files {
		ret = append(ret, ArticleListing{file.Name(), file.Name()})
	}

	return ret;
}
