package lib

const (
	majorsPathURL = "docencia/horarios/indice"
)

type pageMajorsMainPayload struct {
	Semester  string                   `json:"semestre"`
	Container *pageMajorsBaseContainer `json:"queryData"`
}

type pageMajorsBaseContainer struct {
	Container pageMajorsDataContainer `json:"data"`
}

type pageMajorsDataContainer struct {
	Majors []pageMajorsMajor `json:"especialidades_periodo"`
}

type pageMajorsMajor struct {
	Name  string                `json:"especialidad__nombre"`
	ID    int                   `json:"especialidad__id"`
	Plans []pageMajorsMajorPlan `json:"especialidad__planes"`
}

type pageMajorsMajorPlan struct {
	ID   int    `json:"plan__id"`
	Name string `json:"plan__nombre"`
}

func getMajors() (*pageMajorsMainPayload, error) {
	content := &pageMajorsMainPayload{}
	if err := loadPageContent(majorsPathURL, content); err != nil {
		return nil, err
	}
	return content, nil
}
