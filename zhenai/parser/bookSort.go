package parser

import (
	"learning/crawler_signal/engine"
	"regexp"
)

//正则解析式
const (
	bookSort = `<li ><a href="/([^"]+)">([^<]+)</a></li>`
)

//爬取图书的总类别（返回结果为ParseResult）
func ParseBookSort(content []byte) engine.ParseResult {
	//获得正则表达式的对象
	re := regexp.MustCompile(bookSort)
	//获得匹配结果，-1的意思是获取所有的匹配结果
	matches := re.FindAllStringSubmatch(string(content), -1)
	//初始化输出结果
	result := engine.ParseResult{}
	for _, v := range matches {
		result.Items = append(result.Items, "BookSort:"+v[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://www.biikan.com" + "/" + v[1],
			ParserFunc: ParseBookPage, //爬取完总类后，爬取所有的页码
		})
	}
	return result
}
