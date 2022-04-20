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
	suite.localCache = New().(*LocalCache)
	suite.key = "key"
	suite.value = cacheValue{name: "John"}
}

func (suite *LocalCacheTestSuite) TestLocalCache_Set() {
	suite.localCache.(*LocalCache).store[suite.key] = CacheItem{value: suite.value, expireTimer: nil}

	got, _ := suite.localCache.Get(suite.key)

	assert.Equal(suite.T(), suite.value, got)
}

func (suite *LocalCacheTestSuite) TestLocalCache_Get() {
	err := suite.localCache.Set(suite.key, suite.value)

	assert.Nil(suite.T(), err)
}

func TestTestLocalCacheSuite(t *testing.T) {
	suite.Run(t, new(LocalCacheTestSuite))
}
