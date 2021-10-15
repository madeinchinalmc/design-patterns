package template

import "fmt"

type iOpt interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type opt struct {
	iOpt iOpt
}

func (o *opt) genAndSendOTP(otpLength int) error {
	optStr := o.iOpt.genRandomOTP(otpLength)
	o.iOpt.saveOTPCache(optStr)
	message := o.iOpt.getMessage(optStr)
	err := o.iOpt.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOpt.publishMetric()
	return nil
}

type sms struct {
	opt
}

func (s *sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS：generating random otp %s \n", randomOTP)
	return randomOTP
}

func (s *sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(msg string) error {
	fmt.Printf("SMS: sending sms: %s\n", msg)
	return nil
}
func (s *sms) publishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}

type email struct {
	opt
}

func (e *email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Email：generating random otp %s \n", randomOTP)
	return randomOTP
}

func (e *email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache\n", otp)
}

func (e *email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *email) sendNotification(msg string) error {
	fmt.Printf("Email: sending sms: %s\n", msg)
	return nil
}
func (e *email) publishMetric() {
	fmt.Printf("Email: publishing metrics\n")
}

// 客户端代码

func RunApplication() {

	mySms := &sms{}
	o := opt{
		iOpt: mySms,
	}
	_ = o.genAndSendOTP(4)

	myEmail := &email{}
	o = opt{
		iOpt: myEmail,
	}
	_ = o.genAndSendOTP(4)
}
