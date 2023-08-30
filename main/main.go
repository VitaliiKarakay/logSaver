package main

import "time"

/*Необходимо создать сервис который будет принимать и сохранять логи от сервисов которые отправляют SMS и Email сообщения.
Какие данные будут приходить:
{
userID: 2134496917,
phone: '380953071221',
actionID: 324,
actionTitle: 'Good Action',
actionType: 'promoSMS',
message: 'some message',
sender: 'intistele',
status: 'success',
language: 'en',
fullResponce: 'sms_id: 6774560000068401360004',
created: "2023-02-27T00:27:00.031Z",
updated: "2023-02-27T00:27:00.031Z",
messageID: '6774560000068401360004'
}
БД: любая на твое усмотрение
Web: Gin */

type Log struct {
	UserId       int
	Phone        string
	ActionID     int
	ActionTitle  string
	ActionType   string
	Message      string
	Sender       string
	Status       string
	Language     string
	FullResponse string
	Created      time.Time
	Updated      time.Time
	MessageId    string
}

func main() {

}
