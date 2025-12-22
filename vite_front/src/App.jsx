import React, { useState, useEffect, useRef } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { Headphones, Battery, Volume2, Mic, CloudRain, Sparkles, ChevronRight, Star, Send, CheckCircle2, AlertCircle, Mail, Phone, MapPin, Menu, X, Zap, Radio, Waves } from 'lucide-react';

const API_URL = 'http://localhost:8080/api';

// Custom hook for scroll animations
const useScrollAnimation = () => {
  const [isVisible, setIsVisible] = useState(false);
  const ref = useRef(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
        }
      },
      {
        threshold: 0.1,
        rootMargin: '0px 0px -100px 0px'
      }
    );

    if (ref.current) {
      observer.observe(ref.current);
    }

    return () => {
      if (ref.current) {
        observer.unobserve(ref.current);
      }
    };
  }, []);

  return [ref, isVisible];
};

// Navigation Component
const Navigation = ({ currentPage, setCurrentPage }) => {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  
  const pages = [
    { id: 'home', label: 'Home' },
    { id: 'features', label: 'Features' },
    { id: 'specs', label: 'Specs' },
    { id: 'contact', label: 'Contact' }
  ];

  return (
    <nav className="fixed top-0 w-full z-50 border-b border-white/10 bg-slate-950/50 backdrop-blur-xl">
      <div className="max-w-7xl mx-auto px-6 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2 cursor-pointer" onClick={() => setCurrentPage('home')}>
            <Headphones className="w-6 h-6 text-purple-400" />
            <span className="text-2xl font-bold bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent">
              NOVA
            </span>
          </div>
          
          <div className="hidden md:flex items-center gap-8">
            {pages.map(page => (
              <button
                key={page.id}
                onClick={() => setCurrentPage(page.id)}
                className={`text-sm transition-colors ${
                  currentPage === page.id ? 'text-white' : 'text-white/60 hover:text-purple-400'
                }`}
              >
                {page.label}
              </button>
            ))}
          </div>

          <Button
	    onClick = {() => setCurrentPage('contact')}
	    className="hidden md:block bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700"
	  >
            Buy Now
          </Button>

          <button className="md:hidden" onClick={() => setMobileMenuOpen(!mobileMenuOpen)}>
            {mobileMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
          </button>
        </div>

        {mobileMenuOpen && (
          <div className="md:hidden mt-4 pb-4 space-y-3">
            {pages.map(page => (
              <button
                key={page.id}
                onClick={() => {
                  setCurrentPage(page.id);
                  setMobileMenuOpen(false);
                }}
                className={`block w-full text-left px-4 py-2 rounded-lg transition-colors ${
                  currentPage === page.id ? 'bg-purple-600 text-white' : 'text-white/60 hover:bg-white/10'
                }`}
              >
                {page.label}
              </button>
            ))}
          </div>
        )}
      </div>
    </nav>
  );
};

// Home Page
const HomePage = ({ setCurrentPage }) => {
  const [scrollY, setScrollY] = useState(0);

  useEffect(() => {
    const handleScroll = () => setScrollY(window.scrollY);
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const quickFeatures = [
    { icon: Volume2, label: "Hi-Res Audio", description: "Studio quality" },
    { icon: Battery, label: "40H Battery", description: "All-day power" },
    { icon: Mic, label: "Clear Calls", description: "AI noise reduction" }
  ];

  return (
    <div className="pt-32 pb-20 px-6">
      <div className="max-w-7xl mx-auto">
        <div className="grid lg:grid-cols-2 gap-12 items-center">
          <div className="space-y-8 relative z-10">
            <Badge className="bg-purple-500/20 text-purple-300 border-purple-500/50 hover:bg-purple-500/30">
              <Sparkles className="w-3 h-3 mr-1" />
              New Release 2024
            </Badge>
            
            <div className="space-y-4">
              <h1 className="text-6xl md:text-7xl font-bold leading-tight">
                Sound Beyond
                <span className="block bg-gradient-to-r from-purple-400 via-pink-400 to-purple-400 bg-clip-text text-transparent">
                  Limits
                </span>
              </h1>
              <p className="text-xl text-white/70 leading-relaxed max-w-xl">
                Experience premium audio quality with our flagship wireless headphones. 
                Engineered for perfection, designed for you.
              </p>
            </div>

            <div className="flex flex-wrap gap-4">
              <Button 
                size="lg" 
                onClick={() => setCurrentPage('contact')}
                className="bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 text-white shadow-lg shadow-purple-500/50"
              >
                Order Now - $299
                <ChevronRight className="w-4 h-4 ml-2" />
              </Button>
              <Button 
                size="lg" 
                variant="outline" 
                onClick={() => setCurrentPage('features')}
                className="border-white/20 text-black hover:bg-white/10"
              >
                Learn More
              </Button>
            </div>

            <div className="flex flex-wrap gap-4 pt-8">
              {quickFeatures.map((feature, index) => (
                <div key={index} className="flex items-center gap-3 bg-white/5 backdrop-blur-sm rounded-full px-4 py-2 border border-white/10">
                  <feature.icon className="w-4 h-4 text-purple-400" />
                  <div className="text-left">
                    <div className="text-sm font-semibold">{feature.label}</div>
                    <div className="text-xs text-white/50">{feature.description}</div>
                  </div>
                </div>
              ))}
            </div>
          </div>

          <div className="relative lg:h-[600px] flex items-center justify-center">
            <div className="absolute inset-0 bg-gradient-radial from-purple-500/30 via-transparent to-transparent blur-3xl animate-pulse"></div>
            
            <Card className="relative w-full max-w-md bg-gradient-to-br from-white/5 to-white/0 border-white/10 backdrop-blur-xl hover:scale-105 transition-all duration-500 shadow-2xl shadow-purple-500/20"
              style={{
                transform: `rotateY(${scrollY * 0.02}deg) rotateX(${scrollY * -0.01}deg)`
              }}>
              <CardContent className="p-8">
                <div className="relative aspect-square mb-6">
                  <div className="absolute inset-0 flex items-center justify-center">
                    <div className="w-64 h-64 rounded-full bg-gradient-to-br from-purple-600 to-pink-600 flex items-center justify-center relative overflow-hidden">
                      <div className="absolute inset-4 rounded-full bg-slate-950"></div>
                      <Headphones className="w-32 h-32 text-white relative z-10" />
                      <div className="absolute top-1/4 left-1/4 w-2 h-2 rounded-full bg-purple-400 animate-ping"></div>
                      <div className="absolute bottom-1/3 right-1/4 w-2 h-2 rounded-full bg-pink-400 animate-ping delay-500"></div>
                    </div>
                  </div>
                </div>

                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <div>
                      <h3 className="text-2xl font-bold">NOVA Pro</h3>
                      <p className="text-white/60 text-sm">Premium Wireless</p>
                    </div>
                    <div className="flex items-center gap-1">
                      {[...Array(5)].map((_, i) => (
                        <Star key={i} className="w-4 h-4 fill-yellow-400 text-yellow-400" />
                      ))}
                    </div>
                  </div>

                  <div className="grid grid-cols-3 gap-3">
                    <div className="bg-white/5 rounded-lg p-3 text-center border border-white/10">
                      <Battery className="w-5 h-5 mx-auto mb-1 text-green-400" />
                      <div className="text-xs text-white/60">40hrs</div>
                    </div>
                    <div className="bg-white/5 rounded-lg p-3 text-center border border-white/10">
                      <Volume2 className="w-5 h-5 mx-auto mb-1 text-purple-400" />
                      <div className="text-xs text-white/60">Hi-Res</div>
                    </div>
                    <div className="bg-white/5 rounded-lg p-3 text-center border border-white/10">
                      <CloudRain className="w-5 h-5 mx-auto mb-1 text-blue-400" />
                      <div className="text-xs text-white/60">ANC</div>
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

// Animated Feature Section Component
const AnimatedFeature = ({ feature, index, animationType, setCurrentPage }) => {
  const [ref, isVisible] = useScrollAnimation();
  
  const getAnimationClass = () => {
    if (!isVisible) {
      switch (animationType) {
        case 'slideLeft':
          return 'opacity-0 -translate-x-32';
        case 'slideRight':
          return 'opacity-0 translate-x-32';
        case 'slideUp':
          return 'opacity-0 translate-y-32';
        case 'scale':
          return 'opacity-0 scale-75';
        case 'rotate':
          return 'opacity-0 rotate-12 scale-90';
        default:
          return 'opacity-0';
      }
    }
    return 'opacity-100 translate-x-0 translate-y-0 scale-100 rotate-0';
  };

  const iconMap = {
    '🎵': Volume2,
    '🔇': CloudRain,
    '⚡': Zap,
    '🎤': Mic,
    '☁️': Waves,
    '🌈': Sparkles,
    '📡': Radio
  };

  const IconComponent = iconMap[feature.icon] || Sparkles;
  const isEven = index % 2 === 0;

  return (
    <div
      ref={ref}
      className={`transition-all duration-1000 ease-out ${getAnimationClass()}`}
      style={{ transitionDelay: `${index * 100}ms` }}
    >
      <div className={`grid lg:grid-cols-2 gap-12 items-center py-20 ${!isEven ? 'lg:flex-row-reverse' : ''}`}>
        {/* Content Side */}
        <div className={`space-y-6 ${!isEven ? 'lg:order-2' : ''}`}>
          <div className="inline-flex items-center gap-3 bg-purple-500/10 backdrop-blur-sm rounded-full px-6 py-2 border border-purple-500/30">
            <IconComponent className="w-5 h-5 text-purple-400" />
            <span className="text-sm font-semibold text-purple-300">Feature {index + 1}</span>
          </div>

          <h2 className="text-5xl font-bold">
            <span className="bg-gradient-to-r from-white to-white/60 bg-clip-text text-transparent">
              {feature.title}
            </span>
          </h2>

          <p className="text-xl text-white/70 leading-relaxed">
            {feature.description}
          </p>

          <div className="space-y-4 pt-4">
            {feature.details?.map((detail, idx) => (
              <div key={idx} className="flex items-start gap-3">
                <div className="w-6 h-6 rounded-full bg-purple-500/20 flex items-center justify-center flex-shrink-0 mt-1">
                  <CheckCircle2 className="w-4 h-4 text-purple-400" />
                </div>
                <span className="text-white/80">{detail}</span>
              </div>
            ))}
          </div>

          <Button 
            onClick={() => setCurrentPage('specs')}
            className="bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 mt-6"
          >
            Learn More
            <ChevronRight className="w-4 h-4 ml-2" />
          </Button>
        </div>

        {/* Visual Side */}
        <div className={`relative ${!isEven ? 'lg:order-1' : ''}`}>
          <div className="relative aspect-square max-w-lg mx-auto">
            {/* Animated background circles */}
            <div className={`absolute inset-0 bg-gradient-to-br from-purple-500/20 to-pink-500/20 rounded-full blur-3xl ${isVisible ? 'animate-pulse' : ''}`}></div>
            
            <Card className="relative bg-gradient-to-br from-white/5 to-white/0 border-white/10 backdrop-blur-xl overflow-hidden h-full">
              <CardContent className="p-12 flex items-center justify-center h-full">
                <div className="relative w-full h-full flex items-center justify-center">
                  {/* Large Icon */}
                  <div className={`w-48 h-48 rounded-full bg-gradient-to-br from-purple-600 to-pink-600 flex items-center justify-center ${isVisible ? 'animate-bounce' : ''}`}
                    style={{ animationDuration: '3s' }}>
                    <div className="absolute inset-4 rounded-full bg-slate-950/80 backdrop-blur-sm"></div>
                    <IconComponent className="w-24 h-24 text-white relative z-10" />
                  </div>

                  {/* Orbiting particles */}
                  {[...Array(3)].map((_, i) => (
                    <div
                      key={i}
                      className={`absolute w-4 h-4 rounded-full bg-purple-400 ${isVisible ? 'animate-ping' : ''}`}
                      style={{
                        top: `${30 + i * 20}%`,
                        left: `${20 + i * 30}%`,
                        animationDelay: `${i * 0.5}s`,
                        animationDuration: '2s'
                      }}
                    ></div>
                  ))}
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

// Features Page with Scroll Animations
const FeaturesPage = ({ setCurrentPage }) => {
  const [features, setFeatures] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(`${API_URL}/features`)
      .then(res => res.json())
      .then(data => {
        const enhancedFeatures = (data.data || []).map((feature, index) => ({
          ...feature,
          details: getFeatureDetails(index)
        }));
        setFeatures(enhancedFeatures);
        setLoading(false);
      })
      .catch(err => {
        console.error('Error fetching features:', err);
        setLoading(false);
      });
  }, []);

  const getFeatureDetails = (index) => {
    const allDetails = [
      ['Studio-grade 40mm drivers', 'Support for Hi-Res audio codecs', 'Frequency response: 20Hz-40kHz'],
      ['Adaptive noise cancellation', 'Ambient sound mode', 'Wind noise reduction'],
      ['Fast charging support', 'Low battery indicator', 'Power-saving mode'],
      ['Multiple microphones', 'Echo cancellation', 'Voice enhancement'],
      ['Memory foam ear cushions', 'Adjustable headband', 'Lightweight design'],
      ['3D audio positioning', 'Head tracking technology', 'Dolby Atmos support']
    ];
    return allDetails[index] || [];
  };

  const animationTypes = ['slideLeft', 'slideRight', 'slideUp', 'scale', 'rotate', 'slideLeft'];

  return (
    <div className="pt-32 pb-20 px-6">
      <div className="max-w-7xl mx-auto">
        {/* Hero Section */}
        <div className="text-center mb-20">
          <Badge className="bg-purple-500/20 text-purple-300 border-purple-500/50 mb-6">
            <Sparkles className="w-3 h-3 mr-1" />
            Premium Features
          </Badge>
          <h1 className="text-6xl md:text-7xl font-bold mb-6">
            Innovation Meets
            <span className="block bg-gradient-to-r from-purple-400 via-pink-400 to-purple-400 bg-clip-text text-transparent">
              Excellence
            </span>
          </h1>
          <p className="text-xl text-white/60 max-w-2xl mx-auto">
            Discover the cutting-edge technology and thoughtful design that makes NOVA headphones extraordinary
          </p>
        </div>

        {loading ? (
          <div className="text-center py-20">
            <div className="animate-spin w-12 h-12 border-4 border-purple-500 border-t-transparent rounded-full mx-auto"></div>
          </div>
        ) : (
          <div className="space-y-20">
            {features.map((feature, index) => (
              <AnimatedFeature
                key={index}
                feature={feature}
                index={index}
                animationType={animationTypes[index % animationTypes.length]}
                setCurrentPage={setCurrentPage}
              />
            ))}
          </div>
        )}

        {/* CTA Section */}
        <div className="mt-32">
          <Card className="bg-gradient-to-br from-purple-600 to-pink-600 border-0 shadow-2xl shadow-purple-500/50">
            <CardContent className="p-12 text-center space-y-6">
              <h2 className="text-4xl font-bold">Experience All Features</h2>
              <p className="text-white/90 text-lg max-w-2xl mx-auto">
                See how NOVA's premium features work together to deliver an unmatched audio experience
              </p>
              <div className="flex flex-wrap gap-4 justify-center">
                <Button 
                  size="lg" 
                  onClick={() => setCurrentPage('contact')}
                  className="bg-white text-purple-600 hover:bg-white/90"
                >
                  Order Now - $299
                  <ChevronRight className="w-4 h-4 ml-2" />
                </Button>
                <Button 
                  size="lg" 
                  variant="outline" 
                  onClick={() => setCurrentPage('specs')}
                  className="border-white text-black hover:bg-white/10"
                >
                  View Specifications
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

// Specs Page
const SpecsPage = () => {
  const [specs, setSpecs] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch(`${API_URL}/specs`)
      .then(res => res.json())
      .then(data => {
        setSpecs(data.data || []);
        setLoading(false);
      })
      .catch(err => {
        console.error('Error fetching specs:', err);
        setLoading(false);
      });
  }, []);

  return (
    <div className="pt-32 pb-20 px-6">
      <div className="max-w-7xl mx-auto">
        <div className="text-center mb-12">
          <h1 className="text-5xl md:text-6xl font-bold mb-4">
            Technical <span className="bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent">Specifications</span>
          </h1>
          <p className="text-xl text-white/60">Every detail engineered to perfection</p>
        </div>

        {loading ? (
          <div className="text-center py-20">
            <div className="animate-spin w-12 h-12 border-4 border-purple-500 border-t-transparent rounded-full mx-auto"></div>
          </div>
        ) : (
          <div className="grid md:grid-cols-2 gap-6">
            {specs.map((spec, index) => (
              <Card key={index} className="bg-white/5 border-white/10 backdrop-blur-xl">
                <CardHeader>
                  <CardTitle className="text-2xl bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent">
                    {spec.category}
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <div className="space-y-3">
                    {spec.items.map((item, idx) => (
                      <div key={idx} className="flex items-start gap-3 py-2 border-b border-white/10 last:border-0">
                        <div className="w-2 h-2 rounded-full bg-purple-400 mt-2 flex-shrink-0"></div>
                        <span className="text-white/80">{item}</span>
                      </div>
                    ))}
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

// Contact Page
const ContactPage = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    subject: '',
    message: ''
  });
  const [status, setStatus] = useState({ type: '', message: '' });
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    if (!formData.name || !formData.email || !formData.subject || !formData.message) {
      setStatus({ type: 'error', message: 'Please fill in all fields' });
      return;
    }

    setLoading(true);
    setStatus({ type: '', message: '' });

    try {
      const response = await fetch(`${API_URL}/contact`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
      });

      const result = await response.json();

      if (result.success) {
        setStatus({ type: 'success', message: result.message });
        setFormData({ name: '', email: '', subject: '', message: '' });
      } else {
        throw new Error(result.message);
      }
    } catch (error) {
      setStatus({ 
        type: 'error', 
        message: 'Failed to send message. Please try again.' 
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="pt-32 pb-20 px-6">
      <div className="max-w-6xl mx-auto">
        <div className="text-center mb-12">
          <h1 className="text-5xl md:text-6xl font-bold mb-4">
            Get In <span className="bg-gradient-to-r from-purple-400 to-pink-400 bg-clip-text text-transparent">Touch</span>
          </h1>
          <p className="text-xl text-white/60">Have questions? We'd love to hear from you</p>
        </div>

        <div className="grid lg:grid-cols-3 gap-8">
          <div className="space-y-6">
            <Card className="bg-white/5 border-white/10 backdrop-blur-xl">
              <CardHeader>
                <CardTitle className="text-white">Contact Information</CardTitle>
                <CardDescription className="text-white/60">
                  Reach out through any channel
                </CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="flex items-start gap-3">
                  <div className="w-10 h-10 rounded-lg bg-purple-500/20 flex items-center justify-center flex-shrink-0">
                    <Mail className="w-5 h-5 text-purple-400" />
                  </div>
                  <div>
                    <div className="text-sm text-white/60">Email</div>
                    <div className="text-white">support@nova-audio.com</div>
                  </div>
                </div>

                <div className="flex items-start gap-3">
                  <div className="w-10 h-10 rounded-lg bg-purple-500/20 flex items-center justify-center flex-shrink-0">
                    <Phone className="w-5 h-5 text-purple-400" />
                  </div>
                  <div>
                    <div className="text-sm text-white/60">Phone</div>
                    <div className="text-white">0666666666</div>
                  </div>
                </div>

                <div className="flex items-start gap-3">
                  <div className="w-10 h-10 rounded-lg bg-purple-500/20 flex items-center justify-center flex-shrink-0">
                    <MapPin className="w-5 h-5 text-purple-400" />
                  </div>
                  <div>
                    <div className="text-sm text-white/60">Office</div>
                    <div className="text-white">Paris, France</div>
                  </div>
                </div>
              </CardContent>
            </Card>

          </div>

          <div className="lg:col-span-2">
            <Card className="bg-white/5 border-white/10 backdrop-blur-xl">
              <CardHeader>
                <CardTitle className="text-white text-2xl">Send us a Message</CardTitle>
                <CardDescription className="text-white/60">
                  We'll get back to you within 24 hours
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="space-y-6">
                  <div className="grid md:grid-cols-2 gap-6">
                    <div className="space-y-2">
                      <Label htmlFor="name" className="text-white">Full Name</Label>
                      <Input
                        id="name"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        className="bg-white/5 border-white/20 text-white placeholder:text-white/40"
                        placeholder="John Doe"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="email" className="text-white">Email Address</Label>
                      <Input
                        id="email"
                        name="email"
                        type="email"
                        value={formData.email}
                        onChange={handleChange}
                        className="bg-white/5 border-white/20 text-white placeholder:text-white/40"
                        placeholder="john@example.com"
                      />
                    </div>
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="subject" className="text-white">Subject</Label>
                    <Input
                      id="subject"
                      name="subject"
                      value={formData.subject}
                      onChange={handleChange}
                      className="bg-white/5 border-white/20 text-white placeholder:text-white/40"
                      placeholder="Product Inquiry"
                    />
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="message" className="text-white">Message</Label>
                    <Textarea
                      id="message"
                      name="message"
                      value={formData.message}
                      onChange={handleChange}
                      rows={6}
                      className="bg-white/5 border-white/20 text-white placeholder:text-white/40 resize-none"
                      placeholder="Tell us how we can help..."
                    />
                  </div>

                  {status.message && (
                    <Alert className={
                      status.type === 'success' 
                        ? 'bg-green-500/20 border-green-500/50 text-green-300' 
                        : 'bg-red-500/20 border-red-500/50 text-red-300'
                    }>
                      {status.type === 'success' ? (
                        <CheckCircle2 className="h-4 w-4" />
                      ) : (
                        <AlertCircle className="h-4 w-4" />
                      )}
                      <AlertDescription>{status.message}</AlertDescription>
                    </Alert>
                  )}

                  <Button 
                    onClick={handleSubmit}
                    disabled={loading}
                    className="w-full bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 text-white"
                    size="lg"
                  >
                    {loading ? 'Sending...' : (
                      <>
                        Send Message
                        <Send className="w-4 h-4 ml-2" />
                      </>
                    )}
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

// Main App Component
export default function App() {
  const [currentPage, setCurrentPage] = useState('home');

  useEffect(() => {
    window.scrollTo(0, 0);
  }, [currentPage]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-950 via-purple-950 to-slate-950 text-white overflow-x-hidden">
      <div className="fixed inset-0 overflow-hidden pointer-events-none">
        <div className="absolute top-1/4 -left-1/4 w-96 h-96 bg-purple-500 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-pulse"></div>
        <div className="absolute bottom-1/4 -right-1/4 w-96 h-96 bg-pink-500 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-pulse"></div>
      </div>

      <Navigation currentPage={currentPage} setCurrentPage={setCurrentPage} />

      <main className="relative z-10">
        {currentPage === 'home' && <HomePage setCurrentPage={setCurrentPage} />}
        {currentPage === 'features' && <FeaturesPage setCurrentPage={setCurrentPage} />}
        {currentPage === 'specs' && <SpecsPage />}
        {currentPage === 'contact' && <ContactPage />}
      </main>
    </div>
  );
}
