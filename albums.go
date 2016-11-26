package gofycat

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

type Action string

const (
    // None no action 
    None Action = ""
    Add = "add_to"
    Move = "move_contents"
    Remove = "remove_contents"
)


func (c *Cat) GetAlbumFolders() (*http.Response, error) {
    return c.request("GET", AlbumFolders(), nil)
}


// TODO: Check following two functions
func (c *Cat) GetUserAlbumContents(user, album string) (*http.Response, error) {
    return c.request("GET", UserAlbums(user, album), nil)
}

// jesus
func (c *Cat) GetUserAlbumContentsByLinkText(user, linkText string) (*http.Response, error) {
    return c.request("GET", UserAlbumLinkTextContents(user, linkText), nil)
}

func (c *Cat) GetAlbumContents(album string) (*http.Response, error) {
    return c.request("GET", Album(album), nil)
}

func (c *Cat) CreateAlbum(name string) (*http.Response, error) {
    return c.request("POST", Album(name), nil)
}

func (c *Cat) MoveAlbumToFolder(parentid, album string) (*http.Response, error) {
    payload := struct {
        Parent string `json:"parentId"`
    }{parentid}
    js, _ := json.Marshal(payload)
    return c.request("PUT", Album(album), bytes.NewBuffer(js))
}

func (c *Cat) UpdateAlbum(action Action, album, parentid string, a ...string) (*http.Response, error) {
    switch action {
        case Remove:
        
        payload := struct {
            Action string `json:"action"`
            Contents []string `json:"contents"`
        }{fmt.Sprintf("%s", action), a}
        
        js, err := json.Marshal(payload)
        
        if err != nil {
            return nil, err
        }
        return c.request("PATCH", Album(album), bytes.NewBuffer(js))
        
        case Move:
        payload := struct {
            Action string `json:"action"`
            Parent string `json:"parent_id"`
            Gfys []string `json:"gfy_ids"` 
        }{fmt.Sprintf("%s", action), parentid, a}

        js, err := json.Marshal(payload)

        if err != nil {
            return nil, err
        }

        return c.request("PATCH", Album(album), bytes.NewBuffer(js))

        case Add:

        payload := struct {
            Gfys []string `json:"gfy_ids"`
        }{a}

        js, err := json.Marshal(payload)

        if err != nil {
            return nil, err
        }
        return c.request("PATCH", Album(album), bytes.NewBuffer(js))
        default:
        return nil, fmt.Errorf("Invalid action '%s'", action)
    }
}

func (c *Cat) CreateAlbumInFolder(folder, title, description string, contents []string) (*http.Response, error) {
    payload := struct {
        Title string `json:"title"`
        Description string `json:"description"`
        Contents []string `json:"contents"`
    }{title, description, contents}
    js, err := json.Marshal(payload)
    
    if err != nil {
        return nil, err
    }

    return c.request("POST", AlbumFolder(folder), bytes.NewBuffer(js))
}

func (c *Cat) PutAlbumTitle(album, value string) (*http.Response, error) {
    payload := struct {
        Value string `json:"value"`
    }{value}

    js, _ := json.Marshal(payload)

    return c.request("PUT", AlbumTitle(album), bytes.NewBuffer(js))
}

func (c *Cat) PutAlbumDescription(album, value string) (*http.Response, error) {
    payload := struct {
        Value string `json:"value"`
    }{value}

    js, _ := json.Marshal(payload)

    return c.request("PUT", AlbumDescription(album), bytes.NewBuffer(js))
}

func (c *Cat) PutAlbumNsfw(value int, album string) (*http.Response, error) {
    payload := struct {
        Value int `json:"value"`
    }{value}

    js, _ := json.Marshal(payload)

    return c.request("PUT", AlbumNsfw(album), bytes.NewBuffer(js))
}

func (c *Cat) PutPublished(value int, album string) (*http.Response, error) {
    payload := struct {
        Value int `json:"value"`
    }{value}

    js, _ := json.Marshal(payload)

    return c.request("PUT", AlbumPublished(album), bytes.NewBuffer(js))
}