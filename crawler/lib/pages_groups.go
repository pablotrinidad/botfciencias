package lib

import "fmt"

const (
	courseGroupsPathURL = "docencia/horarios/%s/%d/%d"
)

type pageCourseGroupMainPayload struct {
	Container *pageCourseGroupBaseContainer `json:"queryhorarios"`
}

type pageCourseGroupBaseContainer struct {
	Container *pageCourseGroupDataContainer `json:"data"`
}

type pageCourseGroupDataContainer struct {
	Groups []pageCourseGroup `json:"grupos_por_asignatura"`
}

type pageCourseGroup struct {
	Details *pageCourseGroupDetails `json:"grupo__grupo"`
	Staff   []pageCourseGroupStaff  `json:"grupo__profesores"`
}

type pageCourseGroupDetails struct {
	Signups        int                  `json:"grupo__alumnos"`
	Capacity       int                  `json:"grupo__cupo"`
	HasDescription bool                 `json:"grupo__tiene_presentacion"`
	Code           string               `json:"grupo__clave"`
	ID             int                  `json:"grupo__id"`
	Name           *pageGroupCourseName `json:"grupo__asignatura"`
}

type pageGroupCourseName struct {
	Name string `json:"asignatura__nombre"`
}

type pageCourseGroupStaff struct {
	Details      *pageCourseGroupStaffDetails       `json:"profesor__persona"`
	Availability []pageCourseGroupStaffAvailability `json:"profesor__horarios"`
}

type pageCourseGroupStaffDetails struct {
	ID         int    `json:"persona__id"`
	FirstName  string `json:"persona__nombre"`
	MiddleName string `json:"persona__apellido_1"`
	LastName   string `json:"persona__apellido_2"`
}

type pageCourseGroupStaffAvailability struct {
	Location    *pageCourseGroupStaffAvailabilityLocation `json:"profesor_horario__lugar"`
	OnMonday    bool                                      `json:"profesor_horario__lu"`
	OnTuesday   bool                                      `json:"profesor_horario__ma"`
	OnWednesday bool                                      `json:"profesor_horario__mi"`
	OnThursday  bool                                      `json:"profesor_horario__ju"`
	OnFriday    bool                                      `json:"profesor_horario__vi"`
	OnSaturday  bool                                      `json:"profesor_horario__sa"`
	OnSunday    bool                                      `json:"profesor_horario__do"`
	StartTime   string                                    `json:"profesor_horario__hora_inicio"`
	EndTime     string                                    `json:"profesor_horario__hora_termino"`
	Role        *pageCourseGroupStaffRole                 `json:"grupo__cargo"`
}

type pageCourseGroupStaffAvailabilityLocation struct {
	ID   int    `json:"lugar__id"`
	Name string `json:"lugar__nombre"`
}

type pageCourseGroupStaffRole struct {
	Name string `json:"cargo__nombre_corto"`
}

func getGroups(semester string, majorPlanID int, courseID int) (*pageCourseGroupMainPayload, error) {
	content := &pageCourseGroupMainPayload{}
	url := fmt.Sprintf(courseGroupsPathURL, semester, majorPlanID, courseID)
	if err := loadPageContent(url, content); err != nil {
		return nil, err
	}
	return content, nil
}
