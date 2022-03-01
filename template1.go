package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func Template1() {

	now := time.Now()
	targetDate := time.Date(2014, time.December, 31, 12, 13, 24, 0, time.Local)
	oneHourLater := now.Add(1 * time.Hour)

	data := struct {
		Key1            string
		Key2            string
		DateNow         string
		TargetDate      string
		DateNowAdd1Hour string
		Flag1           bool
	}{
		Key1:            "111111",
		Key2:            "2222222",
		DateNow:         now.Format("2006-01-02 15:04:05"),
		TargetDate:      targetDate.Format("2006-01-02 15:04:05"),
		DateNowAdd1Hour: oneHourLater.Format("2006-01-02 15:04:05"),
		Flag1:           true,
	}

	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
