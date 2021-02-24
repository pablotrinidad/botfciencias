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

type Semester uint8

const (
	Elective Semester = iota
	First
	Second
	Third
	Fourth
	Fifth
	Sixth
	Seventh
	Eight
	Ninth
)

type Course struct {
	Semester Semester
	Name     string
	ID       int
}

type Group struct {
	CourseID       int
	Signups        int
	Capacity       int
	HasDescription bool
	Code           string
	ID             int
	Name           string
	Staff          []GroupStaff
}

type StaffRole uint8

const (
	UnknownStaffRole StaffRole = iota
	ProfessorStaffRole
	AssistantStaffRole
	LabAssistantStaffRole
)

type GroupStaff struct {
	ID           int
	FirstName    string
	MiddleName   string
	LastName     string
	Role         StaffRole
	Availability []StaffAvailability
}

type StaffAvailability struct {
	OnMonday    bool
	OnTuesday   bool
	OnWednesday bool
	OnThursday  bool
	OnFriday    bool
	OnSaturday  bool
	OnSunday    bool
	StartTime   string
	EndTime     string
	Location    *ClassLocation
}

type ClassLocation struct {
	ID   int
	Name string
}
