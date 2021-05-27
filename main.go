package main

import (
	"os"
	"path/filepath"

	typeCtors "github.com/readren/functional-go/typectors"
)

func main() {

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := typeCtors.Config{
		GeneratedPackageParentDir: filepath.Join(workingDir, "instantiations"),
		GeneratedPackageName:      "fung",
		TemplatesBaseDir:          "templates",
		TypesDescriptors: []typeCtors.TypeDescriptor{
			{
				TypeConstructorName: "Stream",
				BaseTypeArguments:   typeCtors.TypeArguments{},
				FuncTypeArgumentsForWhichFuncsAreIncluded: []typeCtors.TypeArguments{
					{{Type: "int"}},
					{{Type: "string"}},
				},
			},
			{
				TypeConstructorName: "Stream",
				BaseTypeArguments:   typeCtors.TypeArguments{{Type: "int"}},
				FuncTypeArgumentsForWhichFuncsAreIncluded: []typeCtors.TypeArguments{
					{},
					{{Type: "string"}},
					{{Type: "int"}},
					{{Type: "image.Point", PackagePath: "image"}},
					{{Type: "string"}, {Type: "[]bool"}},
				},
			},
			{
				TypeConstructorName: "Stream",
				BaseTypeArguments:   typeCtors.TypeArguments{{Type: "string"}},
				FuncTypeArgumentsForWhichFuncsAreIncluded: []typeCtors.TypeArguments{
					{},
					{{Type: "string"}},
				},
			},
			{
				TypeConstructorName: "Validate",
				BaseTypeArguments:   typeCtors.TypeArguments{{Type: "int"}, {Type: "string"}},
				FuncTypeArgumentsForWhichFuncsAreIncluded: []typeCtors.TypeArguments{
					{{Type: "string"}},
				},
			},
			{
				TypeConstructorName: "ValiResu",
				BaseTypeArguments:   typeCtors.TypeArguments{{Type: "int"}, {Type: "string"}},
				FuncTypeArgumentsForWhichFuncsAreIncluded: []typeCtors.TypeArguments{
					{{Type: "string"}},
				},
			},
		},
	}

	typeCtors.BuildFungPackage(config)
}
