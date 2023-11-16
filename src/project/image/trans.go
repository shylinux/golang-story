package image

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
	"shylinux.com/x/ice"
	icebergs "shylinux.com/x/icebergs"
	"shylinux.com/x/icebergs/base/nfs"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
)

const (
	WIDTH  = "width"
	HEIGHT = "height"
)

type trans struct {
	resize string `name:"resize height=390 width=390" help:"尺寸"`
	list   string `name:"list path auto" help:"变换"`
}

func (s trans) Resize(m *ice.Message, arg ...string) {
	if strings.HasSuffix(m.Option(nfs.PATH), nfs.PS) {
		msg := m.Cmd(nfs.DIR, m.Option(nfs.PATH), kit.Dict(nfs.DIR_TYPE, nfs.TYPE_CAT))
		web.GoToast(m.Message, "", func(toast func(string, int, int)) []string {
			msg.Table(func(index int, value ice.Maps, head []string) {
				toast(value[nfs.PATH], index, msg.Length())
				Resize(m.Message, value[nfs.PATH], uint(kit.Int(m.Option(HEIGHT))), uint(kit.Int(m.Option(WIDTH))))
			})
			return nil
		})
	} else {
		Resize(m.Message, m.Option(nfs.PATH), uint(kit.Int(m.Option(HEIGHT))), uint(kit.Int(m.Option(WIDTH))))
	}
}
func (s trans) List(m *ice.Message, arg ...string) {
	msg := m.Cmd(nfs.DIR, arg)
	msg.Table(func(index int, value ice.Maps, head []string) {
		m.Push("", value, kit.Split("time,path,size"))
		if strings.HasSuffix(value[nfs.PATH], nfs.PS) {

		} else if f, e := os.Open(value[nfs.PATH]); !m.Warn(e) {
			defer f.Close()
			if img, fmt, err := image.Decode(f); !m.Warn(err) {
				max, min := img.Bounds().Max, img.Bounds().Min
				m.Push(HEIGHT, max.Y-min.Y).Push(WIDTH, max.X-min.X).Push("fmt", fmt).PushButton(s.Resize)
				return
			}
		}
		m.Push(HEIGHT, "").Push(WIDTH, "").Push("fmt", "").PushButton("")
	})
	m.Action(s.Resize).Options(ice.MSG_STATUS, msg.Option(ice.MSG_STATUS))
}

func init() { ice.CodeCmd(trans{}) }
func init() { nfs.ImageResize = Resize }

func Resize(m *icebergs.Message, p string, height, width uint) bool {
	f, e := os.Open(p)
	if m.Warn(e) {
		return false
	}
	defer f.Close()
	img, fmt, err := image.Decode(f)
	if m.Warn(err) {
		return false
	}
	out, err := os.Create(p)
	if m.Warn(err) {
		return false
	}
	defer out.Close()
	switch _img := resize.Resize(width, height, img, resize.Lanczos3); fmt {
	case nfs.PNG:
		png.Encode(out, _img)
	case nfs.JPEG:
		jpeg.Encode(out, _img, nil)
	default:
		m.ErrorNotImplement(fmt)
		return false
	}
	return true
}
