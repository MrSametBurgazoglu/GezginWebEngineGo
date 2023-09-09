package enums

type CssAnimationTimingFunctionType uint8
type CssAnimationDirectionType uint8
type CssAnimationFillModeType uint8
type CssAnimationPlayStateType uint8
type CssFilterType uint8
type CssAlignType uint8
type CssTextAlignType uint8
type CssFloatType uint8
type CssVisibilityType uint8
type CssBackgroundType uint8
type CssBackgroundPositionType uint8
type CssBackgroundRepeatType uint8
type CssBackgroundSizeType uint8
type CssBackgroundAttachmentType uint8
type CssBackgroundBlendModeType uint8
type CssBackgroundClipType uint8
type CssBackgroundOriginType uint8
type CssBorderlineType uint8
type CssBorderStyleType uint8
type CssBorderRadiusType uint8
type CssBorderCollapseType uint8
type CssBorderImageSourceType uint8
type CssBorderImageSliceType uint8
type CssBorderImageWidthType uint8
type CssBorderImageOutsetType uint8
type CssBorderImageRepeatType uint8
type CssPropertyValueType uint8
type CssBoxDecorationType uint8
type CssBoxSizingType uint8
type CssDisplayType uint8
type CssPositionType uint8
type CssFontSizeType uint8
type CssFlexDirectionType uint8

const (
	CSS_ANIMATION_TIMING_FUNCTION_LINEAR CssAnimationTimingFunctionType = iota
	CSS_ANIMATION_TIMING_FUNCTION_EASE
	CSS_ANIMATION_TIMING_FUNCTION_EASE_IN
	CSS_ANIMATION_TIMING_FUNCTION_EASE_OUT
	CSS_ANIMATION_TIMING_FUNCTION_EASE_IN_OUT
	CSS_ANIMATION_TIMING_FUNCTION_STEP_START
	CSS_ANIMATION_TIMING_FUNCTION_STEP_END
)

const (
	CSS_ANIMATION_DIRECTION_NORMAL CssAnimationDirectionType = iota
	CSS_ANIMATION_DIRECTION_REVERSE
	CSS_ANIMATION_DIRECTION_ALTERNATE
	CSS_ANIMATION_DIRECTION_ALTERNATE_REVERSE
)

const (
	CSS_ANIMATION_FILL_MODE_NONE CssAnimationFillModeType = iota
	CSS_ANIMATION_FILL_MODE_FORWARDS
	CSS_ANIMATION_FILL_MODE_BACKWARDS
	CSS_ANIMATION_FILL_MODE_BOTH
)

const (
	CSS_ANIMATION_PLAY_STATE_PAUSED CssAnimationPlayStateType = iota
	CSS_ANIMATION_PLAY_STATE_RUNNING
)

const (
	CSS_ALIGN_EMPTY CssAlignType = iota
	CSS_ALIGN_STRETCH
	CSS_ALIGN_CENTER
	CSS_ALIGN_FLEX_START
	CSS_ALIGN_FLEX_END
	CSS_ALIGN_SPACE_BETWEEN
	CSS_ALIGN_SPACE_AROUND
	CSS_ALIGN_SPACE_EVENLY
	CSS_ALIGN_BASELINE
	CSS_ALIGN_AUTO
)

const (
	CSS_TEXT_ALIGN_EMPTY CssTextAlignType = iota
	CSS_TEXT_ALIGN_CENTER
	CSS_TEXT_ALIGN_JUSTIFY
	CSS_TEXT_ALIGN_LEFT
	CSS_TEXT_ALIGN_RIGHT
)

const (
	CSS_FLOAT_EMPTY CssFloatType = iota
	CSS_FLOAT_LEFT
	CSS_FLOAT_NONE
	CSS_FLOAT_RIGHT
)

const (
	CSS_VISIBILITY_EMPTY CssVisibilityType = iota
	CSS_VISIBILITY_VISIBLE
	CSS_VISIBILITY_HIDDEN
	CSS_VISIBILITY_COLLAPSE
)

const (
	CSS_BACKGROUND_TYPE_URL CssBackgroundType = iota
	CSS_BACKGROUND_TYPE_CONIC_GRADIENT
	CSS_BACKGROUND_TYPE_LINEAR_GRADIENT
	CSS_BACKGROUND_TYPE_RADIAL_GRADIENT
	CSS_BACKGROUND_TYPE_REPEATING_CONIC_GRADIENT
	CSS_BACKGROUND_TYPE_REPEATING_LINEAR_GRADIENT
	CSS_BACKGROUND_TYPE_REPEATING_RADIAL_GRADIENT
)

const (
	CSS_BACKGROUND_POSITION_TYPE_WORD CssBackgroundPositionType = iota
	CSS_BACKGROUND_POSITION_TYPE_PERCENT
	CSS_BACKGROUND_POSITION_TYPE_POS
)

const (
	CSS_BACKGROUND_REPEAT_TYPE_REPEAT CssBackgroundRepeatType = iota
	CSS_BACKGROUND_REPEAT_TYPE_REPEAT_X
	CSS_BACKGROUND_REPEAT_TYPE_REPEAT_Y
	CSS_BACKGROUND_REPEAT_TYPE_NO_REPEAT
	CSS_BACKGROUND_REPEAT_TYPE_ROUND
	CSS_BACKGROUND_REPEAT_TYPE_SPACE
)

const (
	CSS_BACKGROUND_SIZE_TYPE_AUTO CssBackgroundSizeType = iota
	CSS_BACKGROUND_SIZE_TYPE_LENGTH
	CSS_BACKGROUND_SIZE_TYPE_PERCENTAGE
	CSS_BACKGROUND_SIZE_TYPE_COVER
	CSS_BACKGROUND_SIZE_TYPE_CONTAIN
)

const (
	CSS_BACKGROUND_ATTACHMENT_SCROLL CssBackgroundAttachmentType = iota
	CSS_BACKGROUND_ATTACHMENT_FIXED
	CSS_BACKGROUND_ATTACHMENT_LOCAL
)

const (
	CSS_BACKGROUND_BLEND_MODE_NORMAL CssBackgroundBlendModeType = iota
	CSS_BACKGROUND_BLEND_MODE_MULTIPLY
	CSS_BACKGROUND_BLEND_MODE_SCREEN
	CSS_BACKGROUND_BLEND_MODE_OVERLAY
	CSS_BACKGROUND_BLEND_MODE_DARKEN
	CSS_BACKGROUND_BLEND_MODE_LIGHTEN
	CSS_BACKGROUND_BLEND_MODE_COLOR_DODGE
	CSS_BACKGROUND_BLEND_MODE_SATURATION
	CSS_BACKGROUND_BLEND_MODE_COLOR
	CSS_BACKGROUND_BLEND_MODE_LUMINOSITY
)

const (
	CSS_BACKGROUND_CLIP_BORDER_BOX CssBackgroundClipType = iota
	CSS_BACKGROUND_CLIP_PADDING_BOX
	CSS_BACKGROUND_CLIP_CONTENT_BOX
)

const (
	CSS_BACKGROUND_ORIGIN_PADDING_BOX CssBackgroundOriginType = iota
	CSS_BACKGROUND_ORIGIN_BORDER_BOX
	CSS_BACKGROUND_ORIGIN_CONTENT_BOX
)

const (
	CSS_BORDER_LINE_TYPE_MEDIUM CssBorderlineType = iota
	CSS_BORDER_LINE_TYPE_THIN
	CSS_BORDER_LINE_TYPE_THICK
	CSS_BORDER_LINE_TYPE_LENGTH
)

const (
	CSS_BORDER_STYLE_TYPE_NONE CssBorderStyleType = iota
	CSS_BORDER_STYLE_TYPE_HIDDEN
	CSS_BORDER_STYLE_TYPE_DOTTED
	CSS_BORDER_STYLE_TYPE_DASHED
	CSS_BORDER_STYLE_TYPE_SOLID
	CSS_BORDER_STYLE_TYPE_DOUBLE
	CSS_BORDER_STYLE_TYPE_GROOVE
	CSS_BORDER_STYLE_TYPE_RIDGE
	CSS_BORDER_STYLE_TYPE_INSET
	CSS_BORDER_STYLE_TYPE_OUTSET
)

const (
	CSS_BORDER_RADIUS_TYPE_LENGTH CssBorderRadiusType = iota
	CSS_BORDER_RADIUS_TYPE_PERCENTAGE
)

const (
	CSS_BORDER_COLLAPSE_TYPE_EMPTY CssBorderCollapseType = iota
	CSS_BORDER_COLLAPSE_TYPE_SEPARATE
	CSS_BORDER_COLLAPSE_TYPE_COLLAPSE
)

const (
	CSS_BORDER_IMAGE_SOURCE_TYPE_NONE CssBorderImageSourceType = iota
	CSS_BORDER_IMAGE_SOURCE_TYPE_IMAGE
)

const (
	CSS_BORDER_IMAGE_SLICE_TYPE_NUMBER CssBorderImageSliceType = iota
	CSS_BORDER_IMAGE_SLICE_TYPE_PERCENTAGE
	CSS_BORDER_IMAGE_SLICE_TYPE_FILL
)

const (
	CSS_BORDER_IMAGE_WIDTH_TYPE_LENGTH CssBorderImageWidthType = iota
	CSS_BORDER_IMAGE_WIDTH_TYPE_NUMBER
	CSS_BORDER_IMAGE_WIDTH_TYPE_PERCENTAGE
	CSS_BORDER_IMAGE_WIDTH_TYPE_AUTO
)

const (
	CSS_BORDER_IMAGE_OUTSET_TYPE_LENGTH CssBorderImageOutsetType = iota
	CSS_BORDER_IMAGE_OUTSET_TYPE_NUMBER
)

const (
	CSS_BORDER_IMAGE_REPEAT_TYPE_STRETCH CssBorderImageRepeatType = iota
	CSS_BORDER_IMAGE_REPEAT_TYPE_REPEAT
	CSS_BORDER_IMAGE_REPEAT_TYPE_ROUND
	CSS_BORDER_IMAGE_REPEAT_TYPE_SPACE
)

const (
	CSS_PROPERTY_VALUE_TYPE_EMPTY CssPropertyValueType = iota
	CSS_PROPERTY_VALUE_TYPE_AUTO
	CSS_PROPERTY_VALUE_TYPE_LENGTH
	CSS_PROPERTY_VALUE_TYPE_PIXEL
	CSS_PROPERTY_VALUE_TYPE_PERCENTAGE
	CSS_PROPERTY_VALUE_TYPE_NONE
	CSS_PROPERTY_VALUE_TYPE_NORMAL
	CSS_PROPERTY_VALUE_TYPE_MAX_CONTENT
	CSS_PROPERTY_VALUE_TYPE_MIN_CONTENT
	CSS_PROPERTY_VALUE_TYPE_SPAN
	CSS_PROPERTY_VALUE_TYPE_MINMAX
	CSS_PROPERTY_VALUE_TYPE_FIT_CONTENT
)

const (
	CSS_BOX_DECORATION_BREAK_TYPE_SLICE CssBoxDecorationType = iota
	CSS_BOX_DECORATION_BREAK_TYPE_CLONE
	CSS_BOX_DECORATION_BREAK_TYPE_UNSET
)

const (
	CSS_BOX_SIZING_TYPE_CONTENT_BOX CssBoxSizingType = iota
	CSS_BOX_SIZING_TYPE_BORDER_BOX
)

const (
	CSS_DISPLAY_TYPE_BLOCK CssDisplayType = iota
	CSS_DISPLAY_TYPE_INLINE
	CSS_DISPLAY_TYPE_CONTENTS
	CSS_DISPLAY_TYPE_FLEX
	CSS_DISPLAY_TYPE_GRID
	CSS_DISPLAY_TYPE_INLINE_BLOCK
	CSS_DISPLAY_TYPE_INLINE_FLEX
	CSS_DISPLAY_TYPE_INLINE_GRID
	CSS_DISPLAY_TYPE_INLINE_TABLE
	CSS_DISPLAY_TYPE_LIST_ITEM
	CSS_DISPLAY_TYPE_RUN_IN
	CSS_DISPLAY_TYPE_TABLE
	CSS_DISPLAY_TYPE_TABLE_CAPTION
	CSS_DISPLAY_TYPE_TABLE_COLUMN_GROUP
	CSS_DISPLAY_TYPE_TABLE_HEADER_GROUP
	CSS_DISPLAY_TYPE_TABLE_FOOTER_GROUP
	CSS_DISPLAY_TYPE_TABLE_ROW_GROUP
	CSS_DISPLAY_TYPE_TABLE_CELL
	CSS_DISPLAY_TYPE_TABLE_COLUMN
	CSS_DISPLAY_TYPE_TABLE_ROW
	CSS_DISPLAY_TYPE_NONE
)

const (
	CSS_POSITION_TYPE_EMPTY CssPositionType = iota
	CSS_POSITION_TYPE_ABSOLUTE
	CSS_POSITION_TYPE_FIXED
	CSS_POSITION_TYPE_RELATIVE
	CSS_POSITION_TYPE_STATIC
	CSS_POSITION_TYPE_STICKY
)

const (
	CSS_FONT_SIZE_TYPE_EMPTY CssFontSizeType = iota
	CSS_FONT_SIZE_TYPE_MEDIUM
	CSS_FONT_SIZE_TYPE_XX_SMALL
	CSS_FONT_SIZE_TYPE_X_SMALL
	CSS_FONT_SIZE_TYPE_SMALL
	CSS_FONT_SIZE_TYPE_LARGE
	CSS_FONT_SIZE_TYPE_X_LARGE
	CSS_FONT_SIZE_TYPE_XX_LARGE
	CSS_FONT_SIZE_TYPE_SMALLER
	CSS_FONT_SIZE_TYPE_LARGER
	CSS_FONT_SIZE_TYPE_LENGTH
	CSS_FONT_SIZE_TYPE_PERCENTAGE
)

const (
	CSS_FLEX_DIRECTION_EMPTY CssFlexDirectionType = iota
	CSS_FLEX_DIRECTION_COLUMN
	CSS_FLEX_DIRECTION_COLUMN_REVERSE
	CSS_FLEX_DIRECTION_ROW
	CSS_FLEX_DIRECTION_ROW_REVERSE
)
