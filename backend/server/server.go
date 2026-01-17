package server

import (
	"encoding/json"
	"io/fs"
	"net/http"

	"github.com/yhiraki/wakeonlan-webapp/backend/config"
	"github.com/yhiraki/wakeonlan-webapp/backend/wol"
)

type Server struct {
	targets []config.Target
	wolSvc  wol.Service
	mux     *http.ServeMux
}

func NewServer(targets []config.Target, wolSvc wol.Service) *Server {
	s := &Server{
		targets: targets,
		wolSvc:  wolSvc,
		mux:     http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/api/targets", s.handleGetTargets)
	s.mux.HandleFunc("/api/wake", s.handleWake)
}

func (s *Server) MountStatic(fSys fs.FS) {
	s.mux.Handle("/", http.FileServer(http.FS(fSys)))
}

func (s *Server) handleGetTargets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.targets)
}

func (s *Server) handleWake(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		MAC string `json:"mac"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := s.wolSvc.Wake(body.MAC); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
