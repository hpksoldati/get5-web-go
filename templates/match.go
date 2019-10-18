// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/match.gohtml

package templates

import (
	"github.com/FlowingSPDG/get5-web-go/src/models"
	"github.com/FlowingSPDG/get5-web-go/templates/layout"
	"io"
	"strings"
)

// Match generates templates/match.gohtml
func Match(u *models.MatchesPageData) string {
	var _b strings.Builder
	RenderMatch(&_b, u)
	return _b.String()
}

// RenderMatch render templates/match.gohtml
func RenderMatch(_buffer io.StringWriter, u *models.MatchesPageData) {

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

		_buffer.WriteString("<p>Not impalamented yet!</p>")

	}

	layout.RenderBase(_buffer, _body, _menu, _content)
}