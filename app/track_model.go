package app

type Track struct {
	Track    string
	Releases []Release `json:"releases"`
}

type Release struct {
	Name                string           `json:"name"`
	VersionCodes        []string         `json:"versionCodes"`
	ReleaseNotes        []ReleaseNotes   `json:"releaseNotes"`
	Status              string           `json:"status"`
	UserFraction        int              `json:"userFraction"`
	CountryTargeting    CountryTargeting `json:"countryTargeting"`
	InAppUpdatePriority int              `json:"inAppUpdatePriority"`
}
type ReleaseNotes struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}
type CountryTargeting struct {
	Countries          []string `json:"countries"`
	IncludeRestOfWorld bool     `json:"includeRestOfWorld"`
}
