package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	// "strings"
)

func filesBeginning() {
	b := bytes.NewReader([]byte("Different data from io.Reader object, включая кириллицу"))

	data, err := io.ReadAll(b)
	if err != nil {panic(err)}
	fmt.Printf("%s\n", data)
}

func ExamplesWorkWithFiles() {
	dataForFile := []byte("Тестовая строка to be written to file")
	if err := os.WriteFile("test.txt", dataForFile, 0600); err != nil {
		panic(err)
	}
	dataFromFile, err := os.ReadFile("test.txt")
	if err != nil {panic(err)}

	fmt.Printf("initial data == final data: %v\n", bytes.Equal(dataForFile, dataFromFile))
}

func ExampleWorkWithDir() {
	currentDir, err := os.Getwd()
	if err != nil {log.Fatal(err)}
	fmt.Println(currentDir)

	filesFromDir, err := os.ReadDir(".")
	if err != nil {panic(err)}

	for _, file := range filesFromDir {
		info, _ := file.Info()
		fmt.Printf("|-name: %s, size: %d\n", file.Name(), info.Size())
	}
}

func StrangeExample() {
	for i := 1; i < 4; i++ {
		file, err := os.Create(strconv.Itoa(i) + "test.txt")
		if err != nil {log.Fatal(err)}
		defer file.Close()
	}
	// ExampleWorkWithDir()
	os.Rename("2test.txt", "4test.txt")
	for i := 4; i > 1; i-- {
		os.Remove(strconv.Itoa(i) + "test.txt")
	}
	// ExampleWorkWithDir()
}

func ExampleBufioWrite() {
	testFile := "bufio_test.txt"
	file, err := os.Create(testFile)
	if err != nil {log.Fatal(err)}
	defer file.Close()

	w := bufio.NewWriter(file)
	n, err := w.WriteString("Let's write first string\n")
	if err != nil {log.Fatal(err)}
	n2, err := w.WriteString("Let's write second string")
	// everything will be written as one string if not add EOL
	if err != nil {log.Fatal(err)}
	w.Flush()
	fmt.Printf("%d bytes written\n", n + n2)
}

func ExampleBufioRead() {
	testFile := "bufio_test.txt"
	file, err := os.Open(testFile)
	if err != nil {log.Fatal(err)}
	defer file.Close()

	rd := bufio.NewReader(file)
	buf := make([]byte, 10)
	n, err := rd.Read(buf) // read to buf 10 bytes from open file via reader
	if err != nil && err != io.EOF {log.Fatal(err)}
	fmt.Printf("%d bytes are read: %s\n", n, buf)

	s, err := rd.ReadString('\n') // continue to read until EOL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", s)
}

func ExampleBufioScan() {
	file, err := os.Open("scan_test.txt")
	if err != nil {log.Fatal(err)}
	defer file.Close()

	s := bufio.NewScanner(file)
	// reads data line by line
	for s.Scan() { // return true until EOF
		fmt.Printf("%s\n", s.Text()) // data read on this iteration
	}
}

func TestInputOutput() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var sum int
	for {
		// s, err := r.ReadString('\n')
		b, _, err := r.ReadLine()
		if err != nil && err != io.EOF {log.Fatal(err)}
		if err == io.EOF {break}
		// s = strings.TrimRight(s, "\n")
		// fmt.Printf("%v - %T\n", s, s)
		s := string(b)
		if s == "" {break}
		d, err := strconv.Atoi(s)
		// fmt.Println(d)
		if err == nil {sum += d}
	}
	// fmt.Println(sum)
	_, err := w.WriteString(strconv.Itoa(sum))
	if err != nil {log.Fatal(err)}
	w.Flush()
	// fmt.Printf("%d bytes written\n", n)
}

func TestInputOutput2() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {break} // may be unnecessary in some cases
		sd, _ := strconv.Atoi(s)
		sum += sd
	}
	os.Stdout.WriteString(strconv.Itoa(sum))
}

func ExampleCSV() {
	// buf := bytes.NewBuffer(nil)
	buf, er := os.Create("test.csv")
	if er != nil {log.Fatalln("create file error", er)}
	defer buf.Close() // only in case of file

	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil {log.Fatal(err)}
	}
	w.Flush()

	w.WriteAll([][]string{ // we may write several rows at once as well
		{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
		{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	})

	buf.Seek(0, io.SeekStart) // in case of file only
	r := csv.NewReader(buf)
	for i := 1; i < 3; i++ {
		row, err := r.Read() // reading one row
		if err != nil && err != io.EOF {log.Fatal(err)} // should check for EOF
		fmt.Println(row)
	}

	data, err := r.ReadAll() // read remaining till EOF
	if err != nil {log.Fatal(err)}
	for _, row := range data {
		fmt.Println(row)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {return err}
	if info.IsDir() {
		if info.Name() == ".git" {return filepath.SkipDir}
		return nil
	}
	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func ExampleWalk() {
	const dir = "."

	if err := filepath.Walk(dir, walkFunc); err != nil {
		log.Fatalln("Some error", err)
	}
}

func FindCSV(path string, info os.FileInfo, err error) error {
	if err != nil {return err}
	if info.IsDir() {
		if info.Name() == ".git" {return filepath.SkipDir}
		return nil
	}
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	content, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("%s is not csv\n", info.Name())
		return nil
	}
	if len(content) < 10 || len(content[4]) < 10 {return nil}
	res := content[4][2]
	return errors.New(res)

}

func TestWalk() {
	const dir = "/home/andrey/Downloads/task_go_stepik"
	if err := filepath.Walk(dir, FindCSV); err != nil {
		fmt.Println(err)
		// log.Fatalln("Some error", err)
	}
}

func TestFindCSV2() {
	arch, err := zip.OpenReader("/home/andrey/Downloads/task.zip")
	if err != nil {log.Fatalln("error open archive", err)}
	defer arch.Close()
	for _, file := range arch.File {
		content, err := file.Open()
		if err != nil {log.Fatalln("error opening file", file.Name, err)}
		if rows, err := csv.NewReader(content).ReadAll(); err == nil && len(rows) == 10 && len(rows[4]) == 10 {
			fmt.Println(file.Name, rows[4][2])
		}
		content.Close()
	}
}

func findZeroInCSV() {
	// not good - read whole line at once - too much memory required
	dataPath := "/home/andrey/Downloads/task.data"
	dataFile, _ := os.Open(dataPath)
	defer dataFile.Close()
	reader := csv.NewReader(dataFile)
	reader.Comma = ';'
	reader.FieldsPerRecord = 1
	field, _ := reader.Read()
	for i, el := range field {
		if el == "0" {
			fmt.Println(i)
			break
		}
	}
}

func findZero2() {
	dataPath := "/home/andrey/Downloads/task.data"
	dataFile, _ := os.Open(dataPath)
	reader := bufio.NewReader(dataFile)
	defer dataFile.Close()
	for c := 1; ; c++ {
		field, _ := reader.ReadString(';')
		// fmt.Printf("%v, %T\n", field, field)
		// fmt.Println(c)
		if field == "0;" {
			fmt.Println("Found!", c)
			break
		}
		if c > 15000 {
			fmt.Println("Stop it!")
			break
		}
	}
}

func main() {
	// filesBeginning()
	// ExamplesWorkWithFiles()
	// ExampleWorkWithDir()
	// StrangeExample()
	// ExampleBufioWrite()
	// ExampleBufioRead()
	// ExampleBufioScan()
	// TestInputOutput()
	// TestInputOutput2()
	// ExampleCSV()
	// ExampleWalk()
	// TestWalk()
	// TestFindCSV2()
	findZeroInCSV()
	findZero2()
}
