package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const getURL = "https://golang.org"
const postURL = "https://httpbin.org/post"

func GETRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	fmt.Println(resp.Status, resp.Header, resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%s\n", data)
}

type User struct {
	Name	string `json:"name"`
	ID		uint32 `json:"id"`
}

type Output struct {
	JSON struct {
		Name string `json:"name"`
		ID uint32 `json:"id"`
	} `json:"json"`
	URL string `json:"url"`
}

func POSTRequest(url string) {
	var u = User{
		Name: "Alex",
		ID: 1,
	}

	bytesRepr, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepr))  // http.Post awaits io.Reader interface
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(bytesResp))

	var out Output
	err = json.Unmarshal(bytesResp, &out)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", out)
	fmt.Println(out.URL)
}

func PostForm() {
	formData := url.Values{
		"name": {"hello"},
		"surname": {"golang post form"},
	}

	resp, err := http.PostForm(postURL, formData)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result["form"])
}

func RequestWithQueryParams() {
	baseURL := "https://example.com/api/resource"
	params := url.Values{}
	params.Add("param1", "value1")
	params.Add("param2", "value2")
	params.Add("param2", "one_more_value")

	fullURL := baseURL + "?" + params.Encode()
	fmt.Println(fullURL)

	// another one way 
	base2, err := url.Parse("https://example.com")
	if err != nil {log.Fatalln(err)}
	base2.Path += "api/"
	base2.Path += "resource"
	base2.RawQuery = params.Encode()
	fmt.Printf("Encoded URL is %q\n", base2.String())

	fmt.Println(fullURL == base2.String())
}

type Todo struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func ComplexRequest() {
	todo := Todo{
		UserID: 1,
		ID: 2,
		Title: "our title",
		Completed: true,
	}

	jsonReq, err := json.Marshal(todo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonReq))

	baseURL := "https://jsonplaceholder.typicode.com/todos"

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln("error creating request: ", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln("error sending request: ", err)
	}

	defer response.Body.Close()

	fmt.Println("Response status: ", response.Status)

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodyString := string(bodyBytes)
	fmt.Printf("API response as a string: %s\n", bodyString)

	var todoStruct Todo
	err = json.Unmarshal(bodyBytes, &todoStruct)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("API response as a struct: \n%+v\n", todoStruct)
}

func TestQuery() {
	var name, age string
	fmt.Scan(&name, &age)
	baseURL := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	url := baseURL + "?" + params.Encode()
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func main() {
	// GETRequest(getURL)
	// POSTRequest(postURL)
	// PostForm()
	// RequestWithQueryParams()
	// ComplexRequest()
	TestQuery()
}
