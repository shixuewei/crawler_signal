package parser

import (
	"learning/crawler_signal/engine"
	"learning/crawler_signal/model"
	"regexp"
	"strconv"
)

//正则解析式
var authorRe = regexp.MustCompile(
	`<li><span class="gray2">作者：</span><a.*>([^<]+)</a></li>`)
var bookSortRe = regexp.MustCompile(
	`<li><span class="gray2">分类：</span><a.*>([^<]+)</a></li>`)
var languageRe = regexp.MustCompile(
	`<li><span class="gray2">语言：</span><a.*>([^<]+)</a></li>`)
var countryRe = regexp.MustCompile(
	`<li><span class="gray2">国家：</span><a.*>([^<]+)</a></li>`)
var clickRe = regexp.MustCompile(
	`<li><span class="gray2">点击：</span><span class="gray pr20">([^<]+)</span></li>`)
var wordsNumRe = regexp.MustCompile(
	`<li><span class="gray2">字数：</span><span class="gray">([^<]+)</span></li>`)

//爬取每本书的信息
func ParseBookMes(content []byte, name string) engine.ParseResult {
	bookMes := model.BookMes{}
	bookMes.Name = name
	bookMes.Author = extractString(content, authorRe)
	bookMes.BookSort = extractString(content, bookSortRe)
	bookMes.Language = extractString(content, languageRe)
	bookMes.Country = extractString(content, countryRe)
	bookMes.Click, _ = strconv.Atoi(extractString(content, clickRe))
	bookMes.WordsNum, _ = strconv.Atoi(extractString(content, wordsNumRe))

	result := engine.ParseResult{
		Items: []interface{}{bookMes},
	}
	return result
}

//正则匹配函数
func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	//如果匹配到了，长度至少为2，小于2则匹配失败
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
