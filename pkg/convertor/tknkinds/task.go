package tknkinds

import (
	"encoding/json"
	"fmt"
	"strings"

	toto "github.com/in-toto/in-toto-golang/in_toto"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

const (
	stepCmdKeyword      = "command"
	stepArgKeyword      = "arguments"
	stepImgKeyword      = "image"
	stepImgPullPolicy   = "imagepullpolicy"
	stepSecurityContext = "securitycontext"
	stepScriptKeyword   = "script"

	spaceDelimiter = " "
)

//TaskToIntotoLayout :
func TaskToIntotoLayout(task *v1beta1.Task) ([]byte, error) {
	// meta := toto.Metablock{}

	layout := toto.Layout{}
	layout.Type = task.Name
	for _, s := range task.TaskSpec().Steps {
		step := toto.Step{}
		step.Type = s.Name
		step.Name = s.Name
		step.ExpectedCommand = parseStepCmds(&s)
		step.ExpectedMaterials = parseStepMaterial(&s)
		layout.Steps = append(layout.Steps, step)
	}
	return json.MarshalIndent(layout, "", "    ")
}

func parseStepCmds(step *v1beta1.Step) []string {
	cmds := []string{}
	if len(step.Command) != 0 {
		v := fmt.Sprintf("%s : %s", stepCmdKeyword, strings.Join(step.Command, spaceDelimiter))
		cmds = append(cmds, v)
	}
	if step.Script != "" {
		v := fmt.Sprintf("%s: %s", stepScriptKeyword, step.Script)
		cmds = append(cmds, v)
	}
	if len(step.Args) != 0 {
		v := fmt.Sprintf("%s : %s", stepArgKeyword, strings.Join(step.Args, spaceDelimiter))
		cmds = append(cmds, v)
	}
	return cmds
}

func parseStepMaterial(step *v1beta1.Step) [][]string {
	materials := [][]string{}
	v := fmt.Sprintf("%s : %s", stepImgKeyword, step.Image)
	imgMaterial := []string{}

	imgMaterial = append(imgMaterial, v)

	if step.ImagePullPolicy != "" {
		v = fmt.Sprintf("%s : %s", stepImgPullPolicy, step.ImagePullPolicy)
		imgMaterial = append(imgMaterial, v)
	}

	materials = append(materials, imgMaterial)
	if step.SecurityContext != nil {
		secCtxMaterial := []string{}
		if b, err := step.SecurityContext.Marshal(); err != nil {
			v = fmt.Sprintf("%s : %s", stepSecurityContext, string(b))
			secCtxMaterial = append(secCtxMaterial, v)
			materials = append(materials, secCtxMaterial)
		}
	}
	return materials
}
