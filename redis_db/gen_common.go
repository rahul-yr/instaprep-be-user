package redisdb

import (
	"context"
)

var global_ctx = context.Background()
const TIMEOUT_CONTEXT = 5