package sqlutils
// import "fmt"

import "reflect"
import "strings"
import "strconv"
import "github.com/c9s/inflect"
import "database/sql"





func BuildInsertColumnClause(val interface{}) (string, []interface{}) {
	t := reflect.ValueOf(val)
	typeOfT := t.Type()
	tableName := inflect.Tableize(typeOfT.Name())

	var columnNames []string
	var valueFields []string
	var values      []interface{}

	for i := 0; i < t.NumField(); i++ {
		var tag        reflect.StructTag = typeOfT.Field(i).Tag
		var field      reflect.Value = t.Field(i)
		// var fieldType  reflect.Type = field.Type()

		var columnName *string = GetColumnNameFromTag(&tag)
		if columnName == nil {
			continue
		}
		// fieldAttrs := GetColumnAttributesFromTag(&tag)
		columnNames = append(columnNames, *columnName)
		valueFields = append(valueFields, "$" + strconv.Itoa(i + 1) )
		values      = append(values, field.Interface() )
	}
	return "INSERT INTO " + tableName + " (" + strings.Join(columnNames,",") + ") " +
		" VALUES (" + strings.Join(valueFields,",") + ")", values
}


func GetReturningIdFromRows(rows * sql.Rows) (int, error) {
    var id int
    var err error
    rows.Next()
    err = rows.Scan(&id)
	if err != nil {
		return -1, err
	}
    return id, err
}

// id, err := sqlutils.Create(struct pointer)
func Create(db *sql.DB, val interface{}) (int,error) {
	sql , args := BuildInsertColumnClause(val)

	// for pgsql only
	sql += " RETURNING id"

	err := CheckRequired(val)
	if err != nil {
		return -1, err
	}

	rows, err := PrepareAndQuery(db,sql,args...)
	if err != nil {
		return -1, err
	}
	return GetReturningIdFromRows(rows)
}

func Delete(db *sql.DB, val interface{}) (sql.Result, error) {
	sql := "DELETE FROM " + GetTableName(val) + " WHERE id = $1"

	if val.(PrimaryKey) == nil {
		panic("PrimaryKey interface is required.")
	}


	id := val.(PrimaryKey).GetPkId()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

