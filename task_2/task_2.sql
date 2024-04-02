CREATE TABLE IF NOT EXISTS exam (
    exam_ID SERIAL PRIMARY KEY,
    exam_date DATE NOT NULL,
    number_semester SMALLINT NOT NULL,
    exam_period VARCHAR(10) NOT NULL,
    student_ID INTEGER NOT NULL,
    course_ID INTEGER NOT NULL,
    score SMALLINT,
    CONSTRAINT fk_student
        FOREIGN KEY(student_ID)
        	REFERENCES student(student_ID),
    CONSTRAINT fk_course
    	FOREIGN KEY(course_ID)
        	REFERENCES course(course_ID)
    )