package globalUtils

import (
	"errors"
	store2 "github.com/micro/go-micro/v2/store"
	"log"
	"time"
)

const cacheDefaultExpiration = 2 * time.Hour
const cacheMaxListLimit = 200

type Cache struct {
	Store          store2.Store
	expiryDuration time.Duration
	databaseName   string
}

func (c *Cache) DatabaseName() string {
	if c.databaseName == "" {
		log.Printf(glErr.CacheDBNameNotSet())
	}
	return c.databaseName
}

func (c *Cache) SetDatabaseName(databaseName string) {
	c.databaseName = databaseName
}

func (c *Cache) ExpiryDuration() time.Duration {
	if c.expiryDuration == 0 {
		c.expiryDuration = cacheDefaultExpiration
	}
	return c.expiryDuration
}

func (c *Cache) SetExpiryDuration(expiryDuration time.Duration) {
	c.expiryDuration = expiryDuration
}

func (c *Cache) checkDBandPrefix(prefix string) error {
	if c.databaseName == "" {
		return errors.New(glErr.CacheDBNameNotSet())
	}
	if prefix == "" {
		return errors.New(glErr.MissingField("cache prefix"))
	}
	return nil
}

func (c *Cache) GetCacheValue(prefix string, key string) (string, error) {

	err := c.checkDBandPrefix(prefix)
	if err != nil {
		return "", err
	}
	value := ""
	//Check if we have the customer in the cache
	prefixOptions := store2.ReadFrom(c.DatabaseName(), prefix)
	rec1, err := c.Store.Read(key, prefixOptions)
	if err != nil {
		log.Printf(glErr.CacheUnableToReadVal(key, err))
		return "", err
	}
	if len(rec1) > 0 {
		value = string(rec1[0].Value)
	}

	return value, nil
}

func (c *Cache) SetCacheValue(prefix string, key string, value string) error {
	err := c.checkDBandPrefix(prefix)
	if err != nil {
		return err
	}
	rec := store2.Record{
		Key:    key,
		Value:  []byte(value),
		Expiry: c.ExpiryDuration(),
	}
	prefixOptions := store2.WriteTo(c.DatabaseName(), prefix)
	err = c.Store.Write(&rec, prefixOptions)
	if err != nil {
		log.Printf(glErr.CacheUnableToWrite(key, err))
		return err
	}
	return nil
}

func (c *Cache) DeleteCacheValue(prefix string, key string) error {
	err := c.checkDBandPrefix(prefix)
	if err != nil {
		return err
	}
	prefixOptions := store2.DeleteFrom(c.DatabaseName(), prefix)
	err = c.Store.Delete(key, prefixOptions)
	if err != nil {
		log.Printf(glErr.CacheUnableToDeleteVal(key, err))
		return err
	}
	return nil
}

func (c *Cache) ListCacheValues(prefix string, numberOfValues uint, offsetValue uint) ([]string, error) {
	err := c.checkDBandPrefix(prefix)
	if err != nil {
		return nil, err
	}
	if numberOfValues > cacheMaxListLimit {
		log.Printf(glErr.CacheTooManyValuesToList(cacheMaxListLimit))
	}
	prefixOptions := store2.ListFrom(c.DatabaseName(), prefix)
	limit := store2.ListLimit(numberOfValues)
	offset := store2.ListOffset(offsetValue)
	myList, err := c.Store.List(limit, offset, prefixOptions)
	if err != nil {
		log.Printf(glErr.CacheListError(err))
		return nil, err
	}
	//log.Printf("list: %v", myList)
	return myList, nil
}
