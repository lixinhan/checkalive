package template

import (
	"checkalive/checker"
	"strings"
)

type Template struct {
	TemplateString string
	Error string
}
func (template *Template) register() []string{
	return []string{
		"%error%",
		"%name%",
		"%url%"}
}
func (template *Template) ParseUrlMessage(rules checker.UrlCheckRules) string  {
		templateString:=template.TemplateString
	for _,v:= range template.register(){
		if v=="%name%" {
			templateString=strings.Replace(templateString,v,rules.Name,-1)
		}else if v=="%url%"{
			templateString=strings.Replace(templateString,v,rules.Url,-1)
		}else if v=="%error%"{
			templateString=strings.Replace(templateString,v,template.Error,-1)
		}
	}
	return templateString;
}

