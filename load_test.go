package sqlutils
import "testing"

func TestLoad(t *testing.T) {
	db, err := openDB()
	if err != nil {
		t.Fatal(err)
	}

	staff := Staff{Name: "John", Gender: "m", Phone: "1234567"}

	r := Create(db,&staff)
	if r.Error != nil {
		t.Fatal(r.Error)
	}
	if r.Id == -1 {
		t.Fatal("Can not create record")
	}
	t.Logf("staff id: %d", r.Id)

	staff2 := Staff{}
	r = Load(db,&staff2, r.Id)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

	if staff2.Id == 0 {
		t.Fatal("Can not load record")
	}

	r = Delete(db,&staff)
	if r.Error != nil {
		t.Fatal(r.Error)
	}

}
