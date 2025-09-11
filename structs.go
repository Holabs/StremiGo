package stremigo

// AddonCatalog - define addons catalog
// TransportName - required - string, only TransportHttp is currently officially supported
// TransportUrl - required - string, the URL of the addon's manifest.json file
// Manifest - required - object representing the addon's AddonManifest Object
type AddonCatalog struct {
	TransportName string         `json:"transportName"`
	TransportUrl  string         `json:"transportUrl"`
	Manifest      *AddonManifest `json:"manifest"`
}

// AddonCatalogList - list of AddonCatalog objects
// Addons - required - array of AddonCatalog objects
type AddonCatalogList struct {
	Addons []*AddonCatalog `json:"addons"`
}

// AddonManifestBehaviorHints - define configuration properties
type AddonManifestBehaviorHints struct {
	Configurable bool `json:"configurable"`
}

// AddonManifest - define add properties
// ID - required - string, identifier, dot-separated, e.g. "com.stremio.filmon"
// Name - required - string, human readable name
// Description - required - string, human readable description
// Version - required - string, semantic version of the addon
// Resources - required - array of Resource
// Types - required - array of strings, types of content supported by the addon. [ TypeMovie, TypeSeries, TypeChannel, TypeTv ]
// Catalogs - required - array of Catalog objects, lists of movies/series
// Prefixes - optional - array of strings, prefix for the addon's content, e.g. ["com.stremio.filmon.movies", "com.stremio.filmon.series"]
// BehaviorHints - optional - @see AddonManifestBehaviorHints
type AddonManifest struct {
	ID            string                      `json:"id"`
	Version       string                      `json:"version"`
	Name          string                      `json:"name"`
	Logo          string                      `json:"logo"`
	Description   string                      `json:"description"`
	Resources     []*Resource                 `json:"resources"`
	Types         []string                    `json:"types"`
	Catalogs      []*Catalog                  `json:"catalogs"`
	Prefixes      []string                    `json:"idPrefixes,omitempty"`
	BehaviorHints *AddonManifestBehaviorHints `json:"behaviorHints,omitempty"`
}

// Resource - define add properties
// Name - required - string, only supported. [ ResourceCatalog, ResourceMeta, ResourceStream, ResourceSubtitles, ResourceAddonCatalog ]
// Type - optional - array of strings, only supported types. [ TypeMovie, TypeSeries, TypeChannel, TypeTv ]
// Prefixes - optional - array of strings, prefix for the addon's content, e.g. ["com.stremio.filmon.movies", "com.stremio.filmon.series"]
type Resource struct {
	Name     string   `json:"name"`
	Type     []string `json:"types,omitempty"`
	Prefixes []string `json:"idPrefixes,omitempty"`
}

// Catalog - Movies/Series lists
// ID - required - string, the id of the catalog, can be any unique string describing the catalog (unique per addon, as an addon can have many catalogs), for example: if the catalog name is "Favourite Youtube Videos", the id can be "fav_youtube_videos"
// Type - required - string, this is the content type of the catalog. [ TypeMovie, TypeSeries, TypeChannel, TypeTv ]
// Name - required - string, human readable name of the catalog
type Catalog struct {
	ID    string          `json:"id"`
	Type  string          `json:"type"`
	Name  string          `json:"name"`
	Extra []*CatalogExtra `json:"extra,omitempty"`
}

// CatalogExtra - Extra properties for catalogs
// Name - required - string, is the name of the property; this name will be used in the extraProps argument itself. [ CatalogExtraSearched, CatalogExtraGenre, CatalogExtraSkip ]
// IsRequired - optional - boolean, set to true if this property must always be passed
// Options - optional - array of strings, possible values for this property; this is useful for things like genres, where you need the user to select from a pre-set list of options (e.g. { name: "genre", options: ["Action", "Comedy", "Drama"] });
// OptionsLimit - optional - number, the limit of values a user may select from the pre-set options list; by default, this is set to 1
type CatalogExtra struct {
	Name         string   `json:"name"`
	IsRequired   bool     `json:"isRequired,omitempty"`
	Options      []string `json:"options,omitempty"`
	OptionsLimit int      `json:"optionsLimit,omitempty"`
}

// MetaLink
// Name - required - string, human-readable name for the link
// Category - required - string, any unique category name, links are grouped based on their category. For some recommended categories see LinkCategory* constants
// URL - required - string, an external url or meta-link
type MetaLink struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	URL      string `json:"url"`
}

// Video
// ID - required - string, ID of the video
// Title - required - string, title of the video
// Released - required - string, ISO 8601, publish date of the video; for episodes, this should be the initial air date, e.g. "2010-12-06T05:00:00.000Z"
// Thumbnail - optional - string, URL to png of the video thumbnail, in the video's aspect ratio, max file size 5kb
// Streams - optional - array of Stream Objects, in case you can return links to streams while forming meta response, you can pass and array of Stream Objects to point the video to a HTTP URL, BitTorrent, YouTube or any other stremio-supported transport protocol; note that this is exclusive: passing video.streams means that Stremio will not request any streams from other addons for that video; if you return streams that way, it is still recommended to implement the streams resource
// Available - optional - boolean, set to true to explicitly state that this video is available for streaming, from your addon; no need to use this if you've passed streams
// Episode - optional - number, episode number, if applicable
// Season - optional - number, season number, if applicable
// Trailers - optional - array, containing Stream Objects
// Overview - optional - string, video overview/summary
type Video struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Released  string    `json:"released"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	Streams   []*Stream `json:"streams,omitempty"`
	Available bool      `json:"available,omitempty"`
	Episode   int       `json:"episode,omitempty"`
	Season    int       `json:"season,omitempty"`
	Trailers  []*Stream `json:"trailers,omitempty"`
	Overview  string    `json:"overview,omitempty"`
}

// MetaBehaviorHints - all properties are optional
// DefaultVideoId - string, Video.ID of the video to be played by default when the user opens the detail page
type MetaBehaviorHints struct {
	DefaultVideoId string `json:"defaultVideoId,omitempty"`
}

// Meta - metadata for movies/series
// ID - required - string, universal identifier; you may use a prefix unique to your addon, for example yt_id:UCrDkAvwZum-UTjHmzDI2iIw
// Type - required - string, type of the content; e.g. movie, series, channel, tv (see Content Types)
// Name - required - string, name of the content
// Genres - optional - array of strings, genre/categories of the content; e.g. ["Thriller", "Horror"] (warning: this will soon be deprecated in favor of links)
// Poster - optional - string, URL to png of poster; accepted aspect ratios: 1:0.675 (IMDb poster type) or 1:1 (square) ; you can use any resolution, as long as the file size is below 100kb; below 50kb is recommended
// PosterShape - optional - string, can be square (1:1 aspect) or poster (1:0.675) or landscape (1:1.77). If you don't pass this, poster is assumed
// Background - optional - string, the background shown on the stremio detail page ; heavily encouraged if you want your content to look good; URL to PNG, max file size 500kb
// Logo - optional - string, the logo shown on the stremio detail page ; encouraged if you want your content to look good; URL to PNG
// Description - optional - string, a few sentences describing your content
// ReleaseInfo - optional - string, year the content came out ; if it's series or channel, use a start and end years split by a tide - e.g. "2000-2014". If it's still running, use a format like "2000-"
// Director - optional - directors array of names (string) (warning: this will soon be deprecated in favor of links)
// Cast - optional - cast array of names (string) (warning: this will soon be deprecated in favor of links)
// ImdbRating - optional - string, IMDb rating, a number from 0.0 to 10.0 ; use if applicable
// Released - optional - string, ISO 8601, initial release date; for movies, this is the cinema debut, e.g. "2010-12-06T05:00:00.000Z"
// Trailers - optional - array of Stream objects
// Links - optional - array of MetaLink objects, can be used to link to internal pages of Stremio, example usage: array of actor / genre / director links
// Videos - optional - array of Video objects, used for channel and series; if you do not provide this (e.g. for movie), Stremio assumes this meta item has one video, and it's ID is equal to the meta item id
// Runtime - optional - string, human-readable expected runtime - e.g. "120m"
// Language - optional - string, spoken language
// Country - optional - string, official country of origin
// Awards - optional - string, human-readable that describes all the significant awards
// Website - optional - string, URL to official website
// BehaviorHints - optional - @see MetaBehaviorHints
type Meta struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Name          string             `json:"name"`
	Genres        []string           `json:"genres,omitempty"` // Deprecated
	Poster        string             `json:"poster,omitempty"`
	PosterShape   string             `json:"posterShape,omitempty"`
	Background    string             `json:"background,omitempty"`
	Logo          string             `json:"logo,omitempty"`
	Description   string             `json:"description,omitempty"`
	ReleaseInfo   string             `json:"releaseInfo,omitempty"`
	Director      []string           `json:"director,omitempty"` // Deprecated
	Cast          []string           `json:"cast,omitempty"`     // Deprecated
	ImdbRating    string             `json:"imdbRating,omitempty"`
	Released      string             `json:"released,omitempty"`
	Trailers      []*Stream          `json:"trailers,omitempty"`
	Links         []*MetaLink        `json:"links,omitempty"`
	Videos        []*Video           `json:"videos,omitempty"`
	Runtime       string             `json:"runtime,omitempty"`
	Language      string             `json:"language,omitempty"`
	Country       string             `json:"country,omitempty"`
	Awards        string             `json:"awards,omitempty"`
	Website       string             `json:"website,omitempty"`
	BehaviorHints *MetaBehaviorHints `json:"behaviorHints,omitempty"`
}

// MetaPreview - Shorter version of Meta with required poster for movies/series
// ID - required - string, universal identifier; you may use a prefix unique to your addon, for example yt_id:UCrDkAvwZum-UTjHmzDI2iIw
// Type - required - string, type of the content; e.g. movie, series, channel, tv (see Content Types)
// Name - required - string, name of the content
// Poster - optional - string, URL to png of poster; accepted aspect ratios: 1:0.675 (IMDb poster type) or 1:1 (square) ; you can use any resolution, as long as the file size is below 100kb; below 50kb is recommended
// PosterShape - optional - string, can be square (1:1 aspect) or poster (1:0.675) or landscape (1:1.77). If you don't pass this, poster is assumed
// Genres - optional - array of strings, genre/categories of the content; e.g. ["Thriller", "Horror"] (warning: this will soon be deprecated in favor of links)
// ImdbRating - optional - string, IMDb rating, a number from 0.0 to 10.0 ; use if applicable
// ReleaseInfo - optional - string, year the content came out ; if it's series or channel, use a start and end years split by a tide - e.g. "2000-2014". If it's still running, use a format like "2000-"
// Director - optional - directors array of names (string) (warning: this will soon be deprecated in favor of links)
// Cast - optional - cast array of names (string) (warning: this will soon be deprecated in favor of links)
// Links - optional - array of MetaLink objects, can be used to link to internal pages of Stremio, example usage: array of actor / genre / director links
// Description - optional - string, a few sentences describing your content
// Trailers - optional - array of Stream objects
type MetaPreview struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	Poster      string      `json:"poster"`
	PosterShape string      `json:"posterShape,omitempty"`
	Genres      []string    `json:"genres,omitempty"` // Deprecated
	ImdbRating  string      `json:"imdbRating,omitempty"`
	ReleaseInfo string      `json:"releaseInfo,omitempty"`
	Director    []string    `json:"director,omitempty"` // Deprecated
	Cast        []string    `json:"cast,omitempty"`     // Deprecated
	Links       []*MetaLink `json:"links,omitempty"`
	Description string      `json:"description,omitempty"`
	Trailers    []*Stream   `json:"trailers,omitempty"`
}

// MetaList - list of Meta objects
// Metas - required - array of Meta objects
// CacheMaxAge - optional - number, max age of the cache in seconds
// StaleRevalidate - optional - number, max age of the stale cache in seconds
// StaleError - optional - number, max age of the stale error cache in seconds
type MetaList struct {
	Metas           []*Meta `json:"metas"`
	CacheMaxAge     int     `json:"cacheMaxAge,omitempty"`
	StaleRevalidate int     `json:"staleRevalidate,omitempty"`
	StaleError      int     `json:"staleError,omitempty"`
}

// MetaPreviewList - list of MetaPreview objects
// Metas - required - array of MetaPreview objects
// CacheMaxAge - optional - number, max age of the cache in seconds
// StaleRevalidate - optional - number, max age of the stale cache in seconds
// StaleError - optional - number, max age of the stale error cache in seconds
type MetaPreviewList struct {
	Metas           []*MetaPreview `json:"metas"`
	CacheMaxAge     int            `json:"cacheMaxAge,omitempty"`
	StaleRevalidate int            `json:"staleRevalidate,omitempty"`
	StaleError      int            `json:"staleError,omitempty"`
}

// Stream - Stream information. "required/optional" means that one of the fields is required, but not all.
// URL - required/optional - string, direct URL to a video stream - must be an MP4 through https; others supported (other video formats over http/rtmp supported if you set StreamBehaviorHints.NotWebReady)
// YtId - required/optional - string, youtube video ID, plays using the built-in YouTube player
// InfoHash - required/optional - string, info hash of a torrent file, and FileIdx is the index of the video file within the torrent; if FileIdx is not specified, the largest file in the torrent will be selected
// FileIdx - required/optional - number, the index of the video file within the torrent (from InfoHash); if fileIdx is not specified, the largest file in the torrent will be selected
// ExternalUrl - required/optional - string, meta-link or an external url to the video, which should be opened in a browser (webpage), e.g. link to Netflix
// Name - optional - string, name of the stream; usually used for stream quality
// Title - optional - string, description of the stream (warning: this will soon be deprecated in favor of Stream.Description)
// Description - optional - string, description of the stream (previously Stream.Title)
// Sources - optional - array of strings, represents a list of torrent tracker URLs and DHT network nodes. This attribute can be used to provide additional peer discovery options when infoHash is also specified, but it is not required. If used, each element can be a tracker url (tracker:<protocol>://<host>:<port>) where <protocol> can be either http or udp. A DHT node (dht:<node_id/info_hash>) can also be included.
// BehaviorHints - optional - @see StreamBehaviorHints
//
// WARNING: Use of DHT may be prohibited by some private trackers as it exposes torrent activity to a broader network, potentially finding more peers.
type Stream struct {
	URL           string               `json:"url,omitempty"`
	YtId          string               `json:"ytId,omitempty"`
	InfoHash      string               `json:"infoHash,omitempty"`
	FileIdx       int                  `json:"fileIdx,omitempty"`
	ExternalUrl   string               `json:"externalUrl,omitempty"`
	Name          string               `json:"name,omitempty"`
	Title         string               `json:"title,omitempty"`
	Description   string               `json:"description,omitempty"` // Deprecated
	Sources       []string             `json:"sources,omitempty"`
	BehaviorHints *StreamBehaviorHints `json:"behaviorHints,omitempty"`
}

// StreamBehaviorHints - all properties are optional
// CountryWhitelist - array of ISO 3166-1 alpha-3 country codes in lowercase in which the stream is accessible; which hints it's restricted to particular countries
// NotWebReady - bool, applies if the protocol of the url is http(s); needs to be set to true if the URL does not support https or is not an MP4 file
// BingeGroup - string, if defined, addons with the same StreamBehaviorHints.BingeGroup will be chosen automatically for binge watching; this should be something that identifies the stream's nature within your addon: for example, if your addon is called "gobsAddon", and the stream is 720p, the BingeGroup should be "gobsAddon-720p"; if the next episode has a stream with the same BingeGroup, stremio should select that stream implicitly
// VideoHash - string, the calculated OpenSubtitles hash of the video, this will be used when the streaming server is not connected (so the hash cannot be calculated locally), this value is passed to subtitle addons to identify correct subtitles
// VideoSize - number, size of the video file in bytes, this value is passed to the subtitle addons to identify correct subtitles
// Filename - string, filename of the video file, although optional, it is highly recommended to set it when using stream.url (when possible) in order to identify correct subtitles (addon sdk will show a warning if it is not set in this case), this value is passed to the subtitle addons to identify correct subtitles
// ProxyHeaders - only applies to urls; When using this property, you must also set StreamBehaviorHints.NotWebReady to true; This is an object containing request and response which include the headers that should be used for the stream (example value: { "request": { "User-Agent": "Stremio" } })
type StreamBehaviorHints struct {
	// CountryWhitelist - ISO 3166-1 alpha-3 country codes in lowercase
	CountryWhitelist []string `json:"countryWhitelist,omitempty"`
	// NotWebReady - "true" if the stream is not available for web - stream is not mp4 or doesn't support https
	NotWebReady  bool                   `json:"notWebReady,omitempty"`
	BingeGroup   string                 `json:"bingeGroup,omitempty"`
	VideoHash    string                 `json:"videoHash,omitempty"`
	VideoSize    int64                  `json:"videoSize,omitempty"`
	Filename     string                 `json:"filename,omitempty"`
	ProxyHeaders map[string]interface{} `json:"proxyHeaders,omitempty"`
}

// StreamList - list of Stream objects
// Streams - required - array of Stream objects
// CacheMaxAge - optional - number, max age of the cache in seconds
// StaleRevalidate - optional - number, max age of the stale cache in seconds
// StaleError - optional - number, max age of the stale error cache in seconds
type StreamList struct {
	Streams         []*Stream `json:"streams"`
	CacheMaxAge     int       `json:"cacheMaxAge,omitempty"`
	StaleRevalidate int       `json:"staleRevalidate,omitempty"`
	StaleError      int       `json:"staleError,omitempty"`
}

// Subtitles - Stream information
// ID - required - string, unique identifier for each subtitle, if you have more than one subtitle with the same language, the id will differentiate them
// URL - required - string, url to the subtitle file
// Lang - required - string, language code for the subtitle, if a valid ISO 639-2 code is not sent, the text of this value will be used instead
type Subtitles struct {
	ID   string `json:"id"`
	URL  string `json:"url"`
	Lang string `json:"lang"`
}

// SubtitlesList - list of Subtitles objects
// Subtitles - required - array of Subtitles objects
// CacheMaxAge - optional - number, max age of the cache in seconds
// StaleRevalidate - optional - number, max age of the stale cache in seconds
// StaleError - optional - number, max age of the stale error cache in seconds
type SubtitlesList struct {
	Subtitles       []*Subtitles `json:"subtitles"`
	CacheMaxAge     int          `json:"cacheMaxAge,omitempty"`
	StaleRevalidate int          `json:"staleRevalidate,omitempty"`
	StaleError      int          `json:"staleError,omitempty"`
}
