package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	bolt "go.etcd.io/bbolt"
)

var (
	seedsBucket = []byte("seeds")
	seedsKey    = []byte("all")
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
		_ = tx.DeleteBucket([]byte("pages"))
		_ = tx.DeleteBucket([]byte("links"))
		if _, err := tx.CreateBucketIfNotExists(seedsBucket); err != nil {
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

func (s *Store) GetSeeds() ([]string, error) {
	var seeds []string
	err := s.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(seedsBucket).Get(seedsKey)
		if value == nil {
			return nil
		}
		return json.Unmarshal(value, &seeds)
	})
	if err != nil {
		return nil, fmt.Errorf("get seeds: %w", err)
	}
	return seeds, nil
}

func (s *Store) PutSeeds(seeds []string) error {
	unique := make([]string, 0, len(seeds))
	seen := map[string]struct{}{}
	for _, seed := range seeds {
		if _, exists := seen[seed]; exists {
			continue
		}
		seen[seed] = struct{}{}
		unique = append(unique, seed)
	}
	sort.Strings(unique)

	return s.db.Update(func(tx *bolt.Tx) error {
		encoded, err := json.Marshal(unique)
		if err != nil {
			return err
		}
		return tx.Bucket(seedsBucket).Put(seedsKey, encoded)
	})
}
