package httpUtil
import (
"net/http"
"net"
"time"
"strings"
"fmt"
"log"
"io/ioutil"
"encoding/json"
"bytes"
	"github.com/bitly/go-simplejson"
)

var (
	httpClient *http.Client
)

const(
	MaxIdleConns int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout int = 90
	Timeout			= 10
	Keepalive		= 30
)

func init(){
	httpClient = createHttpClient()
}

func createHttpClient() *http.Client{
	client := &http.Client{
		Transport:&http.Transport{
			Proxy:http.ProxyFromEnvironment,
			DialContext:(&net.Dialer{
				Timeout:Timeout*time.Second,
				KeepAlive:Keepalive*time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:	 time.Duration(IdleConnTimeout)* time.Second,
		},
		Timeout:Timeout*time.Second,
	}
	return client
}

func Get(url string,params map[string]string) string{
	if params != nil && len(params) >0{
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%s?",url))
		for k,v := range params{
			sb.WriteString(fmt.Sprintf("%s=%s&",k,v))
		}
		url = sb.String()
		url = url[0:len(url)-1]
	}
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Println("create request fail")
		return ""
	}
	resp,err := httpClient.Do(req)
	if err != nil {
		log.Println("do request fail")
		return ""
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	if status != 200 {
		log.Printf("url:%s request code is %d\n",url,status)
		return ""
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("url:%s read body fail",url)
		return ""
	}
	return string(body)
}

func Post(url ,params string)string{
	req,err := http.NewRequest("POST",url,strings.NewReader(params))
	if err != nil {
		log.Println("create request fail")
		return ""
	}
	resp,err := httpClient.Do(req)
	if err != nil {
		log.Println("do request fail")
		return ""
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	if status != 200 {
		log.Printf("url:%s request code is %d\n",url,status)
		return ""
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("url:%s read body fail",url)
		return ""
	}
	return string(body)
}

func JsonGet(url string,params map[string]interface{}) (*simplejson.Json){
	if params != nil && len(params) >0{
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%s?",url))
		for k,v := range params{
			sb.WriteString(fmt.Sprintf("%s=%s&",k,v))
		}
		url = sb.String()
		url = url[0:len(url)-1]
	}
	req,err := http.NewRequest("GET",url,nil)
	req.Header.Set("Content-Type","application/json")
	if err != nil {
		log.Println("create request fail")
		return nil
	}
	resp,err := httpClient.Do(req)
	if err != nil {
		log.Println("do request fail")
		return nil
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	if status != 200 {
		log.Printf("url:%s request code is %d\n",url,status)
		return nil
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("url:%s read byte fail\n",url)
	}
	json,err := simplejson.NewJson(body)
	if err != nil {
		log.Printf("url:%s simple json unmarch fail\n",url)
		return nil
	}
	return json
}

func JsonPost(url string,params map[string]interface{})*simplejson.Json{
	b,e := json.Marshal(params)
	if e != nil {
		log.Printf("json marshal params fail")
		return nil
	}
	req,err := http.NewRequest("POST",url,bytes.NewReader(b))
	req.Header.Set("Content-Type","application/json")
	if err != nil {
		log.Println("create request fail")
		return nil
	}
	resp,err := httpClient.Do(req)
	if err != nil {
		log.Println("do request fail")
		return nil
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	if status != 200 {
		log.Printf("url:%s request code is %d\n",url,status)
		return nil
	}
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("url:%s read byte fail\n",url)
	}
	json,err := simplejson.NewJson(body)
	if err != nil {
		log.Printf("url:%s simple json unmarch fail\n",url)
		return nil
	}
	return json
}
