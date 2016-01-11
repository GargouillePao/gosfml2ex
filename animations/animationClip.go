package animations

type AnimationClip interface {
	SetOnAnimate(OnAnimateFunc)
	SetOnAnimateEnd(AnimateEndFunc)
	Animate()
	GetNext() AnimationClip
	SetNext(AnimationClip) AnimationClip
	SetAnimationCurve(AnimateCurveFunc)
	SetFrameCount(frame float32)
}

type animationClip struct {
	stepX        float32
	stepY        float32
	step         float32
	startState   []float32
	endState     []float32
	onAnimate    OnAnimateFunc
	onAnimateEnd AnimateEndFunc
	animateCurve AnimateCurveFunc
	nextClip     AnimationClip
}

func NewAnimationClip(start []float32, end []float32, onAnimateFunc OnAnimateFunc, onEndFunc AnimateEndFunc) AnimationClip {
	Slice.Equilongf(&start, &end)
	clip := &animationClip{
		startState:   start,
		endState:     end,
		onAnimate:    onAnimateFunc,
		onAnimateEnd: onEndFunc,
		animateCurve: func(num float32) float32 {
			return num
		},
	}
	return clip
}

func (a *animationClip) SetFrameCount(frame float32) {
	a.step = 1 / frame
}

func (a *animationClip) SetAnimationCurve(curve AnimateCurveFunc) {
	a.animateCurve = curve
}

func (a *animationClip) SetOnAnimate(listener OnAnimateFunc) {
	a.onAnimate = listener
}
func (a *animationClip) SetOnAnimateEnd(listener AnimateEndFunc) {
	a.onAnimateEnd = listener
}
func animate(startState []float32, endState []float32, step float32, listener OnAnimateFunc) {
	listener(Slice.Lerpf(startState, endState, step))
}
func (a *animationClip) Animate() {
	if a.stepX <= 1 {
		a.stepX += a.step
		a.stepY = a.animateCurve(a.stepX)
		animate(a.startState, a.endState, a.stepY, a.onAnimate)
	} else {
		a.AnimateNext()
	}
}
func (a *animationClip) AnimateNext() {
	if a.nextClip != nil {
		a.nextClip.Animate()
	} else {
		if a.onAnimateEnd != nil {
			a.onAnimateEnd()
		}
	}
}
func (a *animationClip) GetNext() AnimationClip {
	return a.nextClip
}
func (a *animationClip) SetNext(clip AnimationClip) AnimationClip {
	a.nextClip = clip
	return a.nextClip
}
