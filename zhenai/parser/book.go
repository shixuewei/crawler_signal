package parser

import (
	"learning/crawler_signal/engine"
	"regexp"
)

//正则解析式
const (
	bookRe = `<h1><a href="([^"]+)" target="_blank">([^<]+)</a></h1>`
)

//爬取每一页书
func ParseBook(content []byte) engine.ParseResult {
	re := regexp.MustCompile(bookRe)
	//只需要前20个匹配项
	matches := re.FindAllStringSubmatch(string(content), 20)
	result := engine.ParseResult{}
	for _, v := range matches {
		name := v[2] //得到每一本书的名字
		result.Items = append(result.Items, "BookName:"+v[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: "https://www.biikan.com" + v[1],
			ParserFunc: func(content []byte) engine.ParseResult {
				//由于此处函数多了一个参数，所以需要包装一下，使用闭包的结构
				return ParseBookMes(content, name) //爬取每一本书的具体内容
			},
		})
	}
	return result
}
