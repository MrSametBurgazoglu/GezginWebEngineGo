package StyleProperty

import (
	"gezgin_web_engine/StyleProperty/enums"
	"gezgin_web_engine/StyleProperty/structs"
	"strconv"
)

const ANIMATION_TIMING_STRING_COUNT = 7

var animationTimingFunctionString = []string{
	"ease",
	"ease-in",
	"ease-in-out",
	"ease-out",
	"linear",
	"step-end",
	"step-start",
}

func setAnimationPlayState(animation *structs.Animation, value string) {
	if value == "paused" {
		animation.AnimationPlayState = enums.CSS_ANIMATION_PLAY_STATE_PAUSED
	} else {
		animation.AnimationPlayState = enums.CSS_ANIMATION_PLAY_STATE_RUNNING
	}
}

func setAnimationFillMode(animation *structs.Animation, value string) {
	switch value {
	case "forwards":
		animation.AnimationFillMode = enums.CSS_ANIMATION_FILL_MODE_FORWARDS
	case "backwards":
		animation.AnimationFillMode = enums.CSS_ANIMATION_FILL_MODE_BACKWARDS
	case "both":
		animation.AnimationFillMode = enums.CSS_ANIMATION_FILL_MODE_BOTH
	default:
		animation.AnimationFillMode = enums.CSS_ANIMATION_FILL_MODE_NONE
	}
}

func setAnimationDirection(animation *structs.Animation, value string) {
	switch value {
	case "reverse":
		animation.AnimationDirection = enums.CSS_ANIMATION_DIRECTION_REVERSE
	case "alternate":
		animation.AnimationDirection = enums.CSS_ANIMATION_DIRECTION_ALTERNATE
	case "alternate-reverse":
		animation.AnimationDirection = enums.CSS_ANIMATION_DIRECTION_ALTERNATE_REVERSE
	default:
		animation.AnimationDirection = enums.CSS_ANIMATION_DIRECTION_NORMAL
	}
}

func setAnimationIterationCount(animation *structs.Animation, value string) {
	count, err := strconv.Atoi(value)
	if err == nil {
		animation.AnimationIterationCount = count
	}
}

func setAnimationDelay(animation *structs.Animation, value string) {
	delay, err := strconv.Atoi(value)
	if err == nil {
		animation.AnimationDelay = delay
	}
}

func setAnimationDuration(animation *structs.Animation, value string) {
	duration, err := strconv.Atoi(value)
	if err == nil {
		animation.AnimationDuration = duration
	}
}
