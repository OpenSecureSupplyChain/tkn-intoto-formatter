package common

import (
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

//ConvertOpt :
type ConvertOpt struct {
	InputFilepath  string
	OutputFilepath string
	Format         string
	StubVal        bool
}

//TknObject :
type TknObject struct {
	GroupKind  string
	ObjectSpec []byte
}

//TknResources :
type TknResources struct {
	PipelineSpecs []*v1beta1.Pipeline
	TaskSpecs     []*v1beta1.Task
}

const (
	//SignAnnotationKey :
	SignAnnotationKey = "ossc.sigstore.tapestry.dev/transparency"
)
