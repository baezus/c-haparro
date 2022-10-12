package main

import (
	"html/template"
	"os"
)

var x = `
<!DOCTYPE html>
<head>
<meta charset = "UTF-8">
<title>Free Poetry!</title>

<link rel="stylesheet" 
type="text/css"
href="horseshoes.css" />

</head>
<body>

<h1>Lucky Chaparro</h1>
<h2>Poems</h2>
<ul class="poems">
	<li><a rel="external" href="https://poets.org/sidereal-time" type="poem" target="_blank">Sidereal Time</a> </br>
	2015</li>
	<li><a rel="external" href="https://thedailydrunk.com/f/quantum-superchalupa" type="poem" target="blank">Quantum Superchalupa</a> </br >
	2020</li>
	<li><a rel="external" href="https://issuu.com/chris_talbot/docs/bkvol11issue3summer" type="poem" target="_blank">Snatching the Mic</a> </br >
	2020</li>
</ul>
</body>
</html>
`

func main() {
	t, err := template.New("hello").Parse(x)
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, "<script>alert('world')</script>")
}
