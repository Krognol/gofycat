package gofycat

type GfycatSearchT struct {
    Gfycats []Gfycat `json:"gfycats"`
    Found int `json:"found"`
    Cursor string `json:"cursor"`
}

func (c *Cat) SearchGfycats(searchText string) (*GfycatSearchT, error) {
    g := &GfycatSearchT{}

    res, err := c.request("GET", GfycatSearch(searchText), nil)

    if err != nil {
        return nil, err
    }
    
    if err = unmarshal(res, g); err != nil {
        return nil, err
    }

    return g, nil
}

func (c *Cat) SearchUserAccount(searchText string) ([]Gfycat, error) {
    var arr []Gfycat

    res, err := c.request("GET", UserSearch(searchText), nil)

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, &arr); err != nil {
        return nil, err
    }

    return arr, nil
}

