CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    student_name VARCHAR(50) NOT NULL,
    city VARCHAR(50),
    starting_year SMALLINT 1NOT NULL
    )

CREATE TABLE IF NOT EXISTS lecturers (
    id SERIAL PRIMARY KEY,
    lecturer_name VARCHAR(100) NOT NULL,
    city VARCHAR(50)
    )

CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL,
    lecturer_id INTEGER NOT NULL,
    credit SMALLINT NOT NULL,
    CONSTRAINT fk_lecturers
        FOREIGN KEY(lecturer_id)
        REFERENCES lecturers(lecturer_id)
    )