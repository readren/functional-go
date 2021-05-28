package typectors

import (
	"embed"
	"fmt"
)

//go:embed templates
var fungTemplatesFS embed.FS

// Contains all the known instances of `TypeConstructor` indexed by its name.
var fungTypeConstructors map[string]TypeConstructor = map[string]TypeConstructor{
	"Recover": {
		{ //0 baseParam
			[]string{},
			[]Template{
				{"Recover", []string{}},               // 0 funcParams
				{"Recover__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "Recover"
			},
		},
	},
	"Errors": {
		{ //0 baseParam
			[]string{},
			[]Template{
				{"Errors", []string{}},               // 0 funcParams
				{"Errors__kType", []string{"kType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "Errors"
			},
		},
		{ //1 baseParam
			[]string{"kType"},
			[]Template{
				{"kType__Errors", []string{"kType"}}, // 0 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "Errors"
			},
		},
	},
	"Func1": {
		{ // 0 baseParam
			[]string{},
			[]Template{
				{"Func1", []string{}},               // 0 funcParams
				{"Func1__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "Func1"
			},
		},
		emptyChapter, // 1 baseParam
		{ //2 baseParam
			[]string{"xType", "yType"},
			[]Template{
				{"xType__yType__Func1", []string{}},               // 0 funcParams
				{"xType__yType__Func1__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("FuncFrom_%s_to_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
			},
		},
	},
	"FlawedFunc1": {
		emptyChapter, // 0 baseParam
		emptyChapter, // 1 baseParam
		{ // 2 baseParam
			[]string{"xType", "yType"},
			[]Template{
				{"xType__yType__FlawedFunc1", []string{}},               // 0 funcParams
				{"xType__yType__FlawedFunc1__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("FlawedFunc1From_%s_to_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
			},
		},
	},
	"Giver1": {
		emptyChapter, //0 baseParam
		{ //1 baseParam
			[]string{"sType"},
			[]Template{
				{"sType__Giver", []string{}},               // 0 funcParams
				{"sType__Giver__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("Giver_%s", baseTypeArguments[0].GetTypeName())
			},
		},
	},
	"Stream": {
		{ //0 baseParam
			[]string{},
			[]Template{
				{"Stream", []string{}},               // 0 funcParams
				{"Stream__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "Stream"
			},
		},
		{ //1 baseParam
			[]string{"eType"},
			[]Template{
				{"eType__Stream", []string{}},                               // 0 funcParams
				{"eType__Stream__aType", []string{"aType"}},                 // 1 funcParams
				{"eType__Stream__aType__bType", []string{"aType", "bType"}}, // 2 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("Stream_%s", baseTypeArguments[0].GetTypeName())
			},
		},
	},
	"Validation": {
		{ //0 baseParams
			[]string{},
			[]Template{
				{"Validation", []string{}}, // 0 funcParams
				{},                         // 1 funcParams
				{"Validation__sType__kType", []string{"sType", "kType"}},                                                                                 // 2 funcParams
				{"Validation__sType__kType__aType", []string{"sType", "kType", "aType"}},                                                                 // 3 funcParams
				{"Validation__sType__kType__aType__bType", []string{"sType", "kType", "aType", "bType"}},                                                 // 4 funcParams
				{"Validation__sType__kType__aType__bType__cType", []string{"sType", "kType", "aType", "bType", "cType"}},                                 // 5 funcParams
				{"Validation__sType__kType__aType__bType__cType__dType", []string{"sType", "kType", "aType", "bType", "cType", "dType"}},                 // 6 funcParams
				{"Validation__sType__kType__aType__bType__cType__dType__eType", []string{"sType", "kType", "aType", "bType", "cType", "dType", "eType"}}, // 7 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "Validation"
			},
		},
		emptyChapter, // 1 baseParams
		{ //2 baseParams
			[]string{"sType", "kType"},
			[]Template{
				{"sType__kType__Validation", []string{}},               // 0 funcParams
				{"sType__kType__Validation__aType", []string{"aType"}}, // 1 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("Validation_%s_by_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
			},
		},
	},
	"ValiResu": {
		{ //0 baseParams
			[]string{},
			[]Template{
				{"ValiResu", []string{}}, // 0 funcParams
				{},
				{"ValiResu__sType__kType", []string{"sType", "kType"}}, // 2 funcParams
				{},
				{"ValiResu__sType__kType__aType__bType", []string{"sType", "kType", "aType", "bType"}},                                                 // 4 funcParams
				{"ValiResu__sType__kType__aType__bType__cType", []string{"sType", "kType", "aType", "bType", "cType"}},                                 // 5 funcParams
				{"ValiResu__sType__kType__aType__bType__cType__dType", []string{"sType", "kType", "aType", "bType", "cType", "dType"}},                 // 6 funcParams
				{"ValiResu__sType__kType__aType__bType__cType__dType__eType", []string{"sType", "kType", "aType", "bType", "cType", "dType", "eType"}}, // 7 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return "ValiResu"
			},
		},
		emptyChapter, //1 baseParams
		{ //2 baseParams
			[]string{"sType", "kType"},
			[]Template{
				{"sType__kType__ValiResu", []string{}},
				{"sType__kType__ValiResu__aType", []string{"aType"}},
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("ValiResu_%s_by_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
			},
		},
	},
}

func BuildFungPackage(config Config) {
	fungPackageBuilder := PackageBuilder{
		TypeConstructorsMap: fungTypeConstructors,
		TemplatesFS:         fungTemplatesFS,
	}
	fungPackageBuilder.Build(config)
}
