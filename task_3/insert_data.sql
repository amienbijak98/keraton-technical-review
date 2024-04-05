-- Memasukkan semua data awal.
INSERT INTO students (student_name, city, starting_year)
VALUES 
    ('Bijak Amien Muttaqie', 'Surabaya', 2019),
    ('Icha Resky', 'Jakarta', 2020),
    ('Budi Santoso', 'Bandung', 2018);

INSERT INTO lecturers (lecturer_name, city)
VALUES 
    ('Dr. I Made Sujana', 'Denpasar'),
    ('Prof. Budi Laksono', 'Yogyakarta'),
    ('Dr. Bagus Prajaka', 'Surabaya');

INSERT INTO courses (course_name, lecturer_id, credit)
VALUES 
    ('Pengantar Ilmu Komputer', 1, 3),
    ('Struktur Data dan Algoritma', 2, 4),
    ('Sistem Manajemen Basis Data', 3, 3);

INSERT INTO course_takens (course_id, student_id, number_semester, mid_exam_date, mid_exam_score, end_exam_date, end_exam_score)
VALUES 
    (1, 1, 1, '2023-05-10', 80, '2023-07-15', 85),
    (2, 2, 2, '2023-06-20', 75, '2023-08-25', 80),
    (3, 3, 1, '2023-05-05', 85, '2023-07-10', 90),
    (1, 2, 1, '2023-05-10', 80, '2023-07-15', 80),
    (2, 3, 2, '2023-06-20', 85, '2023-08-25', 90),
    (3, 1, 1, '2023-05-05', 90, '2023-07-10', 95);