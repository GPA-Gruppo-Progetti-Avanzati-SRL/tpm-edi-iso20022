package reflectx

import (
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/dotnotation"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/common"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-edi-iso20022/iso-20022/messages/pain_013_001_07"
	"github.com/stretchr/testify/require"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestDocument(t *testing.T) {

	d := pain_013_001_07.Document{}
	v := reflect.ValueOf(&d)
	if v.Kind() != reflect.Ptr {
		t.Fatal("must pass a pointer, not a value, to StructScan destination")
	}
	if v.IsNil() {
		t.Fatal("nil pointer passed to StructScan destination")
	}

	// base := Deref(v.Type())

	m := mapper()

	var paths []dotnotation.DotPath

	p, _ := dotnotation.NewPath(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_MsgId)
	paths = append(paths, p)
	p, _ = dotnotation.NewPath(pain_013_001_07.Path_CdtrPmtActvtnReq_GrpHdr_InitgPty_Id_OrgId_Othr_Id)
	paths = append(paths, p)
	fields := m.TraversalsByName(v.Type(), paths)
	// if we are not unsafe and are missing fields, return an error

	values := make([]interface{}, len(paths))

	err := fieldsByTraversal(v, fields, values, true)
	require.NoError(t, err)

	convertAssignRows(values[0], common.Max35TextSample)
	convertAssignRows(values[1], common.Max35TextSample)
}

// isScannable takes the reflect.Type and the actual dest value and returns
// whether or not it's Scannable.  Something is scannable if:
//   - it is not a struct
//   - it implements sql.Scanner
//   - it has no exported fields
func isScannable(t reflect.Type) bool {
	if t.Kind() != reflect.Struct {
		return true
	}

	// it's not important that we use the right mapper for this particular object,
	// we're only concerned on how many exported fields this struct has
	return len(mapper().TypeMap(t).Index) == 0
}

// NameMapper is used to map column names to struct field names.  By default,
// it uses strings.ToLower to lowercase struct field names.  It can be set
// to whatever you want, but it is encouraged to be set before sqlx is used
// as name-to-field mappings are cached after first use on a type.
var NameMapper = NameMapperFunc
var origMapper = reflect.ValueOf(NameMapper)

func NameMapperFunc(s string) string {
	return strings.TrimPrefix(s, "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07")
}

// Rather than creating on init, this is created when necessary so that
// importers have time to customize the NameMapper.
var mpr *common.Mapper

// mprMu protects mpr.
var mprMu sync.Mutex

// mapper returns a valid mapper using the configured NameMapper func.
func mapper() *common.Mapper {
	mprMu.Lock()
	defer mprMu.Unlock()

	if mpr == nil {
		mpr = common.NewMapperFunc("xml", NameMapper)
	} else if origMapper != reflect.ValueOf(NameMapper) {
		// if NameMapper has changed, create a new mapper
		mpr = common.NewMapperFunc("xml", NameMapper)
		origMapper = reflect.ValueOf(NameMapper)
	}
	return mpr
}

// fieldsByName fills a values interface with fields from the passed value based
// on the traversals in int.  If ptrs is true, return addresses instead of values.
// We write this instead of using FieldsByName to save allocations and map lookups
// when iterating over many rows.  Empty traversals will get an interface pointer.
// Because of the necessity of requesting ptrs or values, it's considered a bit too
// specialized for inclusion in reflectx itself.
func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := common.FieldByIndexes(v, traversal)
		if ptrs {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
	return nil
}
