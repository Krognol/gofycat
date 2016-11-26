package gofycat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (g *Cat) GetUsername(username string) (string, error) {
	res, err := g.request("GET", _User(username), nil)
	if err != nil {
		return "", err
	}

	switch {
	case res.StatusCode == 404:
		return fmt.Sprintf("The username '%s' was not found.", username), nil
	case res.StatusCode == 422:
		return fmt.Sprintf("The username '%s' is invalid.", username), nil
	case 200 >= res.StatusCode:
		return fmt.Sprintf("Username '%s' was found.", username), nil
	case res.StatusCode == 401:
		return fmt.Sprintf("Invalid token."), nil
	}
	return "", nil
}

func (g *Cat) EmailIsVerified() (bool, error) {
	res, err := g.request("GET", Verified(), nil)
	if err != nil {
		return false, err
	}

	switch {
	case res.StatusCode == 404:
		return false, nil
	case 200 >= res.StatusCode:
		return true, nil
	case res.StatusCode == 401:
		return false, fmt.Errorf("Invalid token.")
	}
	return false, nil
}

func (g *Cat) SendVerificationEmail() (success bool, err error) {
	res, err := g.request("GET", Verification(), nil)
	if err != nil {
		return false, err
	}

	switch {
	case res.StatusCode == 400:
		return false, fmt.Errorf("Unable to send verification email.")
	case res.StatusCode == 404:
		return false, fmt.Errorf("The user does not have an email registered.")
	case 200 >= res.StatusCode:
		return true, nil
	case res.StatusCode == 401:
		return false, fmt.Errorf("Invalid token.")
	}
	return false, nil
}

func (g *Cat) ResetPassword(usernameOrEmail string) (success bool, err error) {
	payload := struct {
		Value  string `json:"value"`
		Action string `json:"action"`
	}{Value: usernameOrEmail, Action: "send_password_reset_email"}
	js, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	res, err := g.request("PATCH", Users(), bytes.NewBuffer(js))
	if err != nil {
		return false, nil
	}

	switch {
	case 200 >= res.StatusCode:
		return true, nil
	case res.StatusCode == 400:
		return false, fmt.Errorf("Missing parameter action/value.")
	case res.StatusCode == 404:
		return false, fmt.Errorf("Username or email does not exists for this user.")
	case res.StatusCode == 422:
		return false, fmt.Errorf("You have exceeded the maximum amount of reset attempts or the email has not been registered")
	case res.StatusCode == 401:
		return false, fmt.Errorf("Unauthorized access")
	}
	return false, nil
}

type User struct {
	Userid                    string `json:"userid"`
	Username                  string `json:"username"`
	Email                     string `json:"email"`
	Description               string `json:"description"`
	ProfileURL                string `json:"profileUrl"`
	Name                      string `json:"name"`
	Views                     string `json:"views"`
	UploadNotices             bool   `json:"uploadNotices"`
	EmailVerified             bool   `json:"emailVerified"`
	URL                       string `json:"url"`
	CreateDate                string `json:"createDate"`
	ProfileImageURL           string `json:"profileImageUrl"`
	Verified                  bool   `json:"verified"`
	Followers                 string `json:"followers"`
	Following                 string `json:"following"`
	GeoWhitelist              string `json:"geoWhitelist"`
	DomainWhitelist           string `json:"domainWhitelist"`
	IFrameProfileImageVisible bool   `json:"iframeProfileImageVisible"`
	PublishedGfycats          string `json:"publishedGfycats"`
	PublishedAlbums           string `json:"publishedAlbums"`
}

func (g *Cat) GetUser(userid string) (*User, error) {
	u := &User{}
	res, err := g.request("GET", _User(userid), nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 401 {
		return nil, fmt.Errorf("Access token is invalid or has been revoked")
	}
	if err = unmarshal(res, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (g *Cat) GetSelfDetails() (*User, error) {
	u := &User{}
	res, err := g.request("GET", Me(), nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 401 {
		return nil, fmt.Errorf("Unauthorized. Token revoked or invalid.")
	}
	
	if err = unmarshal(res, u); err != nil {
		return nil, err
	}
	return u, nil
}

type Operation struct {
	Op        string   `json:"op"`
	Path      string   `json:"path"`
	Value     string   `json:"value"`
	Whitelist []string `json:"value,-"`
}

type Update struct {
	Operations []*Operation `json:"operations"`
}

func (g *Cat) UpdateDetails(u *Update) (success bool, err error) {
	js, err := json.Marshal(u)
	if err != nil {
		return false, err
	}

	res, err := g.request("PATCH", Me(), bytes.NewBuffer(js))
	if err != nil {
		return false, nil
	}

	switch res.StatusCode {
	case 200:
		return true, nil
	case 400:
		return false, fmt.Errorf("Bad parameter.")
	case 401:
		return false, fmt.Errorf("Unauthorized")
	default:
		return false, fmt.Errorf("Unknown error occured during the request")
	}
}

func (g *Cat) UpdateDomainWhitelist(domains []string) error {
	payload := struct {
		Domains []string `json:"domainWhitelist"`
	}{domains}

	js, _ := json.Marshal(payload)
	_, err := g.request("PUT", DomainWhitelist(), bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	return nil
}

func (g *Cat) GetDomainWhitelist() (*http.Response, error) {
	res, err := g.request("GET", DomainWhitelist(), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (g *Cat) DeleteDomainWhitelist() (*http.Response, error) {
	res, err := g.request("DELETE", DomainWhitelist(), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (g *Cat) UpdateGeoWhitelist(geolocations []string) error {
	payload := struct {
		Domains []string `json:"geoWhitelist"`
	}{geolocations}

	js, _ := json.Marshal(payload)
	_, err := g.request("PUT", GeoWhitelist(), bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	return nil
}

func (g *Cat) GetGeoWhitelist() (*http.Response, error) {
	res, err := g.request("GET",GeoWhitelist(), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (g *Cat) DeleteGeoWhitelist() (*http.Response, error) {
	res, err := g.request("DELETE", GeoWhitelist(), nil)
	if err != nil {
		return nil, err
	}
	return res, err
}

// TODO:
/*
func (g *Cat) UploadProfileImage(f *os.File) error {
	buf, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return err
	}
	res, err := http.Post("https://api.gfycat.com/v1/me/profile_image_url", "image/jpeg", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

}
*/

// TODO:
//func CreateNewUserAccount

func (g *Cat) FollowUser(user string) (int, error) {
	res, err := g.request("PUT", FollowsUser(user), nil)
	return res.StatusCode, err
}

func (g *Cat) UnfollowUser(user string) (int, error) {
	res, err := g.request("DELETE", FollowsUser(user), nil)
	return res.StatusCode, err
} 

func (g *Cat) CheckFollow(user string) (int, error) {
	res, err := g.request("HEAD", FollowsUser(user), nil)
	return res.StatusCode, err
}

func (g *Cat) GetFollows() (*http.Response, error) {
	return g.request("GET", Follows(), nil)
}

func (g *Cat) GetListOfFollowers() (*http.Response, error) {
	return g.request("GET", Followers(), nil)
}

// Gfycat object
type Gfycat struct {
	GfyID string `json:"gfyId"` 
	GfyNumber string `json:"gfyNumber"`
	WebmURL string `json:"webmUrl"`
	GifURL string `json:"gifUrl"`
	MobileURL string `json:"mobileUrl"`
	MobilePosterURL string `json:"mobilePosterUrl"`
	PosterURL string `json:"posterUrl"`
	Thumb360URL string `json:"thumb360Url"`
	Thumb360PosterURL string `json:"thumb360PosterUrl"`
	Thumb100PosterURL string `json:"thumb100PosterUrl"`
	Max5MBGif string `json:"max5mbGif"`
	Max2MBGif string `json:"max2mbGif"`
	MJPGURL string `json:"mjpgUrl"`
	Width int `json:"width"` 
	Height int `json:"height"`
	FrameRate int `json:"frameRate"`
	NumFrames int `json:"numFrames"`
	MP4Size int `json:"mp4Size"`
	WebmSize int `json:"webmSize"`
	GifSize int `json:"gifSize"`
	CreateDate string `json:"createDate"` 
	NSFW string `json:"nsfw"`
	MP4URL string `json:"mp4Url"`
	Likes int `json:"likes"`
	Published int `json:"published"`
	Dislikes int `json:"dislikes"`
	ExtraLemmas string `json:"extraLemmas"`
	MD5 string `json:"md5"`
	Views int `json:"views"`
	Tags interface{} `json:"tags"`
	Username string `json:"userName"`
	GfyName string `json:"gfyName"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UserFeed struct {
	Cursor string `json:"cursor"`
	Gfycats []Gfycat `json:"contents"`
}

func (g *Cat) GetUserFeed(user string) (*UserFeed, error) {
	feed := &UserFeed{}
	res, err := g.request("GET", UserCats(user), nil)
	if err != nil {
		return nil, err
	}
	if err = unmarshal(res, feed); err != nil {
		return nil, err 
	} 
	return feed, nil
}

func (g *Cat) GetPrivateFeed() (*UserFeed, error) {
	feed := &UserFeed{}
	res, err := g.request("GET", MyCats(), nil)
	if err != nil {
		return nil, err
	}
	if err = unmarshal(res, feed); err != nil {
		return nil, err
	}
	return feed, nil
}

func (g *Cat) GetTimeline() (*UserFeed, error) {
	feed := &UserFeed{}
	res, err := g.request("GET", FollowsCats(), nil)
	if err != nil {
		return nil, err
	}
	if err = unmarshal(res, feed); err != nil {
		return nil, err
	}
	return feed, nil
}

func (g *Cat) MyFolders() (*http.Response, error) {
	return g.request("GET", Folders(), nil)
}

func (g *Cat) GetFolder(folderid string) (*http.Response, error) {
	return g.request("GET", Folder(folderid), nil)
}

func (g *Cat) DeleteFolder(folderid string) (*http.Response, error) {
	return g.request("DELETE", Folder(folderid), nil)
}

func (g *Cat) GetFolderName(folderid string) (*http.Response, error) {
	return g.request("GET", FolderName(folderid), nil)
}

func (g *Cat) ChangeFolderName(newName, folderid string) (*http.Response, error) {
	payload := struct {
		Value string `json:"value"`
	}{newName}
	js, _ := json.Marshal(payload)
	return g.request("PUT", FolderName(folderid), bytes.NewBuffer(js))
}

func (g *Cat) MoveFolder(folderid, parentid string) (*http.Response, error) {
	payload := struct {
		ParentID string `json:"parentId"`
	}{parentid}
	js, _ := json.Marshal(payload)
	return g.request("PUT", Folder(folderid), bytes.NewBuffer(js))
}

func (g *Cat) MoveContentInFolder(folderid, parentid string, gfyids []string) (*http.Response, error) {
	payload := struct {
		Action string `json:"action"`
		ParentID string `json:"parent_id"`
		GfyIDs []string `json:"gfy_ids"`
	}{"move_contents", parentid, gfyids}
	js, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return g.request("PATCH", Folder(folderid), bytes.NewBuffer(js))
}

func (g *Cat) CreateFolder(foldername, insidefolder string) (*http.Response, error) {
	payload := struct {
		FolderName string `json:"folderName"`
	}{foldername}

	js, _ := json.Marshal(payload)
	if insidefolder != "" {
		return g.request("POST", Folder(insidefolder), bytes.NewBuffer(js))
	}
	return g.request("POST", Folders(), bytes.NewBuffer(js))
}