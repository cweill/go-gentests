// Code generated by "esc -o bindata/esc.go -pkg=bindata templates"; DO NOT EDIT.

package bindata

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/call.tmpl": {
		name:    "call.tmpl",
		local:   "templates/call.tmpl",
		size:    241,
		modtime: 1663641620,
		compressed: `
H4sIAAAAAAAC/0SOQWrDQAxFryKMFy0YHaDQA3hTSlvatRjLrsCeFo2SEITuHsY4mdWHP2/el/vEs2SG
LtG6dhHuF7FfwA9OLGfW2sgM+c8Ax/JpekoWYYbunKf6eicBI1qLb7RxxJO7Ul4Yehmg5xVeXgHfSWlj
Yy2HvZeIAR5/296PitUbzJB0KU2/K+riTuPX9Z9xLN+kQpOkCMTG7vF85C0AAP//ZQi8iPEAAAA=
`,
	},

	"/templates/function.tmpl": {
		name:    "function.tmpl",
		local:   "templates/function.tmpl",
		size:    2657,
		modtime: 1663641620,
		compressed: `
H4sIAAAAAAAC/7RWTW/bRhA9U79iIjiBWCibuw0fGsQtcmhcyEZzCIpiLQ3VRVdLdXfoQFjsfy/2i1xK
pONLLyK5nI83894MZe0OG6EQlk2ntiRatXRuYe17uGrg+haYc4uFfwXWskc09IUf0LkVwU+EhoTas8ca
7KLyLt8F/Q1sg1sUz6idW1ThWDTAPpsH0t2WwmF/+otAuTPxrKLTEaEJJ2CCsY+brDVXezxzqKwNzx5k
gHc6YnrlXVDt0lOfMx8V92e3HpUv83eu+QEJdUgWoHG9HwErYF16hITh6AJdkXGc3zfU+KZbG3D4Vu+c
O/DjN0NaqP2fMbu1KA0692149gESKO+6Ui0l/9o5xQ8IMUIyTSgm+MpN4Go3kHbW98RRvPStDZBS4Tnk
JTMzLLzQ+6oKjfc/Ez4FARs0nSST83zlil7qfZ9yg9RpZe60blMPvnNFd1rDU9vKM8J8jz98gMf7T/fX
8PNuB5402HKDhgU+m1aDtZ6DVgN76J4iq5GRz+YL/wd3dR1gFRx7ijKtfyWS1kAU1RDpjaXGcLbAn3OM
ivI9lBKlczEI0c2Y+8xxNhxB8R7+emOtrxyCF7FNp1azsIlYeozg/dK4XBOTCPv7VT0G2Y/wnFpn9suF
BkMbg75ORwzGXDv3LmVLCmF/cNmhczaHmFk7lbUsrsFrIGJxDlixjNZDgGEJVROb6eIh5ZvYJbnMr1pQ
X/1ox1zfwrunE6FhH7umQW1fkzDJIMrzXslTOQ315fm9wtClGnpkhIej5ISw1HECl3DVBIEPb7Zcyng8
h2JiDCvRJNbOgTkHqHVkdSrJTT80K2/35haUkLW/ErE83YlmYiFks1qWsQ5oDN9jKgW9BdzC2+c1ZPe3
z8v1KL1Qx64vHrVeF8nqQRF5y43WVXhXDjM4p0PBeby2rSKhOhwG5MefuRc35KymQtd/bWkYnF5j7CF8
RVb1TWESu1ou3EF30mDK8ZEbsS2+0D25V82UvvxQjjCUfZZC4TnRr8bzP+V/o7GRuCX2CfF492/H5aqP
sB4DqktEPXuv0WEGnMD+1kkSRzkCm/AMWv2BUGdBzv9zOtMp+IKiCbyP/2rcYpF1+l8AAAD//1pk1KJh
CgAA
`,
	},

	"/templates/functionmock.tmpl": {
		name:    "functionmock.tmpl",
		local:   "templates/functionmock.tmpl",
		size:    3119,
		modtime: 1663682015,
		compressed: `
H4sIAAAAAAAC/7RWv27cxhOueU8xPtgGaZxX/Qkq7J/lH1xYMiQhLowgWPGGZ0J7S2Z3KOGw2CpFgNRJ
GwSpUvoF8jhK8hjB/uG/453kJs0dOdyZ+Wbm+4Y0ZoVFKRHmRSNzKiu5qfKbubUzY17C0wKWJ8Csnc3c
YzCGXaGmM75Ba1OCF4SaSrlmVxmYWeJc7kr6DOwCcyxvUVk7S7y5LIC905ekmpy8sbO+LVGsdLAltK0R
Cm8B7Q+7uPG04nKNOw6JMf7egfTwtjXGR84F5SredTlb0+B659KhcmV+4IpvkFD5ZB4aV+sRsAGsqYdP
6E0TdIOM4/yuodo13RiPw7V6Ze2G1580qVKuvw3ZjUGh0dpP/b0LEEE511RWFP0zayXfIIQI8WhEsWde
bRO4XPVD2+l7nFH461rrIcXC25DTyRyYwgO9TxLfePezx2cwgAvUjSDd5vnIJT3U+y7lBVKjpD5Vqoo9
uOOSTpWC66oSI6dZkhwdwf3vv97/9Bs4scD9j3/+9cuXUT/fV/nN/7gQqN42Mg8R3dkPnPLPqMGpKa3h
xbraVPIGtyw+yHa44cfpnI+O4Or8zfkSXq1W4EgCOdeo2VBJjyRdHs46Fdm+WEnNXtW12DpLOjeGhT0w
X4AxhJtacEKY51yIuWt6NlWhXezO31nsLCkqBcY41lYK2GVzHXQQOPxOn/EbXGWZH+RAFY7UrRC+i7Re
AFHQTxBEqCeEM4OJtzlGNHCsEwKFtSEI0fFYLa0q2oMjKM7D/R8b4+oD70XsopHpQdhELN4G8H5Ek8W6
F2F3nWZjkF27D+n7wEaeqNa30StyW6M/zJW1z2O2qCn2DRcNWmvaEAcWddIRZglELGwONljfiz5Av7aT
Pbt8chPz7dm+bZkfVUld9aOtvDyB59dbQs1eN0WByjyYMErxny8//P3zH179jwkwqaPklyfQ6e4M76L0
0qASIjZQalr32yBJVliggmhiF6iRgtcEY2RnUM25FNvhWsum9nOJfngZdA3rdazCKp3D08Lrblfhzjxt
zuGFmpRFZNMuMmsBlQps25fluBNz6s49OQFZisz9E7F2T0f6EfMhC7ef+lgb1JqvMdaC7gScwLPbBbTu
z2531lgp66arHpVaDJJlPVPb99XoxeOfDZcMWKt8wa3s80pSKRvshfv4B8uD77qDXPdd/39FvaA77rNL
/z2QZseDI6Grw1dnrwehMeZ4zXWZD761uuE+LfYRzC2LEYZhn0UpcXfQX43nP8r/RGEhMCf2BrE+/b7h
Iu0iLMaAhi/Pfnpfw8MWcAT7vhFU1mIENuLpufoIUQ+CPPwNvMNTcAWFI/AyfJ/a2azl6b8BAAD//2Ql
Md8vDAAA
`,
	},

	"/templates/header.tmpl": {
		name:    "header.tmpl",
		local:   "templates/header.tmpl",
		size:    142,
		modtime: 1663641620,
		compressed: `
H4sIAAAAAAAC/0TMMQ7CMAyF4d2nsDrBQC7BxIK4gkUebYXiViGb9e6OlAi6/bL1voiM1+rQaYFl1ImU
iGo+Q9N1KwXePmRE6g941gspuz3fNkMj0mMkKbKWfatNT4dw65cB3K2AHJO2/DhSzv/6BgAA///GzMM9
jgAAAA==
`,
	},

	"/templates/inline.tmpl": {
		name:    "inline.tmpl",
		local:   "templates/inline.tmpl",
		size:    49,
		modtime: 1663641620,
		compressed: `
H4sIAAAAAAAC/6quTklNy8xLVVDKzMvJzEtVqq1VqK4uSc0tyEksSVVQSk7MyVFS0AOLpual1NYCAgAA
//+q60H/MQAAAA==
`,
	},

	"/templates/inputs.tmpl": {
		name:    "inputs.tmpl",
		local:   "templates/inputs.tmpl",
		size:    177,
		modtime: 1663681639,
		compressed: `
H4sIAAAAAAAC/0yNMaoDMQxE+38KsWz58QECOUCaEMgJFCwvLqwESVsJ3T1YpHAlzWN4416pdSbYOn9O
0y3CfW9wuUKZb2/Ab4PyPF9GarqyOw6qEWbFnbhGMA76h1/I3t7KQzrbLeUTCvJByVFwkJFoKlAOLe5J
5/TiWc/fNwAA//94+RPrsQAAAA==
`,
	},

	"/templates/message.tmpl": {
		name:    "message.tmpl",
		local:   "templates/message.tmpl",
		size:    201,
		modtime: 1663681656,
		compressed: `
H4sIAAAAAAAC/zyN4WqDQBCE//sUiyi0oPsAhT5A/xRpS/9f4mgW9GLuTkNY9t2DB/HXDDPDN6o9BvGg
ckaMbkRJrVmhKgP5ayL+XU8JMUWz+sakCt+bqd4lXYh/cIZsCHvCf48F/O+mFWZ8DPnbzTB7y0Tugvj0
5Zd1B6oG50dQJQ1VmOjjk7hzwc1ICLmXgSoxa16/9XZws7wXqi1l+wwAAP//kC65UskAAAA=
`,
	},

	"/templates/results.tmpl": {
		name:    "results.tmpl",
		local:   "templates/results.tmpl",
		size:    168,
		modtime: 1663641620,
		compressed: `
H4sIAAAAAAAC/1yNTQrCQAyFr/Iosyw9gOBS3HsDoRkJlAy8ma5C7i6pRcFVfr4vee6rVDXBROn7NvoU
AXc+7SUoOqPIhssVy+ODI9y1omjEDHexNTf3NrBkc85a82DstH4jG1MW8uQ4hMbv0385A3/uUd8BAAD/
/7BPz2GoAAAA
`,
	},

	"/templates": {
		name:  "templates",
		local: `templates`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	"templates": {
		_escData["/templates/call.tmpl"],
		_escData["/templates/function.tmpl"],
		_escData["/templates/functionmock.tmpl"],
		_escData["/templates/header.tmpl"],
		_escData["/templates/inline.tmpl"],
		_escData["/templates/inputs.tmpl"],
		_escData["/templates/message.tmpl"],
		_escData["/templates/results.tmpl"],
	},
}
