package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://stageapp.tekioncloud.xyz/api/documentV2/u/docset/search?locale=en_US"
	method := "POST"

	payload := strings.NewReader(`{"sort":[{"field":"createdTime","order":"DESC"},{"field":"id","order":"ASC"}],"filters":[],"searchText":"","groupBy":[],"includeFields":[],"searchableFields":[],"excludeFields":[],"pageInfo":{"start":0,"rows":50}}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("dealerid", "4")
	req.Header.Add("roleid", "4_SuperAdmin")
	req.Header.Add("tekion-api-token", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlF6SXdOREF3UXpOQk0wRkVNelZFTmpNMU5UQXhSREEzUXpjM1FqVkZNalUxTVRrME1EbERSQSJ9.eyJuaWNrbmFtZSI6ImdhdXJhdnByaW50aW5nIiwibmFtZSI6ImdhdXJhdnByaW50aW5nQG1haWxpbmF0b3IuY29tIiwicGljdHVyZSI6Imh0dHBzOi8vcy5ncmF2YXRhci5jb20vYXZhdGFyLzJiMTZmMzNiYjY1YjNjYzZmOTc1MjI2MDlmMjk0ZDgyP3M9NDgwJnI9cGcmZD1odHRwcyUzQSUyRiUyRmNkbi5hdXRoMC5jb20lMkZhdmF0YXJzJTJGZ2EucG5nIiwidXBkYXRlZF9hdCI6IjIwMjMtMDMtMTRUMDg6MzY6MjEuMTM5WiIsImVtYWlsIjoiZ2F1cmF2cHJpbnRpbmdAbWFpbGluYXRvci5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6Ly90ZWtpb24uYXV0aDAuY29tLyIsImF1ZCI6ImxNU3pxYlM3OEVua1hwNXJoWTdaZFZkZHMwQVh2QzRSIiwiaWF0IjoxNjc4NzgyOTgxLCJleHAiOjE2Nzg4MTg5ODEsInN1YiI6ImF1dGgwfDYwMzUwMmQ5NmIzYTYzMDA2OGE4OGYyOCJ9.nuQeyZk4maOhepJUSWSJTVL85A3_Dh88Z_xN_Sm4f5tr9mZrugkZW7ZhJHHIWdsYDXbOCbD2baac5f9VPMQapKJ9BiOfPwKvUmxoMbVGvFSz_2mo0iVxaMsfOHt5Sh4uZpyK3_8fpDsHXNOOAdnV_icKarDsiWkEL6HdF2OQXHBK9-Gr3_qFuLBmCEdZ6Bk0DmqPsIV8HA0spgD5UY_WPbn-f1UNa_S8QGkAWmy9OzKXp_ijD0lZpA7oeJ20WfexVpuiDVZWr3DuhUY3YavGu2BKoHlOE3jW68YdUe7wK9Ks0liSJMMSRTQO3JGdKhdbCsAKJOqvEyNt-cawt1uWlQ")
	req.Header.Add("tenantid", "undefined")
	req.Header.Add("tenantname", "stg-cacargroup")
	req.Header.Add("userid", "32b87bf1-1361-4c29-953f-871580c4b36d")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
