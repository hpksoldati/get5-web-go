// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/helper/loggedinmenu.gohtml

package helper

import (
	"io"
	"strings"
)

// Loggedinmenu generates templates/helper/loggedinmenu.gohtml
func Loggedinmenu() string {
	var _b strings.Builder
	RenderLoggedinmenu(&_b)
	return _b.String()
}

// RenderLoggedinmenu render templates/helper/loggedinmenu.gohtml
func RenderLoggedinmenu(_buffer io.StringWriter) {
	_buffer.WriteString("<li><a href=\"/login\">  <img src=\"/static/img/login_small.png\" height=\"18\"/></a></li>")

}
