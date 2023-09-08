package tools

var (
	DEFALUT_LIMIT = 10
	DEFALUT_PAGE  = 1
)

// InitPage 页面初始化
func InitPage(paramsLimit int, paramsPage int) (limit int, page int) {
	limit = DEFALUT_LIMIT
	if paramsLimit > 0 {
		limit = paramsLimit
	}

	page = DEFALUT_PAGE
	if paramsPage > 0 {
		page = paramsPage
	}

	return limit, page
}

// PageOffset 页面偏移量
func PageOffset(limit int, page int) (offset int) {
	if limit > 0 && page > 0 {
		offset = (page - 1) * limit
	}
	return
}
