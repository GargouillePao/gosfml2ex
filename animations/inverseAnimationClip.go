package animations

type inverseAnimationClip struct {
	animationClip
}

func NewInverseAnimationClip(start []float32, end []float32, onAnimateFunc OnAnimateFunc, onEndFunc AnimateEndFunc) AnimationClip {
	Slice.Equilongf(&start, &end)
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

func (i *inverseAnimationClip) SetAnimationCurve(curve AnimateCurveFunc) {
	i.animateCurve = func(num float32) float32 {
		return curve(1 - num)
	}
}
