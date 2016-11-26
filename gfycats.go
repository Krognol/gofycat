package gofycat

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type Status struct {
    
    // While encoding or not found
    Task string `json:"task"`
    Time int `json:"time"`

    // When complete
    GfyName string `json:"gfyname"`

}

func (c *Cat) CheckStatusOfUpload(name string) (*Status, error) {
    s := &Status{}

    res, err := c.request("GET", Gfycats() + "/fetch/status/"+name, nil)

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, s); err != nil {
        return nil, err
    }

    return s, err
}

type UploadRequest struct {
    GfyName string `json:"gfyname"`
    Secret string `json:"secret"`
}

type Caption struct {
    Text string `json:"text"`
    StartSeconds int `json:"startSeconds"`
    Duration int `json:"duration"`
    FontHeight int `json:"fontHeight"`
    X int `json:"x"`
    Y int `json:"y"`
    FontHeightRelative float32 `json:"fontHeightRelative"`
    XRelatice float32 `json:"xRelative"`
    YRelative float32 `json:"yRelative"`
}

type Cut struct {
    Duration int `json:"duration"`
    Start int `json:"start"`
}

type Crop struct {
    X int `json:"x"`
    Y int `json:"y"`
    W int `json:"w"`
    H int `json:"h"`
}

type UploadFile struct {
    FetchUrl string `json:"fetchUrl"`
    Title string `json:"title"`
    Description string `json:"description"`
    Tags []string `json:"tags"`
    NoMd5 string `json:"noMd5"`
    Private int `json:"private"`
    Nsfw int `json:"nsfw"`
    FetchSeconds int `json:"fetchSeconds"`
    FetchMinutes int `json:"fetchMinutes"`
    FethHours int `json:"fetchHours"`
    Captions []Caption `json:"captions"`
    Cut []Cut `json:"cut"`
    Crop []Crop `json:"crop"`
}

func (c *Cat) UploadFile(file *UploadFile) (*http.Response, error){
    ur := &UploadRequest{}
    
    js, err := json.Marshal(file)
    
    if err != nil {
        return nil, err
    }
    
    res, err := c.request("POST", Gfycats(), bytes.NewBuffer(js))

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, ur); err != nil {
        return nil, err
    }

    r, _ := http.NewRequest("", "", nil)
    mf, _, err := r.FormFile("./"+ur.GfyName)

    if err != nil {
        return nil, err
    }

    return c.request("POST", "https://filedrop.gfycat.com/", mf)
}

func (c *Cat) UpdateGfycatTitle(gfyid, title string) (*http.Response, error) {
    payload := struct {
        Value string `json:"value"`
    }{title}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatTitle(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) DeleteGfycatTitle(gfyid string) (*http.Response, error) {
    return c.request("DELETE", GfycatTitle(gfyid), nil)
}

func (c *Cat) UpdateGfycatTags(gfyid string, tags []string) (*http.Response, error) {
    payload := struct {
        Value []string `json:"value"`
    }{tags}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatTags(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) GetGfycatDomainWhitelist(gfyid string) (*http.Response, error) {
    return c.request("GET", GfycatDomainWhitelist(gfyid), nil)
}

func (c *Cat) UpdateGfycatDomainWhitelist(gfyid string, domains []string) (*http.Response, error) {
    payload := struct {
        DomainWL []string `json:"domainWhitelist"`
    }{domains}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatDomainWhitelist(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) DeleteGfycatDomainWhitelist(gfyid string) (*http.Response, error) {
    return c.request("DELETE", GfycatDomainWhitelist(gfyid), nil)
}

func (c *Cat) GetGfycatGeoWhitelist(gfyid string) (*http.Response, error) {
    return c.request("GET", GfycatGeoWhitelist(gfyid), nil)
}

func (c *Cat) UpdateGfycatGeoWhitelist(gfyid string, geoLocations []string) (*http.Response, error) {
    payload := struct {
        GeoWL []string `json:"geoWhitelist"`
    }{geoLocations}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatGeoWhitelist(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) DeleteGfycatGeoWhitelist(gfyid string) (*http.Response, error) {
    return c.request("DELETE", GfycatGeoWhitelist(gfyid), nil)
}

func (c *Cat) UpdateGfycatDescription(gfyid, desc string) (*http.Response, error) {
    payload := struct {
        Value string `json:"value"`
    }{desc}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatDescription(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) DeleteGfycatDescription(gfyid, desc string) (*http.Response, error) {
    return c.request("DELETE", GfycatDescription(gfyid), nil)
}

func (c *Cat) UpdateGfycatPublished(gfyid, pub string) (*http.Response, error) {
    payload := struct {
        Value string `json:"value"`
    }{pub}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatPublished(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) UpdateGfycatNsfw(gfyid, nsfw string) (*http.Response, error) {
    payload := struct {
        Value string `json:"value"`
    }{nsfw}
    js, _ := json.Marshal(payload)
    return c.request("PUT", GfycatNsfw(gfyid), bytes.NewBuffer(js))
}

func (c *Cat) DeleteGfycat(gfyid string) (*http.Response, error) {
    return c.request("DELETE", _Gfycat(gfyid), nil)
}