package sqlutils
import "database/sql"
import "fmt"

const (
	DriverPg = 1
	DriverMysql = 2
	DriverSqlite = 3
)

// id, err := sqlutils.Create(struct pointer)
func Create(db *sql.DB, val interface{}, driver int) (*Result) {
	sql , args := BuildInsertClause(val)

	// for pgsql only
	if driver == DriverPg {
		sql += " RETURNING id"
	}

	fmt.Println(db)

	err := CheckRequired(val)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	rows, err := PrepareAndQuery(db,sql,args...)
	if err != nil {
		return NewErrorResult(err,sql)
	}

	id, err := GetReturningIdFromRows(rows)

	if err != nil {
		return NewErrorResult(err,sql)
	}

	r := NewResult(sql)
	r.Id = id
	return r
}


