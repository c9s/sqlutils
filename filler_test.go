package sqlutils
import "testing"
import "database/sql"
import _ "github.com/bmizerany/pq"

func openDB() (*sql.DB, error) {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=test sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}


func TestFillRecord(t * testing.T) {
	staff := Staff{}
    db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	// Create Staff
	staff.Id = 1
	staff.Name = "Mary"
	staff.Phone = "1234567"

	r := Create(db,&staff, DriverPg)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	if r.Id == -1 {
		t.Fatal("Primary key failed")
	}
	staff.Id = r.Id


	sql := BuildSelectClause(&staff) + " WHERE id = $1"
	if sql != "SELECT id,name,gender,staff_type,phone FROM staffs WHERE id = $1" {
		t.Fatal(sql)
	}

	stmt , err := db.Prepare(sql)
	rows, err := stmt.Query( r.Id)
	rows.Next()
	err = FillFromRow(&staff,rows)
	if err != nil {
		t.Fatal(err)
	}

	r = Delete(db,&staff)
	if r.Error != nil {
		t.Fatal(r.Error)
	}
}


