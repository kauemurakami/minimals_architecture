package followers

import (
	followers_functions "api-social-media/app/internal/followers/functions"
	"net/http"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followers_functions.FollowUser(w, r)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followers_functions.UnfollowUser(w, r)

}
