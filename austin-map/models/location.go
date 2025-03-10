package models

type Location struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Category    string  `json:"category"`
    Latitude    float64 `json:"latitude"`
    Longitude   float64 `json:"longitude"`
    Description string  `json:"description"`
}