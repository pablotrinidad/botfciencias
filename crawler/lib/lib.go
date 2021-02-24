// Package lib exposes a series of methods for downloading the course offer from UNAM's Faculty
// of Science using the website pagina.fciencias.unam.mx
package lib

func GetMajors() ([]Major, error) {
	pageMajors, err := getMajors()
	if err != nil {
		return nil, err
	}

	majors := make([]Major, len(pageMajors.Container.Container.Majors))
	for i := range pageMajors.Container.Container.Majors {
		raw := pageMajors.Container.Container.Majors[i]
		majors[i] = Major{
			Name:  raw.Name,
			ID:    raw.ID,
			Plans: make([]MajorPlan, len(raw.Plans)),
		}
		for j := range raw.Plans {
			plan := raw.Plans[j]
			majors[i].Plans[j] = MajorPlan{
				Name: plan.Name,
				ID:   plan.ID,
			}
		}
	}

	return majors, nil
}

func GetCourses(semester string, majorPlanID int) ([]Course, error) {
	pageCourse, err := getCourses(semester, majorPlanID)
	if err != nil {
		return nil, err
	}

	var courses []Course
	for _, s := range pageCourse.Container.Container.Semesters {
		for _, c := range s.Courses {
			course := Course{
				ID:       c.Course.ID,
				Name:     c.Course.Name,
				Semester: getSemester(s.Name),
			}
			courses = append(courses, course)
		}
	}
	return courses, nil
}

func GetGroups(semester string, majorPlanID int, courseID int) ([]Group, error) {
	pageGroups, err := getGroups(semester, majorPlanID, courseID)
	if err != nil {
		return nil, err
	}

	var groups []Group
	for _, g := range pageGroups.Container.Container.Groups {
		group := Group{
			CourseID:       courseID,
			Signups:        g.Details.Signups,
			Capacity:       g.Details.Capacity,
			HasDescription: g.Details.HasDescription,
			Code:           g.Details.Code,
			ID:             g.Details.ID,
			Name:           g.Details.Name.Name,
			Staff:          make([]GroupStaff, len(g.Staff)),
		}
		for i := range g.Staff {
			group.Staff[i] = GroupStaff{
				ID:           g.Staff[i].Details.ID,
				FirstName:    g.Staff[i].Details.FirstName,
				MiddleName:   g.Staff[i].Details.MiddleName,
				LastName:     g.Staff[i].Details.LastName,
				Availability: make([]StaffAvailability, len(g.Staff[i].Availability)),
			}

			role := ""
			for j := range g.Staff[i].Availability {
				group.Staff[i].Availability[j] = StaffAvailability{
					OnMonday:    g.Staff[i].Availability[j].OnMonday,
					OnTuesday:   g.Staff[i].Availability[j].OnTuesday,
					OnWednesday: g.Staff[i].Availability[j].OnWednesday,
					OnThursday:  g.Staff[i].Availability[j].OnThursday,
					OnFriday:    g.Staff[i].Availability[j].OnFriday,
					OnSaturday:  g.Staff[i].Availability[j].OnSaturday,
					OnSunday:    g.Staff[i].Availability[j].OnSunday,
					StartTime:   g.Staff[i].Availability[j].StartTime,
					EndTime:     g.Staff[i].Availability[j].EndTime,
				}
				if g.Staff[i].Availability[j].Location != nil {
					group.Staff[i].Availability[j].Location = &ClassLocation{
						ID:   g.Staff[i].Availability[j].Location.ID,
						Name: g.Staff[i].Availability[j].Location.Name,
					}
				}
				role = g.Staff[i].Availability[j].Role.Name
			}
			group.Staff[i].Role = getRole(role)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func getSemester(value string) Semester {
	trans := map[string]Semester{
		"Primer Semestre":  First,
		"Segundo Semestre": Second,
		"Tercer Semestre":  Third,
		"Cuarto Semestre":  Fourth,
		"Quinto Semestre":  Fifth,
		"Sexto Semestre":   Sixth,
		"Séptimo Semester": Seventh,
		"Octavo Semestre":  Eight,
		"Noveno Semester":  Ninth,
	}
	return trans[value]
}

func getRole(value string) StaffRole {
	trans := map[string]StaffRole{
		"Profesor":                ProfessorStaffRole,
		"Ayudante":                AssistantStaffRole,
		"Ayudante de Laboratorio": LabAssistantStaffRole,
	}
	return trans[value]
}
