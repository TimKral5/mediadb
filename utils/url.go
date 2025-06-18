package utils

import "strings"

func ConcatUrls(base string, appendix string, removeTrailingSlash bool) string {
	result := base

	if len(base) > 0 && len(appendix) > 0 && base[len(base)-1] != '/' && appendix[0] != '/' {
		result += "/"
	}

	result += appendix

	if removeTrailingSlash && len(result) > 1 {
		return strings.TrimSuffix(result, "/")
	}

	return result
}
