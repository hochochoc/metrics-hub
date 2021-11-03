package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"metric-hub/config"
	"metric-hub/handler"
	"net/http"
	"time"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	pgClient, err := config.NewEntClient()
	if err != nil {
		log.Fatalf("Fail to initialize pg client: %s", err.Error())
	}
	defer pgClient.Close()
	config.SetClient(pgClient)

	r := mux.NewRouter()
	r.Use(commonMiddleware)
	r.HandleFunc("/process_report", handler.ProcessReportHandler).Methods("POST")
	r.HandleFunc("/process_statistics", handler.ProcessStatisticsHandler)
	r.HandleFunc("/process_outliers", handler.ProcessOutliersHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + viper.GetString("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + viper.GetString("PORT"))
	log.Fatal(srv.ListenAndServe())
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}