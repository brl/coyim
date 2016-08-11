package otr3

import (
	"crypto/rand"
	"io"
	"math/big"
)

func (c *Conversation) rand() io.Reader {
	if c.Rand != nil {
		return c.Rand
	}
	return rand.Reader
}

func randomInto(r io.Reader, b []byte) error {
	if _, err := io.ReadFull(r, b); err != nil {
		return errShortRandomRead
	}
	return nil
}

func randMPI(r io.Reader, b []byte) (*big.Int, error) {
	if err := randomInto(r, b); err != nil {
		return nil, err
	}

	return new(big.Int).SetBytes(b), nil
}

func randSizedMPI(r io.Reader, size int) (*big.Int, error) {
	return randMPI(r, make([]byte, size))
}

func (c *Conversation) randSizedMPI(size int) (*big.Int, error) {
	return randMPI(c.rand(), make([]byte, size))
}

func (c *Conversation) randMPI(buf []byte) (*big.Int, error) {
	return randMPI(c.rand(), buf)
}

func (c *Conversation) randomInto(b []byte) error {
	return randomInto(c.rand(), b)
}
