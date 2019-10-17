package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"
	"html/template"
	_ "html/template"
	"net/http"
	_ "strconv"
	_ "time"
)

func MatchCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchCreateHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchConfigHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchCancelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchCancelHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchRconHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchRconHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchPauseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchPauseHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchUnpauseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchUnpauseHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchAdduserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchAdduserHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

/*
func MatchSendConfigHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchSendConfigHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}
*/

func MatchBackupHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchBackupHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesHandler\nvars : %v", vars)
	//w.WriteHeader(http.StatusOK)
	tpl := template.Must(template.ParseFiles("get5/templates/layout.html", "get5/templates/matches.html")) // template
	fmt.Printf("HomeHandler\nvars : %v\n", vars)

	name := ""
	userid := ""
	loggedin := false
	session, _ := SessionStore.Get(r, SessionData)

	m := &HomeData{
		LoggedIn: false,
		Content:  tpl,
		UserName: name,
		UserID:   userid,
	}

	if _, ok := session.Values["Name"]; ok {
		name, _ = session.Values["Name"].(string)
		loggedin = true
	}

	if _, ok := session.Values["UserID"]; ok {
		userid, _ = session.Values["UserID"].(string)
		loggedin = true
	}

	m.UserID = userid
	m.LoggedIn = loggedin

	fmt.Println(m)

	// テンプレートを描画
	tpl.Execute(w, m)
}

func MatchesWithIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MatchesWithIDHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MyMatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("MyMatchesHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}