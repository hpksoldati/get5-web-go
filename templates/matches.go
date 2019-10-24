// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/matches.gohtml

package templates

import (
	db "github.com/FlowingSPDG/get5-web-go/src/db"
	"github.com/FlowingSPDG/get5-web-go/templates/layout"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Matches generates templates/matches.gohtml
func Matches(u *db.MatchesPageData) string {
	var _b strings.Builder
	RenderMatches(&_b, u)
	return _b.String()
}

// RenderMatches render templates/matches.gohtml
func RenderMatches(_buffer io.StringWriter, u *db.MatchesPageData) {

	_body := func(_buffer io.StringWriter) {

	}

	_menu := func(_buffer io.StringWriter) {
		if u.LoggedIn == true {

			_buffer.WriteString("<li><a id=\"mymatches\" href=\"/mymatches\">My Matches</a></li>")

			_buffer.WriteString("<li><a id=\"match_create\" href=\"/match/create\">Create a Match</a></li>")

			_buffer.WriteString("<li><a id=\"myteams\" href=\"/myteams\">My Teams</a></li>")

			_buffer.WriteString("<li><a id=\"team_create\" href=\"/team/create\">Create a Team</a></li>")

			_buffer.WriteString("<li><a id=\"myservers\" href=\"/myservers\">My Servers</a></li>")

			_buffer.WriteString("<li><a id=\"server_create\" href=\"/server/create\">Add a Server</a></li>")

			_buffer.WriteString("<li><a href=\"/logout\">Logout</a></li>")

		} else {

			_buffer.WriteString("<li><a href=\"/login\"><img src=\"/static/img/login_small.png\" height=\"18\"/></a></li>")

		}
	}

	_content := func(_buffer io.StringWriter) {

		_buffer.WriteString("<div id=\"content\">\n    ")
		if u.MyMatches == true {

			_buffer.WriteString("<h1>Your matches</h1>")

		} else if u.AllMatches == true {

			_buffer.WriteString("<h1>All matches</h1>")

		} else {

			_buffer.WriteString("<h1>Matches for <a href=\"/user/")
			_buffer.WriteString(gorazor.HTMLEscInt(u.Owner.ID))
			_buffer.WriteString("\"> ")
			_buffer.WriteString(gorazor.HTMLEscStr(u.Owner.Name))
			_buffer.WriteString("</a></h1>")

		}
		_buffer.WriteString("\n\n  <table class=\"table table-striped\">\n    <thead>\n      <tr>\n        <th>Match ID</th>\n        <th>Team 1</th>\n        <th>Team 2</th>\n        <th>Status</th>\n      </tr>\n    </thead>\n    <tbody>\n\t")
		n := u.Matches
		_buffer.WriteString("\n\t")
		for i := 0; i < len(n); i++ {
			m := n[i]
			id := n[i].ID
			team1, _ := m.GetTeam1()
			team2, _ := m.GetTeam2()
			status, _ := m.GetStatusString(true)

			_buffer.WriteString("<tr>\n        <td><a href=\"/match/")
			_buffer.WriteString(gorazor.HTMLEscape(id))
			_buffer.WriteString("\"> ")
			_buffer.WriteString(gorazor.HTMLEscape(id))
			_buffer.WriteString(" </a></td>\n\n        <td>\n          ")
			_buffer.WriteString((team1.GetFlagHTML(0.75)))
			_buffer.WriteString("\n          <a href=\"/team/")
			_buffer.WriteString(gorazor.HTMLEscInt(team1.ID))
			_buffer.WriteString("\">")
			_buffer.WriteString(gorazor.HTMLEscStr(team1.Name))
			_buffer.WriteString("</a>\n        </td>\n\n        <td>\n        ")
			_buffer.WriteString((team2.GetFlagHTML(0.75)))
			_buffer.WriteString("\n          <a href=\"/team/")
			_buffer.WriteString(gorazor.HTMLEscInt(team2.ID))
			_buffer.WriteString("\">")
			_buffer.WriteString(gorazor.HTMLEscStr(team2.Name))
			_buffer.WriteString("</a>\n        </td>\n\n        <td>\n          <p>")
			_buffer.WriteString(gorazor.HTMLEscStr(status))
			_buffer.WriteString("</p>\n        </td>\n      </tr>")

		}
		_buffer.WriteString("\n\n    </tbody>\n  </table>\n\n</div>")

	}

	layout.RenderBase(_buffer, _body, _menu, _content)
}
