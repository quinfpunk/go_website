# Go website
A website for the fictive NOVA headphones
# Visualize website
## Backend
Install go (https://go.dev/doc/install)
```
go get github.com/mattn/go-sqlite3
go run main.go
# launch the backend server
```
## Frontend
Install node (https://nodejs.org/en/download) 
Install vite
```
npm install vite
```
Launch the frontend server
```
npm run dev
```

## Backend Architecture
### Page Endpoints:

/home - Landing page with animated product showcase
/features - Dynamic features page that loads content via API
/specs - Technical specifications organized by category
/contacts - Contact form with real-time submission

### API Endpoints:

GET /api/features - Returns product features as JSON
GET /api/specs - Returns technical specifications as JSON
POST /api/contact - Handles contact form submissions

# Todo list
- [x] Split frontend and backend
    - [x] Frontend in JS with a component library
    - [x] Check if component library need a framework to work properly
- [] Change the features page to hae animations on scrolling that reveal each features

