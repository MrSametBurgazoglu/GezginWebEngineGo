package standardHtmlTagVariables

import (
	"gezgin_web_engine/eventSystem"
	"strconv"
)

type StandardHtmlTagVariables struct {
	ContentEditable bool
	Draggable       bool
	Hidden          bool
	Lang            string
	AccessKey       string
	//dir
	Class    []string
	Id       string
	Style    string
	Title    string
	OnClick  string
	TabIndex int
}

func (tag *StandardHtmlTagVariables) SetStandardVariables(name string, value string, inputWidget eventSystem.InputWidget) bool {
	switch name {
	case "id":
		tag.Id = value[1 : len(value)-1]
	case "class":
		tag.Class = append(tag.Class, value[1:len(value)-1])
	case "style":
		tag.Style = value[1 : len(value)-1]
	case "contenteditable":
		tag.ContentEditable = value != "false"
	case "draggable":
		tag.Draggable = value != "false"
	case "hidden":
		tag.Hidden = value != "false"
	case "lang":
		tag.Lang = value
	case "tabindex":
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			tag.TabIndex = valueInt
		}
	case "title":
		tag.Title = value
	case "onclick":
		tag.OnClick = value[1 : len(value)-1]
		eventSystem.SetInput("onclick", inputWidget)
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

func (tag *StandardHtmlTagVariables) GetOnclick() string {
	return tag.OnClick
}
