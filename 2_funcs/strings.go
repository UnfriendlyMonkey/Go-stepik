package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ExampleString() {
    var s string = "Это строка"

    fmt.Printf("Length of the string: %d bytes\n", len(s))

    fmt.Printf("Let's print the 2-nd word in quotes: \"%v\"\n", s[7:])

    // s[3] = 12  // would be an compiler error

    for _, b := range s {
        fmt.Printf("%v ", b)  // print runes (int32) - one for each letter
        // fmt.Printf("%c ", b)  // print letters
    }
    fmt.Print("\n")

    for i := 0; i < len(s); i++ {
        fmt.Printf("%v ", s[i])  // print bytes (uint8) - may be 2 or more for letter
    }
    fmt.Printf("%T ", s[0])
    fmt.Print("\n")
}

func StringMethods() {
    fmt.Println(
        strings.Contains("test", "es"),
        strings.Count("test", "t"),
        strings.HasPrefix("test", "te"),
        strings.HasSuffix("test", "no"),
        strings.Index("test", "e"),
        strings.Join([]string{"hello", "world"}, ":"),
        strings.Repeat("a", 5),
        strings.Replace("blanotfoonotbarnot", "not", "***", -1),  // count less than zero mean all inclusions
        strings.Replace("blanotfoonotbarnot", "not", "***", 2),
        strings.Split("a-b-c-d-e", "-"),
        strings.ToLower("TEst"),
        strings.ToUpper("teST"),
        strings.Trim("tetstet", "te"),
        strings.Trim("tetsltetlwtet", "te"),  // removes only leading and ending code points (not from inside)
    )
}

func ByteSlice() {
    s := "Это байтовый срез"
    bs := []byte(s)
    fmt.Printf("%T - %T", s, bs)
    fmt.Printf("This is how byteSlice looks: %v\n", bs)

    for i := range bs {
        if bs[i] % 2 == 0 {
            bs[i] += 1
            continue
        }
        bs[i] -= 1
    }

    fmt.Printf("This is changed byteSlice as string: %s", bs)
}

func ExampleRune() {
    s := "Это срез рун"
    rs := []rune(s)
    fmt.Printf("%T - %T", s, rs)
    fmt.Printf("This is how runes looks: %v\n", rs)

    for i := range rs {
        if rs[i] == 'р' {
            rs[i] = '*'
        } else {rs[i]++}
    }
    fmt.Printf("Changed rune slice as string: %s\n", string(rs))
}

func UnicodeExample() {
    fmt.Println(unicode.IsDigit('1'))
    fmt.Println(unicode.IsLetter('ѱ'))  // true
    fmt.Println(unicode.IsLower('A'))
    fmt.Println(unicode.IsLower('ѱ'))  // true
    fmt.Println(unicode.IsSpace('\n'))
    fmt.Println(unicode.Is(unicode.Cyrillic, 'ы'))  // true

    fmt.Println(string(unicode.ToLower('F')))
}

func NonAsciiStringLen() {
    var en = "english"
    var ru = "русский"
    fmt.Println(len(en), len(ru))
    fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))
}

func IsStringCorrect() {
    var s string
    // fmt.Scanln(&s)
    s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
    // fmt.Println(s)
    s = strings.TrimRight(s, "\n")
    rs := []rune(s)
    if strings.HasSuffix(s, ".") && unicode.IsUpper(rs[0]) {
        fmt.Println("Right")
    } else {
        fmt.Println("Wrong")
    }
}

func IsPalindrome() {
		var s string
		s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.TrimRight(s, "\n")
		s = strings.ReplaceAll(s, " ", "")
		rs := []rune(s)
		// fmt.Println(rs)
		mid := len(rs) / 2
		right := len(rs) - 1
		// fmt.Println(mid)
		var isPal = true
		for i := 0; i < mid; i++ {
				// fmt.Println(i, rs[i], " - ", rs[right - i], right - i)
				if rs[i] != rs[right - i] {
						isPal = false
						break
				}
		}
		if isPal {fmt.Println("Палиндром")} else {fmt.Println("Нет")}
}

func FindSubstring() {
		var mainS, subS string
		fmt.Scan(&mainS, &subS)
		fmt.Println(strings.Index(mainS, subS))
		// return strings.Index(mainS, subS)
}

func OddLettersOnly() {
		var s string
		reader := bufio.NewReader(os.Stdin)
		s, _ = reader.ReadString('\n')
		s = strings.TrimRight(s, "\n")
		// s = strings.ReplaceAll(s, " ", "")
		rs := []rune(s)
		for i := 1; i < len(rs); i += 2 {
				fmt.Printf("%c", rs[i])
		}
}

func UniqOnly() {
		var s string
		fmt.Scan(&s)
		rs := []rune(s)
		for _, el := range rs {
				if strings.Count(s, string(el)) == 1 {fmt.Printf("%c", el)}
		}
}

func CorrectPassword() {
		var s string
		fmt.Scan(&s)
		rs := []rune(s)
		var correct = true
		if len(rs) < 5 {correct = false}
		if correct {
				for _, el := range rs {
						if !unicode.Is(unicode.Latin, el) && !unicode.Is(unicode.Digit, el) {
								correct = false
								break
						}
				}
		}
		// fmt.Println(unicode.Is(unicode.Latin, rs[0]) || unicode.Is(unicode.Digit, rs[0]))
		if correct {fmt.Println("Ok")} else {fmt.Println("Wrong password")}
}

func main() {
    // ExampleString()
    // StringMethods()
    // ByteSlice()
    // ExampleRune()
    // UnicodeExample()
    // NonAsciiStringLen()
    // IsStringCorrect()
		// IsPalindrome()
		// FindSubstring()
		// OddLettersOnly()
		// UniqOnly()
		CorrectPassword()
}
