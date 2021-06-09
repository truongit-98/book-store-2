package commonservice

import (
	"BookStore/restapi/responses"
	"BookStore/utils"
	"bufio"
	"fmt"
	"github.com/prometheus/common/log"
	"os"
	"path"
	"strings"
)

func LoadCities() ([]byte, error) {
	pathFile := path.Join(utils.GetCurrentPath(), "/data/hanhchinhvn/tinh_tp.json")
	jsonFile, err := os.Open(pathFile)
	if err != nil {
		log.Info(err.Error(), "commonservice/commonservice")
		return make([]byte, 0), responses.ErrSystem
	}
	defer jsonFile.Close()
	var result = make([]byte, 0)
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan(){
	 	result = append(result, scanner.Bytes()...)
	}
	return result, nil
}

func LoadDistrictForCity(cityId string) ([]byte, error) {
	pathFile := path.Join(utils.GetCurrentPath(), fmt.Sprintf("/data/hanhchinhvn/quan-huyen/%s.json", cityId))
	jsonFile, err := os.Open(pathFile)
	if err != nil {
		log.Info(err.Error(), "commonservice/commonservice")
		return make([]byte, 0), responses.ErrSystem
	}
	defer jsonFile.Close()
	var result = make([]byte, 0)
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan(){
		result = append(result, scanner.Bytes()...)
	}
	return result, nil

}

func LoadWardsOfDistrict(districtId string) ([]byte, error) {
	pathFile := path.Join(utils.GetCurrentPath(), fmt.Sprintf("/data/hanhchinhvn/xa-phuong/%s.json", districtId))
	jsonFile, err := os.Open(pathFile)
	if err != nil {
		log.Info(err.Error(), "commonservice/commonservice")
		if strings.Contains(err.Error(), " no such file") {
			return make([]byte, 0), responses.NotExisted
		}
		return make([]byte, 0), responses.ErrSystem
	}
	defer jsonFile.Close()
	var result = make([]byte, 0)
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan(){
		result = append(result, scanner.Bytes()...)
	}
	return result, nil
}



