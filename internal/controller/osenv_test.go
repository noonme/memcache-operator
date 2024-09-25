package controller

import (
	"os"
	"testing"
)

func TestImageForMemcached(t *testing.T) {
	// 模拟环境变量不存在的情况
	os.Unsetenv("MEMCACHED_IMAGE1")
	_, err := imageForMemcached()
	if err == nil {
		t.Errorf("Expected error when environment variable is not set, but got nil")
	}

	// 模拟环境变量存在的情况
	os.Setenv("MEMCACHED_IMAGE1", "redis")
	image, err := imageForMemcached()
	if err != nil {
		t.Errorf("Expected no error when environment variable is set, but got: %v", err)
	}
	if image != "redis" {
		t.Errorf("Expected image to be'redis', but got: %s", image)
	}
}
