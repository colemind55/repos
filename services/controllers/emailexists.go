package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"log"
	 "fmt"
	"time"
	"allsup.assessment/api/services/models"
)

type emailExistsController struct {
	urlPattern *regexp.Regexp
}


func (ex emailExistsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		var ( 
			caseID = ""
			email = "" 
		)

		matches := ex.urlPattern.FindStringSubmatch(r.URL.Path)

		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}


		// get the parameter string values for ?email=&caseID=
		_, err := url.Parse(r.URL.Path)
		if err != nil  {
			log.Fatal(err)
		}
		
		q := r.URL.Query()
		
		if  q["email"] == nil {
			log.Println("Email address is missing")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Email address is missing"))
			return
		}

		email = q.Get("email")

		if q["caseID"] != nil {
			caseID = q.Get("caseID")
		}

		switch r.Method {
		case http.MethodGet:
			ex.get(email, caseID, w)
		case http.MethodPost:
			ex.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	
}

func (ex *emailExistsController) get(email string, caseID string, w http.ResponseWriter) {
	isExists, err := models.Validate(email, caseID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(time.Now().Format(time.RFC822), " Response (", isExists, ") email (", email, ") caseID (", caseID, ")")

	encodeResponseAsJSON(isExists, w)
}

func (ex *emailExistsController) post(w http.ResponseWriter, r *http.Request) {
	Email, err := ex.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	isExists, err := models.Validate(Email.Address, Email.CaseID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(time.Now().Format(time.RFC822), " Response (", isExists, ") email (", Email.Address, ") caseID (", Email.CaseID, ")")
	encodeResponseAsJSON(isExists, w)
}

func (ex *emailExistsController) parseRequest(r *http.Request) (models.Email, error) {
	dec := json.NewDecoder(r.Body)
	var email models.Email
	err := dec.Decode(&email)
	if err != nil {
		return models.Email{}, err
	}
	return email, nil
}

func newEmailExistsController() *emailExistsController {
	return &emailExistsController{
		urlPattern: regexp.MustCompile(`^/emailExists`),
	}
}
