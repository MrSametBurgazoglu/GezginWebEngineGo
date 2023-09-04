package StyleEngine

const cssPropertyCount = 148

var cssPropertiesNameList = []string{
	"accent-color",
	"align-content",
	"align-items",
	"align-self",
	"animation",
	"animation-delay",
	"animation-direction",
	"animation-duration",
	"animation-fill-mode",
	"animation-iteration-count",
	"animation-name",
	"animation-play-state",
	"animation-timing-function",
	"backdrop-filter",
	"backface-visibility",
	"background",
	"background-attachment",
	"background-blend-mode",
	"background-clip",
	"background-color",
	"background-image",
	"background-origin",
	"background-position",
	"background-repeat",
	"background-size",
	"border",
	"border-bottom",
	"border-bottom-color",
	"border-bottom-left-radius",
	"border-bottom-right-radius",
	"border-bottom-style",
	"border-bottom-width",
	"border-collapse",
	"border-color",
	"border-image",
	"border-image-outset",
	"border-image-repeat",
	"border-image-slice",
	"border-image-source",
	"border-image-width",
	"border-left",
	"border-left-color",
	"border-left-style",
	"border-left-width",
	"border-radius",
	"border-right",
	"border-right-color",
	"border-right-style",
	"border-right-width",
	"border-spacing",
	"border-style",
	"border-top",
	"border-top-color",
	"border-top-left-radius",
	"border-top-right-radius",
	"border-top-style",
	"border-top-width",
	"border-width",
	"bottom",
	"color",
	"column-count",
	"column-fill",
	"column-gap",
	"column-rule",
	"column-rule-color",
	"column-rule-style",
	"column-rule-width",
	"column-span",
	"column-width",
	"columns",
	"display",
	"flex",
	"flex-basis",
	"flex-direction",
	"flex-flow",
	"flex-grow",
	"flex-shrink",
	"flex-wrap",
	"font",
	"font-family",
	"font-kerning",
	"font-size",
	"font-stretch",
	"font-style",
	"font-variant",
	"font-variant-caps",
	"font-weight",
	"grid-area",
	"grid-auto-columns",
	"grid-auto-flow",
	"grid-auto-rows",
	"grid-column",
	"grid-column-end",
	"grid-column-gap",
	"grid-column-start",
	"grid-gap",
	"grid-row",
	"grid-row-end",
	"grid-row-gap",
	"grid-row-start",
	"grid-template",
	"grid-template-areas",
	"grid-template-columns",
	"grid-template-rows",
	"height",
	"left",
	"margin",
	"margin-bottom",
	"margin-left",
	"margin-right",
	"margin-top",
	"max-height",
	"max-width",
	"min-height",
	"min-width",
	"opacity",
	"outline",
	"outline-color",
	"outline-offset",
	"outline-style",
	"outline-width",
	"overflow",
	"overflow-wrap",
	"overflow-x",
	"overflow-y",
	"padding",
	"padding-bottom",
	"padding-left",
	"padding-right",
	"padding-top",
	"position",
	"resize",
	"right",
	"text-align",
	"text-align-last",
	"text-decoration",
	"text-decoration-color",
	"text-decoration-line",
	"text-decoration-style",
	"text-decoration-thickness",
	"text-indent",
	"text-justify",
	"text-overflow",
	"text-shadow",
	"text-transform",
	"top",
	"visibility",
	"width",
}

type cssPropertyFunction func(styleProperty *StyleProperty, value string)

var functionList = []cssPropertyFunction{
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	BackgroundColorPropertySetValue,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	ColorPropertySetValue,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	DisplayPropertySetValue,
	nil,
	nil,
	FlexDirectionPropertySetValue,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	HeightPropertySetValue,
	LeftPropertySetValue,
	MarginPropertySetValue,
	MarginBottomPropertySetValue,
	MarginLeftPropertyValue,
	MarginRightPropertySetValue,
	MarginTopPropertySetValue,
	MaxHeightPropertySetValue,
	MaxWidthPropertySetValue,
	MinHeightPropertySetValue,
	MinWidthPropertySetValue,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	PaddingPropertySetValue,
	PaddingBottomPropertySetValue,
	PaddingLeftPropertyValue,
	PaddingRightPropertySetValue,
	PaddingTopPropertySetValue,
	PositionPropertySetValue,
	nil,
	RightPropertySetValue,
	TextAlignPropertySetValue,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	TopPropertySetValue,
	nil,
	WidthPropertySetValue,
}
