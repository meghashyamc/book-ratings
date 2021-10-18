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

var ratingFuncs = map[string](func(book string) (float32, error)){goodreads: GetGoodreadsRating, libraryThing: GetLibraryThingRating}
