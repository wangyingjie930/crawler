package fetcher

/*func FetchSample(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic("request error")
	}
	//添加header
	request.Header.Set("Referer","http://www.zhenai.com/")
	request.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36")
	//这里cookie只有1分钟有效期
	request.Header.Add("cookie","FSSBBIl1UgzbN7NO=5WptxEhYt3d3R5b4.1Q9BvisD_jd04no4OvFu.K.8nKfnKfdk77m1_DHWH7kZQFIjvbjBNOoHCyS5gHaTs0Igfq; sid=01a8a773-3f06-4707-9051-28b26766d056; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1618657552; ec=75yrgAH1-1618736315750-d9dc972dd26fc-1901410320; _exid=6Pcg07Az51%2BcjK%2F65jMPIbWZ0b3lkUbMGhVTMMZlnibgBw8syDjQqVTzg81jrgwlCWwGa7Lf%2F4yEQgJM2XRW3w%3D%3D; _efmdata=h7q8CaYu7hFKflsNcW3NKSF0nzzRMLTcmWIhblIrwIo3K2BaCGXQHNzuQV%2FtbU5%2BAUIgTKJr03gkN7JyCiHBatIgVa3WNEcMMDUNBEjWwDE%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1618741977; FSSBBIl1UgzbN7NP=53oatCCrZczGqqqmgyENzGASoLVe62Mq1RsOyowENRtlouVM7QUXs25BNSGxPZn0K36q7V16QGUnmPNw5lj4733vB4taZsR.zLA26MGMFBFZ3F__FFNieUrjehyT.xdgF9MVlkGweN8Zl_AXD73XBU92741qpuC8GL8UDMgZrUwtiMWhylRxxv84yNoqCr7bDqJNBRB0gxU9bPNhk8KS8AKrOzvUKur.PPnsjRbsbVSL4gHV6oPp9Od3WlT2117zOv_u8VmYT4ziGMHRZkDbOkX")

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d",res.StatusCode)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	return data, nil
}*/
