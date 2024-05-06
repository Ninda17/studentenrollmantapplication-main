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