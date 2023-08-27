/*
 * ./tpm-iso20022-cli gen --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/001.001.03" --xsdfiles "~/iso-20022/schemas/pain.001.001.03.xsd" -o "~/iso-20022/messages/pain/001.001.03"
 * ./tpm-iso20022-cli gen --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/002.001.03" --xsdfiles "~/iso-20022/schemas/pain.002.001.03.xsd" -o "~/iso-20022/messages/pain/002.001.03"
 * ./tpm-iso20022-cli gen --basePkg "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain/008.001.02" --xsdfiles "~/iso-20022/schemas/pain.008.001.02.xsd" -o "~/iso-20022/messages/pain/008.001.02"
 */
package main

import (
	_ "embed"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/tpm-iso20022-cli/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

//go:embed sha.txt
var sha string

//go:embed VERSION
var version string

// appLogo contains the ASCII splash screen
//
//go:embed app-logo.txt
var appLogo []byte

func main() {

	fmt.Println(string(appLogo))
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Sha: %s\n", sha)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cmd.Execute()
}
