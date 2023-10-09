package validation

type Usecase interface {
	CheckReferer(referer string) bool
	CheckMediaURL(url string) bool
}
