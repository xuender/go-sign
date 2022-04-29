package sign

import (
	"bytes"
	"crypto/sha1" // nolint
	"io"
	"os"
)

const (
	sumSize  = 20
	fileMode = 0o666
)

type Sign struct {
	file    string
	secret  []byte
	old     []byte
	sum     []byte
	Error   error
	HasSign bool
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

	_, _ = reader.Seek(0, 0)

	containsReader := NewContainsReader(reader, []byte(Mod.Path))
	// nolint
	hash := sha1.New()
	_, _ = hash.Write(secret)

	if _, ret.Error = io.CopyN(hash, containsReader, fileInfo.Size()-sumSize); ret.Error != nil {
		return ret
	}

	_, _ = hash.Write(secret)

	ret.sum = hash.Sum(nil)
	ret.HasSign = containsReader.contains

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

	file, err := os.OpenFile(p.file, os.O_RDWR|os.O_APPEND, fileMode)
	if err != nil {
		return err
	}
	defer file.Close()

	sum, err := p.Hash(file)
	if err != nil {
		return err
	}

	_, err = file.Write(sum)

	return err
}

func (p *Sign) Hash(reader io.Reader) ([]byte, error) {
	// nolint
	hash := sha1.New()
	_, _ = hash.Write(p.secret)

	if _, err := io.Copy(hash, reader); err != nil {
		return nil, err
	}

	_, _ = hash.Write(p.secret)

	return hash.Sum(nil), nil
}
