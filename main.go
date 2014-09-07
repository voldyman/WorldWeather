package worldweather

import (
	"fmt"
	"html/template"
	"net/http"

	"appengine"
)

var indexTemplate = template.Must(template.ParseFiles("templates/index.tpl", "templates/base.tpl"))
var aboutTemplate = template.Must(template.ParseFiles("templates/about.tpl", "templates/base.tpl"))

var apiManager = NewAPIRequestManager()

func init() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/weather", weatherHandler)

	go apiManager.RunWorker()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := indexTemplate.ExecuteTemplate(w, "base", map[string]string{
		"Title": "Home",
	})

	if err != nil {
		fmt.Fprintf(w, "error occured executing the template index.tpl: %s", err.Error())
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := aboutTemplate.ExecuteTemplate(w, "base", map[string]string{
		"Title": "About",
	})

	if err != nil {
		fmt.Fprintf(w, "error occured executing the template index.tpl: %s", err.Error())
	}
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	r.ParseForm()
	city := r.FormValue("city")
	//output, err := fetchWeatherData(ctx, city)

	out := apiManager.AddItem(ctx, city)

	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//} else {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.Write(<-out)
	//}
}
