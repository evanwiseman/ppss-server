package server

import "net/http"

// Local LAN routes
func LocalRoutes(s *LocalServer, mux *http.ServeMux) {
	// Device Endpoints
	mux.HandleFunc("POST /devices", s.PostDeviceHandler)
	mux.HandleFunc("DELETE /devices/{deviceID}", s.DeleteDeviceByIDHandler)
	mux.HandleFunc("GET /devices", s.GetDevicesHandler)
	mux.HandleFunc("GET /devices/{deviceID}", s.GetDeviceByIDHandler)
}

// Public internet routes
func PublicRoutes(s *PublicServer, mux *http.ServeMux) {

}
