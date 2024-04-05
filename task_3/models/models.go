package models

type Student struct {
	ID           uint          `gorm:"primary_key;autoincrement" json:"id"`
	StudentName  string        `gorm:"type:varchar(50);not null" json:"student_name"`
	City         string        `gorm:"type:varchar(50)" json:"city"`
	StartingYear int16         `gorm:"not null" json:"starting_year"`
	CourseTaken  []CourseTaken `json:"-"`
}

type StudentDetails struct {
	ID           uint                `json:"id"`
	StudentName  string              `json:"student_name"`
	City         string              `json:"city"`
	StartingYear int16               `json:"starting_year"`
	CourseTaken  []CourseTakenDetail `json:"course_taken"`
}

type Lecturer struct {
	ID           uint   `gorm:"primary_key;autoincrement" json:"lecturer_id"`
	LecturerName string `gorm:"type:varchar(100);not null" json:"lecturer_name"`
	City         string `gorm:"type:varchar(50)" json:"city"`
	Course       []Course
}

type Course struct {
	ID          uint          `gorm:"primary_key;autoincrement" json:"course"`
	CourseName  string        `gorm:"type:varchar(100);not null" json:"course_name"`
	LecturerID  uint          `gorm:"not null" json:"lecturer_id"`
	Credit      int16         `gorm:"not null" json:"credit"`
	Lecturer    Lecturer      `json:"-"`
	CourseTaken []CourseTaken `json:"-"`
}

type CourseTaken struct {
	ID             uint    `gorm:"primary_key,autoincrement" json:"-"`
	CourseID       uint    `gorm:"not null" json:"course_id"`
	StudentID      uint    `gorm:"not null" json:"-"`
	NumberSemester int16   `gorm:"not null" json:"number_semester"`
	MidExamDate    string  `gorm:"type:date;not null" json:"mid_exam_date"`
	MidExamScore   int16   `json:"mid_exam_score"`
	EndExamDate    string  `gorm:"type:date;not null" json:"end_exam_date"`
	EndExamScore   int16   `json:"end_exam_score"`
	FinalScore     int8    `json:"final_score"`
	Student        Student `json:"-"`
	Course         Course  `json:"-"`
}

type CourseTakenDetail struct {
	ID             uint   `json:"-"`
	CourseID       uint   `json:"course_id"`
	CourseName     string `json:"course_name"`
	NumberSemester int16  `json:"number_semester"`
	StudentID      uint   `json:"-"`
	LecturerID     uint   `json:"lecturer_id"`
	LecturerName   string `json:"lecturer_name"`
	MidExamDate    string `json:"mid_exam_date"`
	MidExamScore   int16  `json:"mid_exam_score"`
	EndExamDate    string `json:"end_exam_date"`
	EndExamScore   int16  `json:"end_exam_score"`
	FinalScore     int8   `json:"final_score"`
}
