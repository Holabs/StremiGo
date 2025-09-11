package stremigo

import (
	"encoding/json"
	"net/http"
	"strings"
)

var (
	EnabledEndpoints = [6]string{
		PathManifest,
		PathCatalog,
		PathMeta,
		PathStream,
		PathSubtitles,
		PathConfigure,
	}
)

func isEnabledEnpoint(endpoint string) bool {
	for _, tmp := range EnabledEndpoints {
		if tmp == endpoint {
			return true
		}
	}
	return false
}

func setHeaders(w http.ResponseWriter) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

}

func Router(w http.ResponseWriter, r *http.Request, p ProviderInterface) {

	if r.Method == http.MethodOptions {
		setHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	t := ""

	// Unsecured
	if len(parts) == 1 {
		switch parts[0] {
		case "", "login":
			http.Redirect(w, r, "/configure", http.StatusMovedPermanently)
			return
		case PathConfigure:
			p.RenderConfigurePage(w, r, t)
			return
		default:
			if !isEnabledEnpoint(parts[0]) {
				http.Error(w, "Page not found", http.StatusNotFound)
				return
			}
			break
		}
	}

	// Secured

	if p.IsSecured() {

		if len(parts) < 2 {
			http.Error(w, "Neplatná cesta. Očekává se /<token>/<resource>", http.StatusBadRequest)
			return
		}

		t = parts[0]

		// remove token from path
		r.URL.Path = "/" + strings.Join(parts[1:], "/")
	}

	var data any

	switch parts[1] {
	case PathManifest:
		data = p.GetManifest(w, r, t)
		break
	case PathCatalog:
		data = p.GetCatalog(w, r, t)
		break
	case PathMeta:
		data = p.GetMeta(w, r, t)
		break
	case PathStream:
		data = p.GetStream(w, r, t)
		break
	case PathConfigure:
		p.RenderConfigurePage(w, r, t)
		return
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if data == nil {
		return
	}

	setHeaders(w)
	json.NewEncoder(w).Encode(data)
}
