package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/magiconair/properties/assert"
)

func TestDockerHelloWorldExample(t *testing.T) {
	// website::tag::1:: Configure the tag to use on the Docker image.
	tag := "hello_world_app"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// website::tag::2:: Build the Docker image.
	docker.Build(t, "../hello_world_app/docker-hello-world-example", buildOptions)

	// website::tag::3:: Run the Docker image, read the text file from it, and make sure it contains the expected output.
	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Hello, World!", output)
}
