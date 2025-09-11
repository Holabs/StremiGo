package stremigo

// Resources
const (
	ResourceAddonCatalog string = "addon_catalog"
	ResourceCatalog      string = "catalog"
	ResourceMeta         string = "meta"
	ResourceStream       string = "stream"
	ResourceSubtitles    string = "subtitles"
)

// Paths
const (
	PathManifest  string = "manifest.json"
	PathCatalog   string = ResourceCatalog
	PathMeta      string = ResourceMeta
	PathStream    string = ResourceStream
	PathSubtitles string = ResourceSubtitles
	PathConfigure string = "configure"
)

// Content types
const (
	TypeMovie   string = "movie"
	TypeSeries  string = "series"
	TypeChannel string = "channel"
	TypeTv      string = "tv"
)

// Available catalog extra fields
const (
	CatalogExtraSearched string = "search"
	CatalogExtraGenre    string = "genre"
	CatalogExtraSkip     string = "skip"
)

// Available MetaLink.Category options
const (
	LinkCategoryActor    string = "actor"
	LinkCategoryDirector string = "director"
	LinkCategoryWriter   string = "writer"
)

// Prefixes for stream ids - not all
const (
	PrefixImdb    string = "tt"
	PrefixYoutube string = "yt_id:"
)

// Available poster shapes
const (
	PosterShapeSquare    string = "square"
	PosterShapePoster    string = "poster"
	PosterShapeLandscape string = "landscape"
)

const (
	TransportHttp string = "http"
)
