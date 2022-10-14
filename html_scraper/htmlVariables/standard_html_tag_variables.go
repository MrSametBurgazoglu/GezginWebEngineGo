package htmlVariables

type StandardHtmlTagVariables struct {
	ContentEditable bool
	Draggable       bool
	Hidden          bool
	AccessKey       string
	//dir
	Class []string
	Id    string
	Style string
}

func (tag *StandardHtmlTagVariables) SetStandardVariables(name string, value string) bool {
	switch name {
	case "id":
		tag.Id = value
	case "class":
		tag.Class = append(tag.Class, value)
	case "style":
		tag.Style = value
	case "contenteditable":
		tag.ContentEditable = value != "false"
	case "draggable":
		tag.Draggable = value != "false"
	default:
		return false
	}
	return true
}

func (tag *StandardHtmlTagVariables) SetStandardContextVariables(context string) bool {
	switch context {
	case "contenteditable":
		tag.ContentEditable = true
	default:
		return false
	}
	return true
}
