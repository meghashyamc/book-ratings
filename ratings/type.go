package ratings

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
	ratingFuncs              = map[string](func(book string) (float32, error)){libraryThing: GetLibraryThingRating, goodreads: GetGoodreadsRating}
	searchIDEnvVariableNames = map[string]string{libraryThing: "LIBRARY_THING_SEARCH_ID", goodreads: "GOODREADS_SEARCH_ID"}
)
