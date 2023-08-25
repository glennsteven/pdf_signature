package signature

import (
	"bytes"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func Stamp(input, qrcode []byte, pos int) ([]byte, error) {
	output := new(bytes.Buffer)
	wm := model.DefaultWatermarkConfig()
	switch pos {
	case 0:
		wm.Pos = types.TopLeft
	case 1:
		wm.Pos = types.TopCenter
	case 2:
		wm.Pos = types.Right
	case 3:
		wm.Pos = types.Left
	case 5:
		wm.Pos = types.Center
	case 6:
		wm.Pos = types.BottomLeft
	case 7:
		wm.Pos = types.BottomCenter
	case 8:
		wm.Pos = types.BottomRight
	case 9:
		wm.Pos = types.Full
	default:
		wm.Pos = types.Center
	}
	wm.Page = 1
	wm.Mode = model.WMImage
	wm.Image = bytes.NewBuffer(qrcode)
	wm.Diagonal = 0
	wm.Scale = 0.15
	err := api.AddWatermarks(bytes.NewReader(input), output, []string{"1"}, wm, nil)
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
