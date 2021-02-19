package lib

type Major struct {
	Name  string
	ID    int
	Plans []MajorPlan
}

type MajorPlan struct {
	Name string
	ID   int
}

type Course struct {
	Semester int
}
