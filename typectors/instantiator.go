package typectors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type Template struct {
	FileName                         string
	PolymorphicMethodsTypeParameters []string
}

type TypeConstructor struct {
	BaseTypeParameters []string
	// Every type constructor definition is split in many templates each of them containing all the methods that have the same number of type parameters. This field contains said templates indexed by the number of type parameters.
	// The purpose of this sepparation is to avoid the need to parse the go source files.
	Templates       []Template
	TypeNameBuilder func(baseTypeArguments []TypeArgument) string
}

var knowTypeConstructors map[string]TypeConstructor = map[string]TypeConstructor{
	"stream": {
		[]string{"eType"},
		[]Template{
			{"eType__stream", []string{}},
			{"eType__stream__aType", []string{"aType"}},
			{"eType__stream__aType__bType", []string{"aType", "bType"}},
		},
		func(baseTypeArguments []TypeArgument) string {
			elemsTypeName := baseTypeArguments[0].GetTypeName()
			return fmt.Sprintf("Stream_%s", elemsTypeName)
		},
	},
}

//// TypeArgument ////

type TypeArgument struct {
	// the actual type of this type argument expressed in the same way as for variable declarations. For example "int", or "[]image.Point"
	Type string
	// the name to associate to the type in the `Type` field. There should be a one to one relationshipt between types and type names. This field is optional when the `Type` field has a basic native type like "int", but not "[]int" nor "image.Point".
	TypeName string
	// the package where the type in the `Type` field is defined. This field is optional when the `Type` field has a basic native type like "int", but not "[]int" nor "image.Point".
	PackagePath string
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
	MethodTypeArguments TypeArguments `json:"methodTArgs"`
}

func (thisPtr *TemplateArguments) IsEqual(otherPtr *TemplateArguments) bool {
	return thisPtr.TypeConstructorName == otherPtr.TypeConstructorName &&
		thisPtr.BaseTypeArguments.IsEqual(otherPtr.BaseTypeArguments) &&
		thisPtr.MethodTypeArguments.IsEqual(otherPtr.MethodTypeArguments)
}

////

// Knows the arguments needed to incarnate a type.
type TypeIncarnationArguments struct {
	TypeConstructorName string
	BaseTypeArguments   []TypeArgument
	// Specifies for which type arguments are each polymorphic method instantiated. For example, if the template has a polymofphic method "foo" with one type parameter and a method "bar" with two type parameters; and this field value is [ [{"int"}], [{"Point", "image"}], [{"bool"},{"string"}] ]; then the "foo" method would be instanciated two times with type arguments "foo__int(..)" and "foo__Point(..)", and the "bar" method would be instantiated one time with type arguments "bar__bool__string(..)".
	TypeArgumentsForWhichPolymorphicMethodsAreInstantiated [][]TypeArgument
}

type Config struct {
	GeneratedPackageParentDir   string
	GeneratedPackageName        string
	TemplatesFolder             string
	TypeInstantiationsArguments []TypeIncarnationArguments
}

type manager struct {
	config  Config
	tempDir string
	// the set where the `TemplateArguments` of all the already parsed `#dependsOn` directives are accumulated
	requestedDependencies setOfTemplateArgs
	// the set that memorizes which template instantiations where already done
	instantiatedDependencies setOfTemplateArgs
}

func GeneratePackage(config Config) {
	workingDir, err := os.Getwd()
	checkError(err, "unable to get the working directory")

	temporaryDirectoryParent := fmt.Sprintf("%s/temp", workingDir)
	tempDir, err := ioutil.TempDir(temporaryDirectoryParent, "instantiation-*")
	checkError(err, fmt.Sprintf("unable to create a temporary directory inside the directory \"%s\"", temporaryDirectoryParent))
	defer func() { os.Remove(tempDir) }()
	fmt.Printf("temporary working directory: %s\n", tempDir)

	var manager = manager{config, tempDir, setOfTemplateArgs{}, setOfTemplateArgs{}}
	// Instantiate all the templates specified in the `config`
	for _, tia := range config.TypeInstantiationsArguments {
		manager.incarnateType(tia)
	}
	// Instantiate the templates pointed by all the "#dependsOn" directives contained in the already instantianted templates, that aren't already instantiated.
	for {
		missingDependencies := manager.requestedDependencies.diff(manager.instantiatedDependencies)
		if len(missingDependencies) == 0 {
			break
		}
		manager.requestedDependencies = make(setOfTemplateArgs, 0)
		for _, mr := range missingDependencies {
			methodTypeArguments := append(make([][]TypeArgument, 0, 1), mr.MethodTypeArguments)
			tia := TypeIncarnationArguments{mr.TypeConstructorName, mr.BaseTypeArguments, methodTypeArguments}
			manager.incarnateType(tia)
		}
	}

	// Move the generated source files from temp directory to the generated package directory
	generatedPackageDir := fmt.Sprintf("%s/%s", config.GeneratedPackageParentDir, config.GeneratedPackageName)
	existentSrcFiles, err := ioutil.ReadDir(generatedPackageDir)
	checkError(err, fmt.Sprintf("unable to read the files inside the \"%s\" directory", generatedPackageDir))

	fmt.Print("Existent source files: ")
	for _, e := range existentSrcFiles {
		fmt.Printf("%s - ", e.Name())
	}
	fmt.Println()
}

func (managerPtr *manager) incarnateType(tia TypeIncarnationArguments) {
	typeConstructor := knowTypeConstructors[tia.TypeConstructorName]

	// instantiate the unimorphic methods
	typeConstructor.Templates[0].instantiate(tia.TypeConstructorName, typeConstructor, tia.BaseTypeArguments, nil, managerPtr)

	// instantiate the polymorphic methods's template for each of the specified sets of method type arguments
	for _, methodTypeArguments := range tia.TypeArgumentsForWhichPolymorphicMethodsAreInstantiated {
		numberOfMethodsTypeParameters := len(methodTypeArguments)
		// choose the polymorphic methods's template appropiate for the number of type arguments
		polymorphicMethodsTemplate := typeConstructor.Templates[numberOfMethodsTypeParameters]
		// instantiate it
		polymorphicMethodsTemplate.instantiate(tia.TypeConstructorName, typeConstructor, tia.BaseTypeArguments, methodTypeArguments, managerPtr)
	}
}

type codeFile struct {
	fileName string
	content  []byte
}

// Used to obtain the json string after the "#dependesOn" directives
var dependsOnDirectiveRegex = regexp.MustCompile(`(?m)#dependsOn\s*(.+)$`)

// Generates a source file based on this template with the specified type arguments
func (template *Template) instantiate(typeConstructorName string, typeConstructor TypeConstructor, baseTypeArguments []TypeArgument, methodTypeArguments []TypeArgument, managerPtr *manager) {

	templateSrcFile := fmt.Sprintf("%s/%s/%s.go", managerPtr.config.TemplatesFolder, typeConstructorName, template.FileName)
	source, err := ioutil.ReadFile(templateSrcFile)
	checkError(err, fmt.Sprintf("unable to load the template source file %s", templateSrcFile))
	codeFile := codeFile{template.FileName, source}

	dependencies := make(map[string]bool)
	// replace base type parameters with the actual type arguments
	for typeParameterIndex, typeArgument := range baseTypeArguments {
		if len(typeArgument.PackagePath) > 0 {
			dependencies[typeArgument.PackagePath] = true
		}
		typeParameterName := typeConstructor.BaseTypeParameters[typeParameterIndex]
		codeFile.replaceTypeParameterWithTypeArgument(typeParameterName, typeArgument)
	}
	// replace polymorphic methods type parameters with the actual type arguments
	for typeParameterIndex, typeArgument := range methodTypeArguments {
		if len(typeArgument.PackagePath) > 0 {
			dependencies[typeArgument.PackagePath] = true
		}
		typeParameterName := template.PolymorphicMethodsTypeParameters[typeParameterIndex]
		codeFile.replaceTypeParameterWithTypeArgument(typeParameterName, typeArgument)
	}

	// Add this template instantiation to the set of template instantiations that are already done
	managerPtr.instantiatedDependencies.add(&TemplateArguments{typeConstructorName, baseTypeArguments, methodTypeArguments})

	// Obtain the dependencies on template instantiations required by this template. Note that given the type parameters were already replaced by the actual type arguments, the parsed `#dependsOn` directives contain actual types.
	dependenciesMatchs := dependsOnDirectiveRegex.FindAllSubmatch(codeFile.content, -1)
	for _, rm := range dependenciesMatchs {
		fmt.Printf("#dependsOn matchs: %s\n", rm[1]) // TODO remove this line
		var dependency TemplateArguments
		checkError(json.Unmarshal(rm[1], &dependency), fmt.Sprintf("unable to parse the directive: #dependsOn %s", rm[1]))
		managerPtr.requestedDependencies.add(&dependency)
	}

	// Remove the excluded section. This should be done after the dependencies obtention because the excluded section may contain `#dependsOn` directives.
	codeFile.content = regexp.MustCompile(`(?m)^.*#exclude-section-begin(.|\n|\r)*#exclude-section-end.*\n`).ReplaceAll(codeFile.content, []byte{})

	// insert an import clause with the dependencies
	// TODO

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