package sign

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {

	/**
	appid： wxd930ea5d5a258f4f

	mch_id： 10000100

	device_info： 1000

	body： test

	nonce_str： ibuaiVcKdpRxkhJA
	*/
	//m := map[string]string{"appid": "wxd930ea5d5a258f4f", "mch_id": "10000100", "device_info": "1000", "body": "test", "nonce_str": "ibuaiVcKdpRxkhJA"}
	//
	//keys := make([]string, 0, len(m))
	//for k := range m {
	//	keys = append(keys, k)
	//}
	//sort.Strings(keys)
	//var build strings.Builder
	//var count = 0
	//for _, k := range keys {
	//	count++
	//	if m[k] != "" {
	//		if count == len(m) {
	//			build.WriteString(k + "=" + m[k])
	//		} else {
	//			build.WriteString(k + "=" + m[k] + "&")
	//		}
	//	}
	//
	//}
	//result := build.String()
	//
	//stringSignTemp := result + "&key=192006250b4c09247ec02edce69f6a2d"
	//w := md5.New()
	//io.WriteString(w, stringSignTemp)
	//md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	//
	//fmt.Println(strings.ToUpper(md5str2))
	a := 1
	b := 2
	c := 3

	if a != 0 && b != 2 || c == 3 {
		fmt.Println(1)
	}

}
