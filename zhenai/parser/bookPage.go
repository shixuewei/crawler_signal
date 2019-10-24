package parser

import (
	"learning/crawler_signal/engine"
	"regexp"
	"strconv"
)

//正则解析式
const (
	bookPageRe = `<a href="https://www.biikan.com/([^/]+)/[0-9]+.shtml">([^<]+)</a>`
)

//爬取所有的页
func ParseBookPage(content []byte) engine.ParseResult {
	//获得正则表达式的对象
	re := regexp.MustCompile(bookPageRe)
	//获得匹配结果，
	matches := re.FindStringSubmatch(string(content))
	result := engine.ParseResult{}
	//获得每一类图书的总页数
	sum, _ := strconv.Atoi(matches[2])
	for i := 1; i <= sum; i++ {
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://www.biikan.com" + "/" + matches[1] + "/" + strconv.Itoa(i) + ".shtml",
			ParserFunc: ParseBook, //爬取每一页的书
		})
	}
	return result
}
