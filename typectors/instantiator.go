package typectors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// A Tempalte is a source file that contains all the funcs of a `TypeConstructor` that have the same number of both, base type parameters, and func type parameters.
type Template struct {
	FileName           string
	FuncTypeParameters []string
}

// A `TypeConstructor` is a builder of types of the same form. Examples of `TypeConstructor` are `Option`, `List`, `Stream`, `Function1`, `Function2`.
// The definitions that comprises a `TypeConstructor` are grouped by both, the number base type parameters (chapters), and the number of func type parameteres (templates).
type TypeConstructor []Chapter

// A Chapter consists of all the funcs of a `TypeConstructor` that have the same number of base type parameters.
type Chapter struct {
	BaseTypeParameters []string
	// Every chapter is split in many templates, each of them containing all the funcs that have the same number of type parameters. This field contains said templates indexed by the number of type parameters.
	// The purpose of this sepparation is to avoid the need to parse the go source files.
	Templates       []Template
	TypeNameBuilder func(baseTypeArguments TypeArguments) string
}

var emptyChapter = Chapter{[]string{}, []Template{}, nil}

var knowTypeConstructors map[string]TypeConstructor = map[string]TypeConstructor{
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
		emptyChapter, // 0 baseParams
		emptyChapter, // 1 baseParams
		{ //2 baseParams
			[]string{"sType", "kType"},
			[]Template{
				{"sType__kType__Validate", []string{}},
				{"sType__kType__Validate__aType", []string{"aType"}},
			},
			func(baseTypeArguments TypeArguments) string {
				return fmt.Sprintf("Validate_%s", baseTypeArguments[0].GetTypeName())
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
				return fmt.Sprintf("ValiResu_%s", baseTypeArguments[0].GetTypeName())
			},
		},
	},
}

//// TypeArgument ////

type TypeArgument struct {
	// the actual type of this type argument expressed in the same way as for variable declarations. For example "int", or "[]image.Point"
	Type string
	// the package where the type in the `Type` field is defined. This field is optional when the `Type` field has a basic native type like "int", but not "[]int" nor "image.Point".
	PackagePath string
	// the alias of the package. This field is considered only when the `PackagePath` field is defined, and its default value is the last segment of the package path.
	PackageAlias string
	// the name to associate to the type in the `Type` field. There should be a one to one relationshipt between types and type names. This field is optional when the `Type` field has a basic native type like "int", but not "[]int" nor "image.Point".
	TypeName string
}

func (thisPtr *TypeArgument) IsEqual(otherPtr *TypeArgument) bool {
	return thisPtr.Type == otherPtr.Type
}

var nativeBasicTypeRegex = regexp.MustCompile(`^\w+$`)
var validTypeNameRegex = regexp.MustCompile(`^[a-zA-Z]\w*$`)

func (ta *TypeArgument) GetTypeName() string {
	if ta.TypeName != "" {
		if !validTypeNameRegex.Match([]byte(ta.TypeName)) {
			panic(fmt.Errorf("invalid type name \"%s\" : the offender `TypeArgument` is %v", ta.TypeName, ta))
		}
		return ta.TypeName
	} else if nativeBasicTypeRegex.MatchString(ta.Type) {
		return ta.Type
	} else {
		return BuildTypeName(ta.Type)
		// panic(fmt.Errorf("the `TypeArgument.TypeName` field is required when the `TypeArgument.Type` field is not a basic native type : the offender `TypeArgument` is %v", ta))
	}
}

//// TypeArguments ////

type TypeArguments []TypeArgument

func (tas1 TypeArguments) IsEqual(tas2 TypeArguments) bool {
	if len(tas1) != len(tas2) {
		return false
	}
	for i, ta := range tas1 {
		if !ta.IsEqual(&(tas2[i])) {
			return false
		}
	}
	return true
}

//// TemplateArguments ////

type TemplateArguments struct {
	TypeConstructorName string        `json:"typeCtor"`
	BaseTypeArguments   TypeArguments `json:"baseTArgs"`
	MethodTypeArguments TypeArguments `json:"funcTArgs"`
}

func (thisPtr *TemplateArguments) IsEqual(otherPtr *TemplateArguments) bool {
	return thisPtr.TypeConstructorName == otherPtr.TypeConstructorName &&
		thisPtr.BaseTypeArguments.IsEqual(otherPtr.BaseTypeArguments) &&
		thisPtr.MethodTypeArguments.IsEqual(otherPtr.MethodTypeArguments)
}

////

// Knows the arguments needed to incarnate a type.
type TypeDescriptor struct {
	TypeConstructorName string
	BaseTypeArguments   TypeArguments
	// Specifies for which set of type arguments are templates instantiated. For example, if the type constructor method templates had the polymofphic the methods "foo", with one type parameter, and "bar", with two type parameters; and this field value were [ [{"int"}], [{"Point", "image"}], [{"bool"},{"string"}] ]; then the "foo" method would be instanciated two times, the first with type argument "int" ("foo__int(..)") and the second with type argument "image.Point" ("foo__imagePoint(..)"); and the "bar" method would be instantiated one time with the type arguments "bool" and "string" ("bar__bool__string(..)").
	FuncTypeArgumentsForWhichTemplatesAreInstantiated []TypeArguments
}

type Config struct {
	GeneratedPackageParentDir string
	GeneratedPackageName      string
	TemplatesFolder           string
	TypesDescriptors          []TypeDescriptor
}

type manager struct {
	config                   Config
	tempDir                  string
	allDistinctTypeArguments map[string]TypeArgument
	// the set where the `TemplateArguments` of all the already parsed `#dependsOn` directives are accumulated
	requestedDependencies setOfTemplateArgs
	// the set that memorizes which template instantiations where already done
	instantiatedDependencies                 setOfTemplateArgs
	funcsWithNoInternalDependantsAreExcluded bool
}

func GeneratePackage(config Config) {
	workingDir, err := os.Getwd()
	checkError(err, "unable to get the working directory")

	temporaryDirectoryParent := filepath.Join(workingDir, "temp")
	tempDir, err := ioutil.TempDir(temporaryDirectoryParent, "instantiation-*")
	checkError(err, fmt.Sprintf("unable to create a temporary directory inside the directory \"%s\"", temporaryDirectoryParent))
	defer func() { os.RemoveAll(tempDir) }()
	fmt.Printf("temporary working directory: %s\n", tempDir)

	// comonSrcFile := filepath.Join(config.TemplatesFolder, "common.go")
	// checkError(copyFile(comonSrcFile, filepath.Join(tempDir, "common.go")), fmt.Sprintf(`unable to copy the "%s" file to the temporary directory`, comonSrcFile))

	var manager = manager{config, tempDir, map[string]TypeArgument{}, setOfTemplateArgs{}, setOfTemplateArgs{}, false}
	manager.groupAllTypeArgumentsByType()

	// Instantiate all the templates specified in the `config`
	for _, td := range config.TypesDescriptors {
		manager.incarnateType(td)
	}
	manager.funcsWithNoInternalDependantsAreExcluded = true
	// Instantiate the templates pointed by all the "#dependsOn" directives contained in the instantianted templates, that aren't already instantiated.
	for {
		missingDependencies := manager.requestedDependencies.diff(manager.instantiatedDependencies)
		if len(missingDependencies) == 0 {
			break
		}
		manager.requestedDependencies = make(setOfTemplateArgs, 0)
		for _, md := range missingDependencies {
			methodTypeArguments := append(make([]TypeArguments, 0, 1), md.MethodTypeArguments)
			td := TypeDescriptor{md.TypeConstructorName, md.BaseTypeArguments, methodTypeArguments}
			manager.incarnateType(td)
		}
	}

	// Move the generated source files from temporary directory to the one specified in the `Config`.
	generatedPackageDir := filepath.Join(config.GeneratedPackageParentDir, config.GeneratedPackageName)
	checkError(copyDirectory(tempDir, generatedPackageDir), fmt.Sprintf(`unable to copy the generated files from the temporary directory to the destination "%s"`, generatedPackageDir))
}

// Groups all the `TypeArgument` instances contained in all the instances of `TypeDescriptor`, discriminating by the `Type` field and reducing the other fields to the most complete occurrence.
// Panics if two occurrences with the same `Type` field have a difference in the other fields.
func (managerPtr *manager) groupAllTypeArgumentsByType() {
	for _, td := range managerPtr.config.TypesDescriptors {
		// indiscriminately collect the type arguments contained in the `TypeDescriptor` `td`
		typeArguments := make([]TypeArguments, 0, 1+len(td.FuncTypeArgumentsForWhichTemplatesAreInstantiated))
		typeArguments = append(typeArguments, td.FuncTypeArgumentsForWhichTemplatesAreInstantiated...)
		typeArguments = append(typeArguments, td.BaseTypeArguments)
		// for each collected type argument:
		for _, tas := range typeArguments {
			for _, ita := range tas {
				// group it with the others discriminating by tye `Type` field and reducing the other fields to the most complete version.
				managerPtr.registerAndNormalizeTypeArgument(ita)
			}
		}
	}
}

func (managerPtr *manager) registerAndNormalizeTypeArgument(ta TypeArgument) TypeArgument {
	// group it with the others discriminating by tye `Type` field and reducing the other fields to the most complete version.
	mostCompleteTa, found := managerPtr.allDistinctTypeArguments[ta.Type]
	if found {
		if mostCompleteTa.PackagePath != "" && ta.PackagePath != "" && mostCompleteTa.PackagePath != ta.PackagePath ||
			mostCompleteTa.PackageAlias != "" && ta.PackageAlias != "" && mostCompleteTa.PackageAlias != ta.PackageAlias ||
			mostCompleteTa.TypeName != "" && ta.TypeName != "" && mostCompleteTa.TypeName != ta.TypeName {
			panic(fmt.Errorf(`inconsistent type attributes between two type arguments: "%v" and "%v"`, mostCompleteTa, ta))
		}
		if mostCompleteTa.PackagePath == "" && ta.PackagePath != "" {
			mostCompleteTa.PackagePath = ta.PackagePath
		}
		if mostCompleteTa.PackageAlias == "" && ta.PackageAlias != "" {
			mostCompleteTa.PackageAlias = ta.PackageAlias
		}
		if mostCompleteTa.TypeName == "" && ta.TypeName != "" {
			mostCompleteTa.TypeName = ta.TypeName
		}
		managerPtr.allDistinctTypeArguments[ta.Type] = mostCompleteTa
		return mostCompleteTa
	} else {
		managerPtr.allDistinctTypeArguments[ta.Type] = ta
		return ta
	}
}

func (managerPtr *manager) normalizeTemplateArguments(taPtr *TemplateArguments) {
	for btaIndex := range taPtr.BaseTypeArguments {
		taPtr.BaseTypeArguments[btaIndex] = managerPtr.registerAndNormalizeTypeArgument(taPtr.BaseTypeArguments[btaIndex])
	}
	for mtaIndex := range taPtr.MethodTypeArguments {
		taPtr.MethodTypeArguments[mtaIndex] = managerPtr.registerAndNormalizeTypeArgument(taPtr.MethodTypeArguments[mtaIndex])
	}
}

func (managerPtr *manager) incarnateType(td TypeDescriptor) {
	typeConstructor := knowTypeConstructors[td.TypeConstructorName]
	chapter := typeConstructor[len(td.BaseTypeArguments)]
	// // instantiate the unimorphic methods
	// chapter.Templates[0].instantiate(tia.TypeConstructorName, typeConstructor, tia.BaseTypeArguments, nil, managerPtr)

	// instantiate the polymorphic methods's template for each of the specified sets of method type arguments
	for _, funcTypeArguments := range td.FuncTypeArgumentsForWhichTemplatesAreInstantiated {
		numberOfFuncTypeParameters := len(funcTypeArguments)
		// choose the polymorphic funcs's template appropiate for the number of type arguments
		polymorphicFuncsTemplate := chapter.Templates[numberOfFuncTypeParameters]
		// instantiate it
		polymorphicFuncsTemplate.instantiate(td.TypeConstructorName, chapter, td.BaseTypeArguments, funcTypeArguments, managerPtr)
	}
}

type codeFile struct {
	fileName string
	content  []byte
}

type externalDependency struct {
	Path  string
	Alias string
}

// Used to obtain the json string after the "#dependesOn" directives
var dependsOnDirectiveRegex = regexp.MustCompile(`(?m)#dependsOn\s*(.+)$`)
var usesExternalPackageRegex = regexp.MustCompile(`(?m)#usesExternalPackage\s*(.+)$`)
var excludeSectionRegex = regexp.MustCompile(`(?m)^.*#excludeSectionBegin\s(.|\n|\r)*#excludeSectionEnd\s.*\n`)
var importAnchorRegex = regexp.MustCompile(`(?m)^.*#importAnchor\s.*$`)
var startOfFuncsWithNoInternalDependantsRegex = regexp.MustCompile(`(?m)^.*#startOfFuncsWithNoInternalDependants(.|\s)*`)

// Generates a source file based on this template with the specified type arguments
func (template *Template) instantiate(typeConstructorName string, chapter Chapter, baseTypeArguments TypeArguments, methodTypeArguments TypeArguments, managerPtr *manager) {

	templateSrcFile := fmt.Sprintf("%s/%s.go", managerPtr.config.TemplatesFolder, template.FileName)
	source, err := ioutil.ReadFile(templateSrcFile)
	checkError(err, fmt.Sprintf("unable to load the template source file %s", templateSrcFile))
	codeFile := codeFile{template.FileName, source}

	if managerPtr.funcsWithNoInternalDependantsAreExcluded {
		// Remove the section after the `#startOfFuncsWithNoInternalDependants` directive.
		codeFile.content = startOfFuncsWithNoInternalDependantsRegex.ReplaceAll(codeFile.content, []byte{})
	}

	// set where all the external dependencies required by this template instantiation are collected
	externalDependencies := make(map[string]string)

	// Obtain the external dependencies pointed by the `#usesExternalPackage` directives.
	externalDependenciesMatchs := usesExternalPackageRegex.FindAllSubmatch(codeFile.content, -1)
	for _, match := range externalDependenciesMatchs {
		// fmt.Printf("#usesExternalPackage match: %s\n", match[1]) // TODO remove this line
		var ed = externalDependency{}
		checkError(json.Unmarshal(match[1], &ed), fmt.Sprintf("unable to parse the directive: #usesExternalPackage %s", match[1]))
		externalDependencies[ed.Path] = ed.Alias
	}

	// replace base type parameters with the actual type arguments
	for typeParameterIndex, typeArgument := range baseTypeArguments {
		if len(typeArgument.PackagePath) > 0 {
			externalDependencies[typeArgument.PackagePath] = typeArgument.PackageAlias
		}
		typeParameterName := chapter.BaseTypeParameters[typeParameterIndex]
		codeFile.replaceTypeParameterWithTypeArgument(typeParameterName, typeArgument)
	}
	// replace polymorphic methods type parameters with the actual type arguments
	for typeParameterIndex, typeArgument := range methodTypeArguments {
		if len(typeArgument.PackagePath) > 0 {
			externalDependencies[typeArgument.PackagePath] = typeArgument.PackageAlias
		}
		typeParameterName := template.FuncTypeParameters[typeParameterIndex]
		codeFile.replaceTypeParameterWithTypeArgument(typeParameterName, typeArgument)
	}

	// Add this template instantiation to the set of template instantiations that are already done
	managerPtr.instantiatedDependencies.add(&TemplateArguments{typeConstructorName, baseTypeArguments, methodTypeArguments})

	// Collect the dependencies on template instantiations required by this template. Note that given the type parameters were already replaced by the actual type arguments, the parsed `#dependsOn` directives contain actual types.
	internalDependenciesMatchs := dependsOnDirectiveRegex.FindAllSubmatch(codeFile.content, -1)
	for _, match := range internalDependenciesMatchs {
		// fmt.Printf("#dependsOn match: %s\n", match[1]) // TODO remove this line
		var internalDependency TemplateArguments
		checkError(json.Unmarshal(match[1], &internalDependency), fmt.Sprintf("unable to parse the directive: #dependsOn %s", match[1]))
		managerPtr.normalizeTemplateArguments(&internalDependency)
		managerPtr.requestedDependencies.add(&internalDependency)
	}

	// Remove the excluded section. This should be done after the dependencies obtention because the excluded section may contain `#dependsOn` directives.
	codeFile.content = excludeSectionRegex.ReplaceAll(codeFile.content, []byte{})

	// Insert an import clause with the extennal dependencies at the point where the `#importAnchor` directive is located.
	var sb = strings.Builder{}
	if len(externalDependencies) > 0 {
		sb.WriteString("import (\n")
		for ed, alias := range externalDependencies {
			sb.WriteRune('\t')
			if alias != "" {
				sb.WriteString(alias)
				sb.WriteRune('\t')
			}
			sb.WriteRune('"')
			sb.WriteString(ed)
			sb.WriteString("\"\n")
		}
		sb.WriteString(")")
	}
	codeFile.content = importAnchorRegex.ReplaceAllLiteral(codeFile.content, []byte(sb.String()))

	generatedSrcFile := fmt.Sprintf("%s/%s.go", managerPtr.tempDir, codeFile.fileName)
	ioutil.WriteFile(generatedSrcFile, codeFile.content, 0)
}

func (cfp *codeFile) replaceTypeParameterWithTypeArgument(typeParameter string, typeArgument TypeArgument) {
	// replace type declarations
	typeDeclarationRegex := regexp.MustCompile(fmt.Sprintf(`(\b|^)%s(\b|$)`, typeParameter))
	typeDeclarationReplacement := []byte(fmt.Sprintf(`${1}%s${2}`, typeArgument.Type))
	cfp.content = typeDeclarationRegex.ReplaceAll(cfp.content, typeDeclarationReplacement)

	// replace fragments of polymorphic methods identifiers
	identifierFragmentRegex := regexp.MustCompile(fmt.Sprintf(`(\b|_|^)%s(\b|_|$)`, typeParameter))
	identifierFragmentReplacement := fmt.Sprintf(`${1}%s${2}`, typeArgument.GetTypeName())
	cfp.content = identifierFragmentRegex.ReplaceAll(cfp.content, []byte(identifierFragmentReplacement))
	// replace fragments of the generated source file name
	cfp.fileName = identifierFragmentRegex.ReplaceAllString(cfp.fileName, identifierFragmentReplacement)
}

func checkError(err error, msg string) {
	if err != nil {
		panic(fmt.Errorf("%s : %w", msg, err))
	}
}
