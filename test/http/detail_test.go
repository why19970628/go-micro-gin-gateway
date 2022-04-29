package http

import (
	"bytes"
	"errors"
	"github.com/golang/protobuf/proto"
	pb "go-micro-gin-gateway/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

type Params struct {
	Req   proto.Message
	Res   proto.Message
	Token string
	Url   string
}

func TestAppraisalBrandsHandler(t *testing.T) {
	p := &Params{
		Req: &pb.UserDetailRequest{
			UserId: 123456,
		},
		Res:   &pb.UserDetailResponse{},
		Url:   "http://0.0.0.0:9001/UserCommonService/GetUserDetail",
		Token: "",
	}
	ApiTest(p)
}

func ApiTest(p *Params) {
	data, err := proto.Marshal(p.Req)
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(data)
	req, _ := http.NewRequest("POST", p.Url, r)
	req.Header.Add("Content-Type", "application/x-protobuf")
	req.Header.Add("Authorization", p.Token)

	body, err := DoRequest(req)
	if err != nil {
		panic(err)
	}
	res := p.Res
	err = proto.Unmarshal(body, res)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.String())
}

func DoRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(strconv.Itoa(resp.StatusCode))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
