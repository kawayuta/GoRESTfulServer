package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"strconv"
	"io/ioutil"
	"io"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcmoe!Top!")
}


// TodoIndex

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}


// TodoShow

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idParam := ps.ByName("todoId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == ""{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return

}


// TodoCreate

func TodoCreate(w http.ResponseWriter, r * http.Request, ps httprouter.Params) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s%d", r.Host, t.ID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}
