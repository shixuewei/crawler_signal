package fetcher

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

// func determineEncoding(r *bufio.Reader) encoding.Encoding {
// 	bytes, err := r.Peek(1024)
// 	if err != nil {
// 		log.Printf("Fetcher err : %v", err)
// 		return unicode.UTF8
// 	}
// 	e, _, _ := charset.DetermineEncoding(bytes, "")
// 	return e
// }

//获取每个页面的具体内容
func Fetch(url string) ([]byte, error) {
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	//得到resp
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	//注意要对resp.Body进行关闭
	defer resp.Body.Close()

	//如果没取到就返回错误
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	//bodyReader := bufio.NewReader(resp.Body)
	//将gbk编码转变为utf-8
	//e := determineEncoding(bodyReader)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	//返回resp.Body的所有内容
	return ioutil.ReadAll(resp.Body)

}
