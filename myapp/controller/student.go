package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	//validate cookie
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	saveErr := stud.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// no error, success case
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "student added"})
}

// helper function
func extractURLvar(r *http.Request) (int64, error) {
	m := mux.Vars(r)
	sid := m["sid"]
	stdId, idErr := getUserId(sid)
	return stdId, idErr
}

// GetStud function
// to get the info of th e student that is already in the database
func GetStud(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	//grt url parameter
	stdId, idErr := extractURLvar(r)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)
}

//convert string sid to int

func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

// function to update the student info that is already in the database table
func UpdateStud(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	//getting id of the stud to update from url paramater
	old_sid := mux.Vars(r)["sid"]

	//converting string id into int id
	old_stdId, idErr := getUserId(old_sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	//creating student instance
	var stud model.Student

	//getting new student details from request body and convert json object to go objext
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	//calling update method from model package
	updateErr := stud.Update(old_stdId)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, stud)
	}

}

// deleting method
func DeleteStud(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	stdId, idErr := extractURLvar(r)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	students, getErr := model.GetAllStuds()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, students)
}
