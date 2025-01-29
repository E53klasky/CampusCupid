import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import logo from './Assets/CC logo.png';
import './App.css';

// Join Now Page Component
function Join() {
  return (
    <div className="Join-container">
      <h1>Join Campus Cupid</h1>
      <p>Fill in your details to become part of our community.</p>
      <form className="Join-form">
        <div className="Join-form-group">
          <label htmlFor="name">Name</label>
          <input type="text" id="name" name="name" placeholder="Enter your name" />
        </div>
        <div className="Join-form-group">
          <label htmlFor="email">Email</label>
          <input type="email" id="email" name="email" placeholder="Enter your email" />
        </div>
        <button type="submit" className="Join-button">Sign Up</button>
      </form>
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
          <div className="App-logo-section">
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
                    <h1>Discover Your Match</h1>
                    <p>Join our community and find your perfect match today.</p>
                    <div className="App-buttons">
                      <a href="/join" className="App-button">Join Now</a>
                      <a href="/matches" className="App-button">Your Matches</a>
                      <a href="/profile" className="App-button">Your Profile</a>
                    </div>
                  </>
                }
              />
              <Route path="/join" element={<Join />} />
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
