package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// ContactFormData represents the contact form submission
type ContactFormData struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

// Spec represents a product specification
type Spec struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

// Feature represents a product feature
type Feature struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./contacts.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		subject TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	log.Println("✅ Database initialized successfully")
	return nil
}

// CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// API routes
	http.HandleFunc("/api/health", enableCORS(healthHandler))
	http.HandleFunc("/api/features", enableCORS(featuresHandler))
	http.HandleFunc("/api/specs", enableCORS(specsHandler))
	http.HandleFunc("/api/contact", enableCORS(contactHandler))
	http.HandleFunc("/api/contacts", enableCORS(contactsListHandler))

	port := ":8080"
	fmt.Println("🚀 NOVA API Server")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("📡 Server: http://localhost%s\n", port)
	fmt.Println("📊 Database: contacts.db")
	fmt.Println("\n📚 Available Endpoints:")
	fmt.Println("   GET  /api/health       - Health check")
	fmt.Println("   GET  /api/features     - Product features")
	fmt.Println("   GET  /api/specs        - Technical specs")
	fmt.Println("   POST /api/contact      - Submit contact form")
	fmt.Println("   GET  /api/contacts     - List all contacts")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "NOVA API is running",
		Data: map[string]string{
			"version": "1.0.0",
			"status":  "healthy",
		},
	})
}

func featuresHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	features := []Feature{
		{Icon: "🎵", Title: "Hi-Res Audio", Description: "Experience studio-quality sound with high-resolution audio support"},
		{Icon: "🔇", Title: "Active Noise Cancellation", Description: "Block out the world with advanced ANC technology"},
		{Icon: "⚡", Title: "40H Battery Life", Description: "All-day listening with up to 40 hours of playtime"},
		{Icon: "🎤", Title: "Crystal Clear Calls", Description: "AI-powered noise reduction for perfect call quality"},
		{Icon: "☁️", Title: "Cloud Comfort", Description: "Premium memory foam cushions for all-day comfort"},
		{Icon: "🌈", Title: "Spatial Audio", Description: "Immersive 3D audio with head tracking technology"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Data:    features,
	})
}

func specsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	specs := []Spec{
		{
			Category: "Audio",
			Items: []string{
				"Frequency Response: 20Hz - 20kHz",
				"Impedance: 32 Ohm",
				"Driver Size: 40mm",
				"THD: <0.1%",
			},
		},
		{
			Category: "Battery",
			Items: []string{
				"Playtime: 40 hours (ANC off)",
				"Playtime: 30 hours (ANC on)",
				"Charging: USB-C Fast Charge",
				"Charge Time: 2 hours (full)",
				"Quick Charge: 10 min = 5 hours",
			},
		},
		{
			Category: "Connectivity",
			Items: []string{
				"Bluetooth 5.3",
				"Range: 10 meters",
				"Multipoint Connection",
				"Codecs: AAC, SBC, aptX HD",
			},
		},
		{
			Category: "Physical",
			Items: []string{
				"Weight: 250g",
				"Foldable Design",
				"Colors: Black, Silver, Rose Gold",
				"Materials: Aluminum, Leather",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Data:    specs,
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var formData ContactFormData
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	// Validate input
	if formData.Name == "" || formData.Email == "" || formData.Subject == "" || formData.Message == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "All fields are required",
		})
		return
	}

	// Insert into database
	insertSQL := `INSERT INTO contacts (name, email, subject, message) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(insertSQL, formData.Name, formData.Email, formData.Subject, formData.Message)
	if err != nil {
		log.Printf("❌ Error inserting contact: %v", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Failed to save contact information",
		})
		return
	}

	id, _ := result.LastInsertId()
	log.Printf("✉️  New contact saved (ID: %d) from %s <%s>", id, formData.Name, formData.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Thank you for contacting us! We'll get back to you soon.",
		Data: map[string]int64{
			"id": id,
		},
	})
}

func contactsListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query(`SELECT id, name, email, subject, message, created_at FROM contacts ORDER BY created_at DESC`)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Failed to fetch contacts",
		})
		return
	}
	defer rows.Close()

	var contacts []ContactFormData
	for rows.Next() {
		var c ContactFormData
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Subject, &c.Message, &c.CreatedAt)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		contacts = append(contacts, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Data:    contacts,
	})
}
