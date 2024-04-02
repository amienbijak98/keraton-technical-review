CREATE TABLE IF NOT EXISTS student (
    student_ID SERIAL PRIMARY KEY,
    student_name VARCHAR(50) NOT NULL,
    city VARCHAR(50),
    starting_year SMALLINT NOT NULL
    )

CREATE TABLE IF NOT EXISTS lecturer (
    lecturer_ID SERIAL PRIMARY KEY,
    lecturer_name VARCHAR(100) NOT NULL,
    city VARCHAR(50)
    )

CREATE TABLE IF NOT EXISTS course (
    course_ID SERIAL PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL,
    lecturer_ID INTEGER NOT NULL,
    credit SMALLINT NOT NULL,
    CONSTRAINT fk_lecturer
        FOREIGN KEY(lecturer_ID)
        REFERENCES lecturer(lecturer_ID)
    )