package convertor

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/tkn-intoto-formatter/pkg/common"
	"github.com/tkn-intoto-formatter/pkg/convertor/tknkinds"
	"github.com/tkn-intoto-formatter/pkg/parser"
)

//ConvertToIntoto :
func ConvertToIntoto(cOpt common.ConvertOpt) error {

	if cOpt.InputFilepath != "" {
		resources, _ := parser.GetTknResources(cOpt.InputFilepath)
		for _, t := range resources.TaskSpecs {
			taskLayout, err := tknkinds.TaskToIntotoLayout(t)
			if err != nil {
				return errors.Wrapf(err, "error creating intoto layout for task `%s` ", t.Name)
			}
			err = ioutil.WriteFile(cOpt.OutputFilepath, taskLayout, 0655)
			if err != nil {
				return errors.Wrapf(err, "error creating intoto layout for task `%s` ", t.Name)
			}
		}

		for _, p := range resources.PipelineSpecs {
			pipelineLayout, err := tknkinds.PipelineToIntotoLayout(p)
			if err != nil {
				return errors.Wrapf(err, "error creating intoto layout for pipeline `%s` ", p.Name)
			}
			err = ioutil.WriteFile(cOpt.OutputFilepath, pipelineLayout, 0655)
			if err != nil {
				return errors.Wrapf(err, "error creating intoto layout for task `%s` ", p.Name)
			}
		}
	}

	return nil
}
