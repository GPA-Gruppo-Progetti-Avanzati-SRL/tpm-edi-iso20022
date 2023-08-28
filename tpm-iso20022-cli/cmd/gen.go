/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/tpm-iso20022-cli/cmd/generate/golang"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/tpm-iso20022-cli/cmd/generate/golang/model"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/tpm-iso20022-cli/cmd/generate/parser"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/tpm-iso20022-cli/cmd/generate/registry"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	schemaFile        string
	schemaFolder      string
	outFolder         string
	basePackageImport string
	xsdtPackageImport string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generation command for iso-20022 xsd",
	Long: `The command allows to generate code artifacts from iso-20022 xsd files. The generation
accept a single xsd file or a bunch of schemas. The difference is in the fact that the code in tle latter case 
can group common types in a single shared folder of generation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if schemaFile == "" && schemaFolder == "" {
			fmt.Println("Error: at least on of xsdf or xsdd flags has to be set (but not both)")
		}

		schemas, err := findFilesToProcess(schemaFile, schemaFolder)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		messageFolder, err := validateOutputFolder(outFolder)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		fmt.Println("schema files to be worked upon")
		for i, sch := range schemas {
			fmt.Printf("[%d] %s\n", i, sch)
		}

		err = generate(schemas, messageFolder)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
	},
}

func validateOutputFolder(f string) (string, error) {
	messageFolder, ok := util.ResolvePath(f)
	if !ok {
		err := fmt.Errorf("cannot resolve output folder %s", f)
		return "", err
	}

	if fi, ok := util.FileInfo(messageFolder); !ok {
		err := fmt.Errorf("output folder %s doesn't exists", f)
		return "", err
	} else if !fi.IsDir() {
		err := fmt.Errorf("output folder %s is not a folder", f)
		return "", err
	}

	return messageFolder, nil
}

func generate(schemas []string, fld string) error {
	const semLogContext = "gen-cmd::generate"
	var msgs []registry.ISO20022Message
	p := parser.NewParser(parser.Config{Registry: registry.Config{SimpleTypesInCommonPackage: true}})
	for _, xsdFileName := range schemas {
		msgName := strings.TrimSuffix(filepath.Base(xsdFileName), ".xsd")
		log.Trace().Str("msg-name", msgName).Msg(semLogContext)

		b, err := os.ReadFile(xsdFileName)
		if err != nil {
			return err
		}

		msg := registry.ISO20022Message{Name: msgName}
		msgs = append(msgs, msg)

		err = p.Parse(msg, b)
		if err != nil {
			return err
		}
	}

	modelCfg := model.ModelConfig{
		BasePackageImport: basePackageImport,
		XsDtPackageImport: xsdtPackageImport,
	}

	if modelCfg.XsDtPackageImport == "" {
		modelCfg.XsDtPackageImport = modelCfg.BasePackageImport + "/xsdt"
	}

	gm, err := model.NewModel(&modelCfg, msgs, p.TypeRegistry)
	if err != nil {
		return err
	}

	cfg := golang.Config{
		OutFolder:  fld, // filepath.Join(fld, "messages"),
		FormatCode: true,
	}
	err = golang.Generate(&cfg, &gm)
	if err != nil {
		return err
	}

	return nil
}

func findFilesToProcess(files, folder string) ([]string, error) {
	var schemas []string
	var err error
	var ok bool

	if schemaFile != "" {
		schemas = strings.Split(schemaFile, ",")
		for i, f := range schemas {
			schemas[i], ok = util.ResolvePath(f)
			if !ok {
				err = fmt.Errorf("cannot resolve schema file %s", f)
				return nil, err
			}

			if fi, ok := util.FileInfo(schemas[i]); ok {
				if fi.IsDir() {
					err = fmt.Errorf("supposed schema file %s is a directory", schemas[i])
					return nil, err
				}
			} else {
				err = fmt.Errorf("cannot find schema file %s", schemas[i])
				return nil, err
			}
		}
	} else {
		folder, ok = util.ResolvePath(folder)
		if !ok {
			err = fmt.Errorf("cannot resolve schema folder %s", folder)
			return nil, err
		}

		if fi, ok := util.FileInfo(folder); ok {
			if !fi.IsDir() {
				err = fmt.Errorf("specified schema folder %s is not a folder", folder)
				return nil, err
			}
		} else {
			err = fmt.Errorf("cannot find schema folder %s", folder)
			return nil, err
		}

		schemas, err = util.FindFiles(folder, util.WithFindFileType(util.FileTypeFile))
		if err != nil {
			return nil, err
		}
	}

	if len(schemas) == 0 {
		err = fmt.Errorf("no schema files found to process")
		return nil, err
	}

	return schemas, err
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	genCmd.Flags().StringVarP(&schemaFile, "xsdfiles", "f", "", "xsd schema files to process (csv list)")
	genCmd.Flags().StringVarP(&schemaFolder, "xsdfolder", "d", "", "xsd schema folder to process")
	genCmd.MarkFlagsMutuallyExclusive("xsdfiles", "xsdfolder")

	genCmd.Flags().StringVarP(&outFolder, "outFolder", "o", "", "output folder")
	_ = genCmd.MarkFlagRequired("outFolder")

	genCmd.Flags().StringVarP(&basePackageImport, "basePkg", "", model.DefaultModelCfg.BasePackageImport, "base package import")
	genCmd.Flags().StringVarP(&xsdtPackageImport, "xsdtPkg", "", model.DefaultModelCfg.XsDtPackageImport, "xsdt package import (default basePkg + /xsdt)")
}
