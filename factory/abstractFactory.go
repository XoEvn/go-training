package factory

type IRuleConfigParser1 interface {
	Parse(data []byte)
}
type jsonRuleConfigParser1 struct {
}

func (j jsonRuleConfigParser1) Parse(data []byte) {
	panic("implement me")
}

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}
type jsonSystemConfigParser struct {
}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser1
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser1 {
	return jsonRuleConfigParser1{}
}
func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}
