package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

func main() {
	// 创建 Redis 客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 密码
		DB:       0,                // 使用默认的数据库
	})

	// 锁的键名
	lockKey := "mylock"

	// 设置加锁超时时间为 10 秒
	lockTimeout := 30 * time.Second

	lockAcquired, err := acquireLockWithTimeoutAndRetry(client, lockKey, lockTimeout)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if lockAcquired {
		// 成功获取锁，执行需要保护的代码
		fmt.Println("Lock acquired successfully")

		// 设置锁的超时时间
		err := setLockTimeout(client, lockKey, lockTimeout)
		if err != nil {
			fmt.Println("Error setting lock timeout:", err)
		}

		time.Sleep(10 * time.Second)

		defer releaseLock(client, lockKey)
	} else {
		// 获取锁失败，进行相应的错误处理
		fmt.Println("Failed to acquire lock. Please try again later.")
	}
}

// acquireLockWithTimeoutAndRetry 尝试获取锁，设置加锁超时时间，自旋等待
func acquireLockWithTimeoutAndRetry(client *redis.Client, lockKey string, timeout time.Duration) (bool, error) {
	ctx := context.Background()

	// 计算超时时间点
	deadline := time.Now().Add(timeout)

	for {
		// 使用 SETNX 尝试获取锁
		lockAcquired, err := client.SetNX(ctx, lockKey, "1", 0).Result()
		if err != nil {
			return false, err
		}

		// 如果成功获取锁，返回 true
		if lockAcquired {
			return true, nil
		}

		// 检查是否超时
		if time.Now().After(deadline) {
			break
		}

		// 等待一段时间后再重试
		time.Sleep(1 * time.Second) // 可以设置等待的间隔
	}

	// 获取锁失败，超时
	return false, nil
}

// setLockTimeout 设置锁的超时时间
func setLockTimeout(client *redis.Client, lockKey string, timeout time.Duration) error {
	ctx := context.Background()

	// 设置锁的超时时间
	_, err := client.Expire(ctx, lockKey, timeout).Result()
	return err
}

// releaseLock 释放锁
func releaseLock(client *redis.Client, lockKey string) error {
	ctx := context.Background()
	_, err := client.Del(ctx, lockKey).Result()
	return err
}
