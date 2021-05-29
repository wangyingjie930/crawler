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
	request.Header.Add("cookie","FSSBBIl1UgzbN7NO=5WptxEhYt3d3R5b4.1Q9BvisD_jd04no4OvFu.K.8nKfnKfdk77m1_DHWH7kZQFIjvbjBNOoHCyS5gHaTs0Igfq; sid=01a8a773-3f06-4707-9051-28b26766d056; ec=75yrgAH1-1618736315750-d9dc972dd26fc-1901410320; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621675560,1622268128; _exid=lU%2FDmoNdG8We6fIOd4ITAKqoRa9mXf7MjgGt%2B2bZFKOCqFuZS6calNx5QaW1QgoWsIwWkOj7gYPy1KBjTugeTQ%3D%3D; _efmdata=h7q8CaYu7hFKflsNcW3NKSF0nzzRMLTcmWIhblIrwIo3K2BaCGXQHNzuQV%2FtbU5%2B4xF2pOuPa9xWlbOrbBBXZTk5WpDH1FWhGUa6yplMKDU%3D; FSSBBIl1UgzbN7NP=53hEVlbcnmqZqqqmyzyy.rGCeBTHdtVsWz8HRhX4AFCGvGv.r1kB4MXamfSB21TnHHKGIEhi2w1v1IlKd3uL6JxFns6hlf.gGliHGvy8HXek5F8t0D7hzRK8G5i2Au5Bh016SDVC8_WnPu.fKqOrmdIRBTzQrmgWaUC3vSJ80HmNu3MRyyuItXg2Rd_CkO2e7Cq96Ra3FoQLn9VDpE5YYBpNfb1NiyZBDm3o9QdJFYGRi.7sob0PU_08DNt_OfsVqE; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1622283280")

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
