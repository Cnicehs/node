package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./urls.txt")
	defer f.Close()
	b, _ := io.ReadAll(f)
	urls := strings.Split(string(b), "\n")
	results := make([]string, len(urls))
	for i := 0; i < len(urls); i++ {
		result := GetUrlString(strings.TrimSpace(urls[i]))
		if result != "" {
			result = DecodeBase64(result)
			if result != "" {
				results = append(results, strings.TrimSpace(result))
			}
		}
	}
	resultsCombine := strings.Join(results, "\n")
	resultsCombine = strings.TrimSpace(resultsCombine)
	out, _ := os.Create("./docs/index.html")
	defer out.Close()
	out.WriteString(base64.StdEncoding.EncodeToString([]byte(resultsCombine)))
}

func GetUrlString(u string) string {
	req, e := http.NewRequest("GET", u, nil)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	ret, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	return string(ret)
}

func DecodeBase64(s string) string {
	b, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	return string(b)
}
