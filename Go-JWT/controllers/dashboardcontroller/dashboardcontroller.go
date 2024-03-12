package dashboardcontroller

import (
	"github.com/BerkatPS/go-jwt-testing/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// membuat data statis untuk ditampilkan di response json

	data := []map[string]interface{}{
		{
			"id":      1,
			"makanan": "Susu",
			"minuman": "SGM",
		},
		{
			"id":      2,
			"makanan": "Susu",
			"minuman": "SGM",
		}, {
			"id":      3,
			"makanan": "Susu",
			"minuman": "SGM",
		},
	}
	helper.Writejson(w, http.StatusOK, data)
}
