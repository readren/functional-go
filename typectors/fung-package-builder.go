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
				{"Recover__aType", []string{"aType"}}, // 0 funcParams
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
				{},                                   // 0 funcParams
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
		emptyChapter, //0 baseParam
		emptyChapter, //1 baseParam
		{ //2 baseParam
			[]string{"xType", "yType"},
			[]Template{
				{"xType__yType__Func1", []string{}},               // 0 funcParams
				{"xType__yType__Func1__aType", []string{"aType"}}, // 0 funcParams
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("FuncFrom_%s_to_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
			},
		},
	},
	"Giver1": {
		emptyChapter, //0 baseParam
		{ //1 baseParam
			[]string{"sType"},
			[]Template{
				{"sType__Giver", []string{}},
				{"sType__Giver__aType", []string{"aType"}},
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
				{},
				{"Stream__aType", []string{"aType"}},
			},
			func(baseTypeArguments TypeArguments) string {
				return "Stream"
			},
		},
		{ //1 baseParam
			[]string{"eType"},
			[]Template{
				{"eType__Stream", []string{}},
				{"eType__Stream__aType", []string{"aType"}},
				{"eType__Stream__aType__bType", []string{"aType", "bType"}},
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("Stream_%s", baseTypeArguments[0].GetTypeName())
			},
		},
	},
	"Validate": {
		{ //0 baseParams
			[]string{},
			[]Template{
				{},
				{"Validate__aType__bType", []string{"aType", "bType"}},
				{"Validate__aType__bType__cType", []string{"aType", "bType", "cType"}},
				{"Validate__aType__bType__cType__dType", []string{"aType", "bType", "cType", "dType"}},
				{"Validate__aType__bType__cType__dType__eType", []string{"aType", "bType", "cType", "dType", "eType"}},
			},
			func(baseTypeArguments TypeArguments) string {
				return "Validate"
			},
		},
		emptyChapter, // 1 baseParams
		{ //2 baseParams
			[]string{"sType", "kType"},
			[]Template{
				{"sType__kType__Validate", []string{}},
				{"sType__kType__Validate__aType", []string{"aType"}},
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("Validate_%s_idx_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
			},
		},
	},
	"ValiResu": {
		{ //0 baseParams
			[]string{},
			[]Template{
				{},
				{"ValiResu__aType__bType", []string{"aType", "bType"}},
				{"ValiResu__aType__bType__cType", []string{"aType", "bType", "cType"}},
				{"ValiResu__aType__bType__cType__dType", []string{"aType", "bType", "cType", "dType"}},
				{"ValiResu__aType__bType__cType__dType__eType", []string{"aType", "bType", "cType", "dType", "eType"}},
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
				return fmt.Sprintf("ValiResu_%s_idx_%s", baseTypeArguments[0].GetTypeName(), baseTypeArguments[1].GetTypeName())
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
