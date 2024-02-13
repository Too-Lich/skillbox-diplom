package main

import (
	"diplom/internal/mms"
	"diplom/internal/sms"
)

func main() {

	sms.CheckSms()
	mms.CheckMMS()

}
