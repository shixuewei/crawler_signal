package main

import (
	"learning/crawler_signal/engine"
	"learning/crawler_signal/zhenai/parser"
)

//单机版爬虫
func main() {
	//启动爬虫引擎
	engine.Run(engine.Request{
		Url:        "https://www.biikan.com",
		ParserFunc: parser.ParseBookSort, //先爬取书的类别
	})
}
