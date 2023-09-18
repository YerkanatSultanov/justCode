package main

import (
	"fmt"
	"justCode/lecture2"
)

func main() {
	allEmployees := []interface{}{
		lecture2.CEO{Name: "Yerkanat", Surname: "Sultanov", Age: 20, Address: "Nazarbaeva 193", Salary: 800000, HiringDate: "17.04.2021", JobStatus: "Main Office"},
		lecture2.SeniorDevOpsEngineer{Name: "Syrkhan", Surname: "Madiyev", Age: 20, Address: "Aksai-2", Salary: 600000, HiringDate: "20.04.2021", JobStatus: "Home office"},
		lecture2.MiddleCloudEngineer{Name: "Adil", Surname: "Murziyev", Age: 20, Address: "Kaskelen", Salary: 300000, HiringDate: "23.04.2021", JobStatus: "Home office"},
		lecture2.Intern{Name: "Askar", Surname: "Bono", Age: 16, Address: "Satpayev", Salary: 50000, HiringDate: "30.04.2021", JobStatus: "Main office"},
	}

	for _, employee := range allEmployees {
		switch emp := employee.(type) {
		case lecture2.CEO:
			fmt.Printf("Name: %s %s, Age: %d, Salary: $%d, JobStatus: %s\n", emp.Name, emp.Surname, emp.Age, emp.Salary, emp.JobStatus)
		case lecture2.MiddleCloudEngineer:
			fmt.Printf("Name: %s %s, Age: %d, Salary: $%d, JobStatus: %s\n", emp.Name, emp.Surname, emp.Age, emp.Salary, emp.JobStatus)
		case lecture2.SeniorDevOpsEngineer:
			fmt.Printf("Name: %s %s, Age: %d, Salary: $%d, JobStatus: %s\n", emp.Name, emp.Surname, emp.Age, emp.Salary, emp.JobStatus)
		case lecture2.Intern:
			fmt.Printf("Name: %s %s, Age: %d, Salary: $%d, JobStatus: %s\n", emp.Name, emp.Surname, emp.Age, emp.Salary, emp.JobStatus)
		}
	}

}
