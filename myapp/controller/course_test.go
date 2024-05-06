package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	url := "http://localhost:8081/course"
	var jsonStr = []byte(`{"courseid":4,"coursename":"Solidity"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	expRes := `{
		"status": "course added"
	  }`

	assert.JSONEq(t, expRes, string(body))
}

func TestGetCourse(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8081/course/4")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)

	expRes := `{
		"courseid": 4,
		"coursename": "Solidity"
	  }`

	assert.JSONEq(t, expRes, string(body))
}

func TestDeleteCourse(t *testing.T) {
	url := "http://localhost:8081/course/4"
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	expRes := `{
		"status": "deleted"
	  }`

	assert.JSONEq(t, expRes, string(body))
}

func TestCourseNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8081/course/4")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error":"status not found"}`
	assert.JSONEq(expResp, string(body))
}
