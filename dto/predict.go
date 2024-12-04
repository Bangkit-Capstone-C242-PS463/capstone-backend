package dto

type PredictRequest struct {
	UserStory string `json:"user_story"`
}

type PredictManualRequest struct {
	Symptoms []string `json:"symptoms"`
}

type PredictResponse struct {
	PredictedDisease   string   `json:"predicted_disease"`
	IdentifiedSymptoms []string `json:"identified_symptoms,omitempty"`
}
