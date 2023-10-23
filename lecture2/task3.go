package lecture2

import (
	"fmt"
)

type Employee interface {
	GetSalaryPerYear() (string, int)
	GetJobStatus() string
	GetHiringDate() string
	GetDetails()
}

func (ceo CEO) GetSalaryPerYear() (string, int) {
	name := ceo.Name + " " + ceo.Surname
	salary := ceo.Salary

	return name, salary
}

func (ceo CEO) GetJobStatus() string {
	return ceo.JobStatus
}

func (ceo CEO) GetHiringDate() string {
	return ceo.HiringDate
}

func (ceo CEO) GetDetails() {
	fmt.Printf("Employee name: %v %v, age: %v, with job status: %v", ceo.Name, ceo.Surname, ceo.Age, ceo.JobStatus)
}

func (mid MiddleCloudEngineer) GetSalaryPerYear() (string, int) {
	name := mid.Name + " " + mid.Surname
	salary := mid.Salary

	return name, salary
}

func (mid MiddleCloudEngineer) GetJobStatus() string {
	return mid.JobStatus
}

func (mid MiddleCloudEngineer) GetHiringDate() string {
	return mid.HiringDate
}

func (mid MiddleCloudEngineer) GetDetails() {
	fmt.Printf("Employee name: %v %v, age: %v, with job status: %v", mid.Name, mid.Surname, mid.Age, mid.JobStatus)
}
func (sen SeniorDevOpsEngineer) GetSalaryPerYear() (string, int) {
	name := sen.Name + " " + sen.Surname
	salary := sen.Salary

	return name, salary
}

func (sen SeniorDevOpsEngineer) GetJobStatus() string {
	return sen.JobStatus
}

func (sen SeniorDevOpsEngineer) GetHiringDate() string {
	return sen.HiringDate
}

func (sen SeniorDevOpsEngineer) GetDetails() {
	fmt.Printf("Employee name: %v %v, age: %v, with job status: %v", sen.Name, sen.Surname, sen.Age, sen.JobStatus)
}

func (intern Intern) GetSalaryPerYear() (string, int) {
	name := intern.Name + " " + intern.Surname
	salary := intern.Salary

	return name, salary
}

func (intern Intern) GetJobStatus() string {
	return intern.JobStatus
}

func (intern Intern) GetHiringDate() string {
	return intern.HiringDate
}

func (intern Intern) GetDetails() {
	fmt.Printf("Employee name: %v %v, age: %v, with job status: %v", intern.Name, intern.Surname, intern.Age, intern.JobStatus)
}
