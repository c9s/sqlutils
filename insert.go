package sqlutils
// import "fmt"

import "reflect"
import "strings"
import "strconv"
import "database/sql"

func BuildInsertClause(val interface{}) (string, []interface{}) {
	t := reflect.ValueOf(val).Elem()
	typeOfT := t.Type()
	tableName := GetTableName(val)

	var columnNames []string
	var valueFields []string
	var values      []interface{}

	var fieldId int = 1

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}


		var attributes = GetColumnAttributesFromTag(&tag)

		// if it's a serial column (with auto-increment, we can simply skip)
		if _, ok := attributes["serial"] ; ok {
			continue
		}

		// TODO: see if we can skip null columns, or simply skip Id column

		columnNames = append(columnNames, *columnName)
		valueFields = append(valueFields, "$" + strconv.Itoa(fieldId) )
		values      = append(values, field.Interface() )
		fieldId++
	}
	return "INSERT INTO " + tableName + " ( " + strings.Join(columnNames,", ") + " ) " +
		" VALUES ( " + strings.Join(valueFields,", ") + " )", values
}

func GetReturningIdFromRows(rows * sql.Rows) (int64, error) {
    var id int64
    var err error
    rows.Next()
    err = rows.Scan(&id)
	if err != nil {
		return -1, err
	}
    return id, err
}


