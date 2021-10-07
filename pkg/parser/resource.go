package parser

import (
	"fmt"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
)

// parseTknPipeline parses pipeline
func parseTknPipeline(r []byte) (*v1beta1.Pipeline, error) {
	var pipeline v1beta1.Pipeline
	if r == nil {
		return nil, nil
	}

	_, _, err := scheme.Codecs.UniversalDeserializer().Decode(r, nil, &pipeline)
	if err != nil {
		fmt.Printf("error parsing `pipeline' object: %v", err)
		return nil, err
	}

	return &pipeline, nil
}

// parseTknTask parses Task
func parseTknTask(r []byte) (*v1beta1.Task, error) {
	if r == nil {
		return nil, nil
	}
	var task v1beta1.Task
	_, _, err := scheme.Codecs.UniversalDeserializer().Decode(r, nil, &task)
	if err != nil {
		fmt.Printf("error parsing `task' object: %v", err)
		return nil, err
	}
	return &task, nil
}
