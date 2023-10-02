package utils

import (
	"fyne.io/fyne/v2/canvas"
)

func ImageOverlaps(image1, image2 *canvas.Image) bool {
    
    if image1.Position().X < image2.Position().X+image2.Size().Width &&
        image1.Position().X+image1.Size().Width > image2.Position().X {
      
        if image1.Position().Y < image2.Position().Y+image2.Size().Height &&
            image1.Position().Y+image1.Size().Height > image2.Position().Y {
            return true
        }
    }
    return false
}
