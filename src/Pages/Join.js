import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import logo from './Assets/CC logo.png';
import './App.css';
import Join from './pages/Join';

function App() {
  return (
    <Router>
      <div className="App">
        {/* Header */}
        <header className="App-header">
          <div className="App-header-logo">
            <img src={logo} alt="Campus Cupid Logo" />
            <h1>Campus Cupid</h1>
          </div>
          <nav className="App-nav">
            <Link to="/">Home</Link>
            <Link to="/join">Join Now</Link>
            <Link to="/profile">Profile</Link>
          </nav>
        </header>

        <div className="App-body">
          <Routes>
            <Route
              path="/"
              element={
                <>
                  <h1>Welcome to Campus Cupid!</h1>
                  <p>
                    Campus Cupid is the perfect platform to connect with like-minded individuals. Whether you're looking for a soulmate or a companion, your journey starts here.
                  </p>
                  <div className="App-buttons">
                    <Link to="/join" className="App-button">
                      Join Now
                    </Link>
                  </div>
                </>
              }
            />
            <Route path="/join" element={<Join />} />
          </Routes>
        </div>

        {/* Footer */}
        <footer className="App-footer">
          <p>&copy; 2025 Campus Cupid. All rights reserved.</p>
          <div className="App-footer-links">
            <a href="#privacy">Privacy Policy</a>
            <a href="#terms">Terms of Service</a>
            <a href="#contact">Contact Us</a>
          </div>
        </footer>
      </div>
    </Router>
  );
}

export default App;
