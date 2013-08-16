package main

import (
	restful "github.com/emicklei/go-restful"
	"github.com/jmcvetta/randutil"
	"log"
	"net/http"
)

type Crew struct {
	Name string
	Rank string
}

type Foo struct {
	Crew []Crew
}

type Bar struct {
	Rand string
}

func getFoo(request *restful.Request, response *restful.Response) {
	f := Foo{
		Crew: []Crew{
			Crew{"Kirk", "Captain"},
			Crew{"McCoy", "Doctor"},
		},
	}
	response.WriteEntity(&f)

}

func getBar(request *restful.Request, response *restful.Response) {
	var b Bar
	b.Rand, _ = randutil.AlphaString(8)
	response.WriteEntity(&b)
}

func main() {
	baseUrl := ":8080"
	var ws restful.WebService
	ws.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/bar").To(getBar).
		Writes(Bar{}))
	ws.Route(ws.GET("/foo").To(getFoo).
		Writes(Foo{}))
	restful.Add(&ws)
	log.Printf("Starting webserver on %v...", baseUrl)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
