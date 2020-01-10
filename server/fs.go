package server

import (
	"os"
	"path/filepath"
)

type FST struct {
	Name string
	Path *string
}

func (srv *Server) IterOnFileSystemPaths(mapFunc func(fst FST) bool) bool {
	for _, cfg := range []FST{} {
		if !mapFunc(cfg) {
			return false
		}
	}
	return true
}

func (srv *Server) PrepareFileSystem() bool {
	return srv.IterOnFileSystemPaths(func(cfg FST) bool {
		var err error
		if *cfg.Path, err = filepath.Abs(*cfg.Path); err != nil {
			srv.Logger.Debug("find absolute path error", "error", err, "creating-path", *cfg.Path)
			return false
		}
		*cfg.Path += "/"
		if err = os.MkdirAll(*cfg.Path, os.ModePerm); err != nil {
			srv.Logger.Debug("create path directory error", "error", err, "creating-path", *cfg.Path)
			return false
		}

		srv.Logger.Info("host path", "name", cfg.Name, "creating-path", *cfg.Path)
		return true
	})
}

func (srv *Server) DropFileSystem() bool {
	return srv.IterOnFileSystemPaths(func(cfg FST) bool {
		if err := os.RemoveAll(*cfg.Path); err != nil {
			srv.Logger.Debug("drop path directory error", "error", err, "dropping-path", *cfg.Path)
			return false
		}
		return true
	})
}
