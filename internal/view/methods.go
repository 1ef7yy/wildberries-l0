package view

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (v *view) GetData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		v.Logger.Info("Id is empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Id is empty"))
		return
	}

	data, err := v.domain.GetDataByID(id)
	if err != nil {
		v.Logger.Error("Error getting data: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error: "+err.Error())
		return
	}
	if data.OrderUid == "" {
		v.Logger.Info(fmt.Sprintf("Data with id: %s not found", id))
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Data with id: %s not found", id)
		return
	}

	resp, err := json.Marshal(&data)
	if err != nil {
		v.Logger.Error("Error marshaling data: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error: "+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (v *view) RestoreCache() error {
	return v.domain.RestoreCache()
}
