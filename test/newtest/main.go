package main

import (
	"fmt"
	"os"
)

func imageForMemcached() (string, error) {

	var imageEnvVar = "MEMCACHED_IMAG"
	// os.Setenv("MEMCACHED_IMAGE1", "redis")
	image, found := os.LookupEnv(imageEnvVar)
	if !found {
		return "", fmt.Errorf("Unable to find %s environment variable with the image", imageEnvVar)
	}
	return image, nil
}
func main() {
	os.Setenv("MEMCACHED_IMAG", "redis")
	image, err := imageForMemcached()
	fmt.Println(image, err)
}
