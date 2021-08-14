// logos.go
package logos

import (
	"strings"
)

// This portion of fnufetch uses significant portions of 'pfetch' source code
// and, as such, must provide the copyright notice bundled with the software.

/*


The MIT License (MIT)

Copyright (c) 2016-2019 Dylan Araps

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


*/

var SystemLogos map[string][]string = map[string][]string{
	"Arch Linux": []string{
		`${c6}       /\`,
		`${c6}      /  \`,
		`${c6}     /\   \`,
		`${c4}    /      \`,
		`${c4}   /   ,,   \`,
		`${c4}  /   |  |  -\`,
		`${c4} /_-''    ''-_\`,
	},
	"Ubuntu": []string{
		"  ${c3}         _",
		"  ${c3}     ---(_)",
		"  ${c3} _/  ---  \\",
		"  ${c3}(_) |   |",
		"   ${c3} \\  --- _/",
		"     ${c3} ---(_)",
		"        ",
	},
}

var DefaultLogo []string = []string{
	"  ${c4}     ___     ",
	"  ${c4}    (${c7}.. ${c4}|",
	"  ${c4}    (${c5}<> ${c4}|",
	"  ${c4}   / ${c7}__  ${c4}\\",
	"  ${c4}  ( ${c7}/  \\ ${c4}/|",
	"  ${c5} _${c4}/\\ ${c7}__)${c4}/${c5}_${c4})",
	`  ${c5} \/${c4}-____${c5}\/`,
}

var ColorPallette []string = []string{
	"${c0}", "\u001b[0m\u001b[38;5;248m",
	"${c1}", "\u001b[0m\u001b[31;1m",
	"${c2}", "\u001b[0m\u001b[32;1m",
	"${c3}", "\u001b[0m\u001b[33;1m",
	"${c4}", "\u001b[0m\u001b[34;1m",
	"${c5}", "\u001b[0m\u001b[35;1m",
	"${c6}", "\u001b[0m\u001b[36;1m",
	"${c7}", "\u001b[0m\u001b[37;1m",
}

func ColorReplace(logo []string) []string {
	for line, _ := range logo {
		logo[line] = strings.NewReplacer(ColorPallette...).Replace(logo[line])
	}

	return append(logo, "\r")
}

func GetLogo(name string) []string {
	logo, ok := SystemLogos[name]

	if !ok {
		return ColorReplace(DefaultLogo)
	}

	return ColorReplace(logo)
}
