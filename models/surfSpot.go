package models

type SurfSpotList struct {
	PhotoURL         string   `json:"photo_url"`
	Title            string   `json:"destination"`
	Location         string   `json:"destination_state"`
	PeakSeasonBegins string   `json:"peak_season_begins"`
	PeakSeasonEnds   string   `json:"peak_season_ends"`
	SurfBreak        []string `json:"surf_break"`
	DifficultyLevel  int      `json:"difficulty_level"`
	ForecastURL		string		`json:"Magic Seaweed Link"`
}
