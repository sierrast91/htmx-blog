package route

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"htmx-blog-app/internal/utils"
)

type handler func(w http.ResponseWriter, r *http.Request)

var (
	db   *sql.DB
	tmpl *template.Template
)

func NewRouter(store *sql.DB, template *template.Template) *http.ServeMux {

	db = store
	tmpl = template

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))
	router.HandleFunc("GET /assets/", logMiddle(http.StripPrefix("/assets/", fs).ServeHTTP))

	router.HandleFunc("GET /", logMiddle(indexHandler))
	router.HandleFunc("GET /blogs", logMiddle(blogsHandler))
	router.HandleFunc("GET /blogs/{id}", logMiddle(blogHandler))

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if err := render(w, "index.html", nil); err != nil {
		utils.Err(w, err, http.StatusInternalServerError)
		return
	}
}
func blogsHandler(w http.ResponseWriter, r *http.Request) {
	bs := dirToBlogList("./blogs/")
	if err := render(w, "blogs.html", bs); err != nil {
		utils.Err(w, err, http.StatusInternalServerError)
	}
}
func blogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.Err(w, err, http.StatusInternalServerError)
	}
	bs := dirToBlogList("./blogs/")
	if err := render(w, "blog.html", bs[id-1]); err != nil {
		utils.Err(w, err, http.StatusInternalServerError)
	}
}

func render(w http.ResponseWriter, name string, data interface{}) error {
	tmpl := template.Must(template.ParseGlob("templates/*/*.html"))
	return tmpl.ExecuteTemplate(w, name, data)
}

func logMiddle(next handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next(w, r)
	}
}
