package worker

import "fmt"

const (
	QUEUE_CRITICAL = "critical"
	QUEUE_DEFAULT  = "default"
)

const (
	TASK_SEND_MATCHED_EMAIL            = "task:send_matched_email"
	TASK_CALCULATE_USER_ATTRACTIVENESS = "task:calculate_user_attractiveness"
)

type SendMatchEmailTaskPayload struct {
	UserId        string
	MatchedUserId string
	CorrelationId string
}

type CalculateUserAttractivenessTaskPayload struct {
	Userid        string
	Response      string
	CorrelationId string
}

const EmailBody = `
Dear %s,

Congratulations! ❤️ You have a new match on our platform. ❤️

Here are the details of your match:
Name: %s

We encourage you to reach out to your match and start a conversation. Who knows where it might lead? ❤️

Thank you for using our platform. We wish you all the best in your journey to find a match.

Best regards,
[Your Name]
[Your Position]
`

func EmailBodyBuilder(userName, matchedUserName string) string {
	return fmt.Sprintf(EmailBody, userName, matchedUserName)
}
