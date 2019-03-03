package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {

	http.HandleFunc("/v1/contacts", contactsHandleFunc)
	http.ListenAndServe(":8080", nil)
}

func contactsHandleFunc(w http.ResponseWriter, r *http.Request) {
	regExpContactID := regexp.MustCompile(`/contacts/[0-9]+`)
	regExpTel := regexp.MustCompile(`[0-9]{11}`)

	c := contact{}
	c.Dir = "./data/"
	//check if here can be empty string
	if rawTel := regExpContactID.FindString(r.URL.Path); rawTel != "" {
		//validate
		arrRawTel := strings.Split(rawTel, "/")
		c.Tel = arrRawTel[len(arrRawTel)-1]
		if !regExpTel.MatchString(c.Tel) {
			w.WriteHeader(404)
			return
		}

		//Obtain contact
		if r.Method == "GET" {

			err := c.showContact()
			if err != nil {
				w.WriteHeader(500)
				return
			}
			enc, err := json.Marshal(c)
			if err != nil {
				w.WriteHeader(500)
				return
			}
			w.Write(enc)

			//Upd contact
		} else if r.Method == "POST" {
			c, err := getContactFromRequest(r)
			if err != nil {
				w.WriteHeader(400)
				return
			}
			err = c.updateContact()
			if err != nil {
				w.WriteHeader(500)
				return
			}
			//Delete contact
		} else if r.Method == "DELETE" {
			c, err := getContactFromRequest(r)
			if err != nil {
				w.WriteHeader(400)
				return
			}
			err = c.deleteContact()
			if err != nil {
				w.WriteHeader(500)
				return
			}
		}
		//Obtain contacts
	} else if r.Method == "GET" {
		contacts, err := indexContact(c.Dir)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		enc, err := json.Marshal(contacts)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Write(enc)
		//Create contact
	} else if r.Method == "POST" {
		c, err := getContactFromRequest(r)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		c.createContact()
	}
}

func getContactFromRequest(r *http.Request) (c contact, err error) {
	text, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return
	}
	dec := json.NewDecoder(strings.NewReader(string(text)))
	err = dec.Decode(&c)
	if err != nil {
		return
	}
	return
}
