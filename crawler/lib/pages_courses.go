package lib

const (
	planCoursesPathURL = "docencia/horarios/indiceplan/%s/%d"
)

type pagePlanCourseMainPayload struct {
	Container *pagePlanCourseBaseContainer `json:"querygruposplan"`
}

type pagePlanCourseBaseContainer struct {
	Container *pagePlanCourseDataContainer `json:"data"`
}

type pagePlanCourseDataContainer struct {
	Semesters []pagePlanCourseSemester `json:"grupos_por_plan"`
}

type pagePlanCourseSemester struct {
	Name    string                                  `json:"plan__bloque"`
	Courses []pagePlanCourseSemesterCourseContainer `json:"plan__grupos_bloque"`
}

type pagePlanCourseSemesterCourseContainer struct {
	Course *pagePlanCourseSemesterCourse `json:"asignatura__asignatura"`
}

type pagePlanCourseSemesterCourse struct {
	ID   int    `json:"asignatura__id"`
	Name string `json:"asignatura__nombre"`
}

func getCourses(semester string, majorPlanID int) (*pagePlanCourseMainPayload, error) {
	content := &pagePlanCourseMainPayload{}
	if err := loadPageContent(majorsPathURL, content); err != nil {
		return nil, err
	}
	return content, nil
}
