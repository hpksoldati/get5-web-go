// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/myservers.gohtml

package templates

import (
	db "github.com/FlowingSPDG/get5-web-go/src/db"
	"github.com/FlowingSPDG/get5-web-go/templates/layout"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Myservers generates templates/myservers.gohtml
func Myservers(u *db.MyserversPageData) string {
	var _b strings.Builder
	RenderMyservers(&_b, u)
	return _b.String()
}

// RenderMyservers render templates/myservers.gohtml
func RenderMyservers(_buffer io.StringWriter, u *db.MyserversPageData) {

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

		_buffer.WriteString("<div id=\"content\">\n\n<ul class=\"list-group\">\n\n    ")
		if len(u.Servers) == 0 {

			_buffer.WriteString("<li class=\"list-group-item\">No servers found.</li>")

		} else {

			_buffer.WriteString("<table class=\"table table-striped\">\n    <thead>\n      <tr>\n        <th>Server ID</th>\n        <th>Display Name</th>\n        <th>IP Address</th>\n        <th>Port</th>\n        <th>Status</th>\n        <th></th>\n      </tr>\n    </thead>\n    <tbody>\n\n    ")
			for i := 0; i < len(u.Servers); i++ {
				s := u.Servers[i]

				_buffer.WriteString("<tr>\n        <td>")
				_buffer.WriteString(gorazor.HTMLEscInt(s.Id))
				_buffer.WriteString("</td>\n        <td>")
				_buffer.WriteString(gorazor.HTMLEscStr(s.Display_name))
				_buffer.WriteString("</td>\n        <td>")
				_buffer.WriteString(gorazor.HTMLEscStr(s.Ip_string))
				_buffer.WriteString("</td>\n        <td>")
				_buffer.WriteString(gorazor.HTMLEscInt(s.Port))
				_buffer.WriteString("</td>\n            ")
				if s.In_use == true {

					_buffer.WriteString("<td>In use</td>")

				} else {

					_buffer.WriteString("<td>Free</td>")

				}
				_buffer.WriteString("\n        <td>\n          <a href=\"/server/")
				_buffer.WriteString(gorazor.HTMLEscInt(s.Id))
				_buffer.WriteString("/edit\" class=\"btn btn-primary btn-xs\">Edit</a>\n          ")
				if s.In_use == false {

					_buffer.WriteString("<a href=\"/server/")
					_buffer.WriteString(gorazor.HTMLEscInt(s.Id))
					_buffer.WriteString("/delete\" class=\"btn btn-danger btn-xs\">Delete</a>")

				}
				_buffer.WriteString("\n        </td>\n      </tr>")

			}
			_buffer.WriteString("\n\n    </tbody>\n  </table>")

		}
		_buffer.WriteString("\n</div>")

		_buffer.WriteString("<script>\n    $(document).ready(function () {\n    $(\"#myservers\").parent().addClass(\"active\"); })\n</script>")

	}

	layout.RenderBase(_buffer, _body, _menu, _content)
}
