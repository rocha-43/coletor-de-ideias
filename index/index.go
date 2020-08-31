package index

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK! -index")

	if r.Method != "GET" {
		fmt.Printf("METHOD ERROR - just method GET avaliable...")
	}

	fp := path.Join("templates", "index.html")
	temp, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := temp.Execute(w, ""); err != nil {
		fmt.Printf("SERVER ERROR - template execution error...")
	}
}
