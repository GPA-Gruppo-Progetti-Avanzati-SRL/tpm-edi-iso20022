package generate

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type importResolver struct {
	imports []importRef
}

type importRef struct {
	path  string
	alias string
}

func (ir *importResolver) Resolve(pkg string) *importRef {
	foundIndex := -1
	for ix, i := range ir.imports {
		if i.alias == pkg {
			return &i
		} else {
			n := strings.LastIndex(i.path, "/")
			if n >= 0 && i.path[n+1:] == pkg {
				foundIndex = ix
			}
		}
	}

	if foundIndex >= 0 {
		return &ir.imports[foundIndex]
	}

	return nil
}

func ParseImports(f *ast.File) ([]importRef, error) {

	imports := make([]importRef, 0)
	for _, i := range f.Imports {
		ir := importRef{path: strings.ReplaceAll(i.Path.Value, "\"", "")}
		if i.Name != nil {
			ir.alias = i.Name.Name
		}
		imports = append(imports, ir)
	}
	return imports, nil
}

type FieldInfo struct {
	Pos     token.Pos
	End     token.Pos
	FilePos token.Position
	Name    string
	Pointer bool
	Slice   bool
	X       string
	Sel     string
	QType   string
}

func (fi *FieldInfo) String() string {
	var stb strings.Builder

	stb.WriteString(fi.Name)

	if fi.Slice {
		stb.WriteString(" []")
	} else {
		stb.WriteString("   ")
	}

	if fi.Pointer {
		stb.WriteString(" *")
	} else {
		stb.WriteString(" ")
	}

	addDot := false
	if fi.X != "" {
		stb.WriteString(fi.X)
		addDot = true
	}

	if fi.Sel != "" {
		if addDot {
			stb.WriteRune('.')
		}

		stb.WriteString(fi.Sel)
	}

	return stb.String()
}

func parseField(fs *token.FileSet, field *ast.Field, resolver importResolver) FieldInfo {

	fi := FieldInfo{}
	if len(field.Names) > 0 {
		fi.Name = field.Names[0].Name
	}

	fi.FilePos = fs.Position(field.Pos())
	fi.Pos = field.Pos()
	fi.End = field.End()
	parseFieldKind(&fi, field.Type)
	/*
		switch tt := field.Type.(type) {
		case *ast.StarExpr:
			fi.Pointer = true
			switch ttt := tt.X.(type) {
			case *ast.SelectorExpr:
				fi.Sel = ttt.Sel.Name
				fi.X = ttt.X.(*ast.Ident).Name
			case *ast.Ident:
				fi.X = ttt.Name
			}
		case *ast.SelectorExpr:
			fi.Sel = tt.Sel.Name
			fi.X = tt.X.(*ast.Ident).Name
		case *ast.Ident:
			fi.X = tt.Name
		case *ast.ArrayType:
			fmt.Sprintf("element of array?  %#v", tt.Elt)
		case *ast.FuncType:
			panic(fmt.Sprintf("func type paremeter not supported %#v", field.Type))
		default:
			panic(fmt.Sprintf("unrecognized fieldInfo %#v", field.Type))
		}
	*/

	if fi.Sel != "" {
		i := resolver.Resolve(fi.X)
		if i == nil {
			panic(fmt.Sprintf("Cannot find import for %s", fi.X))
		}

		fi.QType = i.path + "/" + fi.Sel
	} else {
		fi.QType = fi.X
	}

	return fi
}

func parseFieldKind(fi *FieldInfo, fieldType ast.Expr) {
	switch tt := fieldType.(type) {
	case *ast.StarExpr:
		fi.Pointer = true
		parseFieldKind(fi, tt.X)
		/*
			switch ttt := tt.X.(type) {
			case *ast.SelectorExpr:
				fi.Sel = ttt.Sel.Name
				fi.X = ttt.X.(*ast.Ident).Name
			case *ast.Ident:
				fi.X = ttt.Name
			}
		*/
	case *ast.ArrayType:
		fi.Slice = true
		parseFieldKind(fi, tt.Elt)

	case *ast.SelectorExpr:
		fi.Sel = tt.Sel.Name
		fi.X = tt.X.(*ast.Ident).Name
	case *ast.Ident:
		fi.X = tt.Name

	case *ast.FuncType:
		panic(fmt.Sprintf("func type paremeter not supported %#v", fieldType))
	default:
		panic(fmt.Sprintf("unrecognized fieldInfo %#v", fieldType))
	}

}

func Scan(fn string) error {

	var fs *token.FileSet
	fs = token.NewFileSet()

	log.Info().Str("filename", fn).Msg("start processing file")

	var f *ast.File
	var err error

	f, err = parser.ParseFile(fs, fn, nil, 0)
	if err != nil {
		return err
	}

	imports, err := ParseImports(f)
	if err != nil {
		return err
	}

	importresolver := importResolver{imports: imports}

	for _, node := range f.Decls {
		switch node.(type) {

		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)

					fmt.Printf("Struct: name=%s\n", typeSpec.Name.Name)

					switch typeSpec.Type.(type) {
					case *ast.StructType:
						structType := typeSpec.Type.(*ast.StructType)
						for _, field := range structType.Fields.List {
							fi := parseField(fs, field, importresolver)
							fmt.Printf("field: %s - %s\n", fi.Name, fi.String())
							/*							i := field.Type.(*ast.Ident)
														fieldType := i.Name

														for _, name := range field.Names {
															fmt.Printf("\tField: name=%s type=%s\n", name.Name, fieldType)
														}
							*/
						}

					}
				}
			}
		}
	}

	// ast.Print(fs, f)

	return nil
}
