// Package ping provides a controller to handle the ping endpoint (/ping) in the playerService.
// This endpoint is used for checking if the server is alive.

// The ping package includes the Ping function and the PingResp structure.

// The Ping function handles GET requests on the "/ping" endpoint. It returns a JSON response with the message "pong" to indicate that the server is alive and responsive.

// The PingResp structure defines a single field "Ping" to hold the ping value response. It is used to serialize the response into JSON format.

package ping