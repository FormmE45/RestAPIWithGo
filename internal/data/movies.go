package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` //Using '-' tag directive to ignore the field when parsing to json
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`           //Using 'omitempty' tag directive to ignore the field if empty
	Runtime   Runtime   `json:"runtime,omitempty,string"` //Changing json type by using string directive
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
	//All field are need to be exported (start with a capital letter) in order to be visible to encoding/json package.
	//If not the field won't be included when encoding a struct to JSON
}
