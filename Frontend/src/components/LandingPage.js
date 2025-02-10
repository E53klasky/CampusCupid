import React from "react";
import { Container, Typography, Button, Grid, Card, CardContent } from "@mui/material";
import FavoriteIcon from "@mui/icons-material/Favorite";
import PeopleIcon from "@mui/icons-material/People";
import SecurityIcon from "@mui/icons-material/Security";
import { styled } from "@mui/system";
import { useNavigate } from "react-router-dom"; // Import useNavigate

// Tinder color palette
const colors = {
  orange: "#FF7158",
  magenta: "#FD2B7B",
  gray: "#424242",
  light: "#FFFFFF",
  background: "#F8F8F8",
};

const HeroSection = styled("div")({
  textAlign: "center",
  padding: "100px 20px",
  background: `linear-gradient(135deg, ${colors.orange} 0%, ${colors.magenta} 100%)`,
  "& h2": {
    color: colors.light,
    fontWeight: 400,
    marginBottom: "1.5rem",
    fontFamily: "'Great Vibes', cursive",
    fontSize: "4.5rem",
  },
  "& h5": {
    color: colors.light,
    fontWeight: 400,
    marginBottom: "2rem",
    fontFamily: "'Poppins', sans-serif",
  },
});

const FeaturesSection = styled("div")({
  padding: "80px 20px",
  backgroundColor: colors.background,
  "& .MuiCard-root": {
    transition: "transform 0.2s ease-in-out",
    "&:hover": {
      transform: "translateY(-5px)",
    },
  },
  "& .MuiCardContent-root": {
    textAlign: "center",
    padding: "2rem",
  },
  "& .MuiSvgIcon-root": {
    fontSize: "3rem",
    marginBottom: "1rem",
    color: colors.magenta,
  },
});

const CallToAction = styled("div")({
  textAlign: "center",
  padding: "80px 20px",
  background: colors.gray,
  "& h4": {
    color: colors.light,
    fontWeight: 700,
    marginBottom: "2rem",
    fontFamily: "'Poppins', sans-serif",
  },
});

const StyledButton = styled(Button)({
  backgroundColor: colors.magenta,
  color: colors.light,
  padding: "12px 32px",
  borderRadius: "30px",
  fontSize: "1.1rem",
  fontFamily: "'Poppins', sans-serif",
  textTransform: "none",
  boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
  "&:hover": {
    backgroundColor: colors.orange,
    transform: "translateY(-2px)",
    boxShadow: "0 6px 8px rgba(0, 0, 0, 0.15)",
  },
});

const LandingPage = () => {
  const navigate = useNavigate(); // Initialize useNavigate

  return (
    <div>
      <HeroSection>
        <Typography variant="h2" gutterBottom>
          Campus Cupid
        </Typography>
        <Typography variant="h5" paragraph>
          Join millions of people connecting through love
        </Typography>
        <StyledButton 
          variant="contained" 
          size="large"
          onClick={() => navigate("/login")} // Navigate to login
        >
          Get Started
        </StyledButton>
      </HeroSection>

      <FeaturesSection>
        <Container>
          <Typography 
            variant="h4" 
            align="center" 
            gutterBottom 
            sx={{ 
              fontFamily: "'Poppins', serif",
              color: colors.gray,
              marginBottom: "3rem"
            }}
          >
            Why Choose Us?
          </Typography>
          <Grid container spacing={4} justifyContent="center">
            <Grid item xs={12} sm={4}>
              <Card elevation={2}>
                <CardContent>
                  <FavoriteIcon />
                  <Typography variant="h6" gutterBottom sx={{ fontFamily: "'Poppins', sans-serif" }}>
                    Smart Matching
                  </Typography>
                  <Typography sx={{ color: colors.gray }}>
                    Our advanced algorithm helps you find the right match
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
            <Grid item xs={12} sm={4}>
              <Card elevation={2}>
                <CardContent>
                  <PeopleIcon />
                  <Typography variant="h6" gutterBottom sx={{ fontFamily: "'Poppins', sans-serif" }}>
                    Large Community
                  </Typography>
                  <Typography sx={{ color: colors.gray }}>
                    Join a diverse and growing user base looking for love
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
            <Grid item xs={12} sm={4}>
              <Card elevation={2}>
                <CardContent>
                  <SecurityIcon />
                  <Typography variant="h6" gutterBottom sx={{ fontFamily: "'Poppins', sans-serif" }}>
                    Safe & Secure
                  </Typography>
                  <Typography sx={{ color: colors.gray }}>
                    Your data and privacy are our top priority
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Container>
      </FeaturesSection>

      <CallToAction>
        <Typography variant="h4" gutterBottom>
          LOVE STARTS HERE
        </Typography>
        <StyledButton 
          variant="contained" 
          size="large"
          onClick={() => navigate("/login")} // Navigate to login
          sx={{ 
            backgroundColor: colors.orange,
            "&:hover": {
              backgroundColor: colors.magenta,
            }
          }}
        >
          Sign Up Now
        </StyledButton>
      </CallToAction>
    </div>
  );
};

export default LandingPage;
