package redmine

import "time"

type Issue struct {
	ID      int `json:"id"`
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"project"`
	Tracker struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tracker"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	Priority struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"priority"`
	Author struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"author"`
	AssignedTo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"assigned_to"`
	Subject         string  `json:"subject"`
	Description     string  `json:"description"`
	StartDate       string  `json:"start_date"`
	DueDate         string  `json:"due_date"`
	DoneRatio       int     `json:"done_ratio"`
	SpentHours      float64 `json:"spent_hours"`
	TotalSpentHours float64 `json:"total_spent_hours"`
	// CustomFields    []struct {
	// 	ID       int    `json:"id"`
	// 	Name     string `json:"name"`
	// 	Value    string `json:"value"`
	// } `json:"custom_fields"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	ClosedOn  time.Time `json:"closed_on"`
}

type Issues []Issue
