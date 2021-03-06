package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	. "webcita/src/github.com/lss/webapp/viewmodel"
)

func main() {

	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		requestedFile := r.URL.Path[1:]
		t := templates[requestedFile + ".html"]
		context := NewBase()
		if t != nil {
			err := t.Execute(w, context)
			if err != nil {
				log.Println(err)
			}
		}else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8383", nil)
}


func populateTemplates() map[string] *template.Template {
	result := make(map[string] *template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+ "/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")

	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template ' " + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}

		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result


}

