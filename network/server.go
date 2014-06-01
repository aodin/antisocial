package network

import (
	"encoding/json"
	"fmt"
	"github.com/aodin/aspect"
	"github.com/aodin/aspect/postgis"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type Server struct {
	config Config
	db     *aspect.DB
	home   *template.Template
}

func (s *Server) API(w http.ResponseWriter, r *http.Request) {
	// Was there a requested lat, lng?
	lat, lng := requestLatLng(r)
	if lat != 0.0 && lng != 0.0 {
		s.Search(w, r, lat, lng)
		return
	}

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) >= 3 && parts[2] != "" {
		s.Hood(w, r)
		return
	}

	stmt := aspect.Select(
		Hoods.C["id"],
		Hoods.C["name"],
		Hoods.C["population"],
		Hoods.C["housing"],
		Hoods.C["area"],
		Hoods.C["crimes"],
		Hoods.C["311"],
		Hoods.C["foreclosures"],
		Hoods.C["licenses"],
		Hoods.C["score"],
		Hoods.C["rank"],
		postgis.GeoJSON(Hoods.C["geom"]),
	).OrderBy(Hoods.C["name"])

	// Get all neighborhoods
	var hoods []Hood
	if err := s.db.QueryAll(stmt, &hoods); err != nil {
		log.Println(err)
	}

	// JSON dump the data
	b, err := json.Marshal(hoods)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

func Favicon(w http.ResponseWriter, r *http.Request) {
	w.Write(nil)
}

func (s *Server) Root(w http.ResponseWriter, r *http.Request) {
	// Display a list of neighborhoods
	stmt := aspect.Select(
		Hoods.C["id"],
		Hoods.C["name"],
		Hoods.C["population"],
		Hoods.C["housing"],
		Hoods.C["area"],
		Hoods.C["crimes"],
		Hoods.C["311"],
		Hoods.C["foreclosures"],
		Hoods.C["licenses"],
		Hoods.C["score"],
		Hoods.C["rank"],
		postgis.GeoJSON(Hoods.C["geom"]),
	).OrderBy(Hoods.C["name"])

	// Get all neighborhoods
	var hoods []Hood
	if err := s.db.QueryAll(stmt, &hoods); err != nil {
		log.Println(err)
	}

	s.home.Execute(w, s.config)

	// for _, hood := range hoods {
	// 	w.Write([]byte(hood.String() + "\n"))
	// }
}

func (s *Server) Search(w http.ResponseWriter, r *http.Request, lat, lng float64) {
	// Find the neighborhood that this point is within
	p := postgis.Point{lat, lng}
	stmt := aspect.Select(
		Hoods.C["id"],
		Hoods.C["name"],
		Hoods.C["population"],
		Hoods.C["housing"],
		Hoods.C["area"],
		Hoods.C["crimes"],
		Hoods.C["311"],
		Hoods.C["foreclosures"],
		Hoods.C["licenses"],
		Hoods.C["score"],
		Hoods.C["rank"],
		postgis.GeoJSON(Hoods.C["geom"]),
	).Where(postgis.Within(Hoods.C["geom"], p))

	var hood Hood
	if err := s.db.QueryOne(stmt, &hood); err != nil {
		log.Println(err)
	}
	// JSON dump the data
	b, err := json.Marshal(hood)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

func requestLatLng(r *http.Request) (lat, lng float64) {
	latS := r.FormValue("lat")
	lngS := r.FormValue("lng")
	var err error
	if lat, err = strconv.ParseFloat(latS, 64); err != nil {
		return
	}
	if lng, err = strconv.ParseFloat(lngS, 64); err != nil {
		return
	}
	return
}

func (s *Server) Hood(w http.ResponseWriter, r *http.Request) {
	// What id was requested?
	// TODO There has to be a better way...
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) < 3 || parts[2] == "" {
		http.NotFound(w, r)
		return
	}

	// Cast to an integer
	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Get the json data for that neighborhood
	stmt := aspect.Select(
		Hoods.C["id"],
		Hoods.C["name"],
		Hoods.C["population"],
		Hoods.C["housing"],
		Hoods.C["area"],
		Hoods.C["crimes"],
		Hoods.C["311"],
		Hoods.C["foreclosures"],
		Hoods.C["licenses"],
		Hoods.C["score"],
		Hoods.C["rank"],
		postgis.GeoJSON(Hoods.C["geom"]),
	).Where(Hoods.C["id"].Equals(id))

	var hood Hood
	if err := s.db.QueryOne(stmt, &hood); err != nil {
		log.Println(err)
	}
	// JSON dump the data
	b, err := json.Marshal(hood)
	if err != nil {
		log.Println(err)
	}
	w.Write(b)
}

func (s *Server) Rank(w http.ResponseWriter, r *http.Request) {
	// Allow the user to rank the neighborhoods by various stats
}

func (s *Server) ListenAndServe() error {
	address := fmt.Sprintf(":%d", s.config.Port)
	log.Printf("Server running on address %s\n", address)
	return http.ListenAndServe(address, nil)
}

func NewServer(config Config, db *aspect.DB) *Server {
	homePath := filepath.Join(config.TemplateDir, "home.html")
	s := &Server{
		config: config,
		db:     db,
		home:   template.Must(template.ParseFiles(homePath)),
	}
	// Serve the static files
	staticURL := "/static/"
	http.Handle(
		staticURL,
		http.StripPrefix(staticURL, http.FileServer(http.Dir(config.StaticDir))),
	)
	http.HandleFunc("/", s.Root)
	http.HandleFunc("/favicon.ico", Favicon)
	http.HandleFunc("/api/", s.API)
	return s
}
