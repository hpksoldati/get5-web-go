package main

import (
	"./get5"
	_ "fmt"
	"github.com/go-ini/ini"
	_ "github.com/gorilla/mux"
	_ "log"
	_ "net/http"
)

type Config struct {
	SteamAPIKey string
	DefaultPage string
	SQLHost     string
	SQLUser     string
	SQLPass     string
	SQLPort     int
	SQLDBName   string
	HOST        string
}

var (
	STATIC_DIR = "./static"
	HOST       = "localhost:8081"
	Cnf        Config
	SQLHost    string
	SQLUser    string
	SQLPass    string
	SQLPort    int
	SQLDBName  string
)

func init() {
	c, _ := ini.Load("config.ini")
	Cnf = Config{
		HOST:      c.Section("GET5").Key("HOST").MustString(""),
		SQLHost:   c.Section("sql").Key("host").MustString(""),
		SQLUser:   c.Section("sql").Key("user").MustString(""),
		SQLPass:   c.Section("sql").Key("pass").MustString(""),
		SQLPort:   c.Section("sql").Key("port").MustInt(3306),
		SQLDBName: c.Section("sql").Key("database").MustString(""),
	}
	HOST = Cnf.HOST
	SQLHost = Cnf.SQLHost
	SQLUser = Cnf.SQLUser
	SQLPass = Cnf.SQLPass
	SQLPort = Cnf.SQLPort
	SQLDBName = Cnf.SQLDBName
}

func main() {
	/*
		r := mux.NewRouter()
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

		//s := r.Host("get5.flowing.tokyo").Subrouter()
		r.HandleFunc("/", get5.HomeHandler).Methods("GET")
		r.HandleFunc("/login", get5.LoginHandler).Methods("GET")
		r.HandleFunc("/logout", get5.LogoutHandler).Methods("GET")

		r.HandleFunc("/match/create", get5.MatchCreateHandler)             // GET/POST
		r.HandleFunc("/match/{matchID}", get5.MatchHandler)                // ?
		r.HandleFunc("/match/{matchID}/config", get5.MatchConfigHandler)   // ?
		r.HandleFunc("/match/{matchID}/cancel", get5.MatchCancelHandler)   // ?
		r.HandleFunc("/match/{matchID}/rcon", get5.MatchRconHandler)       // ?
		r.HandleFunc("/match/{matchID}/pause", get5.MatchPauseHandler)     // ?
		r.HandleFunc("/match/{matchID}/unpause", get5.MatchUnpauseHandler) // ?
		r.HandleFunc("/match/{matchID}/adduser", get5.MatchAdduserHandler) // ?
		//r.HandleFunc("/match/{matchID}/sendconfig", get5.MatchSendConfigHandler) // ?
		r.HandleFunc("/match/{matchID}/backup", get5.MatchBackupHandler).Methods("GET") // GET

		r.HandleFunc("/match/{matchID}/finish", get5.MatchFinishHandler).Methods("POST")                                             // POST
		r.HandleFunc("/match/{matchID}/map/{mapNumber}/start", get5.MatchMapStartHandler).Methods("POST")                            // POST
		r.HandleFunc("/match/{matchID}/map/{mapNumber}/update", get5.MatchMapUpdateHandler).Methods("POST")                          // POST
		r.HandleFunc("/match/{matchID}/map/{mapNumber}/finish", get5.MatchMapFinishHandler).Methods("POST")                          // POST
		r.HandleFunc("/match/{matchID}/map/{mapNumber}/player/{steamid64}/update", get5.MatchMapPlayerUpdateHandler).Methods("POST") // POST

		r.HandleFunc("/matches", get5.MatchesHandler)                // ?
		r.HandleFunc("/matches/{userID}", get5.MatchesWithIDHandler) // ?
		r.HandleFunc("/mymatches", get5.MyMatchesHandler)            // ?

		r.HandleFunc("/team/create", get5.TeamCreateHandler)              // GET/POST
		r.HandleFunc("/team/{teamID}", get5.TeamHandler).Methods("GET")   // GET
		r.HandleFunc("/team/{teamID}/edit", get5.TeamEditHandler)         // GET/POST
		r.HandleFunc("/team/{teamID}/delete", get5.TeamDeleteHandler)     // ?
		r.HandleFunc("/teams/{userID}", get5.TeamsHandler).Methods("GET") // GET
		r.HandleFunc("/myteams", get5.MyTeamsHandler).Methods("GET")      // GET

		r.HandleFunc("/server/create", get5.ServerCreateHandler)                           // GET/POST
		r.HandleFunc("/server/{serverid}/edit", get5.ServerEditHandler)                    // GET/POST
		r.HandleFunc("/server/{serverid}/delete", get5.ServerDeleteHandler).Methods("GET") // GET
		r.HandleFunc("/myservers", get5.MyServersHandler)                                  // ?

		r.HandleFunc("/user/{userID}", get5.UserHandler)
		r.Methods("GET", "POST")
		http.Handle("/", r)
		fmt.Println("RUNNING")
		log.Fatal(http.ListenAndServe(HOST, nil))
	*/
	get5.MySQLGetGameServerData(SQLHost, SQLUser, SQLPass, SQLDBName, SQLPort)
}