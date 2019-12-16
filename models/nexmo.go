package models

type RequestNexmoSMS struct {
	API_KEY    string `json:"api_key"`
	API_SECRET string `json:"api_secret"`
	To         string `json:"to"`
	From       string `json:"from"`
	Text       string `json:"text"`
}
