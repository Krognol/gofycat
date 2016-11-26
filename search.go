package gofycat

func (c *Cat) SearchGfycats(searchText string) ([]Gfycat, error) {
    var arr []Gfycat

    res, err := c.request("GET", GfycatSearch(searchText), nil)

    if err != nil {
        return nil, err
    }

    if err = unmarshal(res, &arr); err != nil {
        return nil, err
    }

    return arr, nil
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

