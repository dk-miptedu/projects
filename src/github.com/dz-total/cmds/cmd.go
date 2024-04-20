package cmds

import (
	"FinalTaskAppGoBasic/configs"
	"FinalTaskAppGoBasic/dbs"
	"FinalTaskAppGoBasic/handlers"
	"FinalTaskAppGoBasic/logs"
	"FinalTaskAppGoBasic/prometeus"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Cmd() {
	fmt.Printf("initi logrus... ")
	logs.InitializeLogging("./configs")
	fmt.Println("logrus Init Done.")

	configApp, err := configs.LoadConfig("./configs/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	dbs.Connect(configApp.DataBase)
	dbs.Migrate()
	r := mux.NewRouter()
	r.Use(prometeus.PrometheusMiddleware)
	r.Path("/metrics").Handler(promhttp.Handler())
	r.HandleFunc("/transactions", handlers.HandleTransactions)
	r.HandleFunc("/transactions", handlers.HandleTransactions).Methods("GET", "POST")
	r.HandleFunc("/transactions/{id}", handlers.HandleTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", handlers.HandleTransactions).Methods("PUT, DELETE") // Для PUT и DELETE
	r.HandleFunc("/users", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	http.Handle("/", r)

	port := fmt.Sprintf("%d", configApp.Restapi.Port)
	if configApp.Restapi.Port == 0 {
		port = "8080"
	}

	logs.Log.Infof("======= Starting server on port %s =======", port)
	fmt.Println("Server is running on port " + port)
	logs.Log.Fatal(http.ListenAndServe(":"+port, nil))

}
