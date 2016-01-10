package animations

import (
	//sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	sfutils "sfmlex/utils"
	"time"
)

var Vector = sfutils.Utils().Vector

type Animation struct {
	clip     AnimationClip
	clipHead AnimationClip
	start    bool
	delay    float64
}

func NewAnimation(delay float64) *Animation {
	return &Animation{start: false, delay: delay}
}

func (a *Animation) AddClip(clip AnimationClip) {
	if a.clip == nil {
		a.clip = clip
		a.clipHead = clip
	} else {
		for {
			if a.clip.GetNext() == nil {
				a.clip.SetNext(clip)
				a.clip = a.clipHead
				fmt.Println(a.clip)
				return
			} else {
				a.clip = a.clip.GetNext()
			}
		}
	}

}
func (a *Animation) GetFirstClip() AnimationClip {
	return a.clip
}
func (a *Animation) Animate() {
	if a.start {
		if a.clip != nil {
			a.clip.Animate()
		}

	}
}
func (a *Animation) Play() {
	time.AfterFunc(time.Duration(a.delay*time.Duration.Seconds(1)), func() {
		a.start = true
	})
}
func (a *Animation) Stop() {
	a.start = false
}
