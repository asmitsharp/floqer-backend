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
	"github.com/joho/godotenv"
	"github.com/magicx-ai/groq-go/groq"
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

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

var salaryData []SalaryData

func main() {
	loadCSVData()
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/salaries", getSalaries)
	r.GET("/api/salaries/:year", getSalariesByYear)
	r.POST("/api/salaries/chat", handleChat)

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

func handleChat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	summary := `{'AI Architect': 252551.24137931035, 'AI Developer': 135466.7894736842, 'AI Engineer': 162989.9406779661, 'AI Product Manager': 141766.66666666666, 'AI Programmer': 62042.0, 'AI Research Engineer': 82963.0, 'AI Research Scientist': 101296.0, 'AI Scientist': 121949.0, 'AI Software Engineer': 151169.4, 'AWS Data Architect': 258000.0, 'Admin & Data Analyst': 61805.5, 'Analytics Engineer': 160215.25522041763, 'Analytics Engineering Manager': 399880.0, 'Applied Data Scientist': 97004.23076923077, 'Applied Machine Learning Engineer': 141726.33333333334, 'Applied Machine Learning Scientist': 104289.57142857143, 'Applied Research Scientist': 66666.0, 'Applied Scientist': 190034.39030023094, 'Autonomous Vehicle Technician': 82777.5, 'Azure Data Engineer': 100000.0, 'BI Analyst': 118243.75, 'BI Data Analyst': 67900.8947368421, 'BI Data Engineer': 60000.0, 'BI Developer': 105512.3, 'Bear Robotics': 167500.0, 'Big Data Architect': 126751.0, 'Big Data Developer': 117000.0, 'Big Data Engineer': 66251.41666666667, 'Business Data Analyst': 73567.76190476191, 'Business Intelligence': 144810.2142857143, 'Business Intelligence Analyst': 114066.41884816754, 'Business Intelligence Data Analyst': 83209.5, 'Business Intelligence Developer': 100667.87719298246, 'Business Intelligence Engineer': 139744.65322580645, 'Business Intelligence Lead': 137379.5, 'Business Intelligence Manager': 137102.5, 'Business Intelligence Specialist': 132261.5, 'CRM Data Analyst': 40000.0, 'Cloud Data Architect': 250000.0, 'Cloud Data Engineer': 131617.75, 'Cloud Database Engineer': 143538.46153846153, 'Compliance Data Analyst': 45000.0, 'Computational Biologist': 190384.25, 'Computer Vision Engineer': 172677.0625, 'Computer Vision Software Engineer': 77760.6, 'Consultant Data Engineer': 84269.5, 'Data Analyst': 107920.17462932455, 'Data Analyst Lead': 71500.0, 'Data Analytics Associate': 91000.0, 'Data Analytics Consultant': 86690.625, 'Data Analytics Engineer': 78839.4, 'Data Analytics Lead': 195536.73913043478, 'Data Analytics Manager': 134608.4193548387, 'Data Analytics Specialist': 98811.125, 'Data Architect': 161082.08083140876, 'Data DevOps Engineer': 65638.0, 'Data Developer': 97747.63333333333, 'Data Engineer': 146567.13483796295, 'Data Infrastructure Engineer': 204192.04545454544, 'Data Integration Developer': 140580.75, 'Data Integration Engineer': 121272.4, 'Data Integration Specialist': 93675.34782608696, 'Data Lead': 159940.3076923077, 'Data Management Analyst': 92538.09090909091, 'Data Management Consultant': 92500.0, 'Data Management Specialist': 104828.72727272728, 'Data Manager': 108645.45238095238, 'Data Modeler': 130851.41071428571, 'Data Modeller': 83052.0, 'Data Operations Analyst': 96432.91666666667, 'Data Operations Associate': 77599.33333333333, 'Data Operations Engineer': 133431.25, 'Data Operations Manager': 123000.0, 'Data Operations Specialist': 91403.33333333333, 'Data Pipeline Engineer': 172500.0, 'Data Product Manager': 156899.44444444444, 'Data Product Owner': 127292.75, 'Data Quality Analyst': 83502.41666666667, 'Data Quality Engineer': 100350.6, 'Data Quality Manager': 82852.83333333333, 'Data Reporting Analyst': 70250.0, 'Data Science': 162381.53136531365, 'Data Science Analyst': 101000.0, 'Data Science Consultant': 106658.2048192771, 'Data Science Director': 177416.5, 'Data Science Engineer': 142334.06896551725, 'Data Science Lead': 167585.96153846153, 'Data Science Manager': 193961.31967213115, 'Data Science Practitioner': 133650.0, 'Data Science Tech Lead': 375000.0, 'Data Scientist': 154143.68810386473, 'Data Scientist Lead': 136153.0, 'Data Specialist': 91182.94186046511, 'Data Strategist': 102400.41666666667, 'Data Strategy Manager': 138750.0, 'Data Visualization Analyst': 120000.0, 'Data Visualization Engineer': 122975.0, 'Data Visualization Specialist': 112515.21428571429, 'Decision Scientist': 166094.63157894736, 'Deep Learning Engineer': 190807.5, 'Deep Learning Researcher': 124163.0, 'Director of Business Intelligence': 187500.0, 'Director of Data Science': 218775.33333333334, 'ETL Developer': 122251.28571428571, 'ETL Engineer': 98190.66666666667, 'Encounter Data Management Professional': 72812.5, 'Finance Data Analyst': 141933.66666666666, 'Financial Data Analyst': 90375.0, 'Head of Data': 211860.19672131148, 'Head of Data Science': 178387.25, 'Head of Machine Learning': 299758.4285714286, 'Insight Analyst': 50090.642857142855, 'Lead Data Analyst': 67077.5, 'Lead Data Engineer': 139230.33333333334, 'Lead Data Scientist': 105139.54545454546, 'Lead Machine Learning Engineer': 88312.0, 'ML Engineer': 193520.9696969697, 'ML Ops Engineer': 156400.0, 'MLOps Engineer': 146846.4347826087, 'Machine Learning Developer': 79662.8, 'Machine Learning Engineer': 188639.36758661186, 'Machine Learning Infrastructure Engineer': 167395.27272727274, 'Machine Learning Manager': 149166.5, 'Machine Learning Modeler': 173350.0, 'Machine Learning Operations Engineer': 136527.77777777778, 'Machine Learning Research Engineer': 66157.0, 'Machine Learning Researcher': 149644.57142857142, 'Machine Learning Scientist': 190077.84782608695, 'Machine Learning Software Engineer': 188440.26666666666, 'Machine Learning Specialist': 55000.0, 'Manager Data Management': 80507.0, 'Managing Director Data Science': 280000.0, 'Marketing Data Analyst': 144327.0, 'Marketing Data Engineer': 66970.0, 'Marketing Data Scientist': 90000.0, 'NLP Engineer': 130002.26666666666, 'Power BI Developer': 64781.0, 'Principal Data Analyst': 122500.0, 'Principal Data Architect': 38154.0, 'Principal Data Engineer': 158924.33333333334, 'Principal Data Scientist': 194089.6, 'Principal Machine Learning Engineer': 129430.0, 'Product Data Analyst': 69436.16666666667, 'Prompt Engineer': 205093.58823529413, 'Quantitative Research Analyst': 51000.0, 'Research Analyst': 113149.09756097561, 'Research Engineer': 190177.977124183, 'Research Scientist': 192011.02268431, 'Robotics Engineer': 140416.66666666666, 'Robotics Software Engineer': 196625.0, 'Sales Data Analyst': 60000.0, 'Software Data Engineer': 111627.66666666667, 'Staff Data Analyst': 79917.0, 'Staff Data Scientist': 134500.0, 'Staff Machine Learning Engineer': 185000.0}
job_titles_by_employment_type: {'CT': 19, 'FL': 14, 'FT': 153, 'PT': 14}
total_employ: 16494
avg_salary_year_wise: { '2020': '$102,251', '2021': '$99,922', '2022': '$134,350', '2023': '$153,733', '2024': '$150,643'}
salary_distribution_statistics: {'count': 16494.0, 'mean': 149713.57572450588, 'std': 68516.1369182901, 'min': 15000.0, '25%': 101517.5, '50%': 141300.0, '75%': 185900.0, 'max': 800000.0}
percentage_of_remote_jobs: 32.79980599005699
number_of_people_by_company_size: {'M': 15268, 'L': 1038, 'S': 188}`

	// Generate prompt for LLM, including the salary data
	prompt := fmt.Sprintf("You are a Data Analyst, Based on the following ML Engineer salary data:\n\n%s\n\nAnswer this question: %s", summary, req.Message)

	// Call Groq API
	response, err := callGroqAPI(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from LLM"})
		return
	}

	c.JSON(http.StatusOK, ChatResponse{Response: response})
}

func callGroqAPI(prompt string) (string, error) {
	// Load API Key from environment variable
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("GROQ_API_KEY")
	fmt.Println(apiKey)
	if apiKey == "" {
		return "", fmt.Errorf("GROQ_API_KEY environment variable not set")
	}

	// Initialize Groq client
	cli := groq.NewClient(apiKey, &http.Client{})

	// Prepare the request
	req := groq.ChatCompletionRequest{
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Model:       groq.ModelIDLLAMA370B, // Ensure the model ID matches your desired model
		MaxTokens:   8000,
		Temperature: 0.7,
		TopP:        0.9,
		NumChoices:  1,
		Stream:      false,
	}

	// Send the request
	resp, err := cli.CreateChatCompletion(req)
	if err != nil {
		return "", fmt.Errorf("error occurred: %v", err)
	}

	// Return the response content
	return resp.Choices[0].Message.Content, nil
}
