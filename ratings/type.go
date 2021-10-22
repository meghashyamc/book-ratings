package ratings

import "os"

const (
	goodreads    = "Goodreads"
	libraryThing = "Library Thing"
	reqType      = "request_type"
	url          = "url"
	paramKey     = "param_key"
	paramVal     = "param_val"
	headerKey    = "header_key"
	headerVal    = "header_val"
)

var (
	ratingFuncs       = map[string](func(book string) (float32, error)){libraryThing: GetLibraryThingRating, goodreads: GetGoodreadsRating}
	searchIDsForSites = map[string]string{libraryThing: os.Getenv("LIBRARY_THING_SEARCH_ID"), goodreads: os.Getenv("GOODREADS_SEARCH_ID")}
)
