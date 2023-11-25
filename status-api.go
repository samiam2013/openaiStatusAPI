package openaistatusapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type componentsStatus struct {
	Page struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		TimeZone  string `json:"time_zone"`
		UpdatedAt string `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		ID                 string      `json:"id"`
		Name               string      `json:"name"`
		Status             string      `json:"status"`
		CreatedAt          string      `json:"created_at"`
		UpdatedAt          string      `json:"updated_at"`
		Position           int         `json:"position"`
		Description        string      `json:"description"`
		Showcase           bool        `json:"showcase"`
		StartDate          string      `json:"start_date"`
		GroupID            interface{} `json:"group_id"`
		PageID             string      `json:"page_id"`
		Group              bool        `json:"group"`
		OnlyShowIfDegraded bool        `json:"only_show_if_degraded"`
	} `json:"components"`
}

func getComponentsStatus() (componentsStatus, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET",
		"https://status.openai.com/api/v2/components.json", nil)
	if err != nil {
		return componentsStatus{},
			errors.Join(fmt.Errorf("error creating request"), err)
	}
	csRaw, err := client.Do(req)
	if err != nil {
		return componentsStatus{},
			errors.Join(fmt.Errorf("error getting components status"), err)
	}
	var cs componentsStatus
	if err := json.NewDecoder(csRaw.Body).Decode(&cs); err != nil {
		return componentsStatus{},
			errors.Join(fmt.Errorf("error decoding components status"), err)
	}
	return cs, nil
}

// make an iota-based enum for components
type ComponentName string

const (
	API               ComponentName = "API"
	ManagementConsole ComponentName = "Management Console"
)

func GetComponent(component ComponentName) (string, error) {
	componentsStatus, err := getComponentsStatus()
	if err != nil {
		return "", errors.Join(fmt.Errorf("error getting components status"), err)
	}
	for _, comp := range componentsStatus.Components {
		if comp.Name == string(component) {
			return comp.Status, nil
		}
	}
	return "", errors.New("component not found")
}
