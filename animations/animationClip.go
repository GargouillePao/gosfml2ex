package animations

import (
	sf "bitbucket.org/krepa098/gosfml2"
)

type AnimationClip interface {
	SetOnAnimate(listener func(interface{}))
	SetOnAnimateEnd(listener func())
	Animate()
	GetNext() AnimationClip
	SetNext(AnimationClip) AnimationClip
	SetAnimationCurve(func(float32) float32)
	SetFrameCount(frame float32)
}

type animationClip struct {
	stepX        float32
	stepY        float32
	step         float32
	startState   interface{}
	endState     interface{}
	onAnimate    func(interface{})
	onAnimateEnd func()
	animateCurve func(float32) float32
	nextClip     AnimationClip
}

func NewSingleAnimationClip(start interface{}, end interface{}, onAnimateFunc func(interface{}), onEndFunc func()) AnimationClip {
	return &animationClip{
		startState:   start,
		endState:     end,
		onAnimate:    onAnimateFunc,
		onAnimateEnd: onEndFunc,
		animateCurve: func(num float32) float32 {
			return num
		},
	}
}

func NewMultiAnimationClip(start interface{}, end interface{}, onAnimateFunc func(interface{}), nextClip AnimationClip) AnimationClip {
	return &animationClip{
		startState: start,
		endState:   end,
		onAnimate:  onAnimateFunc,
		nextClip:   nextClip,
		animateCurve: func(num float32) float32 {
			return num
		},
	}
}

func (a *animationClip) SetFrameCount(frame float32) {
	a.step = 1 / frame
}

func (a *animationClip) SetAnimationCurve(curve func(float32) float32) {
	a.animateCurve = curve
}

func (a *animationClip) SetOnAnimate(listener func(interface{})) {
	a.onAnimate = listener
}
func (a *animationClip) SetOnAnimateEnd(listener func()) {
	a.onAnimateEnd = listener
}
func animate(startState interface{}, endState interface{}, step float32, listener func(interface{})) {
	var resultV1 float32
	var resultV2 sf.Vector2f
	var resultV3 sf.Vector3f
	var resultColor sf.Color
	var err error

	resultV1, err = Vector.LerpV1(startState, endState, step)
	if err == nil {
		listener(resultV1)
		return
	}
	resultV2, err = Vector.LerpV2(startState, endState, step)
	if err == nil {
		listener(resultV2)
		return
	}

	resultV3, err = Vector.LerpV3(startState.(sf.Vector3f), endState.(sf.Vector3f), step)
	if err == nil {
		listener(resultV3)
		return
	}
	resultColor, err = Vector.LerpV4(startState.(sf.Color), endState.(sf.Color), step)
	if err == nil {
		listener(resultColor)
		return
	}
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
