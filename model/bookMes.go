package model

//所爬取的书的信息
type BookMes struct {
	//书名
	Name string
	//作者
	Author string
	//书的类别
	BookSort string
	//书籍语言
	Language string
	//哪国出版
	Country string
	//点击数
	Click int
	//总字数
	WordsNum int
}
