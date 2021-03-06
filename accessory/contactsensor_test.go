package accessory

import (
	"testing"
)

func TestContactSensor(t *testing.T) {
	info := Info{
		Name:             "Sensor1",
		SerialNumber:     "009",
		Manufacturer:     "Achme",
		Model:            "Accessory",
		FirmwareRevision: "1.0.0",
	}

	cs := NewContactSensor(info)

	// fetch services
	serv := cs.GetServices()

	// assert two services
	if has, should := len(serv), 2; should != has {
		t.Fatalf("contact sensor services has %v we expected %v", has, should)
	}

	if has, should := serv[0].GetType(), "3E"; should != has {
		t.Fatalf("Type expected '%s' has %v", should, has)
	}

	if has, should := serv[1].GetType(), "80"; should != has {
		t.Fatalf("Type expected '%s' has %v", should, has)
	}

	stics := serv[0].GetCharacteristics()
	if has, should := len(stics), 8; should != has {
		t.Fatalf("service[0] has %v charateristic(s) we expected %v", has, should)
	}

	stics = serv[1].GetCharacteristics()
	if has, should := len(stics), 6; should != has {
		t.Fatalf("service[1] has %v charateristic(s) we expected %v", has, should)
	}
}
