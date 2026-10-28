package main

import (
	"fmt"
	"strings"
)

func countMarketers(jobTitles []string) int {
	count := 0

	for _, jobTitle := range jobTitles {
		if strings.ToLower(jobTitle) == "marketer" {
			count++
		}
	}

	return count
}

func lastWorkExperience(workExperiences []string) string {
	if len(workExperiences) == 0 {
		return ""
	}

	return workExperiences[len(workExperiences)-1]
}

func main() {
	jobTitles := []string{"Marketer", "Engineer", "marketer", "Designer"}
	experiences := []string{"Intern", "Frontend Developer", "Backend Engineer"}

	fmt.Println(countMarketers(jobTitles))
	fmt.Println(lastWorkExperience(experiences))
}
