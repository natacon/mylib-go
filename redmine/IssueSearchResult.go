package redmine

import "time"

type IssueSearchResult struct {
	Issues []struct {
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
		FixedVersion struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"fixed_version"`
		Subject      string `json:"subject"`
		Description  string `json:"description"`
		DueDate      string `json:"due_date"`
		DoneRatio    int    `json:"done_ratio"`
		CustomFields []struct {
			ID       int      `json:"id"`
			Name     string   `json:"name"`
			Multiple bool     `json:"multiple,omitempty"`
			Value    []string `json:"value"`
		} `json:"custom_fields"`
		CreatedOn time.Time `json:"created_on"`
		UpdatedOn time.Time `json:"updated_on"`
	} `json:"issues"`
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}
