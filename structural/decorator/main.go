package main

import "fmt"

type notificationResp struct {
	err           error
	destinationId string
}

type notification interface {
	send() notificationResp
}

type smsNotification struct {
	destinationId string
	msg           string
}

func (sn smsNotification) send() notificationResp {
	fmt.Println("Sending SMS notification")
	fmt.Println(sn.msg)

	if len(sn.destinationId) > 0 {
		fmt.Println("Delivered to: ", sn.destinationId)
	}

	return simulateResponse(sn.destinationId)
}

type slackNotification struct {
	slackUser    string
	notification notification
}

func (s slackNotification) send() notificationResp {
	resp := s.notification.send()

	if len(resp.destinationId) > 0 {
		fmt.Println("Sending to slack user: ", s.slackUser)
	}

	return simulateResponse(resp.destinationId)
}

func main() {
	sms := smsNotification{
		destinationId: "111111111111",
		msg:           "Message for SMS",
	}

	slack := slackNotification{
		slackUser:    "slack-user-sq",
		notification: sms,
	}

	slack.send()
}

func simulateResponse(dest string) notificationResp {
	return notificationResp{
		nil,
		dest,
	}
}
