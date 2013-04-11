Gatsby SQLUtils
=================

Gatsby SQLUtils package helps you build SQL.


Usage
-----

Import from GitHub:

```go
import "github.com/c9s/go-sqlutils" sqlutils
```

Define your struct with json spec or with field tag:

```go
type Staff struct {
	Id        int `json:"id"`
	Name      string `json:"name" field:",required"`
	Gender    string `json:"gender"`
	StaffType string `json:"staff_type"` // valid types: doctor, nurse, ...etc
	Phone     string `json:"phone"`
}
```

To build select clause depends on the struct fields:

```go
sql := sqlutils.BuildSelectClause(Staff{})
// sql = " SELECT id, name, gender, staff_type, phone FROM staffs"
```

To build where clause from map:

```
sql, args := sqlutils.BuildWhereClauseWithAndOp(map[string]interface{} {
    "name": "John"
})
// sql = " WHERE name = $1 AND id = $2"
```

To create new record:

```go
staff := Staff{Name:"Mary"}
Create(db,&staff)
```

To update struct object:

```go
staff.Name = "NewName"
rows, err := Update(db,&staff)
```




