package lib

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type entryPayload struct {
	Semester  string     `json:"semestre"`
	QueryData *queryData `json:"queryData"`
}

type queryData struct {
	Data timeTable `json:"data"`
}

type timeTable struct {
	Majors []major `json:"especialidades_periodo"`
}

type major struct {
	Name  string `json:"especialidad__nombre"`
	ID    int    `json:"especialidad__id"`
	Plans []plan `json:"especialidad__planes"`
}

type plan struct {
	ID   int    `json:"plan__id"`
	Name string `json:"plan__nombre"`
}

type planSchedulePayload struct {
	Container *planDataContainer `json:"querygruposplan"`
}

type planDataContainer struct {
	Container *pspContainer `json:"data"`
}

type pspContainer struct {
	Container []*semester `json:"grupos_por_plan"`
}

type semester struct {
	Name    string             `json:"plan__bloque"`
	Courses []*courseContainer `json:"plan__grupos_bloque"`
}

type courseContainer struct {
	Course *course `json:"asignatura__asignatura"`
}

type course struct {
	ID   int    `json:"asignatura__id"`
	Name string `json:"asignatura__nombre"`
}

type courseOfferPayload struct {
	Container *scheduleContainer `json:"queryhorarios"`
}

type scheduleContainer struct {
	Container *scheduleDataContainer `json:"data"`
}

type scheduleDataContainer struct {
	Groups []courseGroup `json:"grupos_por_asignatura"`
}

type courseGroup struct {
	Details *groupDetails `json:"grupo__grupo"`
	Staff   []groupStaff  `json:"grupo__profesores"`
}

type groupDetails struct {
	Signups        int              `json:"grupo__alumnos"`
	Capacity       int              `json:"grupo__cupo"`
	HasDescription bool             `json:"grupo__tiene_presentacion"`
	Code           string           `json:"grupo__clave"`
	ID             int              `json:"grupo__id"`
	Name           *groupCourseName `json:"grupo__asignatura"`
}

type groupCourseName struct {
	Name string `json:"asignatura__nombre"`
}

type groupStaff struct {
	Details      *staffDetails       `json:"profesor__persona"`
	Availability []staffAvailability `json:"profesor__horarios"`
}

type staffDetails struct {
	ID        int    `json:"persona__id"`
	Name      string `json:"persona__nombre"`
	FirstName string `json:"persona__apellido_1"`
	LastName  string `json:"persona__apellido_2"`
}

type staffAvailability struct {
	Location    *staffAvailabilityLocation `json:"profesor_horario__lugar"`
	OnMonday    bool                       `json:"profesor_horario__lu"`
	OnTuesday   bool                       `json:"profesor_horario__ma"`
	OnWednesday bool                       `json:"profesor_horario__mi"`
	OnThursday  bool                       `json:"profesor_horario__ju"`
	OnFriday    bool                       `json:"profesor_horario__vi"`
	OnSaturday  bool                       `json:"profesor_horario__sa"`
	OnSunday    bool                       `json:"profesor_horario__do"`
	StartTime   string                     `json:"profesor_horario__hora_inicio"`
	EndTime     string                     `json:"profesor_horario__hora_termino"`
	Role        *staffRole                 `json:"grupo__cargo"`
}

type staffRole struct {
	Name string `json:"cargo__nombre_corto"`
}

type staffAvailabilityLocation struct {
	ID   int    `json:"lugar__id"`
	Name string `json:"lugar__nombre"`
}

const (
	baseURL = "https://pagina.fciencias.unam.mx/"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func Run() error {
	data, err := foo("docencia/horarios/indice")
	if err != nil {
		return err
	}
	content := &entryPayload{}
	if err := json.Unmarshal([]byte(data), content); err != nil {
		return err
	}

	plan := content.QueryData.Data.Majors[0].Plans[0]
	planData, err := foo(fmt.Sprintf("docencia/horarios/indiceplan/%s/%d", content.Semester, plan.ID))
	if err != nil {
		return err
	}

	planContent := &planSchedulePayload{}
	if err := json.Unmarshal([]byte(planData), planContent); err != nil {
		return err
	}

	course := planContent.Container.Container.Container[0].Courses[4].Course
	courseData, err := foo(fmt.Sprintf("docencia/horarios/20212/%d/%d", plan.ID, course.ID))
	if err != nil {
		return err
	}
	courseOffer := &courseOfferPayload{}
	if err := json.Unmarshal([]byte(courseData), courseOffer); err != nil {
		return err
	}

	fmt.Println(courseOffer)
	fmt.Println(planContent)
	fmt.Println(content)

	return nil
}

func foo(path string) (string, error) {
	res, err := http.Get(baseURL + path)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data string

	doc.Find("head script").Each(func(i int, selection *goquery.Selection) {
		if len(selection.Nodes) != 1 {
			return
		}
		node := selection.Nodes[0]
		selection = selection.FilterNodes(node)
		if _, ok := selection.Attr("data-drupal-selector"); !ok {
			return
		}
		data = node.FirstChild.Data
	})
	return data, nil
}
