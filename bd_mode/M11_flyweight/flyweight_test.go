package M11_flyweight

import "testing"

func ExampleGetImageFlyweight() {

	viewer := NewImageViewer("image1.pag")
	viewer.Display()

}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}
