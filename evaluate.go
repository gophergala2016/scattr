package main

import (
	"github.com/robertkrimen/otto"
)

func evaluateScript(src string, payload string) (string, error) {
	javaScript := otto.New()
	var evalFunc otto.Value
	javaScript.Set("rule", func(call otto.FunctionCall) otto.Value {
		evalFunc = call.Argument(0)
		return otto.UndefinedValue()
	})

	javaScript.Run(src)
	arg, err := javaScript.ToValue(payload)

	if err != nil {
		return "", err
	}

	ret, err := evalFunc.Call(otto.NullValue(), arg)

	if err != nil {
		return "", err
	}

	return ret.ToString()
}