package main

import "fmt"

type Notifer interface {
	Send(msg string)
}

type Application struct {
	notifer Notifer
}

func (a *Application) SetNotifer(notifer Notifer) {
	a.notifer = notifer
}

func (a Application) DoSomething() {
	a.notifer.Send("I'm doing something")
}

type EmailNotifer struct{}

func (n EmailNotifer) Send(msg string) {
	fmt.Println("Send by email : ", msg)
}

type FacebookDecorator struct {
	warapper Notifer
}

func (fb FacebookDecorator) Send(msg string){
	fb.warapper.Send(msg)
	fmt.Println("Send by facebook", msg)
}

type SlackDecorator struct {
	wrapper Notifer
}

func (sl SlackDecorator) Send(msg string){
	sl.wrapper.Send(msg)
	fmt.Println("Send by slackL : ", msg)
}

func main(){
	var stack Notifer
	stack = EmailNotifer{}
	facebookEnable := true
	slackEnable := true
	if facebookEnable {
		stack = FacebookDecorator{warapper: stack}
	}
	if slackEnable {
		stack = SlackDecorator{wrapper: stack}
	}

	app := Application{}
	app.SetNotifer(stack)
	app.DoSomething()

}