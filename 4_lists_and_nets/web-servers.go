package main

import (
	// "fmt"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is Get-request")
	// name := r.URL.Query().Get("name")
	name := r.FormValue("name")
	greeting := fmt.Sprintf("Hello, %s!", name)
	w.Write([]byte(greeting))
	w.WriteHeader(200)  // Optional
}

func handlePOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is Post-request")
	bytesBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		w.Write([]byte("Bad request body"))
		return
	}
	fmt.Println(string(bytesBody))
	w.Write([]byte("OK!"))
	w.WriteHeader(http.StatusOK)
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println(r.Proto)
	fmt.Println(r.Cookies())

	fmt.Println("RawQuery: ", r.URL.String())
	fmt.Println("Name: ", r.URL.Query().Get("name"))  // parametr or empty string
	fmt.Println("IsExist: ", r.URL.Query().Has("name"))  // if this parameter exists

	switch r.Method {
	case http.MethodGet:
		handleGET(w, r)
	case http.MethodPost:
		handlePOST(w, r)	
	default:
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)  // return 405
	}
	// w.Write([]byte("Hello from Golang!"))
}

func startServer() {
	http.HandleFunc("/", handlerMain)

	fmt.Println("Server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}

var count = 0

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(fmt.Sprint(count)))
		// fmt.Fprint(w, count)  // better, but doesn't pass tests
	case http.MethodPost:
		in := r.FormValue("count")
		// in := r.Form.Get("count")  // nearly the same for this case
		counter, err := strconv.Atoi(in)
		if err != nil {
			// http.Error(w, "это не число", http.StatusBadRequest) // doesn't pass tests
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
		}
		count += counter
	default:
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)  // return 405
	}
}

func testServer() {
	http.HandleFunc("/count", countHandler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatalln("Error starting server: ", err)
	}
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func MuxServer() {
	serverMux := http.NewServeMux()

	th1123 := &timeHandler{format: time.RFC1123}
	serverMux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	serverMux.Handle("/time/rfc3339", th3339)

	log.Println("Listening...")
	err := http.ListenAndServe(":8080", serverMux)  // serverMux instead of nil in this case
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func main() {
	// startServer()
	// testServer()
	MuxServer()
}
