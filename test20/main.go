package main

import (
	"io"
	"os"
	"strings"
	"errors"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(buf []byte) (n int, err error)  {
	rot13 := map[byte]byte{
		'A':'N',
		'B':'O',
		'C':'P',
		'D':'Q',
		'E':'R',
		'F':'S',
		'G':'T',
		'H':'U',
		'I':'V',
		'J':'W',
		'K':'X',
		'L':'Y',
		'M':'Z',
		'N':'A',
		'O':'B',
		'P':'C',
		'Q':'D',
		'R':'E',
		'S':'F',
		'T':'G',
		'U':'H',
		'V':'I',
		'W':'J',
		'X':'K',
		'Y':'L',
		'Z':'M',
		'a':'n',
		'b':'o',
		'c':'p',
		'd':'q',
		'e':'r',
		'f':'s',
		'g':'t',
		'h':'u',
		'i':'v',
		'j':'w',
		'k':'x',
		'l':'y',
		'm':'z',
		'n':'a',
		'o':'b',
		'p':'c',
		'q':'d',
		'r':'e',
		's':'f',
		't':'g',
		'u':'h',
		'v':'i',
		'w':'j',
		'x':'k',
		'y':'l',
		'z':'m',
		'0':'9',
		'1':'8',
		'2':'7',
		'3':'6',
		'4':'5',
		'5':'4',
		'6':'3',
		'7':'2',
		'8':'1',
		'9':'0',
		' ':' ',
		'~':'~',
		'!':'!',
		'@':'@',
		'#':'#',
		'$':'$',
		'%':'%',
		'^':'^',
		'&':'&',
		'*':'*',
		'(':'(',
		')':')',
		'_':'_',
		'+':'+',
		'`':'`',
		'-':'-',
		'=':'=',
		'?':'?',
		',':',',
		'.':'.',
	}

	l, err := rot.r.Read(buf)
	if err != nil{
		return 0, errors.New("Read something wrong")
	}

	for i, v := range buf{
		if v == byte(0){
			return i, nil
		}
		buf[i] = rot13[v]
	}

	return l, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr? Lrf,Zl rznvy vf whaobwvna@dd.pbz")
	//s1 := strings.NewReader("You cracked the code? Yes,My email is junbojian@qq.com")

	r := rot13Reader{s}
	//r1 := rot13Reader{s1}

	io.Copy(os.Stdout, &r)
	fmt.Println("")
	//io.Copy(os.Stdout, &r1)
}
