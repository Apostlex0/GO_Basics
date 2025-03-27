package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myurl string = "https://www.learncodeonline.in:3000/learn?coursename=reactjs&paymentid=1234"

func main() {
	fmt.Println("Welcome to urlhandling in golang")
	fmt.Println("URL is: ", myurl)

	// URL parsing
	result, _ := url.Parse(myurl)
	fmt.Println("Result is: ", result.Scheme)
	fmt.Println("Result is: ", result.Host)
	fmt.Println("Result is: ", result.Path)
	fmt.Println("Result is: ", result.RawQuery)
	fmt.Println("Result is: ", result.Port())

	qparams := result.Query()
	fmt.Printf("Query params are: %T\n", qparams)

	fmt.Println("Query params are: ", qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, value := range qparams {
		fmt.Println("Value is: ", value)
	}
	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.learncodeonline.in",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsofUrl.String()
	fmt.Println("Another URL is: ", anotherURL)

	PerformPostJsonRequest()

}

func PerformPostJsonRequest() {
	const myurl = "https://localhost:8000/post"

	requestBody := strings.NewReader(`
	   {
		"coursename": "Lets go with golang "
		"paymentid": 1234
		"username": "hitesh"
		"platform": "learnCodeOnline.in"
	   }
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "https://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")


	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
