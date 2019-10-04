package dom

func SetStyle(node Node, key, value string) {
	SetStyles(node, map[string]string{key: value})
}

func SetStyles(node Node, styles map[string]string) {
	styleObj := node.exposeObject().Get("style")
	for key := range styles {
		styleObj.Set(key, styles[key])
	}
}
