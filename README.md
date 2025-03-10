# Austin Interactive Map

An interactive web application that displays various points  around Austin, Texas. Users can explore all the spots that I visited on my recent trip there!

## Features

- Interactive map of Austin, TX using OpenStreetMap
- Filterable locations by category
- Clickable markers with location details
- Responsive sidebar with location listings
- Category-based filtering system
- Smooth map transitions and popups

## Technologies Used

### Backend
- **Go (Golang)** - Server-side programming language
- **Gorilla Mux** - HTTP router and URL matcher for building Go web servers

### Frontend
- **HTML5** - Structure and content
- **CSS3** - Styling and layout
- **JavaScript** - Client-side interactivity
- **Leaflet.js** - Open-source JavaScript library for interactive maps
- **OpenStreetMap** - Free, open-source map data

## Project Structure 

austin-map-app/
├── main.go # Main application entry point
├── handlers/
│ └── locations.go # Location data and API handlers
├── models/
│ └── location.go # Location data structure
└── static/
├── index.html # Frontend interface
└── style.css # CSS styling


## Getting Started

1. Clone the repository
2. Make sure you have Go installed on your system
3. Install dependencies:
   ```bash
   go mod init austin-map
   go get github.com/gorilla/mux
   ```
4. Run the application:
   ```bash
   go run main.go
   ```
5. Open your browser and visit `http://localhost:8080`

## API Endpoints

- `GET /api/locations` - Returns a list of all locations in JSON format

## Current Locations Featured

- Hotdoddy burgers
- Austin Tennis Academy (Recreation)
- University of Texas at Austin (Education)
- Comedy Mothership (Recreation)
- Lady Bird Lake (Recreation)
- Honey Baked Ham (Shopping)
- Terry Black's Barbeque (Recreation)

## Future Enhancements

- Database integration for dynamic location management
- User authentication system
- Ability to add new locations
- Search functionality
- Routing capabilities
- Additional location details and photos
- Real-time updates for events or traffic

## License

This project is open source and available under the MIT License.