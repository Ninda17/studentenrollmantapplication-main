package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

// adding course
func Addcourse(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	var cors model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cors); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	saveErr := cors.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	// no error, success case
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "course added"})
}

// helper function
func extractingURLvar(r *http.Request) (string, error) {
	m := mux.Vars(r)
	return m["cid"], nil
}

// get Course function
func GetCourse(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	courseId, idErr := extractingURLvar(r)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadGateway, idErr.Error())
		return
	}

	c := model.Course{CourseID: courseId}
	getErr := c.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "status not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func getCourseId(userIdParam string) (string, error) {
	return userIdParam, nil
}

// function to update
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	old_cid := mux.Vars(r)["cid"]

	//converting string id into int id
	old_cid, idErr := getCourseId(old_cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	//creating course instance
	var course model.Course

	//getting new course details from request body and convert json object to go objext
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	//calling update method from model package
	updateErr := course.Update(old_cid)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "course not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, course)
	}
}

// deletig method
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}
	courseId, idErr := extractingURLvar(r)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := model.Course{CourseID: courseId}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	// validate cookie
	if !VerifyCookie(w, r) {
		return
	}

	courses, getErr := model.GetAllCourse()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
