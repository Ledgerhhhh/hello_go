package try_designPatterns

import (
	"fmt"
	"testing"
)

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func (r RealImage) LoadImage() {
	fmt.Println("loading ....")
}

func (r RealImage) Display() {
	fmt.Println(r.filename)
}

type proxyImage struct {
	RealImage *RealImage
	filename  string
}

func (p *proxyImage) Display() {
	if p.RealImage == nil {
		p.RealImage = &RealImage{
			filename: p.filename,
		}
		p.RealImage.LoadImage()
	}
	p.RealImage.Display()

}

func TestRealImg(t *testing.T) {
	proxy := proxyImage{
		RealImage: nil,
		filename:  "proxy",
	}
	proxy.Display()
	proxy.Display()
	proxy.Display()
	proxy.Display()
}
