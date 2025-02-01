package dto

type SwipeRequest struct {
	SwipedUserId string `json:"swiper_user_id" binding:"required"`
	Preference   string `json:"preference" binding:"required" enum:"YES,NO"`
}
