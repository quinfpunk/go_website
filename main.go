package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ContactFormData represents the contact form submission
type ContactFormData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
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

func main() {
	// Serve static pages
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/features", featuresHandler)
	http.HandleFunc("/contacts", contactsHandler)
	http.HandleFunc("/specs", specsHandler)
	
	// API endpoints
	http.HandleFunc("/api/contact", contactAPIHandler)
	http.HandleFunc("/api/specs", specsAPIHandler)
	http.HandleFunc("/api/features", featuresAPIHandler)
	
	// Redirect root to home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
		http.NotFound(w, r)
	})

	port := ":8080"
	fmt.Printf("🚀 Server starting on http://localhost%s\n", port)
	fmt.Println("📱 Pages available:")
	fmt.Println("   - http://localhost:8080/home")
	fmt.Println("   - http://localhost:8080/features")
	fmt.Println("   - http://localhost:8080/contacts")
	fmt.Println("   - http://localhost:8080/specs")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, getPageHTML("home", getHomeContent()))
}

func featuresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, getPageHTML("features", getFeaturesContent()))
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, getPageHTML("contacts", getContactsContent()))
}

func specsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, getPageHTML("specs", getSpecsContent()))
}

// API Handlers
func contactAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var formData ContactFormData
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: save to database or save to a Notion or send email
	log.Printf("Contact form submitted: %+v", formData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Thank you for contacting us! We'll get back to you soon.",
	})
}

func specsAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Specs list
	specs := []Spec{
		{Category: "Audio", Items: []string{"Frequency Response: 20Hz - 20kHz", "Impedance: 32 Ohm", "Driver Size: 40mm"}},
		{Category: "Battery", Items: []string{"Playtime: 40 hours", "Charging: USB-C Fast Charge", "Charge Time: 2 hours"}},
		{Category: "Connectivity", Items: []string{"Bluetooth 5.3", "Range: 10 meters", "Multipoint Connection"}},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(specs)
}

func featuresAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Features list 
	features := []Feature{
		{Icon: "🎵", Title: "Hi-Res Audio", Description: "Experience studio-quality sound with high-resolution audio support"},
		{Icon: "🔇", Title: "Active Noise Cancellation", Description: "Block out the world with advanced ANC technology"},
		{Icon: "⚡", Title: "40H Battery Life", Description: "All-day listening with up to 40 hours of playtime"},
		{Icon: "🎤", Title: "Crystal Clear Calls", Description: "AI-powered noise reduction for perfect call quality"},
		{Icon: "☁️", Title: "Cloud Comfort", Description: "Premium memory foam cushions for all-day comfort"},
		{Icon: "🌈", Title: "Spatial Audio", Description: "Immersive 3D audio with head tracking technology"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(features)
}

func getPageHTML(currentPage, content string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>NOVA - %s</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.2/gsap.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/gsap/3.12.2/ScrollTrigger.min.js"></script>
    %s
</head>
<body>
    <nav>
        <div class="logo">NOVA</div>
        <ul class="nav-links">
            <li><a href="/home" class="%s">Home</a></li>
            <li><a href="/features" class="%s">Features</a></li>
            <li><a href="/specs" class="%s">Specs</a></li>
            <li><a href="/contacts" class="%s">Contact</a></li>
        </ul>
    </nav>
    
    %s
    
    %s
</body>
</html>`, 
		strings.Title(currentPage), 
		getStyles(),
		getActiveClass(currentPage, "home"),
		getActiveClass(currentPage, "features"),
		getActiveClass(currentPage, "specs"),
		getActiveClass(currentPage, "contacts"),
		content,
		getScripts(currentPage))
}

func getActiveClass(current, page string) string {
	if current == page {
		return "active"
	}
	return ""
}

func getStyles() string {
	return `<style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            overflow-x: hidden;
            background: #0a0a0a;
            color: #fff;
            min-height: 100vh;
        }

        .gradient-bg {
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
            opacity: 0.15;
            filter: blur(80px);
            animation: gradientShift 15s ease infinite;
            z-index: 0;
        }

        @keyframes gradientShift {
            0%, 100% { transform: translate(0, 0) scale(1); }
            33% { transform: translate(5%, 10%) scale(1.1); }
            66% { transform: translate(-5%, -10%) scale(0.9); }
        }

        nav {
            position: fixed;
            top: 0;
            width: 100%;
            padding: 1.5rem 4rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            z-index: 1000;
            backdrop-filter: blur(10px);
            background: rgba(10, 10, 10, 0.8);
            border-bottom: 1px solid rgba(255, 255, 255, 0.1);
        }

        .logo {
            font-size: 1.5rem;
            font-weight: 700;
            letter-spacing: -0.5px;
            background: linear-gradient(135deg, #667eea 0%, #f093fb 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .nav-links {
            display: flex;
            gap: 3rem;
            list-style: none;
        }

        .nav-links a {
            color: rgba(255, 255, 255, 0.6);
            text-decoration: none;
            font-size: 0.95rem;
            transition: all 0.3s;
            position: relative;
            padding: 0.5rem 0;
        }

        .nav-links a:hover {
            color: #fff;
        }

        .nav-links a.active {
            color: #fff;
        }

        .nav-links a.active::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            height: 2px;
            background: linear-gradient(135deg, #667eea 0%, #f093fb 100%);
        }

        .container {
            max-width: 1400px;
            margin: 0 auto;
            padding: 6rem 2rem 2rem;
            position: relative;
            z-index: 10;
            min-height: 100vh;
        }

        .page-title {
            font-size: 3.5rem;
            font-weight: 800;
            margin-bottom: 1rem;
            background: linear-gradient(135deg, #fff 0%, rgba(255, 255, 255, 0.6) 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            opacity: 0;
        }

        .page-subtitle {
            font-size: 1.25rem;
            color: rgba(255, 255, 255, 0.6);
            margin-bottom: 3rem;
            opacity: 0;
        }

        /* Home Page Styles */
        .hero {
            min-height: 100vh;
            display: flex;
            align-items: center;
            padding-top: 5rem;
        }

        .hero-content {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 4rem;
            align-items: center;
        }

        .hero-text h1 {
            font-size: 5rem;
            font-weight: 800;
            line-height: 1.1;
            margin-bottom: 1.5rem;
            background: linear-gradient(135deg, #fff 0%, rgba(255, 255, 255, 0.6) 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .hero-text p {
            font-size: 1.25rem;
            line-height: 1.8;
            color: rgba(255, 255, 255, 0.7);
            margin-bottom: 2rem;
        }

        .cta-buttons {
            display: flex;
            gap: 1rem;
        }

        .btn {
            padding: 1rem 2.5rem;
            border: none;
            border-radius: 50px;
            font-size: 1rem;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s ease;
            text-decoration: none;
            display: inline-block;
        }

        .btn-primary {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: #fff;
            box-shadow: 0 10px 40px rgba(102, 126, 234, 0.4);
        }

        .btn-primary:hover {
            transform: translateY(-2px);
            box-shadow: 0 15px 50px rgba(102, 126, 234, 0.6);
        }

        .btn-secondary {
            background: rgba(255, 255, 255, 0.1);
            color: #fff;
            backdrop-filter: blur(10px);
            border: 1px solid rgba(255, 255, 255, 0.2);
        }

        .btn-secondary:hover {
            background: rgba(255, 255, 255, 0.15);
            transform: translateY(-2px);
        }

        .product-card {
            background: rgba(255, 255, 255, 0.03);
            backdrop-filter: blur(20px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 2rem;
            padding: 3rem;
            box-shadow: 0 30px 80px rgba(0, 0, 0, 0.4);
        }

        .product-image-wrapper {
            position: relative;
            width: 100%;
            height: 400px;
            display: flex;
            align-items: center;
            justify-content: center;
            margin-bottom: 2rem;
        }

        .glow {
            position: absolute;
            width: 300px;
            height: 300px;
            background: radial-gradient(circle, rgba(102, 126, 234, 0.4) 0%, transparent 70%);
            border-radius: 50%;
            filter: blur(60px);
            animation: pulse 4s ease-in-out infinite;
        }

        @keyframes pulse {
            0%, 100% { transform: scale(1); opacity: 0.5; }
            50% { transform: scale(1.2); opacity: 0.8; }
        }

        .headphones {
            width: 280px;
            height: 280px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            border-radius: 50%;
            position: relative;
            z-index: 2;
            display: flex;
            align-items: center;
            justify-content: center;
            box-shadow: 0 20px 60px rgba(102, 126, 234, 0.4);
        }

        .headphones::before {
            content: '';
            position: absolute;
            width: 240px;
            height: 240px;
            background: #0a0a0a;
            border-radius: 50%;
        }

        .headphones::after {
            content: '🎧';
            font-size: 6rem;
            position: absolute;
            z-index: 3;
        }

        .floating {
            animation: floating 3s ease-in-out infinite;
        }

        @keyframes floating {
            0%, 100% { transform: translateY(0); }
            50% { transform: translateY(-20px); }
        }

        /* Features Page Styles */
        .features-grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
            gap: 2rem;
            margin-top: 3rem;
        }

        .feature-card {
            background: rgba(255, 255, 255, 0.03);
            backdrop-filter: blur(20px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 1.5rem;
            padding: 2.5rem;
            transition: all 0.3s ease;
            opacity: 0;
        }

        .feature-card:hover {
            transform: translateY(-10px);
            background: rgba(255, 255, 255, 0.05);
            border-color: rgba(102, 126, 234, 0.3);
        }

        .feature-icon {
            font-size: 3rem;
            margin-bottom: 1.5rem;
            display: block;
        }

        .feature-card h3 {
            font-size: 1.5rem;
            margin-bottom: 1rem;
            color: #fff;
        }

        .feature-card p {
            font-size: 1rem;
            line-height: 1.6;
            color: rgba(255, 255, 255, 0.6);
        }

        /* Specs Page Styles */
        .specs-container {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
            gap: 2rem;
            margin-top: 3rem;
        }

        .spec-category {
            background: rgba(255, 255, 255, 0.03);
            backdrop-filter: blur(20px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 1.5rem;
            padding: 2.5rem;
            opacity: 0;
        }

        .spec-category h3 {
            font-size: 1.75rem;
            margin-bottom: 1.5rem;
            background: linear-gradient(135deg, #667eea 0%, #f093fb 100%);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .spec-item {
            padding: 1rem 0;
            border-bottom: 1px solid rgba(255, 255, 255, 0.1);
            color: rgba(255, 255, 255, 0.8);
            font-size: 1rem;
        }

        .spec-item:last-child {
            border-bottom: none;
        }

        /* Contact Page Styles */
        .contact-container {
            max-width: 800px;
            margin: 3rem auto;
        }

        .contact-form {
            background: rgba(255, 255, 255, 0.03);
            backdrop-filter: blur(20px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 1.5rem;
            padding: 3rem;
            opacity: 0;
        }

        .form-group {
            margin-bottom: 1.5rem;
        }

        .form-group label {
            display: block;
            margin-bottom: 0.5rem;
            color: rgba(255, 255, 255, 0.8);
            font-weight: 500;
        }

        .form-group input,
        .form-group textarea {
            width: 100%;
            padding: 1rem;
            background: rgba(255, 255, 255, 0.05);
            border: 1px solid rgba(255, 255, 255, 0.2);
            border-radius: 0.5rem;
            color: #fff;
            font-size: 1rem;
            font-family: inherit;
            transition: all 0.3s;
        }

        .form-group input:focus,
        .form-group textarea:focus {
            outline: none;
            border-color: #667eea;
            background: rgba(255, 255, 255, 0.08);
        }

        .form-group textarea {
            min-height: 150px;
            resize: vertical;
        }

        .form-message {
            margin-top: 1rem;
            padding: 1rem;
            border-radius: 0.5rem;
            display: none;
        }

        .form-message.success {
            background: rgba(34, 197, 94, 0.2);
            border: 1px solid rgba(34, 197, 94, 0.4);
            color: rgb(134, 239, 172);
        }

        .form-message.error {
            background: rgba(239, 68, 68, 0.2);
            border: 1px solid rgba(239, 68, 68, 0.4);
            color: rgb(252, 165, 165);
        }

        @media (max-width: 968px) {
            nav {
                padding: 1.5rem 2rem;
            }

            .nav-links {
                gap: 1.5rem;
            }

            .hero-content {
                grid-template-columns: 1fr;
            }

            .hero-text h1 {
                font-size: 3rem;
            }

            .page-title {
                font-size: 2.5rem;
            }

            .features-grid,
            .specs-container {
                grid-template-columns: 1fr;
            }
        }
    </style>`
}

func getHomeContent() string {
	return `<div class="gradient-bg"></div>
    <section class="hero">
        <div class="container">
            <div class="hero-content">
                <div class="hero-text">
                    <h1>Sound Beyond Limits</h1>
                    <p>Experience premium audio quality with our flagship wireless headphones. Engineered for perfection, designed for you.</p>
                    <div class="cta-buttons">
                        <a href="/contacts" class="btn btn-primary">Order Now - $299</a>
                        <a href="/features" class="btn btn-secondary">Learn More</a>
                    </div>
                </div>

                <div class="hero-product">
                    <div class="product-card">
                        <div class="product-image-wrapper">
                            <div class="glow"></div>
                            <div class="headphones floating"></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>`
}

func getFeaturesContent() string {
	return `<div class="gradient-bg"></div>
    <div class="container">
        <h1 class="page-title">Premium Features</h1>
        <p class="page-subtitle">Discover what makes NOVA headphones extraordinary</p>
        
        <div class="features-grid" id="featuresGrid">
            <!-- Features will be loaded dynamically -->
        </div>
    </div>`
}

func getSpecsContent() string {
	return `<div class="gradient-bg"></div>
    <div class="container">
        <h1 class="page-title">Technical Specifications</h1>
        <p class="page-subtitle">Every detail engineered to perfection</p>
        
        <div class="specs-container" id="specsContainer">
            <!-- Specs will be loaded dynamically -->
        </div>
    </div>`
}

func getContactsContent() string {
	return `<div class="gradient-bg"></div>
    <div class="container">
        <h1 class="page-title">Get In Touch</h1>
        <p class="page-subtitle">Have questions? We'd love to hear from you</p>
        
        <div class="contact-container">
            <form class="contact-form" id="contactForm">
                <div class="form-group">
                    <label for="name">Full Name</label>
                    <input type="text" id="name" name="name" required>
                </div>
                
                <div class="form-group">
                    <label for="email">Email Address</label>
                    <input type="email" id="email" name="email" required>
                </div>
                
                <div class="form-group">
                    <label for="subject">Subject</label>
                    <input type="text" id="subject" name="subject" required>
                </div>
                
                <div class="form-group">
                    <label for="message">Message</label>
                    <textarea id="message" name="message" required></textarea>
                </div>
                
                <button type="submit" class="btn btn-primary">Send Message</button>
                
                <div class="form-message" id="formMessage"></div>
            </form>
        </div>
    </div>`
}

func getScripts(page string) string {
	baseScript := `<script>
        gsap.registerPlugin(ScrollTrigger);

        // Page load animations
        gsap.to('.page-title', {
            opacity: 1,
            y: 0,
            duration: 1,
            ease: 'power3.out',
            delay: 0.2
        });

        gsap.to('.page-subtitle', {
            opacity: 1,
            y: 0,
            duration: 1,
            ease: 'power3.out',
            delay: 0.4
        });
    </script>`

	switch page {
	case "home":
		return `<script>
            gsap.registerPlugin(ScrollTrigger);

            gsap.to('.hero-text', {
                opacity: 1,
                y: 0,
                duration: 1.2,
                ease: 'power3.out',
                delay: 0.3
            });

            gsap.to('.hero-product', {
                opacity: 1,
                x: 0,
                duration: 1.2,
                ease: 'power3.out',
                delay: 0.6
            });

            const card = document.querySelector('.product-card');
            card.addEventListener('mousemove', (e) => {
                const rect = card.getBoundingClientRect();
                const x = e.clientX - rect.left;
                const y = e.clientY - rect.top;
                
                const centerX = rect.width / 2;
                const centerY = rect.height / 2;
                
                const rotateX = (y - centerY) / 20;
                const rotateY = (centerX - x) / 20;
                
                gsap.to(card, {
                    rotationX: rotateX,
                    rotationY: rotateY,
                    duration: 0.5,
                    ease: 'power2.out',
                    transformPerspective: 1000
                });
            });

            card.addEventListener('mouseleave', () => {
                gsap.to(card, {
                    rotationX: 0,
                    rotationY: 0,
                    duration: 0.5,
                    ease: 'power2.out'
                });
            });
        </script>`

	case "features":
		return baseScript + `<script>
            fetch('/api/features')
                .then(res => res.json())
                .then(features => {
                    const grid = document.getElementById('featuresGrid');
                    features.forEach(feature => {
                        const card = document.createElement('div');
                        card.className = 'feature-card';
                        card.innerHTML = ` + "`" + `
                            <span class="feature-icon">${feature.icon}</span>
                            <h3>${feature.title}</h3>
                            <p>${feature.description}</p>
                        ` + "`" + `;
                        grid.appendChild(card);
                    });

                    gsap.to('.feature-card', {
                        opacity: 1,
                        y: 0,
                        duration: 0.8,
                        stagger: 0.15,
                        ease: 'power3.out'
                    });
                });
        </script>`

	case "specs":
		return baseScript + `<script>
            fetch('/api/specs')
                .then(res => res.json())
                .then(specs => {
                    const container = document.getElementById('specsContainer');
                    specs.forEach(spec => {
                        const category = document.createElement('div');
                        category.className = 'spec-category';
                        category.innerHTML = ` + "`" + `
                            <h3>${spec.category}</h3>
                            ${spec.items.map(item => ` + "`" + `<div class="spec-item">${item}</div>` + "`" + `).join('')}
                        ` + "`" + `;
                        container.appendChild(category);
                    });

                    gsap.to('.spec-category', {
                        opacity: 1,
                        y: 0,
                        duration: 0.8,
                        stagger: 0.15,
                        ease: 'power3.out'
                    });
                });
        </script>`

	case "contacts":
		return baseScript + `<script>
            gsap.to('.contact-form', {
                opacity: 1,
                y: 0,
                duration: 1,
                ease: 'power3.out',
                delay: 0.4
            });

            document.getElementById('contactForm').addEventListener('submit', async (e) => {
                e.preventDefault();
                
                const formData = {
                    name: document.getElementById('name').value,
                    email: document.getElementById('email').value,
                    subject: document.getElementById('subject').value,
                    message: document.getElementById('message').value
                };

                const messageDiv = document.getElementById('formMessage');
                
                try {
                    const response = await fetch('/api/contact', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(formData)
                    });

                    const result = await response.json();

                    if (result.success) {
                        messageDiv.className = 'form-message success';
                        messageDiv.textContent = result.message;
                        messageDiv.style.display = 'block';
                        document.getElementById('contactForm').reset();
                    } else {
                        throw new Error(result.message);
                    }
                } catch (error) {
                    messageDiv.className = 'form-message error';
                    messageDiv.textContent = 'Failed to send message. Please try again.';
                    messageDiv.style.display = 'block';
                }
            });
        </script>`

	default:
		return baseScript
	}
}
