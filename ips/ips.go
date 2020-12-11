package ips

import (
	"context"
	"fmt"
	"github.com/imakiri/playground/protos"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

type FaceDetecter struct {
	protos.UnimplementedFaceDetecterServer
}

func NewFaceDetector() *FaceDetecter {
	return &FaceDetecter{}
}

func (e *FaceDetecter) Detect(_ context.Context, dr *protos.DetectionRequest) (*protos.DetectionResponse, error) {
	rImg := dr.GetImg()

	fmt.Print("Detection in process\n")

	img, err := gocv.IMDecode(rImg, gocv.IMReadColor)
	if err != nil {
		return &protos.DetectionResponse{Response: &protos.DetectionResponse_Err{Err: &protos.Error{Error: protos.Errors_IncorrectImage}}}, err
	}
	defer img.Close()

	imgGrey := gocv.NewMat()
	defer imgGrey.Close()

	gocv.CvtColor(img, &imgGrey, gocv.ColorBGRToGray)

	harrcascade := "ips\\haarcascades\\haarcascade_frontalface_alt.xml"
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(harrcascade)
	defer classifier.Close()

	//rects := classifier.DetectMultiScaleWithParams(img, 1.5, 2, 0, image.Point{}, image.Point{X: img.Size()[0], Y: img.Size()[1]})
	rects := classifier.DetectMultiScale(img)

	for _, r := range rects {
		gocv.Rectangle(&img, r.Add(image.Point{
			X: 1,
			Y: 1,
		}), color.RGBA{B: 255, G: 255}, 1)
		gocv.Rectangle(&img, r, color.RGBA{R: 255}, 1)
	}

	data, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return &protos.DetectionResponse{Response: &protos.DetectionResponse_Err{Err: &protos.Error{Error: protos.Errors_InternalServiceError}}}, err
	}

	return &protos.DetectionResponse{Response: &protos.DetectionResponse_Img{Img: &protos.Image{Data: data}}}, nil
}
