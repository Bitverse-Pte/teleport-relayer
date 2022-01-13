package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type CacheFileWriter struct {
	homeDir       string
	cacheDir      string
	cacheFilename string
}

func NewCacheFileWriter(homeDir, cacheDir, cacheFilename string) *CacheFileWriter {
	return &CacheFileWriter{
		homeDir:       homeDir,
		cacheDir:      cacheDir,
		cacheFilename: cacheFilename,
	}
}

func (w *CacheFileWriter) Write(height uint64) error {
	cacheDataObj := &Data{}
	cacheDataObj.LatestHeight = height

	cacheDataWriteBytes, err := json.Marshal(cacheDataObj)
	if err != nil {
		return err
	}

	cacheDir := path.Join(w.homeDir, w.cacheDir)
	filename := path.Join(cacheDir, w.cacheFilename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// And the home folder doesn't exist
		if _, err := os.Stat(w.homeDir); os.IsNotExist(err) {
			// Create the home folder
			if err = os.Mkdir(w.homeDir, os.ModePerm); err != nil {
				return err
			}
		}
		// Create the home config folder
		if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
			// Create the home folder
			if err = os.Mkdir(cacheDir, os.ModePerm); err != nil {
				return err
			}
		}
		// Then create the file...
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}

	} else {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}
	}
	return nil
}

func (w *CacheFileWriter) LoadCache() *Data {
	// If the file exists, the initial height is the latest_height in the file
	filename := path.Join(w.homeDir, w.cacheDir, w.cacheFilename)
	file, err := os.Open(filename)
	if err != nil {
		// logger.Fatal("read cache file err: ", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		// logger.Fatal("read cache file err: ", err)
	}

	cacheData := &Data{}
	if err = json.Unmarshal(content, cacheData); err != nil {
		// logger.Fatal("read cache file unmarshal err: ", err)
	}

	return cacheData
}
