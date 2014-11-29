package main

import (
	"errors"
	"net/http"
)

type action func(rw http.ResponseWriter, r *http.Request) error

type appController struct{}

func (c *appController) action(a action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			switch err {
			case Err400:
				http.Error(rw, err.Error(), http.StatusBadRequest)
			case Err404:
				http.Error(rw, err.Error(), http.StatusNotFound)
			default:
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}
	})
}

var (
	Err404 = errors.New("Not found")
	Err400 = errors.New("Bad input")
	Err500 = errors.New("Internal Server Error")
)
