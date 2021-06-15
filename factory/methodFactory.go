package factory


type IRuleConfigParserFactory interface {
	CreateParser()IRuleConfigParser
}

type yamlRuleConfigParserFactory struct {

}

func(y yamlRuleConfigParserFactory)CreateParser()IRuleConfigParser{
	return yamlRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct {

}


func (j jsonRuleConfigParserFactory)CreateParser()IRuleConfigParser{
	return jsonRuleConfigParser{}
}

func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory{
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}