package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SalaryData struct {
	WorkYear          int    `json:"work_year"`
	ExperienceLevel   string `json:"experience_level"`
	EmploymentType    string `json:"employment_type"`
	JobTitle          string `json:"job_title"`
	Salary            int    `json:"salary"`
	SalaryCurrency    string `json:"salary_currency"`
	SalaryInUSD       int    `json:"salary_in_usd"`
	EmployeeResidence string `json:"employee_residence"`
	RemoteRatio       int    `json:"remote_ratio"`
	CompanyLocation   string `json:"company_location"`
	CompanySize       string `json:"company_size"`
}

var salaryData []SalaryData

func main() {
	loadCSVData()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/salaries", getSalaries)
	r.GET("/api/salaries/:year", getSalariesByYear)

	r.Run(":8080")
}

func loadCSVData() {
	file, err := os.Open("salaries.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records[1:] {
		workYear, _ := strconv.Atoi(record[0])
		salary, _ := strconv.Atoi(record[4])
		salaryInUSD, _ := strconv.Atoi(record[6])
		remoteRatio, _ := strconv.Atoi(record[8])

		salaryData = append(salaryData, SalaryData{
			WorkYear:          workYear,
			ExperienceLevel:   record[1],
			EmploymentType:    record[2],
			JobTitle:          record[3],
			Salary:            salary,
			SalaryCurrency:    record[5],
			SalaryInUSD:       salaryInUSD,
			EmployeeResidence: record[7],
			RemoteRatio:       remoteRatio,
			CompanyLocation:   record[9],
			CompanySize:       record[10],
		})
	}
	fmt.Println(records)
}

func getSalaries(c *gin.Context) {
	c.JSON(http.StatusOK, salaryData)
}

func getSalariesByYear(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	yearData := []SalaryData{}
	for _, data := range salaryData {
		if data.WorkYear == year {
			yearData = append(yearData, data)
		}
	}
	c.JSON(http.StatusOK, yearData)
}
