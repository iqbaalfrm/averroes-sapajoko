package http

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Berhasil bool        `json:"berhasil"`
	Pesan    string      `json:"pesan"`
	Data     interface{} `json:"data,omitempty"`
	Kesalahan interface{} `json:"kesalahan,omitempty"`
}

func TulisJSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func ResponSukses(w http.ResponseWriter, status int, pesan string, data interface{}) {
	TulisJSON(w, status, APIResponse{
		Berhasil: true,
		Pesan:    pesan,
		Data:     data,
	})
}

func ResponGagal(w http.ResponseWriter, status int, pesan string, kesalahan interface{}) {
	TulisJSON(w, status, APIResponse{
		Berhasil:  false,
		Pesan:     pesan,
		Kesalahan: kesalahan,
	})
}
