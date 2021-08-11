package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <- chan - canal somento -leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {

			resp, erro := http.Get(url)
			fmt.Println(erro, "ok")
			fmt.Println(resp, resp.Body, resp.Request.URL)

			html, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			fmt.Println(html)
			r, _ := regexp.Compile("<title>(.*?)<\\title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.google.com")
	t2 := titulo("https://www.amazon.com.br")

	fmt.Println("Primeiros:", <-t1, '|', <-t2)
	fmt.Println("Segundos:", <-t1, '|', <-t2)
}
