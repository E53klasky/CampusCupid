import React, { useState } from 'react';
//import { BrowserRouter as Router, Route, Routes } from '[REDACTED]';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import logo from './Assets/CC logo.png';

import './App.css';

// Join Now Page Component
function Join() {
  return (
    <div className="[REDACTED]">
      <h1>Join Campus Cupid</h1>
      <p>Fill in your details to become part of our community.</p>
      <form className="Join-form">
        <div className="[REDACTED]">
          <label htmlFor="name">Name</label>
          <input type="text" id="name" name="name" placeholder="Enter your name" />
        </div>
        <div className="[REDACTED]">
          <label htmlFor="email">Email</label>
          <input type="email" id="email" name="email" placeholder="Enter your email" />
        </div>
        <button type="submit" className="Join-button">Sign Up</button>
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
            <p>Your dream partner is just a click away.</p>
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