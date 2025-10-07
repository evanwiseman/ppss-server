package server

import "net/http"

// Local LAN routes
func LocalRoutes(s *LocalServer, mux *http.ServeMux) {
	mux.HandleFunc("POST /devices", s.PostDeviceHandler)
	mux.HandleFunc("DELETE /devices/{deviceID}", s.DeleteDeviceByIDHandler)
}

// Public internet routes
func PublicRoutes(s *PublicServer, mux *http.ServeMux) {

}
