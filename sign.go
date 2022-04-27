package gosign

import (
	"bytes"
	"crypto/sha1" // nolint
	"io"
	"os"
)

const (
	sumSize  = 20
	FileMode = 0o666
	DirMode  = 0o700
)

type Sign struct {
	file      string
	secret    []byte
	old       []byte
	sum       []byte
	Error     error
	Hasgosign bool
}

func NewSign(file string, secret []byte) *Sign {
	ret := &Sign{file: file, secret: secret}

	fileInfo, err := os.Stat(file)
	if err != nil {
		ret.Error = err

		return ret
	}

	if fileInfo.Size() < sumSize {
		ret.Error = ErrSignFailed

		return ret
	}

	reader, err := os.Open(file)
	if err != nil {
		ret.Error = err

		return ret
	}
	defer reader.Close()

	ret.old = make([]byte, sumSize)
	if _, ret.Error = reader.ReadAt(ret.old, fileInfo.Size()-sumSize); ret.Error != nil {
		return ret
	}

	if _, ret.Error = reader.Seek(0, 0); ret.Error != nil {
		return ret
	}

	containsReader := NewContainsReader(reader, []byte("github.com/xuender/gosign"))
	// nolint
	hash := sha1.New()
	if _, ret.Error = hash.Write(secret); ret.Error != nil {
		return ret
	}

	if _, ret.Error = io.CopyN(hash, containsReader, fileInfo.Size()-sumSize); ret.Error != nil {
		return ret
	}

	if _, ret.Error = hash.Write(secret); ret.Error != nil {
		return ret
	}

	ret.sum = hash.Sum(nil)
	ret.Hasgosign = containsReader.contains

	return ret
}

func (p *Sign) Check() error {
	if bytes.Equal(p.old, p.sum) {
		return nil
	}

	return ErrSignFailed
}

func (p *Sign) Sign() error {
	if p.Check() == nil {
		return ErrSigned
	}

	sum, err := p.Hash()
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(p.file)
	if err != nil {
		return err
	}

	writer, err := os.OpenFile(p.file, os.O_RDWR|os.O_APPEND, FileMode)
	if err != nil {
		return err
	}
	defer writer.Close()

	if _, err = writer.Seek(fileInfo.Size(), 0); err != nil {
		return err
	}

	_, err = writer.Write(sum)

	return err
}

func (p *Sign) Hash() ([]byte, error) {
	reader, err := os.Open(p.file)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	// nolint
	hash := sha1.New()

	if _, err := hash.Write(p.secret); err != nil {
		return nil, err
	}

	if _, err := io.Copy(hash, reader); err != nil {
		return nil, err
	}

	if _, err := hash.Write(p.secret); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
