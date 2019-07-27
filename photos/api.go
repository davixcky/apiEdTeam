package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	JSONURL = "https://jsonplaceholder.typicode.com/photos"
)

type Photos struct{}

var client = &http.Client{}

// Consume la api
func (*Photos) run(method string, id int) (resp *http.Response, code int, err error) {
	var url string
	if id == 0 || strings.ToUpper(method) == "POST" {
		url = fmt.Sprintf("%s", JSONURL)
	} else {
		url = fmt.Sprintf("%s/%d", JSONURL, id)
	}

	req, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	return resp, resp.StatusCode, nil
}

//Realiza la peticion segun su verbo (GET, PUT, POST, DELETE)
func (p *Photos) GetOne(id int) (resp *http.Response, code int, err error) {
	return p.run("GET", id)
}

func (p *Photos) GetAll() (resp *http.Response, code int, err error) {
	return p.run("GET", 0)
}

func (p *Photos) Delete(id int) (resp *http.Response, code int, err error) {
	return p.run("DELETE", id)
}

func (p *Photos) Update(id int) (resp *http.Response, code int, err error) {
	return p.run("PUT", id)
}

func (p *Photos) Post() (resp *http.Response, code int, err error) {
	return p.run("POST", 0)
}
