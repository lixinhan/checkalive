package dispatcher

import (
	"checkalive/checker"
	"checkalive/config"
	"checkalive/reporter"
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
	for _,v:= range dispatcher.Config.Reporter.Dingtalk{
		reporter:=reporter.DingtalkGroupMessage{}
		reporter.SetAccessToken(v.AccessToken);
		reporter.SetSecret(v.Secret);
		reporter.SendMessage(message)
	}
}


