package api

import "github.com/carolineeey/multifinance-api/api"

func main() {
	api := api.NewClient(
		"0.0.0.0:8080",
	)
}
