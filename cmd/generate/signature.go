package main

import (
	"fmt"
	"strings"
)

type signatureType string

var (
	typedReturnSignature  signatureType = "func(*tfschema.ResourceData) (*%s, error)"
	singleReturnSignature signatureType = "func(*tfschema.ResourceData, *%s) error"
	boolReturnSignature   signatureType = "func(*tfschema.ResourceData, *%s) (bool, error)"
)

func (s signatureType) WrapFn(ns string) string {
	var suffix string

	switch s {
	case typedReturnSignature:
		suffix = "typedReturn"
	case singleReturnSignature:
		suffix = "singleReturn"
	case boolReturnSignature:
		suffix = "boolReturn"
	}

	return fmt.Sprintf("_clientWrap_%s_%s", suffix, ns)
}

type templateTypeDef struct {
	n             string
	signatureType signatureType
}

func (t *templateTypeDef) TFName() string {
	name := t.n
	return "tfschema." + strings.ToUpper(string(name[0])) + name[1:] + "Func"
}

func (t *templateTypeDef) Name(ns string) string {
	return fmt.Sprintf("_%sFunc_%s", t.n, ns)
}

func (t *templateTypeDef) DeclareName(ns string) string {
	return t.n + " " + t.Name(ns)
}

func (t *templateTypeDef) FieldAssignWrapFn(ns string) string {
	return fmt.Sprintf("%s: %s(\"%s\", r.%s)", t.n, t.signatureType.WrapFn(ns), t.Name(ns), t.n)
}

func (t *templateTypeDef) Signature(client string) string {
	return fmt.Sprintf(string(t.signatureType), client)
}
