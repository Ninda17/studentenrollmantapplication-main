package model

import (
	"myapp/datastore/postgres"
)

type Course struct {
	CourseID   string `json:"courseid"`
	CourseName string `json:"coursename"`
}

const queryInsertCourse = "INSERT INTO course (cid, coursename) VALUES ($1, $2);"

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertCourse, c.CourseID, c.CourseName)
	return err
}

//searching course using courseID

const queryGetCourse = "SELECT * FROM course WHERE cid = $1;"

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCourse, c.CourseID).Scan(&c.CourseID, &c.CourseName)
}

// creating new reciver function to update
const queryUpdateCourse = "UPDATE course SET cid = $1, coursename = $2 WHERE cid = $3 RETURNING cid;"

func (c *Course) Update(oldID string) error {
	err := postgres.Db.QueryRow(queryUpdateCourse, c.CourseID, c.CourseName, oldID).Scan(&c.CourseID)
	return err
}

//creating delete query

const queryDeleteCourse = "DELETE FROM course WHERE cid = $1 RETURNING cid;"

func (c *Course) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteCourse, c.CourseID).Scan(&c.CourseID); err != nil {
		return err
	}
	return nil
}

//getall query
//get all course

func GetAllCourse() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * from course")
	if getErr != nil {
		return nil, getErr
	}

	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.CourseID, &c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}
