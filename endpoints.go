package gofycat

import (
	"io"
	"net/http"
)

func (c *Cat) request(method, url string, body io.Reader) (*http.Response, error) {
	var err error
	c.auth, err = c.authenticate(c.Type, c.ClientID, c.ClientSecret)
	
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", "Bearer " + c.auth.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

var Base = func() string { return "https://api.gfycat.com/v1/" }

var Authentication = func() string { return Base() + "oauth/token" }

var Users = func() string { return Base() + "users/" }

var _User = func(user string) string { return Users() + user }

var UserCats = func(user string) string { return _User(user) + "/gfycats" }


// ======= DON'T USE THESE ENDPOINTS ========
var Me = func() string { return Base() + "me"}

var MyCats = func() string { return Me() + "/gfycats" }

var DomainWhitelist = func() string { return Me() + "/domain-whitelist" }

var GeoWhitelist = func() string { return Me() + "/geo-whitelist" }

var Verified = func() string { return Me() + "/email_verified" }

var Verification = func() string { return Me() + "/send_verification_email" }

var Follows = func() string { return Me() + "/follows" }

var FollowsUser = func(user string) string { return Follows() + "/" + user}

var Followers = func() string { return Me() + "/followers" }

var FollowsCats = func() string { return Follows() + "/gfycats" }

var Folders = func() string { return Me() + "/folders" }

var Folder = func(folderid string) string { return Folders() + "/" + folderid }

var FolderName = func(folderid string) string { return Folder(folderid) + "/name" }

var BookmarkFolders = func() string { return Me() + "/bookmark-folders" }

var BookmarkFolder = func(id string) string { return BookmarkFolders() + "/" + id }

var Bookmarks = func() string { return Me() + "/bookmarks" }

var BookmarkGfy = func(id string) string { return Bookmarks() + "/" + id }

var BookmarkFolderContents = func(fid, gid string) string { return BookmarkFolder(fid) + "/contents/" + gid }

var AlbumFolders = func() string { return Me() + "/album-folders" }

var Albums = func() string { return Me() + "/albums" }

var Album = func(aid string) string { return Albums() + "/" + aid}

var AlbumFolder = func(fid string) string { return AlbumFolders() + "/" + fid }

var AlbumTitle = func(aid string) string { return Album(aid) + "/title" }

var AlbumDescription = func(aid string) string { return Album(aid) + "/description" }

var AlbumNsfw = func(aid string) string { return Album(aid) + "/nsfw" }

var AlbumPublished = func(aid string) string { return Album(aid) + "/published" }

var AlbumOrder = func(aid string) string { return Album(aid) + "/order" } 

var _Gfycat = func(gid string) string { return Gfycats() + "/" + gid }

var GfycatTitle = func(gid string) string { return _Gfycat(gid) + "/title" }

var GfycatDescription = func(gid string) string { return _Gfycat(gid) + "/description" }

var GfycatTags = func(gid string) string { return _Gfycat(gid) + "/tags" }

var GfycatPublished = func(gid string) string { return _Gfycat(gid) + "/published" }

var GfycatNsfw = func(gid string) string { return _Gfycat(gid) + "/nsfw" }

var GfycatDomainWhitelist = func(gid string) string { return _Gfycat(gid) + "/domain-whitelist" }

var GfycatGeoWhitelist = func(gid string) string { return _Gfycat(gid) + "/geo-whitelist" }



// ========= END 'DON'T USE THESE ENDPOINTS' ========

var UserAlbums = func(uid, aid string) string { return Users() + "/" + uid + "/albums/" + aid }

var UserAlbumLinkTextContents = func(uid, lt string) string { return Users() + uid + "/album_link/" + lt }

var Gfycats = func() string { return Base() + "gfycats" }

var TrendingGfycats = func() string { return Gfycats() + "/trending" }

var TrendingGfycatsTag = func(tag string) string { return TrendingGfycats() + "?" + tag}

var Tags = func() string { return Base() + "tags" }

var TrendingTags = func() string { return Tags() + "/trending" }

var TrendingTagsPopulated = func() string { return TrendingTags() + "/populated" }

var TestApi = func() string { return "https://api.gfycat.com/v1test" }

var GfycatSearch = func(text string) string { return TestApi() + "/gfycats/search?search_text="+text}

var UserSearch = func(text string) string { return TestApi() + "/me/gfycats/search?search_text=" + text}
