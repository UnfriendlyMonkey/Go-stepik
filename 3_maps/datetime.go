package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func datetimeBasics() {
	now := time.Now()

	certainTime := time.Date(
		1976,
		time.January,
		6,
		10,
		14,
		43,
		55,
		time.UTC,
	)

	unixTime := time.Unix(250000, 12)
	fmt.Println(now.Format("02-01-2006 15:04:05"))
	fmt.Println(certainTime.Format("06-01-1976 22:22:22"))
	fmt.Println(unixTime.Format("1917-01-01 13:13:13"))
}

func parsingTime() {
	outFormat := "02-01-2006 15:04:05"
	firstTime, _ := time.Parse("2006/01/02 15-04", "2023/12/20 17-57")
	fmt.Println(firstTime)
	fmt.Println(firstTime.Format(outFormat))

	loc, err := time.LoadLocation("Asia/Yekaterinburg")
	if err != nil {log.Fatal(err)}
	secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
	if err != nil {log.Fatal(err)}
	fmt.Println(secondTime, secondTime.Format(outFormat))

	fmt.Println(secondTime.Month(), secondTime.YearDay(), secondTime.Weekday(), secondTime.Unix())

	fmt.Println(firstTime.After(secondTime))
	fmt.Println(firstTime.Before(secondTime))
	fmt.Println(firstTime.Equal(secondTime))

	future := firstTime.Add(time.Hour * 12)
	past := firstTime.AddDate(-1, -2, -3)
	fmt.Println(future.Sub(past))
}

func convert() {
	// var a string
	// fmt.Scan(&a)
	a := "1986-04-06T05:20:00+06:00"
	outFormat := "Mon Jan 02 15:04:05 -0700 2006"
	// dt, err := time.Parse("2006-01-02T15:04:05-07:00", a)
	dt, err := time.Parse(time.RFC3339, a)
	if err != nil{fmt.Println(err)}
	// fmt.Println(dt)
	fmt.Println(dt.Format(outFormat))
	fmt.Println(dt.Format(time.UnixDate))
}

func reschedule() {
	// var a string
	reader := bufio.NewReader(os.Stdin)
	buf, _ := reader.ReadString('\n')
	buf = strings.TrimRight(buf, "\n")
	dt, err := time.Parse(time.DateTime, buf)
	if err != nil{log.Fatal(err)}
	if dt.Hour() > 12 {
		dt = dt.Add(time.Hour * 24)
	}
	fmt.Println(dt.Format(time.DateTime))
}

func exampleDuration() {
	now := time.Now()
	past := now.AddDate(0, 0, -1)
	future := now.AddDate(0, 0, 1)
	fmt.Println(time.Since(past).Round(time.Second))
	fmt.Println(time.Until(future).Round(time.Minute))
	dur, err := time.ParseDuration("1h3m44s")
	if err != nil{log.Fatal(err)}
	fmt.Println(dur.Round(time.Minute).Minutes())
}

func findDateDiff() {
	inFormat := "02.01.2006 15:04:05"
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	inDates := scaner.Text()
	fmt.Println(inDates)
	ds := strings.Split(inDates, ",")
	date1, _ := time.Parse(inFormat, ds[0])
	date2, _ := time.Parse(inFormat, ds[1])
	fmt.Println(date1, date2)
	fmt.Println(date1.Before(date2))
	var diff time.Duration
	if date1.Before(date2) {
		diff = date2.Sub(date1)
	} else {
		diff = date1.Sub(date2)
	}
	fmt.Println(diff)
}

const now = 1589570165

func timeAdding() {
	scaner := bufio.NewScanner(os.Stdin) // use fmt.Scanf next time!
	scaner.Scan()
	inStr := scaner.Text()
	// inStr := "12 мин. 13 сек."
	inStr = strings.Replace(inStr, "мин.", "m", -1)
	inStr = strings.Replace(inStr, "сек.", "s", -1)
	inStr = strings.ReplaceAll(inStr, " ", "")
	fmt.Println(inStr)
	inDur, err := time.ParseDuration(inStr)
	if err != nil {log.Fatal(err)}
	outTime := time.Unix(now, 0).Add(inDur).UTC()
	// fmt.Println(time.Unix(now, 0))
	// fmt.Println(time.Unix(now, 0).Add(inDur))
	fmt.Println(outTime.Format(time.UnixDate))
}

func timeAdding2() {
	var min, sec int64
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	fmt.Println(min, sec)
	// duration, err := time.ParseDuration(fmt.Sprintf("%sm%ss", min, sec))
	outTimestamp := now + min * 60 + sec
	outTime := time.Unix(outTimestamp, 0).UTC()
	fmt.Println(outTime.Format(time.UnixDate))
}

func main() {
	// datetimeBasics()
	// convert()
	// parsingTime()
	// reschedule()
	// exampleDuration()
	// findDateDiff()
	// timeAdding()
	timeAdding2()
}
