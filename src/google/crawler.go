package google

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func save() error {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println("err")
		return err
	}

	fmt.Println(db)

	return nil
}

func getHtml() (string, error) {
	url := "http://www.google.com"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

func Crawler() (string, error) {
	html, err := getHtml()

	if err != nil {
		return html, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(html))

	list := htmlquery.Find(doc, "//a")
}
