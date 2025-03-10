package handlers

import (
    "austin-map/models"
    "encoding/json"
    "net/http"
)

// Sample data - in a real application, this would come from a database
var locations = []models.Location{
    {
        ID:          1,
        Name:        "Hotdoddy",
        Category:    "landmark",
        Latitude:    30.267073,
        Longitude:   -97.743074,
        Description: "Great burger joint in the heart of Austin",
    },
    {
        ID:          2,
        Name:        "Austin Tennis Academy",
        Category:    "recreation",
        Latitude:    30.304430,
        Longitude:   -97.948242,
        Description: "Top Tennis academy in Austin",
    },
    {
        ID:          3,
        Name:        "University of Texas at Austin",
        Category:    "education",
        Latitude:    30.2849,
        Longitude:   -97.7341,
        Description: "The University of Texas main campus",
    },
    {
        ID:          4,
        Name:        "Comedy Mothership",
        Category:    "recreation",
        Latitude:    30.2674,
        Longitude:   -97.7396,
        Description: "Great comedy club in Austin",
    },
    {
        ID:          5,
        Name:        "Lady Bird Lake",
        Category:    "recreation",
        Latitude:    30.2477,
        Longitude:   -97.7181,
        Description: "Recreation area with hiking and biking trails",
    },
    {
        ID:          6,
        Name:        "Honey Baked Ham",
        Category:    "shopping",
        Latitude:    30.373038,
        Longitude:   -97.724051,
        Description: "Best ham breakfast tacos I have ever had",
    },
    {
        ID:          7,
        Name:        "Terry Black's Barbeque",
        Category:    "recreation",
        Latitude:    30.260059,
        Longitude:   -97.755302,
        Description: "Best barbeque I have ever had, and argubly the best in Austin",
    },
}

func GetLocations(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(locations)
}