package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEnroll(t *testing.T) {
	url := "http://localhost:8081/enroll"
	var jsonStr = []byte(`{"stdid":1223,"cid":""csc101}`)

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
		"status": "enrolled"
	  }`

	assert.JSONEq(t, expRes, string(body))
}

func TestGetEnroll(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8081/enroll/1223/2")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)

	expRes := `{
		"stdid": 1223,
		"cid": "2",
		"date": "2024-04-30T07:04:26Z"
	  }`

	assert.JSONEq(t, expRes, string(body))
}

func TestDeleteEnroll(t *testing.T) {
	url := "http://localhost:8081/enroll/1223/2"
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

func TestEnrollNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8081/enroll/1224/2")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{
		"error": "No such enrollment"
	  }`
	assert.JSONEq(expResp, string(body))
}
