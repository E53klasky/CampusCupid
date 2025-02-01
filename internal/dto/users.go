package dto

type CreateUserRequest struct {
	Email     string  `json:"email" binding:"required,email" faker:"email"`
	Name      string  `json:"name" binding:"required" faker:"name"`
	Password  string  `json:"password" binding:"required" faker:"word"`
	Gender    string  `json:"gender" binding:"required" faker:"oneof: male, female, binary"`
	Age       int8    `json:"age" binding:"required,gte=18,lte=100" faker:"oneof: 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30"`
	Latitude  float64 `json:"latitude" binding:"required" faker:"lat"`
	Longitude float64 `json:"longitude" binding:"required" faker:"long"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Age      int8   `json:"age"`
}

type DiscoverV2Request struct {
	Age struct {
		LessThan    *int32 `json:"less_than"`
		GreaterThan *int32 `json:"greater_than"`
	} `json:"age"`
	Gender string `json:"gender,omitempty"`
}

type DiscoverV2Response struct {
	UserID   string  `json:"user_id"`
	Name     string  `json:"name"`
	Gender   string  `json:"gender"`
	Age      int32   `json:"age"`
	Distance float64 `json:"distance"`
}
