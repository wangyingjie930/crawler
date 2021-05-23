package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)


var rateLimiter = time.Tick(200 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	newUrl := strings.Replace(url, "http://","https://",1)
	request, err := http.NewRequest(http.MethodGet, newUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	//添加header
	request.Header.Set("Referer","http://www.zhenai.com/")
	request.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36")
	//这里cookie只有1分钟有效期
	request.Header.Add("cookie","FSSBBIl1UgzbN7NO=5WptxEhYt3d3R5b4.1Q9BvisD_jd04no4OvFu.K.8nKfnKfdk77m1_DHWH7kZQFIjvbjBNOoHCyS5gHaTs0Igfq; sid=01a8a773-3f06-4707-9051-28b26766d056; ec=75yrgAH1-1618736315750-d9dc972dd26fc-1901410320; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621675560; _exid=Lm5JAZ96WD2DVgtSxoPa10u7Qoy94kOkoRlOSkpyZeC87Ew086M8dhEd%2BNYBR5PL925IWBmF0C9ggExah8Mr3Q%3D%3D; _efmdata=h7q8CaYu7hFKflsNcW3NKSF0nzzRMLTcmWIhblIrwIo3K2BaCGXQHNzuQV%2FtbU5%2Bv6cGev%2FP4dT8od3qjPUlqoj%2BH%2F9q%2BYWHS8nxofa3CZg%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621771301; FSSBBIl1UgzbN7NP=53tNltCcjpcQqqqmybpNhKGO3UnGRj24guIM5KHhMXV0red2hEWHxlpGPKpR10Uo7v3reM_2Gs8mN3sdqSWy_nHsc_HPTFPDrQEzNpifnPRp8.Y5Zw9.TWC9Ots.pPFGitCSWCNhP_dor6H3giNBfQ1TKTVz2fCI1MX2HkKLGCvXWvVMbiHLS0zI3y4m4dPa68AOxM37_zH8adBUzdiB4wtz3yQReg1oMhIgGHWOAQZNwaS8q0cpU_6ELGabWxBmw3")

	resp , err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//fmt.Printf("wrong status code: %d",resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	uft8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	return  ioutil.ReadAll(uft8Reader)
}


func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error : %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
