package animations

type inverseAnimationClip struct {
	animationClip
}

func NewInverseAnimationClip(start interface{}, end interface{}, onAnimateFunc func(interface{}), onEndFunc func()) AnimationClip {
	animationClip := animationClip{
		startState:   start,
		endState:     end,
		onAnimate:    onAnimateFunc,
		onAnimateEnd: onEndFunc,
		animateCurve: func(num float32) float32 {
			return num
		},
	}
	return &inverseAnimationClip{animationClip}
}

func (i *inverseAnimationClip) SetAnimationCurve(curve func(float32) float32) {
	i.animateCurve = func(num float32) float32 {
		return curve(1 - num)
	}
}
