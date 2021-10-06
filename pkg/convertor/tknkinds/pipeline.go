package tknkinds

import (
	"encoding/json"
	"fmt"
	"strings"

	toto "github.com/in-toto/in-toto-golang/in_toto"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

//PipelineToIntotoLayout :
func PipelineToIntotoLayout(pipeline *v1beta1.Pipeline) ([]byte, error) {
	// meta := toto.Metablock{}
	layout := toto.Layout{}
	layout.Type = pipeline.Name
	for _, task := range pipeline.Spec.Tasks {
		step := toto.Step{}
		step.Type = task.TaskRef.Name
		step.Name = task.Name
		step.ExpectedCommand = parsePipelineTaskCommands(&task)
		step.ExpectedMaterials = parsePipelineTaskMaterial(&task)
		layout.Steps = append(layout.Steps, step)
	}
	return json.MarshalIndent(layout, "", "    ")
}

func parsePipelineTaskCommands(pipelineTask *v1beta1.PipelineTask) []string {
	commands := []string{}

	// for _, cond := range pipelineTask.Conditions{
	// 	cond.
	// }
	for _, raf := range pipelineTask.RunAfter {
		commands = append(commands, fmt.Sprintf("run-after=%s", raf))
	}
	return commands
}

func parsePipelineTaskMaterial(pipelineTask *v1beta1.PipelineTask) [][]string {
	materials := [][]string{}
	inputResources := []string{}
	outputResources := []string{}
	workspaces := []string{}
	if pipelineTask.Resources != nil {
		for _, resIn := range pipelineTask.Resources.Inputs {
			inputResources = append(inputResources, fmt.Sprintf("in-resource-name=%s", resIn.Name),
				fmt.Sprintf("in-resource=%s", resIn.Resource),
				fmt.Sprintf("in-resource-from=%s", strings.Join(resIn.From, ",")))
		}
		materials = append(materials, inputResources)
	}
	if pipelineTask.Resources != nil {
		for _, resOut := range pipelineTask.Resources.Outputs {
			outputResources = append(outputResources, fmt.Sprintf("out-resource-name=%s", resOut.Name),
				fmt.Sprintf("out-resource=%s", resOut.Resource))
		}
		materials = append(materials, outputResources)
	}
	for _, w := range pipelineTask.Workspaces {
		workspaces = append(workspaces, fmt.Sprintf("workspace-name=%s", w.Name),
			fmt.Sprintf("workspace-subpath=%s", w.SubPath),
			fmt.Sprintf("workspace=%s", w.Workspace))
	}
	materials = append(materials, workspaces)
	return materials
}
