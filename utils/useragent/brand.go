package useragent

import (
    "io/ioutil"
    "tracking/utils/log"
    "encoding/json"
)

const filePath string = "Brand.json"

var brandsMap = make(map[string]string)

func init()  {
    loadBrand()
}

func loadBrand() {
    bytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        log.Logger().Errorf("[useragent][brand]Read Brand File Error: %v\n", err)
    }
    if err := json.Unmarshal(bytes, &brandsMap); err != nil {
        log.Logger().Errorf("[useragent][brand]JSON Unmarshar Brand Error:%v\n", err)
    }
    log.Logger().Infof("BrandsMap Len: %v\n", len(brandsMap))
}

func getBrand(model string) string {
    if brandsMap[model] != "" {
        return brandsMap[model]
    }
    return "Unknown"
}
