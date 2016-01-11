package animations

const (
	Circling = iota
	Pingpong
)

type loopAnimationClip struct {
	animationClip
	loopCount int
	loopStep  int
	loopType  int
}

func NewLoopAnimation(start []float32, end []float32, loopCount int, loopType int, onAnimateFunc OnAnimateFunc, onEndFunc AnimateEndFunc) AnimationClip {
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
	return &loopAnimationClip{animationClip: animationClip, loopCount: loopCount, loopType: loopType}
}

func (l *loopAnimationClip) Animate() {
	if l.stepX <= 1 {
		l.stepX += l.step
		l.stepY = l.animateCurve(l.stepX)
		animate(l.startState, l.endState, l.stepY, l.onAnimate)
	} else {
		if l.loopStep < l.loopCount {
			l.PingPongAnimation()
			l.loopStep++
		} else {
			if l.loopCount < 0 {
				l.PingPongAnimation()
			} else {
				l.AnimateNext()
			}
		}
	}
}

func (l *loopAnimationClip) PingPongAnimation() {
	animationCurve := l.animateCurve
	if l.loopType == Pingpong {
		l.SetAnimationCurve(func(num float32) float32 {
			return animationCurve(1 - num)
		})
	}
	l.stepX = 1 - l.stepX
}
