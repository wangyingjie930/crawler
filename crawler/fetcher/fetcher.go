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
	request.Header.Add("cookie","FSSBBIl1UgzbN7NO=5WptxEhYt3d3R5b4.1Q9BvisD_jd04no4OvFu.K.8nKfnKfdk77m1_DHWH7kZQFIjvbjBNOoHCyS5gHaTs0Igfq; sid=01a8a773-3f06-4707-9051-28b26766d056; ec=75yrgAH1-1618736315750-d9dc972dd26fc-1901410320; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1621675560,1622268128; _exid=OeLqsGCLmmBbWdFHsGrxwtyr4CyS4m1kL3M6nMgBV6jNfohM02i9B6AHTW4oeQCO%2Fv8Hrh8%2F4s%2FRrNwBtCGH2Q%3D%3D; _efmdata=h7q8CaYu7hFKflsNcW3NKSF0nzzRMLTcmWIhblIrwIo3K2BaCGXQHNzuQV%2FtbU5%2BSR9dTt4PHlRx5a5%2BT21q8cnq9V47cW7y9ZOckAbNQYk%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1622303565; FSSBBIl1UgzbN7NP=53hxj9bcnFLZqqqmyXHriFG0e.8mhjPi640NiQ_UjC1EixG5o1wXo7nIfKC6AlZ2HReHLwvnhZHYLRevbQHPG8eoRo54KRRrzWw4MxM8MSL3b9NYDdF2442CiR2OOq.YJfyCiK9TO5.g3U84p5HK7Qlb1gA_RjcJal9_UIhPWp5QUvCBHm9HBGaMevWgWavxtNg3.n9YnuK2rxTttqdMsd.VgS2R.I933_xEEVcLNFGCIyvrDV15mvx9pr73RuxMDG")

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
