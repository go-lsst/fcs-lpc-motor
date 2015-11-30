package main

import (
	"fmt"
	"log"
	"net/http"
)

func runWebServer() error {
	srv := &webServer{
		motor:  NewMotor("134.158.125.223:502"),
		params: make([]Param, len(params)),
	}
	copy(srv.params, params)

	http.Handle("/lsst-fcs-motors", srv)
	log.Printf("listening on http://localhost:7070/lsst-fcs-motors ...\n")
	err := http.ListenAndServe(":7070", nil)
	return err
}

type webServer struct {
	motor  Motor
	params []Param
}

func (srv *webServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w,
		`%s
<title>LSST FCS Motors</title>
`,
		css,
	)

	fmt.Fprintf(w, "<body>\n\n")
	for i, motor := range []Motor{
		NewMotor("134.158.125.223:502"),
		NewMotor("134.158.125.224:502"),
	} {
		fmt.Fprintf(w, "<div class=\"motor-%d\">\n", i+1)
		fmt.Fprintf(w, "<div class=\"header\"><h1>Motor-%d (%s)</h1></div>\n", i+1, motor.Address)
		fmt.Fprintf(w, "<table border=\"1\" style=\"width:100%%\">\n")
		fmt.Fprintf(w, "\t<tr><th>Parameter</th><th>Title</th><th>Value</th></tr>\n")
		for _, p := range srv.params {
			fmt.Fprintf(w, "\t<tr>\n")
			fmt.Fprintf(
				w,
				"\t\t<td>%02d.%03d</td><td>%s</td> ",
				p.menu, p.index, p.title,
			)
			o, err := motor.read(NewParameter(p.mbreg()))
			if err != nil {
				fmt.Fprintf(w, "<td>err=%v</td>\n", err)
				fmt.Fprintf(w, "\t</tr>\n")
				continue
			}
			fmt.Fprintf(
				w,
				"<td><pre><code>%s ==> %6d</code></pre></td>\n",
				displayBytes(o), codec.Uint16(o),
			)
			fmt.Fprintf(w, "\t</tr>\n")
		}
		fmt.Fprintf(w, "</table>\n</div>\n\n")
	}
	fmt.Fprintf(w, "</body>\n")
}

const css = `<style>
div.header {
    background-color:black;
    color:white;
    text-align:center;
    padding:5px;
}

div.motor-1 {
    width:45%;
    float:left;
    padding:5px; 
}

div.motor-2 {
    width:45%;
    float:right;
    padding:5px; 
}

#footer {
    background-color:black;
    color:white;
    clear:both;
    text-align:center;
    padding:5px; 
}
</style>
`
