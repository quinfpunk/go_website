# Go website
A website for the fictive NOVA headphones
# Visualize website
```
# install go see (https://go.dev/doc/install)
go run main.go
# on you browser look at the site: http://localhost:8080
```
## Backend Architecture
### Page Endpoints:

/home - Hero section landing page with animated product showcase
/features - Dynamic features page that loads content via API
/specs - Technical specifications organized by category
/contacts - Contact form with real-time submission

### API Endpoints:

GET /api/features - Returns product features as JSON
GET /api/specs - Returns technical specifications as JSON
POST /api/contact - Handles contact form submissions

# Todo list
- [] Split frontend and backend
    - [] Frontend in JS with a component library
    - [] Check if component library need a framework to work properly

