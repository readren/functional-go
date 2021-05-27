package typectors

import (
	"embed"
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

// A `TypeConstructor` is a builder of types of the same form. Examples of `TypeConstructor` are `Option`, `List`, `Stream`, `Func1`, `Func2`.
// The definitions that comprises a `TypeConstructor` are grouped by both, the number of base type parameters, and the number of func type parameteres.
type TypeConstructor []Chapter

// A `Chapter` consists of all the funcs of a `TypeConstructor` that have the same number of base type parameters.
// The funcs contained by a chapter are grouped by the number of func type arguments. Each of these groups is named a "template"
// The name "chapter" comes from the analogy with "book", where the entire book is a type constructor, and the pages are the templates.
type Chapter struct {
	BaseTypeParameters []string
	// Every chapter is split in many templates, each of them containing all the funcs that have the same number of type parameters. This field contains said templates indexed by the number of type parameters.
	// The purpose of this sepparation is to avoid the need to parse the go source files.
	Templates []Template
	// Not being used
	TypeNameBuilder func(baseTypeArguments TypeArguments) string
}

// Given we use a slice to contain the chapters in order to index them by the number of base type arguments; some elements of the slice may be empty.
var emptyChapter = Chapter{[]string{}, []Template{}, nil}

type PackageBuilder struct {
	TypeConstructorsMap map[string]TypeConstructor
	TemplatesFS         embed.FS
}

//// TypeArgument ////

type TypeArgument struct {
	// the actual type of this type argument expressed in the same way as for variable declarations. For example "int", or "[]image.Point"
	Type string
	// the package where the type in the `Type` field is defined. This field is optional when the `Type` field has a basic native type like "int", but not "[]int" nor "image.Point".
	PackagePath string
	// the alias of the package. This field is considered only when the `PackagePath` field is defined. Its default value is the last segment of the package path.
	PackageAlias string
	// the name to associate to the type in the `Type` field. There should be a one to one relationshipt between types and type names. This field is optional. When omited the name is generated automatically.
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

// An list of `TypeArgument`s.
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

// Specifies a `Template` instantiation.
// The which `Template` is pointed is specified by the type constructor name and the length of the base and func type arguments lists.
// The with which actual type arguments is the template instantiatted is specified by the content of said type argument's lists.
type TemplateArguments struct {
	TypeConstructorName string        `json:"typeCtor"`
	BaseTypeArguments   TypeArguments `json:"baseTArgs"`
	FuncTypeArguments   TypeArguments `json:"funcTArgs"`
}

func (thisPtr *TemplateArguments) IsEqual(otherPtr *TemplateArguments) bool {
	return thisPtr.TypeConstructorName == otherPtr.TypeConstructorName &&
		thisPtr.BaseTypeArguments.IsEqual(otherPtr.BaseTypeArguments) &&
		thisPtr.FuncTypeArguments.IsEqual(otherPtr.FuncTypeArguments)
}

////

// Knows the information that the `incarnateType` function needs to incarnate a single type (or a singleton companion object when the `BaseTypeArguments` list is empty).
type TypeDescriptor struct {
	// The name that identifies the `TypeConstructor`. For example: "Stream", or "Func1".
	TypeConstructorName string
	// The list of actual base type arguments.
	// The length of this list must be either zero or match the number of type parameters of the type constructor. When the length is zero this `TypeDescriptor` does not describe a type but a companion singleton object with static functions.
	// For example, to construct a type based on the `Stream` type constructor, the length of this field should be one. And for the case of the `Func1` type constructor, the length should be two.
	BaseTypeArguments TypeArguments
	// Specifies which `func`s of the type constructor are included in the generated type. The inclusion criteria is not by `func` name but by the number and type of `func`s type argiments.
	// This field specifies the set of `func` type arguments lists for which are `func`s are included in the definition of the generated type. For example, given a type constructor with one unimorphic method "uni" and two polymofphic methods "foo<A>" and "bar<A,B>", where "foo" has one type parameter `A`, and "bar" has two type parameters `B` and `C`; if this field value were [ [{"int"}], [{"Point", "image"}], [{"bool"},{"string"}] ]; then the unimorphic method "uni" wouln't be included (because the set does not contain an empty arguments list), two versions of the "foo" method would be included, the first with type argument "int" ("foo__int(..)") and the second with type argument "image.Point" ("foo__imagePoint(..)"); and one version of the "bar" method would be included with the type arguments "bool" and "string" ("bar__bool__string(..)").
	FuncTypeArgumentsForWhichFuncsAreIncluded []TypeArguments
}

type Config struct {
	GeneratedPackageParentDir string
	GeneratedPackageName      string
	// path to the directory that is the direct parent of all the type construtors directories, relative to the root of the `embed.FS` specified in `PackageBuilder.TemplatesFS`.
	// Each type constructor directory contained there should have the same name than the corresponding key in the `PackageBuilder.TypeConstructorsMap`, and value of `TypeDescriptor.TypeConstructorName`.
	TemplatesBaseDir string
	// Specifies which types to instantiate, based on the provided type constructors.
	TypesDescriptors []TypeDescriptor
}

func (packageBuilder PackageBuilder) Build(config Config) {
	workingDir, err := os.Getwd()
	checkError(err, "unable to get the working directory")

	temporaryDirectoryParent := filepath.Join(workingDir, "temp")
	tempDir, err := ioutil.TempDir(temporaryDirectoryParent, "instantiation-*")
	checkError(err, fmt.Sprintf("unable to create a temporary directory inside the directory \"%s\"", temporaryDirectoryParent))
	defer func() { os.RemoveAll(tempDir) }()
	fmt.Printf("temporary working directory: %s\n", tempDir)

	// comonSrcFile := filepath.Join(config.TemplatesFolder, "common.go")
	// checkError(copyFile(comonSrcFile, filepath.Join(tempDir, "common.go")), fmt.Sprintf(`unable to copy the "%s" file to the temporary directory`, comonSrcFile))

	var manager = manager{
		packageBuilder,
		config,
		tempDir,
		map[string]TypeArgument{},
		setOfTemplateArgs{},
		setOfTemplateArgs{},
		false,
		make(map[string]setOfTemplateArgs),
	}
	manager.groupAllTypeArgumentsByType()

	// Instantiate all the templates specified in the `config`
	for _, td := range config.TypesDescriptors {
		manager.incarnateType(td)
	}
	// Instantiate the templates pointed by all the "#dependsOn" directives contained in the instantianted templates, that aren't already instantiated.
	for {
		missingDependencies := manager.requestedDependencies.diff(manager.instantiatedDependencies)
		if len(missingDependencies) == 0 {
			break
		}
		manager.requestedDependencies = make(setOfTemplateArgs, 0)
		for _, md := range missingDependencies {
			methodTypeArguments := append(make([]TypeArguments, 0, 1), md.FuncTypeArguments)
			td := TypeDescriptor{md.TypeConstructorName, md.BaseTypeArguments, methodTypeArguments}
			manager.incarnateType(td)
		}
		manager.funcsWithNoInternalDependantsAreExcluded = true
	}

	// Move the generated source files from temporary directory to the one specified in the `Config`.
	generatedPackageDir := filepath.Join(config.GeneratedPackageParentDir, config.GeneratedPackageName)
	checkError(copyDirectory(tempDir, generatedPackageDir), fmt.Sprintf(`unable to copy the generated files from the temporary directory to the destination "%s"`, generatedPackageDir))

	// print dependencies relationship report
	manager.printInternalDependencisReport()
}

// Contains information and state that is shared between the `PackageBuilder.Build` function and other functions it calls.
type manager struct {
	PackageBuilder           PackageBuilder
	config                   Config
	tempDir                  string
	allDistinctTypeArguments map[string]TypeArgument
	// the set where the `TemplateArguments` of all the already parsed `#dependsOn` directives are accumulated
	requestedDependencies setOfTemplateArgs
	// the set that memorizes which template instantiations where already done
	instantiatedDependencies                 setOfTemplateArgs
	funcsWithNoInternalDependantsAreExcluded bool
	dependenciesGroupedByDependent           map[string]setOfTemplateArgs
}

func (managerPtr *manager) printInternalDependencisReport() {
	for dependent, dependants := range managerPtr.dependenciesGroupedByDependent {
		fmt.Printf("\n\"%s\" depends on:\n", dependent)
		for _, dependant := range dependants {
			var sb strings.Builder
			for _, bta := range dependant.BaseTypeArguments {
				sb.WriteString(bta.GetTypeName())
				sb.WriteString("__")
			}
			sb.WriteString(dependant.TypeConstructorName)
			for _, fta := range dependant.FuncTypeArguments {
				sb.WriteString("__")
				sb.WriteString(fta.GetTypeName())
			}
			sb.WriteString(".go")
			fmt.Printf("\t%v\n", sb.String())
		}
	}
}

// Groups all the `TypeArgument` instances contained in all the instances of `TypeDescriptor`, discriminating by the `Type` field and reducing the other fields to the most complete occurrence.
// Panics if two occurrences with the same `Type` field have a difference in the other fields.
func (managerPtr *manager) groupAllTypeArgumentsByType() {
	for _, td := range managerPtr.config.TypesDescriptors {
		// indiscriminately collect the type arguments contained in the `TypeDescriptor` `td`
		typeArguments := make([]TypeArguments, 0, 1+len(td.FuncTypeArgumentsForWhichFuncsAreIncluded))
		typeArguments = append(typeArguments, td.FuncTypeArgumentsForWhichFuncsAreIncluded...)
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

// Generates the golang source files that compose a type (`Stream<int>`, or `Validate<image.Point, string>`, or ...); provided the type constructor (`Stream`, or `Validate`, or ...),  the base type arguments (`int` for stream, or `image.Point` and `string` for validate, or ...), and a set of func type arguments lists.
func (managerPtr *manager) incarnateType(td TypeDescriptor) {
	// Obtain the `TypeConstructor` specified in the received `TypeDescriptor`.
	typeConstructor := fungTypeConstructors[td.TypeConstructorName]
	// Pick the chapter with the number of base type parameters specified in the received `TypeDescriptor`.
	chapter := typeConstructor[len(td.BaseTypeArguments)]
	// Instantiate the templates of the chapter for each of the sets of func type arguments specified in the received `TypeDescriptor`
	for _, funcTypeArguments := range td.FuncTypeArgumentsForWhichFuncsAreIncluded {
		numberOfFuncTypeParameters := len(funcTypeArguments)
		// choose the polymorphic funcs's template appropiate for the number of type arguments
		template := chapter.Templates[numberOfFuncTypeParameters]
		// instantiate it
		template.instantiate(td.TypeConstructorName, chapter, td.BaseTypeArguments, funcTypeArguments, managerPtr)
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

	templateSrcFile := fmt.Sprintf("%s/%s/%s.go", managerPtr.config.TemplatesBaseDir, typeConstructorName, template.FileName)
	source, err := managerPtr.PackageBuilder.TemplatesFS.ReadFile(templateSrcFile)
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
	internalDependenciesSet := make(setOfTemplateArgs, 0)
	for _, match := range internalDependenciesMatchs {
		// fmt.Printf("#dependsOn match: %s\n", match[1]) // TODO remove this line
		var internalDependency TemplateArguments
		checkError(json.Unmarshal(match[1], &internalDependency), fmt.Sprintf("unable to parse the directive: #dependsOn %s", match[1]))
		managerPtr.normalizeTemplateArguments(&internalDependency)
		managerPtr.requestedDependencies.add(&internalDependency)
		// memorize the dependency relationship for the final report
		internalDependenciesSet.add(&internalDependency)
	}
	managerPtr.dependenciesGroupedByDependent[codeFile.fileName] = internalDependenciesSet

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

func (managerPtr *manager) normalizeTemplateArguments(taPtr *TemplateArguments) {
	for btaIndex := range taPtr.BaseTypeArguments {
		taPtr.BaseTypeArguments[btaIndex] = managerPtr.registerAndNormalizeTypeArgument(taPtr.BaseTypeArguments[btaIndex])
	}
	for mtaIndex := range taPtr.FuncTypeArguments {
		taPtr.FuncTypeArguments[mtaIndex] = managerPtr.registerAndNormalizeTypeArgument(taPtr.FuncTypeArguments[mtaIndex])
	}
}

func checkError(err error, msg string) {
	if err != nil {
		panic(fmt.Errorf("%s : %w", msg, err))
	}
}
