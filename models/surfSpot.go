package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SurfSpot struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Destination string             `bson:"destination" json:"destination"`
	Address     string             `bson:"address" json:"address"`
	Country     string             `bson:"country" json:"country"`
	Difficulty  int                `bson:"difficulty" json:"difficulty"`
	SurfBreak   []string           `bson:"surf_break" json:"surf_break"`
	SeasonStart string             `bson:"season_start" json:"season_start"`
	SeasonEnd   string             `bson:"season_end" json:"season_end"`
	Photo       string             `bson:"photo" json:"photo"`
	Link        string             `bson:"link" json:"link"`
	Geocode     string             `bson:"geocode" json:"geocode"`
	Saved       bool               `bson:"saved" json:"saved"`
}
