package engine

//Request变量定义的涵义：
//1.发送网页请求时，会返回页面中相关的url,所有需要Url变量
//2.在获取到第一个请求页面时，会执行广度优先搜索的算法，所以需要一个函数变量，
//来对下一个变量进行操作，所以定义了ParserFunc func([]byte) ParseResult：
//其输入为所获页面的内容，返回结果为ParseResult(爬取内容)
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//ParseResult变量定义的涵义：
//在获取相应的页面的内容时，必然获得本页的信息，及Items,
//同时也会得到下一个页面的请求，其为Request类型
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//定义一个空函数，当所获的页面没有所需要执行的操作时，调用此函数
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
