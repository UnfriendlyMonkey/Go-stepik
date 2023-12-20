package main

import (
	// "encoding/json"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"

	// "io"
	"log"
	"os"
)

type myStruct struct {
	Name	string
	Age		int
	Status bool
	Values []int
}

type user struct {
	Name string
	Email string
	Status bool
	Language []byte
}

type student struct {
	LastName string
	FirstName string
	MiddleName string
	Birthday string
	Address string
	Phone string
	Rating []int // the only necessary for the task. may as well omit all others
}

type group struct {
	ID int
	Number string
	Year int
	Students []student // the only necessary for the task. may as well omit all others
}

func ExampleMarshal() []byte {
	s := myStruct{"John Connor", 35, true, []int{15, 11, 37}}
	fmt.Println(s)

	data, err := json.Marshal(s)
	if err != nil {log.Fatal(err)}

	fmt.Printf("%s - %T\n", data, data)
	return data
}

func ExampleUnmarshal() {
	data := ExampleMarshal()
	var s myStruct
	if err := json.Unmarshal(data, &s); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v - %T\n", s, s)

	corData := bytes.Trim(data, "{")
	if !json.Valid(corData) {
		fmt.Println("invalid json!", string(corData))
	}
}

func toAndFrom() {
	newUser := user{"Alex", "email@email.email", true, []byte("ru"),}
	data, err := json.Marshal(newUser)
	if err != nil {log.Fatal(err)}
	newUser.Language = []byte("en")
	err = json.Unmarshal(data, &newUser)
	if err != nil {log.Fatal(err)}
	fmt.Println(string(newUser.Language))
}

func findAverageMark() {
	// path := "./students.json"
	// file, err := os.Open(path)
	// if err != nil {log.Fatal(err)}
	// reader := json.NewDecoder(file)
	reader := json.NewDecoder(os.Stdin)
	var groupData group
	err := reader.Decode(&groupData)
	if err != nil{log.Fatal(err)}
	fmt.Printf("%v - %T\n", groupData, groupData)
	studCount := len(groupData.Students)
	var sum, count int
	for _, el := range groupData.Students {
		fmt.Println(el.Rating)
		count += len(el.Rating)
		for _, item := range el.Rating {
			sum += item
		}
	}
	res := float64(sum) / float64(count)
	avg := float64(count) / float64(studCount)
	fmt.Println(sum, count, res, avg)
	var average = map[string]float64{"Average": avg}
	newPath := "./avg_res.json"
	file, err := os.Create(newPath)
	if err != nil {log.Fatal(err)}
	defer file.Close()
	output, err := json.MarshalIndent(average, "", "    ")
	if err != nil {log.Fatal(err)}
	// wri := json.NewEncoder(file)
	// wri.Encode(groupData)
	fmt.Println(string(output))
	os.Stdout.Write(output)
	writer := bufio.NewWriter(file)
	_, err = writer.Write(output)
	if err != nil {log.Fatal(err)}
}

type annotatedStruct struct {
	Name string `json:"name"` // json key would be name (not Name)
	Age int `json:",omitempty"` // would be omitted if empty upon marshalling
	Status bool `json:"-"` // would be always ignored upon marshalling
	comment string // non-exporting fields would be always ignored when marshalling
}

func annotations() {
	src := annotatedStruct{"John", 0, true, "I'll be back"}
	// data, err := json.Marshal(src)
	// if err != nil {log.Fatal(err)}
	// fmt.Printf("%s\n", data) // would include only Name
	dst := annotatedStruct{}
	buf := new(bytes.Buffer) // instead of real file

	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)

	enc.Encode(src)
	dec.Decode(&dst)

	fmt.Print(dst) // would include all fields (even with annotations) except non-exported
}

func main() {
	// ExampleMarshal()
	// ExampleUnmarshal()
	// toAndFrom()
	// findAverageMark()
	annotations()
}
