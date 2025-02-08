import React, { useState } from 'react';
//import { BrowserRouter as Router, Route, Routes } from '[REDACTED]';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import logo from './Assets/CC logo.png';

import './App.css';

// Join Now Page Component
function Join() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    dateOfBirth: '',
    gender: '',
    interestedIn: '',
    lookingFor: [],
    interests: '',
    sexualOrientation: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleCheckboxChange = (e) => {
    const { value, checked } = e.target;
    setFormData(prevState => ({
      ...prevState,
      lookingFor: checked 
        ? [...prevState.lookingFor, value]
        : prevState.lookingFor.filter(item => item !== value)
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Form submitted:', formData);
    // Add your form submission logic here
  };

  return (
    <div className="join-container">
      <h1>Join Campus Cupid</h1>
      <p>Fill in your details to become part of our community.</p>
      <form className="join-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="name">Name</label>
          <input 
            type="text" 
            id="name" 
            name="name" 
            value={formData.name}
            onChange={handleChange}
            placeholder="Enter your name" 
            required 
          />
        </div>

        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input 
            type="email" 
            id="email" 
            name="email" 
            value={formData.email}
            onChange={handleChange}
            placeholder="Enter your email" 
            required 
          />
        </div>

        <div className="form-group">
          <label htmlFor="dateOfBirth">Date of Birth</label>
          <input 
            type="date" 
            id="dateOfBirth" 
            name="dateOfBirth"
            value={formData.dateOfBirth}
            onChange={handleChange}
            required 
          />
        </div>

        <div className="form-group">
          <label htmlFor="gender">Gender</label>
          <select 
            id="gender" 
            name="gender"
            value={formData.gender}
            onChange={handleChange}
            required
          >
            <option value="">Select gender</option>
            <option value="male">Male</option>
            <option value="female">Female</option>
            <option value="non-binary">Non-binary</option>
            <option value="other">Other</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="interestedIn">Interested In</label>
          <select 
            id="interestedIn" 
            name="interestedIn"
            value={formData.interestedIn}
            onChange={handleChange}
            required
          >
            <option value="">Select preference</option>
            <option value="men">Men</option>
            <option value="women">Women</option>
            <option value="everyone">Everyone</option>
          </select>
        </div>

        <div className="form-group">
          <label>Looking For</label>
          <div className="checkbox-group">
            <label>
              <input 
                type="checkbox" 
                name="lookingFor" 
                value="friendship"
                onChange={handleCheckboxChange}
              /> Friendship
            </label>
            <label>
              <input 
                type="checkbox" 
                name="lookingFor" 
                value="dating"
                onChange={handleCheckboxChange}
              /> Dating
            </label>
            <label>
              <input 
                type="checkbox" 
                name="lookingFor" 
                value="relationship"
                onChange={handleCheckboxChange}
              /> Relationship
            </label>
          </div>
        </div>

        <div className="form-group">
          <label htmlFor="interests">Interests</label>
          <textarea 
            id="interests" 
            name="interests"
            value={formData.interests}
            onChange={handleChange}
            placeholder="Tell us about your interests (e.g., movies, music, sports)"
            rows="4"
          />
        </div>

        <div className="form-group">
          <label htmlFor="sexualOrientation">Sexual Orientation</label>
          <select 
            id="sexualOrientation" 
            name="sexualOrientation"
            value={formData.sexualOrientation}
            onChange={handleChange}
            required
          >
            <option value="">Select orientation</option>
            <option value="straight">Straight</option>
            <option value="gay">Gay</option>
            <option value="lesbian">Lesbian</option>
            <option value="bisexual">Bisexual</option>
            <option value="pansexual">Pansexual</option>
            <option value="asexual">Asexual</option>
            <option value="other">Other</option>
          </select>
        </div>

        <button type="submit" className="join-button">Sign Up</button>
      </form>
    </div>
  );
}
// Profile Page Component
function Profile() {
  const [uploadedImage, setUploadedImage] = useState(null);

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setUploadedImage(reader.result);
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="[REDACTED]">
      <h1>Your Profile</h1>
      <div className="[REDACTED]">
        <h2>Your Profile Picture</h2>
        <div className="[REDACTED]">
        
        </div>
        
        <div className="[REDACTED]">
          <input
            type="file"
            id="[REDACTED]"
            accept="image/*"
            onChange={handleImageChange}
            className="[REDACTED]"
          />
          <label htmlFor="[REDACTED]" className="[REDACTED]">Change Image</label>
        </div>
      </div>
    </div>
  );
}

function App() {
  return (
    <Router>
      <div className="App">
        {/* Header Section */}
        <header className="App-header">
          <h1>Campus Cupid</h1>
        </header>

        <div className="App-main">
          {/* Left Branding Section */}
          <div className="[REDACTED]">
            <img src={logo} className="App-logo" alt="Campus Cupid Logo" />
            <h2>Welcome to Campus Cupid!</h2>
            <p>Your dream partner is just a click away</p>
          </div>

          {/* Main Content Section */}
          <div className="App-content">
            <Routes>
              <Route
                path="/"
                element={
                  <>
                    <h1>Love Starts Here</h1>
                    <div className="[REDACTED]">
                      <a href="/join" className="App-button">Join Now</a>
                      <a href="/matches" className="App-button">Your Matches</a>
                      <a href="/profile" className="App-button">Your Profile</a>
                    </div>
                  </>
                }
              />
              <Route path="/join" element={<Join />} />
              <Route path="/profile" element={<Profile />} />
            </Routes>
          </div>
        </div>

        {/* Footer Section */}
        <footer className="App-footer">
          <p>&copy; 2025 Campus Cupid. All rights reserved.</p>
        </footer>
      </div>
    </Router>
  );
}

export default App;