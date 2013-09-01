package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func main() {
	getOldCatalog()
	http.HandleFunc("/", goGoCatalog)
	fmt.Println("Serving catalog...")
	if os.Getenv("PORT") != "" {
		fmt.Println(os.Getenv("PORT"))
	} else {
		fmt.Println("wtf")
	}
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func getOldCatalog() {

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

		fmt.Printf("%s - %s - %s\n", num, title, fac)

	})

	// resp, err := http.Get("http://www.heinz.cmu.edu/academic-resources/course-results/index.aspx")

	// if err != nil {
	// 	panic(err)
	// }

	// s := strings.NewReader("<a></a>")
	// a, err := html.Parse(s)

	// fmt.Println(a)

	// fmt.Println(resp)
	// fmt.Println('\n')
	// fmt.Println(resp.Body)

	// thing, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	// thing2 := strings.NewReader(string(thing))
	// doc, err := html.Parse(thing2)

	// fmt.Printf("%s", doc)

}

func goGoCatalog(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(res, "hello, world")
}
