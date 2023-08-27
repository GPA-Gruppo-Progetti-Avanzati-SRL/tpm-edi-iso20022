package generate_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/cmds/tpm-iso20022-cli/generate"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScan(t *testing.T) {

	fn := "../../../messages/pain_013_01_07/pain_013_01_07.go"
	err := generate.Scan(fn)
	require.NoError(t, err)

}
