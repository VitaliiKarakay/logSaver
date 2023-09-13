package store_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"logSaver/pkg/store"
)

type StoreSuite struct {
	suite.Suite
	Store *store.DB
}

func RunSuite(t *testing.T) {
	suite.Run(t, new(StoreSuite))
}
