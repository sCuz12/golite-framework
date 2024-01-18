package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)


func CheckForMaintenance(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		//maintenance logic
		value, exists := os.LookupEnv("MAINTENANCE_MODE")

		if !exists {
			// Proceed with normal handling
			next.ServeHTTP(w, r)
			return
		}

		inMaintenanceMode,err := strconv.ParseBool(value)

		if err != nil {
            fmt.Fprintf(w, "Error parsing MAINTENANCE_MODE variable: %v", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

		if !inMaintenanceMode {
			next.ServeHTTP(w, r)
			return
		}
		
		// If in maintenance mode, return a maintenance message
        w.WriteHeader(http.StatusServiceUnavailable)
        w.Write([]byte("The site is currently down for maintenance"))
	})
}