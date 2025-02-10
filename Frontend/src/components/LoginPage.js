import React, { useState } from "react";
import { Container, TextField, Button, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";

function LoginPage() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = () => {
    console.log("Logging in with", email, password);
    navigate("/"); // Redirect to Landing Page after login
  };

  return (
    <Container style={{ textAlign: "center", marginTop: "50px" }}>
      <Typography variant="h5">Login</Typography>
      <TextField 
        label="Email" 
        fullWidth 
        margin="normal" 
        value={email} 
        onChange={(e) => setEmail(e.target.value)}
      />
      <TextField 
        label="Password" 
        type="password" 
        fullWidth 
        margin="normal" 
        value={password} 
        onChange={(e) => setPassword(e.target.value)}
      />
      <Button 
        variant="contained" 
        color="primary" 
        style={{ marginTop: "20px" }}
        onClick={handleLogin}
      >
        Login
      </Button>
    </Container>
  );
}

export default LoginPage;
