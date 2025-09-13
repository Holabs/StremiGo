package stremigo

import "net/http"

type ProviderInterface interface {
	GetManifest(w http.ResponseWriter, r *http.Request, token string) *AddonManifest

	GetCatalog(w http.ResponseWriter, r *http.Request, token string) *MetaPreviewList

	GetMeta(w http.ResponseWriter, r *http.Request, token string) *Meta

	GetStream(w http.ResponseWriter, r *http.Request, token string) *StreamList

	GetSubtitles(w http.ResponseWriter, r *http.Request, token string) *SubtitlesList

	RenderConfigurePage(w http.ResponseWriter, r *http.Request, token string)

	IsSecured() bool
}
