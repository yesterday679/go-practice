package cache_service

import (
	"learn/gin_blog/pkg/errors"
	"strconv"
	"strings"
)

type Article struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return errors.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		errors.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.TagID))
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
