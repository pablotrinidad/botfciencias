package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pablotrinidad/botfciencias/crawler/lib"
)

const (
	semester = "20212"
)

func main() {
	majors, err := lib.GetMajors()
	if err != nil {
		fmt.Printf("Failed fetching majors; %v", err)
		os.Exit(1)
	}

	// Retrieve ALL courses for all plans in all majors (IN PARALLEL)
	courses, elapsed, err := fetchCoursesByPlan(majors)
	if err != nil {
		fmt.Printf("Failed fetching some plan courses; %v", err)
		os.Exit(1)
	}
	fmt.Printf("Completed courses fetch in %s\n", elapsed)

	// Retrieve ALL groups for all retrieved courses (IN PARALLEL)
	groups, elapsed, err := fetchGroups(courses)
	if err != nil {
		fmt.Printf("Failed fetching some course groups; %v", err)
		os.Exit(1)
	}
	fmt.Printf("Completed groups fetch (%d) in %s", len(groups), elapsed)
	fmt.Println("JEJEJEJEJE")
}

func fetchCoursesByPlan(majors []lib.Major) (map[int][]lib.Course, time.Duration, error) {
	type planCourse struct {
		course lib.Course
		plan   int
	}

	coursesByPlan := map[int][]lib.Course{}
	plansRequests := make([]func(), 0)

	coursesChan := make(chan planCourse)
	errorsChan := make(chan error)

	for m := range majors {
		major := majors[m]
		for p := range major.Plans {
			plan := major.Plans[p]
			fn := func() {
				courses, err := lib.GetCourses(semester, plan.ID)
				if err != nil {
					errorsChan <- err
					return
				}

				for c := range courses {
					course := courses[c]
					data := planCourse{
						course: course,
						plan:   plan.ID,
					}
					coursesChan <- data
				}
			}
			plansRequests = append(plansRequests, fn)
		}
	}

	var elapsed time.Duration
	go func() {
		elapsed = callConcurrent(plansRequests)
		close(coursesChan)
		close(errorsChan)
	}()

	for result := range coursesChan {
		planID := result.plan
		course := result.course
		coursesByPlan[planID] = append(coursesByPlan[planID], course)
	}

	for err := range errorsChan {
		return nil, elapsed, err
	}

	return coursesByPlan, elapsed, nil
}

func fetchGroups(coursesByPlan map[int][]lib.Course) ([]lib.Group, time.Duration, error) {
	groupRequests := make([]func(), 0)
	groupsChan := make(chan lib.Group)
	errorsChan := make(chan error)

	for p := range coursesByPlan {
		planID := p
		courses := coursesByPlan[p]
		coursesRequests := make([]func(), len(courses))

		for i := range courses {
			courseID := courses[i].ID
			fn := func() {
				groups, err := lib.GetGroups(semester, planID, courseID)
				if err != nil {
					errorsChan <- err
					return
				}

				for g := range groups {
					group := groups[g]
					groupsChan <- group
				}
			}
			coursesRequests[i] = fn
		}

		groupRequests = append(groupRequests, coursesRequests...)
	}

	fmt.Println("------------------------------------")
	fmt.Println(len(groupRequests))
	fmt.Println("------------------------------------")

	var elapsed time.Duration
	go func() {
		elapsed = callConcurrent(groupRequests[:50])
		close(groupsChan)
		close(errorsChan)
	}()

	var groups []lib.Group
	for group := range groupsChan {
		fmt.Printf("%d, %s, %d\n", group.ID, group.Name, group.CourseID)
		groups = append(groups, group)
	}

	return groups, elapsed, nil
}
