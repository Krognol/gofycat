package gofycat

import (
	"net/http"
	"encoding/json"
	"bytes"
)

func (c *Cat) GetBookmarks() (*http.Response, error) {
    return c.request("GET", BookmarkFolders(), nil)
}

func (c *Cat) GetBookmarkFolderContent(folderid string) (*http.Response, error) {
    return c.request("GET", BookmarkFolder(folderid), nil)
}

func (c *Cat) GfyIsBookmarked(gfy string) (*http.Response, error) {
    return c.request("GET", BookmarkGfy(gfy), nil)
}

func (c *Cat) BookmarkGfy(id string) (*http.Response, error) {
    return c.request("PUT", BookmarkGfy(id), nil)
}

func (c *Cat) BookmarkGfyInFolder(folder, gfy string) (*http.Response, error) {
    return c.request("PUT", BookmarkFolderContents(folder, gfy), nil)
}

func (c *Cat) UnbookmarkGfyInFolder(folder, gfy string) (*http.Response, error) {
    return c.request("DELETE", BookmarkFolderContents(folder, gfy), nil)
}

func (c *Cat) MoveToFolder(folder, parent string, gfys []string) (*http.Response, error) {
    payload := struct {
        Action string `json:"action"`
        Parent string `json:"parent_id"`
        Gfys []string `json:"gfy_ids"`
    }{"move_contents", parent, gfys}
    js, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }
    return c.request("PATCH", BookmarkFolder(folder), bytes.NewBuffer(js))
}

func (c *Cat) UnbookmarkGfy(gfy string) (*http.Response, error) {
    return c.request("DELETE", BookmarkGfy(gfy), nil)
}