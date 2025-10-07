package server

import "net/http"

// Local LAN routes
func LocalRoutes(s *LocalServer, mux *http.ServeMux) {
	mux.HandleFunc("POST /devices", s.PostDeviceHandler)
	mux.HandleFunc("DELETE /device/{serial_number}", s.DeleteDeviceHandler)
}

// Public internet routes
func PublicRoutes(s *PublicServer, mux *http.ServeMux) {

}
