package shortid

import "testing"

func TestID(t *testing.T) {
	for i := 0; i < 50; i++ {
		id, err := ID(192)
		t.Logf("Generated ID: %v", id)

		if err != nil {
			t.Fatalf("Error while generating unique ID: %v", err)
		}

		if len(id) != 32 {
			t.Fatalf("Expected a string of length 32 but got %v", len(id))
		}
	}

	id, err := ID(7)
	if id != "" || err == nil {
		t.Fatal("Should not be able to generate ID for given bits that is not divisible by 8")
	}
}

func TestUUID(t *testing.T) {
	for i := 0; i < 50; i++ {
		uuid, err := UUID()
		t.Logf("Generated UUID: %v", uuid)

		if err != nil {
			t.Fatalf("Error while generating unique ID: %v", err)
		}

		if len(uuid) != 22 {
			t.Fatalf("Expected a string of length 22 but got %v", len(uuid))
		}
	}
}
