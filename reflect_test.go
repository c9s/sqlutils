package sqlutils
import "testing"
import "sort"

type fooRecord struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Internal int    `json:-`
}

func TestColumnNameMap(t *testing.T) {
	columns := GetColumnMap( &fooRecord{ Id: 3, Name: "Mary" } )
	t.Log(columns)

	if len(columns) == 0 {
		t.Fail()
	}
}


func TestColumnNamesParsing(t * testing.T) {
	columns := ParseColumnNames( &fooRecord{Id:3, Name: "Mary"} )

	// sort.Strings(columns)
	t.Log(columns)

	i := sort.SearchStrings(columns, "Internal")
	if columns[i] == "Internal" {
		t.Fail()
	}

	if len(columns) != 3 {
		t.Fail()
	}

	columns = ParseColumnNames( &fooRecord{Id:4, Name: "John"} )
	t.Log(columns)
	if len(columns) != 3 {
		t.Fail()
	}
}

