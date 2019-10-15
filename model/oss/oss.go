package oss

import (
	"errors"
	"github.com/tecbot/gorocksdb"
)

var engine Engine

var ErrNotExist = errors.New("object not exists")

type ByteObject interface {
	Data() []byte
	Free()
}

type Engine interface {
	Get([]byte) (ByteObject, error)
	Put([]byte, []byte) error
	Delete([]byte) error
	Close() error
}

type RocksDBEngine struct {
	DB *gorocksdb.DB
	wOpt *gorocksdb.WriteOptions
	rOpt *gorocksdb.ReadOptions
}

func (e *RocksDBEngine) Get(b []byte) (ByteObject, error) {
	obj, err := e.DB.Get(e.rOpt, b)
	if err != nil {
		return nil, err
	}
	if !obj.Exists() {
		return nil, ErrNotExist
	}
	return obj, nil
}

func (e *RocksDBEngine) Put(k []byte, v []byte) error {
	return e.DB.Put(e.wOpt, k, v)
}

func (e *RocksDBEngine) Delete(k []byte) error {
	return e.DB.Delete(e.wOpt, k)
}

func (e *RocksDBEngine) Close() error {
	e.DB.Close()
	return nil
}

func NewRocksDB(opts *gorocksdb.Options, name string) (Engine, error) {
	e := new(RocksDBEngine)
	var err error
	e.DB, err = gorocksdb.OpenDb(opts, name)
	if err != nil {
		return nil, err
	}
	e.wOpt = gorocksdb.NewDefaultWriteOptions()
	e.rOpt = gorocksdb.NewDefaultReadOptions()
	return e, nil
}

func RegisterEngine(e Engine) error {
	engine = e
	return nil
}







