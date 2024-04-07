package controller

import (
	"encoding/json"
	"myapp/model"
	"net/http"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	//create variable stud of type student
	var stud model.Student
	//read the request body and create a decoder object
	decoder := json.NewDecoder(r.Body)
	//store the json object data to stud variable
	if err := decoder.Decode(&stud); err != nil {

		// w.Write([]byte("Invalid json data"))

		response, _ := json.Marshal(map[string]string{"error": "invalid json body"})
		w.Header().Set("Contnt-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	//defer the closing of request body until the function returns
	defer r.Body.Close()

	//call the Create() using student object, stud
	saveErr := stud.Create()
	if saveErr != nil {
		// w.Write([]byte("Database error"))
		response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		w.Header().Set("Contnt-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	// no error
	// w.Write([]byte("response success"))
	response, _ := json.Marshal(map[string]string{"status": "student added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

}
