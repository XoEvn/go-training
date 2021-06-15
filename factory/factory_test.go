package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{

		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v ,want %v", got, tt.want)
			}
		})
	}
}


func Test_jsonConfigParserFactory_CreateRuleParser(t *testing.T) {
	tests := []struct {
		name string
		want IRuleConfigParser
	}{
		{
			name: "json",
			want: jsonRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateRuleParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRuleParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		want ISystemConfigParser
	}{
		{
			name: "json",
			want: jsonSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := jsonConfigParserFactory{}
			if got := j.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
