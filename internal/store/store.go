package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/groovy-sky/aws-docs/internal/model"
	bolt "go.etcd.io/bbolt"
)

var (
	pagesBucket = []byte("pages")
	linksBucket = []byte("links")
)

type Store struct {
	db *bolt.DB
}

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("create metadata directory: %w", err)
	}
	database, err := bolt.Open(path, 0o600, nil)
	if err != nil {
		return nil, fmt.Errorf("open metadata db: %w", err)
	}
	store := &Store{db: database}
	if err := store.db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(pagesBucket); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists(linksBucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		_ = database.Close()
		return nil, fmt.Errorf("initialize metadata db: %w", err)
	}
	return store, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) GetPage(rawURL string) (*model.PageRecord, error) {
	var record *model.PageRecord
	err := s.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(pagesBucket).Get([]byte(rawURL))
		if value == nil {
			return nil
		}
		decoded := model.PageRecord{}
		if err := json.Unmarshal(value, &decoded); err != nil {
			return err
		}
		record = &decoded
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("get page %s: %w", rawURL, err)
	}
	return record, nil
}

func (s *Store) PutPage(record model.PageRecord) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		encoded, err := json.Marshal(record)
		if err != nil {
			return err
		}
		return tx.Bucket(pagesBucket).Put([]byte(record.URL), encoded)
	})
}

func (s *Store) GetLinks(rawURL string) ([]string, error) {
	var links []string
	err := s.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(linksBucket).Get([]byte(rawURL))
		if value == nil {
			return nil
		}
		return json.Unmarshal(value, &links)
	})
	if err != nil {
		return nil, fmt.Errorf("get links for %s: %w", rawURL, err)
	}
	return links, nil
}

func (s *Store) PutLinks(rawURL string, links []string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		encoded, err := json.Marshal(links)
		if err != nil {
			return err
		}
		return tx.Bucket(linksBucket).Put([]byte(rawURL), encoded)
	})
}
