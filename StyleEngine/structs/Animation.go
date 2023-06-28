package structs

import (
	"gezgin_web_engine/StyleEngine/enums"
)

type cubicBezierFunction struct {
	values [4]int
}

type stepsFunction struct {
	value   int
	isStart bool
}

type cssAnimationTimingFunction struct {
	cubicBezier cubicBezierFunction
	steps       stepsFunction
	timingType  enums.CssAnimationTimingFunctionType
}

type Animation struct {
	AnimationNameInherit           bool
	AnimationDelayInherit          bool
	AnimationDurationInherit       bool
	AnimationIterationCountInherit bool
	AnimationDirectionInherit      bool
	AnimationFillModeInherit       bool
	AnimationPlayStateInherit      bool
	AnimationTimingFunctionInherit bool

	AnimationName                  string
	AnimationDelay                 int
	AnimationDuration              int
	AnimationIterationCount        int
	WhichAnimationTimingPartActive int // 0, 1, 2
	AnimationDirection             enums.CssAnimationDirectionType
	AnimationFillMode              enums.CssAnimationFillModeType
	AnimationPlayState             enums.CssAnimationPlayStateType
	AnimationTimingFunction        cssAnimationTimingFunction
}
