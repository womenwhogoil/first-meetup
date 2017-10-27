package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	port = "8080"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	const html = `
<html>
	<body>
		<img src={{.ImageURL}}>
		<p>
			<b>{{.Text}}</b>
		</p>
	</body>
</html>`
	t, err := template.New("webpage").Parse(html)
	if err != nil {
		log.Fatal(err)
	}

	wwgPage := struct {
		ImageURL string
		Text     string
	}{
		ImageURL: "https://secure.meetupstatic.com/photos/event/2/c/1/8/global_463271288.jpeg",
		Text:     "Welcome to Women Who Go IL!",
	}

	if err = t.Execute(w, wwgPage); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("Started server at http://127.0.0.1:%+v\n", port)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":"+port, nil)
}
