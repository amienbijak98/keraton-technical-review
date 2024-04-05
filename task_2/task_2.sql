CREATE TABLE IF NOT EXISTS course_taken (
    id SERIAL PRIMARY KEY,
    course_id INTEGER NOT NULL,
    student_id INTEGER NOT NULL,
    number_semester SMALLINT NOT NULL,
    mid_exam_date DATE NOT NULL,
    mid_exam_score SMALLINT,
    end_exam_date DATE NOT NULL,
    end_exam_score SMALLINT,
    final_score SMALLINT GENERATED ALWAYS AS ((mid_exam_score * 0.5) + (end_exam_score * 0.5)) STORED,
    CONSTRAINT fk_students
        FOREIGN KEY(student_id)
        	REFERENCES students(id),
    CONSTRAINT fk_courses
    	FOREIGN KEY(course_id)
        	REFERENCES courses(id)
    )