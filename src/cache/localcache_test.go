package localcache

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LocalCacheTestSuite struct {
	suite.Suite
	localCache Cache
	key        string
	value      cacheValue
}

type cacheValue struct {
	name string
}

func (suite *LocalCacheTestSuite) SetupTest() {
	suite.localCache = New()
}

func (suite *LocalCacheTestSuite) TestLocalCache_Set() {
	key := "key"

	_ = suite.localCache.Set(key, cacheValue{name: "John"})
	got, _ := suite.localCache.Get(key)

	assert.Equal(suite.T(), cacheValue{name: "John"}, got)
}

func (suite *LocalCacheTestSuite) TestLocalCache_Get() {
	key := "key"

	error := suite.localCache.Set(key, cacheValue{name: "John"})

	assert.Nil(suite.T(), error)
}

func TestTestLocalCacheSuite(t *testing.T) {
	suite.Run(t, new(LocalCacheTestSuite))
}
