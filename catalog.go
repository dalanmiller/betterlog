package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
	"os"
)

type Classes struct {
	classList []Class
}

type Class struct {
	title string
	num string
	faculty string
}

func main() {
	courses := getOldCatalog()
	fmt.Printf("# of Courses - %v", len(courses))

	b, err := json.Marshal(courses)
	if err != nil {
		fmt.Println("error:", err)
	}



	file, err := os.Open("classes.json")
    if err != nil {
        // handle the error here
        os.Create("classes.json")
    }

    file.Write(b)

    file.Close()
}

func getOldCatalog() []Class {

	classList := make([]Class)
	var doc *goquery.Document
	var e error

	doc, e = goquery.NewDocument("http://www.heinz.cmu.edu/academic-resources/course-results/index.aspx")
	if e != nil {
		panic(e)
	}

	doc.Find("table.course-list tr").Each(func(i int, s *goquery.Selection) {
		var num, title, fac string

		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch {
			case i == 0:
				num = s.Find("a").Text()
			case i == 1:
				title = s.Find("a").Text()
			case i == 2:
				fac = s.Find("a").Text()
			}

		})
		classList = append(classList, Class{num: num, title: title, faculty: fac})
	})


	cl := Classes{classList:classList}
	return cl

}
