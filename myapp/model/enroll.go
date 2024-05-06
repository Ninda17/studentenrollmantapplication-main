package model

import "myapp/datastore/postgres"

type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"cid"`
	Date_Enrolled string `json:"date"`
}

const (
	queryEnrollStd = "INSERT INTO enroll(std_id, course_id,date_enrolled) VALUES($1, $2, $3) RETURNING std_id;"
)

func (e *Enroll) EnrollStud() error {
	row := postgres.Db.QueryRow(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled)
	err := row.Scan(&e.StdId)
	return err
}

// get enrollment
const queryGetEnroll = "SELECT std_id, course_id, date_enrolled FROM enroll WHERE std_id=$1 and course_id = $2;"

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll, e.StdId, e.CourseID).Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
}

// get all enrolls
func GetAllEnrolls() ([]Enroll, error) {
	rows, getErr := postgres.Db.Query("SELECT std_id, course_id,date_enrolled from enroll;")
	if getErr != nil {
		return nil, getErr
	}
	//create a slice of type course
	enrolls := []Enroll{}

	for rows.Next() {
		var e Enroll
		dbErr := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if dbErr != nil {
			return nil, dbErr
		}
		enrolls = append(enrolls, e)
	}
	rows.Close()
	return enrolls, nil
}

// delete
const queryDeleteEnroll = "DELETE FROM enroll WHERE std_id=$1 and course_id =$2 RETURNING std_id;"

func (e *Enroll) Delete() error {
	row := postgres.Db.QueryRow(queryDeleteEnroll, e.StdId, e.CourseID)
	err := row.Scan(&e.StdId)
	return err
}
