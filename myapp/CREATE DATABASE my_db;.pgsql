create DATABASE my_db;

CREATE TABLE student (
StdId int NOT NULL,
FirstName varchar(45) NOT NULL,
LastName varchar(45) DEFAULT NULL,
Email varchar(45) NOT NULL,
PRIMARY KEY (StdId),
UNIQUE (Email)
)

CREATE TABLE course (
cid varchar(32) NOT NULL,
coursename varchar(45) NOT NULL,
PRIMARY KEY (cid)
)

CREATE TABLE admin (
FirstName varchar(45) NOT NULL,
LastName varchar(45) DEFAULT NULL,
Email varchar(45) NOT NULL,
Password varchar(45) NOT NULL,
PRIMARY KEY (Email)
)

CREATE TABLE enroll (
std_id int NOT NULL,
course_id varchar(45) NOT NULL,
date_enrolled varchar(45) DEFAULT NULL,
PRIMARY KEY (std_id, course_id),
CONSTRAINT course_fk FOREIGN KEY (course_id) REFERENCES course
(cid) ON DELETE CASCADE ON UPDATE CASCADE,
CONSTRAINT std_fk FOREIGN KEY (std_id) REFERENCES student (StdId) ON
DELETE CASCADE ON UPDATE CASCADE
)