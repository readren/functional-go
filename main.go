package main

import (
	"fmt"
	"os"
	"path/filepath"

	typeCtors "github.com/readren/functional-go/typectors"
)

func main() {

	if false {
		s := typeCtors.BuildTypeName("map[string][3]image.Point")
		fmt.Println(s)

	} else {

		workingDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		config := typeCtors.Config{
			GeneratedPackageParentDir: filepath.Join(workingDir, "instantiations"),
			GeneratedPackageName:      "fung",
			TemplatesFolder:           filepath.Join(workingDir, "typectors", "templates"),
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

		typeCtors.GeneratePackage(config)
	}
}
