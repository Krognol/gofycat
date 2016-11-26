package gofycat

import (
	"net/http"
	"net/url"
)

// Count is 10 by default
func (c *Cat) GetTrendingGfycats(tag, cursor string) ([]Gfycat, error) {
    var arr []Gfycat
    var res *http.Response
    var err error
    
    url := TrendingGfycats() + "?count=10"

    if tag != "" {
        url += "&"+tag
    }  

    if cursor != "" {
        url += "&cursor="+cursor
    }

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, &arr); err != nil {
        return nil, err
    }

    return arr, nil
}

func (c *Cat) GetTrendingTags() ([]string, error) {
    var arr []string
    res, err := c.request("GET", TrendingTags(), nil)

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, &arr); err != nil {
        return nil, err
    }

    return arr, nil
}

func (c *Cat) GetTrendingTagsPopulated(values url.Values) ([]Gfycat, error) {
    var arr []Gfycat

    url := TrendingTagsPopulated() + values.Encode()

    res, err := c.request("GET", url, nil)

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, &arr); err != nil {
        return nil, err
    } 

    return arr, nil
}
