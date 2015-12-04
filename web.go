package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-lsst/ncs/drivers/m702"
	"golang.org/x/net/websocket"
)

func runWebServer() error {
	srv := newWebServer()
	http.Handle("/lsst-fcs-motors", srv)
	http.Handle("/lsst-fcs-motors/data", websocket.Handler(srv.dataHandler))
	log.Printf("listening on http://%s/lsst-fcs-motors ...\n", srv.Addr)
	err := http.ListenAndServe(srv.Addr, nil)
	return err
}

type client struct {
	srv   *webServer
	ws    *websocket.Conn
	datac chan []byte
}

func (c *client) run() {
	defer func() {
		c.srv.unregister <- c
		c.ws.Close()
	}()
	//c.ws.SetReadLimit(maxMessageSize)
	//c.ws.SetReadDeadline(time.Now().Add(pongWait))
	//c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for data := range c.datac {
		err := websocket.Message.Send(c.ws, string(data))
		if err != nil {
			log.Printf("error sending data to [%v]: %v\n", c.ws, err)
			break
		}
	}
}

type webServer struct {
	Motors []motorStatus
	Addr   string
	tmpl   *template.Template
	params []m702.Parameter

	clients    map[*client]bool // registered clients
	register   chan *client
	unregister chan *client

	datac chan []motorStatus
}

func newWebServer() *webServer {
	srv := &webServer{
		Motors: []motorStatus{
			{Addr: "134.158.125.223:502"},
			{Addr: "134.158.125.224:502"},
		},
		Addr:       "clrinfopc07.in2p3.fr:7070",
		tmpl:       template.Must(template.New("fcs").Parse(displayTmpl)),
		params:     make([]m702.Parameter, len(params)),
		clients:    make(map[*client]bool),
		register:   make(chan *client),
		unregister: make(chan *client),
		datac:      make(chan []motorStatus),
	}
	copy(srv.params, params)

	go srv.run()

	return srv
}

func (srv *webServer) run() {
	go func() {
		tick := time.NewTicker(5 * time.Second)
		srv.publishData()
		for range tick.C {
			srv.publishData()
		}
	}()

	buf := new(bytes.Buffer)
	for {
		select {
		case c := <-srv.register:
			srv.clients[c] = true
		case c := <-srv.unregister:
			if _, ok := srv.clients[c]; ok {
				delete(srv.clients, c)
				close(c.datac)
			}
		case data := <-srv.datac:
			buf.Reset()
			err := json.NewEncoder(buf).Encode(data)
			if err != nil {
				log.Printf("error marshalling data: %v\n", err)
				continue
			}
			for c := range srv.clients {
				select {
				case c.datac <- buf.Bytes():
				default:
					close(c.datac)
					delete(srv.clients, c)
				}
			}
		}
	}
}

var global = 0

func (srv *webServer) publishData() {
	status := make([]motorStatus, len(srv.Motors))
	copy(status, srv.Motors)

	global++
	for i := range status {
		data := &status[i]
		//motor := m702.New(data.Addr)
		for _, p := range srv.params {
			/* FIXME
			err := motor.ReadParam(&p)
			if err != nil {
				log.Printf("error reading Pr-%v: %v\n", p, err)
				continue
			}
			*/
			var err error
			p.Data[2] = byte(i)
			p.Data[3] = byte(global)
			data.Params = append(data.Params, newMotorData(p, err))
		}
	}

	log.Printf("pushing data... (global=%d)\n", global)
	srv.datac <- status
}

func (srv *webServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("accepting new connection from %v...\n", r.Host)
	srv.tmpl.Execute(w, srv)
}

func (srv *webServer) dataHandler(ws *websocket.Conn) {
	c := &client{
		srv:   srv,
		datac: make(chan []byte, 256),
		ws:    ws,
	}
	srv.register <- c
	c.run()
}

const displayTmpl = `
<html>

<head>
<title>LSST FCS Motors</title>

<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>

<style>

header {
    background-color:black;
    color:white;
    text-align:center;
    padding:5px;
}
nav {
    line-height:100%;
    background-color:#eeeeee;
    height:100%;
    width:20%;
    float:left;
    padding:5px;
}
section {
    width:75%;
    float:left;
    padding:10px;
}

div.header {
    background-color:black;
    color:white;
    text-align:center;
    padding:5px;
}

div.motor {
    width:70%;
    float:left;
    padding:5px; 
}

div.motor-1 {
    width:45%;
    float:right;
    padding:5px; 
	display: none;
}

#footer {
    background-color:black;
    color:white;
    clear:both;
    text-align:center;
    padding:5px; 
}
</style>

<script type="text/javascript">
$(document).ready(function() {
	var sock = null;
	var wsuri = "ws://{{.Addr}}/lsst-fcs-motors/data";
	var data = [ ];

	var imotor = 0;
	{{range $idx, $m := .Motors}}
	$("#butn-motor-{{$idx}}").click(function() {
		imotor = {{$idx}};
		display(imotor);
	});
	{{end}}

	window.onload = function() {
		console.log("onload");
		sock = new WebSocket(wsuri);
		sock.onopen = function() {
			console.log("connected to " + wsuri);
		}
		sock.onclose = function(e) {
			console.log("connection closed ("+e.code+")");
			sock.close();
			sock = null;
		}
		sock.onmessage = function(e) {
			var obj = JSON.parse(e.data);
			console.log("new data: "+obj);
			// console.log("json: "+e.data);
			data = obj;
			display(imotor);
		}
	};

	var display = function(i) {
		var motor = data[i];
		var hdr = $("#motor-header");
		hdr.text("Motor-"+i+" ("+motor["Addr"]+")");
		var params = motor["Params"];
		var table = $("#motor-table");
		table.empty();
		table.append("<tr><th>Parameter</th><th>Title</th><th>Value</th></tr>");
		for (var j=0; j < params.length; j++) {
			var p = params[j];
			var index = p["Index"];
			var title = p["Title"];
			var value = p["Value"];
			var row = "<tr class=\"data\"><td>"+index+"</td><td>"+title+"</td><td>"+value+"</td></tr>";
			table.append(row);
		}
	};
});
</script>
</head>

<body>

<header>
<h1>LSST FCS Motors</h1>
</header>

<nav>
{{range $i, $m := .Motors}}<button id="butn-motor-{{$i}}">Motor-{{$i}} ({{$m.Addr}})</button>{{end}}
</nav>

<section>

<div class="motor">
  <div class="header"><h1 id="motor-header">Motor</h1></div>
  <table border="1" style="width:100%" id="motor-table">
    <tbody>
      <tr><th>Parameter</th><th>Title</th><th>Value</th></tr>
	</tbody>
  </table>
</div>

</section>

</body>

</html>
`

type motorStatus struct {
	Addr   string
	Params []motorData
}

type motorData struct {
	Index string
	Title string
	Value string
}

func newMotorData(p m702.Parameter, err error) motorData {
	data := motorData{
		Index: fmt.Sprintf("<pre><code>%v</code></pre>", p),
		Title: p.Title,
	}
	if err != nil {
		data.Value = fmt.Sprintf("err=%v", err)
	} else {
		data.Value = fmt.Sprintf(
			"<pre><code>%s ==> %6d</code></pre>",
			displayBytes(p.Data[:]),
			codec.Uint32(p.Data[:]),
		)
	}

	return data
}
