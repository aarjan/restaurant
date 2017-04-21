package main

import (
	"net/http"

	"encoding/json"

	"fmt"

	"github.com/aarjan/restaurant/config"
	"github.com/aarjan/restaurant/models"
)

type result struct {
	Meta
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "hello", http.StatusInternalServerError)
}
func VATGet(w http.ResponseWriter, r *http.Request) {
	res := result{}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vats := []models.VAT{}
	err := config.DB.Find(&vats).Error
	// if err != nil {
	// 	res = wrapper(http.StatusInternalServerError, err.Error(), nil)
	// 	json.NewEncoder(w).Encode(res)
	// }

	res = wrapper(http.StatusOK, "ok", vats)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println(err)
	}

}

func wrapper(c int, m string, i interface{}) result {

	return result{
		Meta: Meta{
			Code:    c,
			Message: m,
		},
		Data: i,
	}

}
