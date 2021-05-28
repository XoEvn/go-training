package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{

}

func(J jsonRuleConfigParser)Parse(data[]byte){
	panic("implement me")
}

type yamlRuleConfigParser struct {

}

func(Y yamlRuleConfigParser)Parse(data[]byte){
	panic("implement me")
}

func NewRuleConfigParser(t string) IRuleConfigParser{
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}

	}
	return nil

}



