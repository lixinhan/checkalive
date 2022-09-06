package dispatcher

import (
	"checkalive/checker"
	"checkalive/config"
	"checkalive/recipient"
	"checkalive/template"
)

type Dispatcher struct {
	Config config.Config
}


func (dispatcher *Dispatcher) Url(rules checker.UrlCheckRules)  {
	checker:=checker.UrlChecker{}
	checker.SetUrl(rules.Url)
	err,_:=checker.Check()
	if(err!=nil){
		template:=template.Template{
			dispatcher.Config.Template,
			err.Error(),
		}
		dispatcher.sendMessage(template.ParseUrlMessage(rules))
	}

}
func (dispatcher *Dispatcher) sendMessage(message string )  {
	for _,v:= range dispatcher.Config.Recipient.Dingtalk{
		recipient:=recipient.DingtalkGroupMessage{}
		recipient.SetAccessToken(v.AccessToken);
		recipient.SetSecret(v.Secret);
		recipient.SendMessage(message)
	}
}


