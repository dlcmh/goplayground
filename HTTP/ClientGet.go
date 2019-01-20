package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// go run ClientGet.go http://www.golang.com
// The response header is
// HTTP/2.0 200 OK
// Connection: close
// Alt-Svc: quic=":443"; ma=2592000; v="44,43,39,35"
// Content-Type: text/html; charset=utf-8
// Date: Sun, 20 Jan 2019 08:19:06 GMT
// Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
// Vary: Accept-Encoding
// Via: 1.1 google
//
//
// got body

// go run ClientGet.go https://www.briefnotes.net
// The response header is
// HTTP/1.1 200 OK
// Content-Length: 3910
// Connection: keep-alive
// Content-Type: text/html; charset=UTF-8
// Date: Sun, 20 Jan 2019 08:19:45 GMT
// Server: nginx
//
//
// got body

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	url, err := url.Parse(os.Args[1])
	checkError(err)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)

	// only accept UTF-8
	request.Header.Add("Accept-Charset", "utf-8;q=1, ISO-8859-1;q=0")
	checkError(err)

	response, err := client.Do(request)
	checkError(err)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))

	chSet := getCharSet(response)
	if chSet != "utf-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	// var buf [512]byte
	// reader := response.Body
	// fmt.Println("got body")
	// for {
	// 	n, err := reader.Read(buf[0:])
	// 	fmt.Print(string(buf[0:n]))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(0)
	// 	}
	// 	fmt.Print(string(buf[0:n]))
	// }

	os.Exit(0)
}

func getCharSet(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "utf-8"
	}
	idx := strings.Index(contentType, "charset=")
	if idx == -1 {
		//guess
		return "utf-8"
	}
	chSet := strings.Trim(contentType[idx+8:], " ")
	return strings.ToLower(chSet)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
