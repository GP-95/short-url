package url

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/GP-95/short-url/internal/cache"
	"github.com/GP-95/short-url/internal/db"
)

func addNewUrl(c context.Context, url string) (string, error) {
	ok, hash, err := checkIfUrlExists(url)
	if err != nil {
		return "", err
	}

	if ok {
		slog.Info("Url already exists: " + url)
		return hash, nil
	}

	hash, err = createUniqueCode(url)
	if err != nil {
		return "", err
	}

	err = db.DB.SaveUrlAndHash(url, hash)
	if err != nil {
		return "", err
	}

	err = cache.REDIS.SaveCodeAndUrl(c, hash, url)
	if err != nil {
		slog.Warn(err.Error())
	}

	return hash, nil
}

func getCodeUrl(c context.Context, code string) (string, error) {
	url, err := cache.REDIS.GetUrlByCode(c, code)
	if err != nil {
		slog.Info(err.Error())
	}

	if len(url) != 0 {
		slog.Info("URL was already in cache: " + url)
		cache.REDIS.ResetCodeTTL(c, code)
		return url, nil
	}

	url, err = db.DB.FindUrlByHash(code)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	if len(url) != 0 {
		err = cache.REDIS.SaveCodeAndUrl(c, code, url)
		if err != nil {
			slog.Warn(err.Error())
		}
	}

	return url, nil
}

func checkIfUrlExists(url string) (bool, string, error) {
	// Here we could also check cache first, if something like RediSearch was used
	hash, err := db.DB.FindHashByUrl(url)

	if err != nil {
		return false, "", err
	}

	return len(hash) != 0, hash, nil
}
