package main

import (
	"strings"

	"github.com/pkg/errors"
)

type packageVars struct {
	Client          string
	NS              string
	ImportedPackage *importedPackage
}

func makeVars(args []string) (*packageVars, error) {
	if len(args) == 0 {
		return nil, errors.New("no args provided")
	}

	clientName := args[0]
	imported, err := makeImportedPackage(clientName, args[1:])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ns string
	if imported != nil {
		ns = imported.Name
	} else {
		ns = strings.ToLower(clientName)
	}

	return &packageVars{
		Client:          clientName,
		NS:              ns,
		ImportedPackage: imported,
	}, nil
}

type importedPackage struct {
	Name   string
	Source string
}

func makeImportedPackage(name string, args []string) (*importedPackage, error) {
	chunks := strings.Split(name, ".")
	if len(chunks) > 2 {
		return nil, errors.New("invalid client name provided " + name)
	}

	if len(chunks) == 1 {
		return nil, nil
	}

	if len(args) < 1 {
		return nil, errors.New("missing the import path for client package " + chunks[0])
	}

	return &importedPackage{
		Name:   chunks[0],
		Source: args[0],
	}, nil
}
