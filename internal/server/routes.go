package server

import "net/http"

// Local LAN routes
func LocalRoutes(s *LocalServer, mux *http.ServeMux) {
	// Device Endpoints
	mux.HandleFunc("POST /devices", s.PostDeviceHandler)
	mux.HandleFunc("PUT /devices", s.PutDevicesHandler)
	mux.HandleFunc("GET /devices", s.GetDevicesHandler)
	mux.HandleFunc("GET /devices/{deviceID}", s.GetDeviceByIDHandler)
	mux.HandleFunc("DELETE /devices/{deviceID}", s.DeleteDeviceByIDHandler)

	// WDLM Endpoints
	mux.HandleFunc("POST /wdlms", s.PostWdlmHandler)
	mux.HandleFunc("PUT /wdlms", s.PutWdlmHandler)
	mux.HandleFunc("GET /wdlms", s.GetWdlmsHandler)
	mux.HandleFunc("GET /wdlms/{wdlmID}", s.GetWdlmByIDHandler)
	mux.HandleFunc("DELETE /wdlms/{wdlmID}", s.DeleteWdlmByID)

	// Admin Endpoints
	mux.HandleFunc("POST /admin/reset/devices", s.ResetDevicesHandler)
}

// Public internet routes
func PublicRoutes(s *PublicServer, mux *http.ServeMux) {

}
