package functionaltests

import ports_v1 "ports_microservice/external/proto/ports"

func generateCreatePortReq() *ports_v1.CreatePortRequest {
	return &ports_v1.CreatePortRequest{
		Name:        "Great Port",
		City:        "St. Petersburg",
		Country:     "Russia",
		Alias:       []string{"alias", "aliases"},
		Regions:     []string{"Len", "Oblast"},
		Coordinates: []float32{23, 123},
		Province:    "Len",
		Timezone:    "Moscow",
	}
}
