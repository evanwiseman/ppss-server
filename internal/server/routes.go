package server

import "net/http"

// Local LAN routes
func LocalRoutes(s *LocalServer, mux *http.ServeMux) {
	// Device Endpoints
	mux.HandleFunc("POST /devices", s.PostDeviceHandler)
	mux.HandleFunc("GET /devices", s.GetDevicesHandler)
	mux.HandleFunc("GET /devices/{deviceID}", s.GetDeviceByIDHandler)
	mux.HandleFunc("PUT /devices", s.PutDevicesHandler)
	mux.HandleFunc("DELETE /devices/{deviceID}", s.DeleteDeviceByIDHandler)

	// WDLM Endpoints

	// Admin Endpoints
	mux.HandleFunc("POST /admin/reset/devices", s.ResetDevicesHandler)
}

// Public internet routes
func PublicRoutes(s *PublicServer, mux *http.ServeMux) {

}
