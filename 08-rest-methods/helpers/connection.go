package connections

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GET() (string, error) {
	res, err := http.Get("https://httpbin.org/get?cc=IN")
	if err != nil {
		return "not ok", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "not ok", err
	}
	return string(body), nil
}

func POST() (string, error) {
	type Person struct {
		Name string
		Age  int
	}

	data, err := json.Marshal(Person{Name: "Sushant", Age: 30})
	if err != nil {
		return "not ok", err
	}
	res, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "not ok", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "not ok", err
	}
	return string(body), nil
}

func PUT() (string, error) {
	type Employee struct {
		Name        string
		Designation string
	}
	data, err := json.Marshal(Employee{
		Name:        "Sushant",
		Designation: "Senior Software Engineer",
	})
	if err != nil {
		return "not ok", err
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, "https://httpbin.org/put", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "not ok", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "not ok", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "not ok", err
	}
	return string(body), nil
}

func DELETE() (string, error) {

	type Movie struct {
		Name   string
		Rating int
	}

	movie := Movie{
		Name:   "Scott Pilgrim vs the World",
		Rating: 10,
	}

	client := &http.Client{}
	data, err := json.Marshal(movie)
	if err != nil {
		return "not ok", err
	}
	req, err := http.NewRequest(http.MethodDelete, "https://httpbin.org/delete", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "not ok", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "not ok", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "not ok", err
	}

	return string(body), nil
}
