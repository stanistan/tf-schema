// This generates structs and conversion functions
// so that you can use a "typed" terraform provider.
//
// Usage:
//
// 	go run . foo.Bar github.com/some/source/foo
// 	go run . Bar
//
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"

	"github.com/pkg/errors"
)

const tfSchema = "github.com/stanistan/tf-schema/v2"

func main() {
	packageName := os.Getenv("GOPACKAGE")
	if packageName == "" {
		die(errors.New("missing package name in GOPACKAGE"))
	}

	vars, err := makeVars(os.Args[1:])
	if err != nil {
		die(errors.Wrapf(err, "could not make vars from args %q", os.Args[1:]))
	}

	filename := fmt.Sprintf("generated_%s_schema.go", vars.NS)
	f, err := os.Create(filename)
	die(err)
	defer f.Close()

	var buf bytes.Buffer
	die(packageTemplate.Execute(&buf, struct {
		*packageVars
		Self       string
		Args       []string
		Package    string
		TypeDefs   map[string]*templateTypeDef
		Signatures map[string]signatureType
	}{
		packageVars: vars,
		Self:        tfSchema,
		Args:        os.Args[1:],
		Package:     os.Getenv("GOPACKAGE"),
		TypeDefs: map[string]*templateTypeDef{
			"Configure": {"ConfigureContext", typedReturnSignature},
			"Create":    {"CreateContext", singleReturnSignature},
			"Read":      {"ReadContext", singleReturnSignature},
			"Delete":    {"DeleteContext", singleReturnSignature},
			"Update":    {"UpdateContext", singleReturnSignature},
		},
		Signatures: map[string]signatureType{
			"singleReturn": singleReturnSignature,
			"typedReturn":  typedReturnSignature,
		},
	}))

	p, err := format.Source(buf.Bytes())
	die(errors.Wrapf(err, "could not format source"))

	_, err = f.Write(p)
	die(errors.Wrapf(err, "could not write the file"))
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
