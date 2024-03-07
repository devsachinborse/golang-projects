package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Hanlder returned wrong status code: got %v want %v", status,
			http.StatusOK)
	}

}

//test create todo

func TestCreateTodo(t *testing.T) {
	jsonStr := []byte(`{"title":"Test Todo", "status":"pending" }`)

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Hanlder returned wrong status code: got %v want %v", status,
			http.StatusOK)
	}

}

func TestUpdateTodo(t *testing.T) {
	jsonStr := []byte(`{"title":"Update todo", "status":"pending"}`)

	req, err := http.NewRequest("PUT", "/todos/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Hanlder returned wrong status code: got %v want %v", status,
			http.StatusOK)
	}
}

func TestDeleteTodo(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/todos/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wroug status code: got %v want %v", status, http.StatusOK)
	}
}
