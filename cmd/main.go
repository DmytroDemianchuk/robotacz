package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html",
		"../templates/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func thx(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html",
		"../templates/create.html", "../templates/thx.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "thx", nil)
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/thx", thx)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()

	srv := gin.New()

	cfg, err := config.Parse()
	if err != nil {
		logrus.Fatalf("error psring config: %s", err.Error())
	}

	db, err := database.CreateConn(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.SSLMode)
	if err != nil {
		logrus.Fatalf("failed to connection db: %s", err.Error())
	}

	// init deps

	peopleRepository := repository.NewPeople(db)
	peopleService := service.NewPeople(peopleRepository)
	peopleTransport := rest.NewPeople(peopleService)

	srv.GET("/musics", peopleTransport.List)
	srv.GET("/music/:id", peopleTransport.Get)
	srv.POST("/music", peopleTransport.Create)
	srv.PUT("/music/:id", peopleTransport.Update)
	srv.DELETE("/music/:id", peopleTransport.Delete)

	if err := srv.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		logrus.Fatalf("error occured while running http server %s", err.Error())
	}
}
