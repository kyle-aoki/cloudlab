package syntax

import (
	"cloud/pkg/args"
)

func FindAndExecute() {
	traverse(SyntaxTree)
}

func traverse(commandMap map[string]any) {
	if val, ok := commandMap[args.PollOrEmpty()]; ok {
		switch val.(type) {
		case Command:
			val.(Command).fn()
		default:
			traverse(val.(map[string]any))
		}
	} else {
		helpText()
	}
}
