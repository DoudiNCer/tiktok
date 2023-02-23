package common

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var CacheManager = cache.New(24*time.Hour, time.Duration(time.Monday))
var KeyAddUser = "key_add_user"
