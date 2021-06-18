package redisdb

import (
	"log"
	"os"
	"sync"
	"context"
	"time"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

var (
	redisOneClient  *redis.Client
	redisOneOnce sync.Once
)

type RedisOneOps struct{
	// TODO
}

// This method is responsible for creating a singleton redis instance
// Required environment varaibles are
// 		{host_env}
// 		{password_env}
func GetRedisOneInstance() *redis.Client {
	redisOneOnce.Do(func() {
		redisOneClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOSTNAME"),
			Password: os.Getenv("REDIS_PASSWORD"), // no password set
			DB:       0,                           // use default DB
		})
		log.Println("[redis_operations] : created singleton redisOne database instance..")
		// <-- thread safe
	})
	return redisOneClient
}


// This method only store data in redis and returns nil if succesful else returns error object
func (helpers *RedisOneOps) StoreKV(key string, value string, ttl time.Duration) error {
	// get the redis instance
	rdb := GetRedisOneInstance()
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	err := rdb.Set(timeout, key, value, ttl).Err()
	// check operation errors
	if err != nil {
		return err
	}
	// return nil if no error
	return nil
}

// This method only stores composite data structure in redis and returns nil if succesful else returns error object
func (helpers *RedisOneOps) StoreJSON(key string, value interface{}, ttl time.Duration) error {
	// get the redis instance
	rdb := GetRedisOneInstance()
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rdb.Set(timeout, key, jsonBytes, ttl).Err()
	// check operation errors
	if err != nil {
		return err
	}
	// return nil if no error
	return nil
}

// This method returns the value if key present in redis else returns error object.
func (helpers *RedisOneOps) GetKV(key string) (string, error) {
	// get the redis instance
	rdb := GetRedisOneInstance()
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	val2, err := rdb.Get(timeout, key).Result()
	if err != nil {
		// It can be one of two cases.
		// Either key is not present or some other operation exception
		return "", err
	}
	return val2, nil
}

// This method returns the byte value of string if key present in redis else returns error object.
//
// needs to unmarshal it before using it
func (helpers *RedisOneOps) GetJSON(key string, v interface{}) error {
	// get the redis instance
	rdb := GetRedisOneInstance()
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	val2, err := rdb.Get(timeout, key).Result()
	if err != nil {
		// It can be one of two cases.
		// Either key is not present or some other operation exception
		return err
	}
	err = json.Unmarshal([]byte(val2), v)
	if err != nil {
		return err
	}
	return nil
}

// This method returns the value if key present in redis else returns ""
// returns error obejct if any other error
func (helpers *RedisOneOps) DeleteKey(key string) error {
	// get the redis instance
	rdb := GetRedisOneInstance()
	// common code
	timeout, cancel := context.WithTimeout(global_ctx, TIMEOUT_CONTEXT*time.Second)
	defer cancel()
	// end
	err := rdb.Del(timeout, key).Err()
	if err == redis.Nil {
		return nil
	} else if err == nil {
		return nil
	} else {
		return err
	}
}


