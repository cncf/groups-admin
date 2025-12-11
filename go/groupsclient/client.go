package groupsclient

import (
	"encoding/json"
	. "fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// TokenResponse structure to hold the JSON token response
type TokenResponse struct {
	Token string `json:"token"`
}

// GroupsClient struct to hold client configuration
type GroupsClient struct {
	BaseURL string
	Token   string
	Client  *http.Client
}
type Org struct {
	ID                  int    `json:"id"`
	Object              string `json:"object"`
	Created             string `json:"created"`
	Updated             string `json:"updated"`
	Title               string `json:"title"`
	Domain              string `json:"domain"`
	ParentGroupID       int    `json:"parent_group_id"`
	LoggedOutWikiPageID int    `json:"logged_out_wiki_page_id"`
	DefaultTimezone     string `json:"default_timezone"`
	DisableSignup       bool   `json:"disable_signup"`
	DisablePlusOne      bool   `json:"disable_plus_one"`
	GaCode              string `json:"ga_code"`
	LoginPageText       string `json:"login_page_text"`
	NoAccountText       string `json:"no_account_text"`
	SsoProvider         string `json:"sso_provider"`
	SsoClientID         string `json:"sso_client_id"`
	SsoClientSecret     string `json:"sso_client_secret"`
	SsoDomain           string `json:"sso_domain"`
}
type User struct {
	ID                      int    `json:"id"`
	Object                  string `json:"object"`
	Created                 string `json:"created"`
	Updated                 string `json:"updated"`
	Email                   string `json:"email"`
	FullName                string `json:"full_name"`
	UserName                string `json:"user_name"`
	Timezone                string `json:"timezone"`
	Status                  string `json:"status"`
	ProfilePhotoURL         string `json:"profile_photo_url"`
	PostPref                string `json:"post_pref"`
	PerPagePref             string `json:"per_page_pref"`
	AllowFacebookLogin      bool   `json:"allow_facebook_login"`
	AllowGoogleLogin        bool   `json:"allow_google_login"`
	AllowSsoLogin           bool   `json:"allow_sso_login"`
	CsrfToken               string `json:"csrf_token"`
	TwoFactorEnabled        bool   `json:"two_factor_enabled"`
	RecoveryCodes           string `json:"recovery_codes"`
	DontMungeMessageID      bool   `json:"dont_munge_message_id"`
	AboutMe                 string `json:"about_me"`
	AboutFormat             string `json:"about_format"`
	Location                string `json:"location"`
	Website                 string `json:"website"`
	TimePref                string `json:"time_pref"`
	DatePref                string `json:"date_pref"`
	MondayStart             bool   `json:"monday_start"`
	ProfilePrivacy          string `json:"profile_privacy"`
	DefaultMessageView      string `json:"default_message_view"`
	TopicsSortDir           string `json:"topics_sort_dir"`
	TopicSortDir            string `json:"topic_sort_dir"`
	MessagesSortDir         string `json:"messages_sort_dir"`
	ExpandedMessagesSortDir string `json:"expanded_messages_sort_dir"`
	SearchSort              string `json:"search_sort"`
	SearchSortDir           string `json:"search_sort_dir"`
	PhotosOrderBy           string `json:"photos_order_by"`
	PhotosSortDir           string `json:"photos_sort_dir"`
	AlbumOrderBy            string `json:"album_order_by"`
	AlbumSortDir            string `json:"album_sort_dir"`
	DefaultCalendarView     string `json:"default_calendar_view"`
	DefaultHashtagView      string `json:"default_hashtag_view"`
	DefaultRsvpView         string `json:"default_rsvp_view"`
	HomePage                string `json:"home_page"`
}
type MemberInfo struct {
	ID                  int       `json:"id"`
	Object              string    `json:"object"`
	Created             string    `json:"created"`
	Updated             string    `json:"updated"`
	UserID              int       `json:"user_id"`
	GroupID             int       `json:"group_id"`
	GroupName           string    `json:"group_name"`
	Status              string    `json:"status"`
	PostStatus          string    `json:"post_status"`
	EmailDelivery       string    `json:"email_delivery"`
	MessageSelection    string    `json:"message_selection"`
	AutoFollowReplies   bool      `json:"auto_follow_replies"`
	MaxAttachmentSize   string    `json:"max_attachment_size"`
	ApprovedPosts       int       `json:"approved_posts"`
	ModStatus           string    `json:"mod_status"`
	PendingMsgNotify    string    `json:"pending_msg_notify"`
	PendingSubNotify    string    `json:"pending_sub_notify"`
	SubNotify           string    `json:"sub_notify"`
	StorageNotify       string    `json:"storage_notify"`
	SubGroupNotify      string    `json:"sub_group_notify"`
	MessageReportNotify string    `json:"message_report_notify"`
	AccountNotify       string    `json:"account_notify"`
	ModPermissions      string    `json:"mod_permissions"`
	OwnerMsgNotify      string    `json:"owner_msg_notify"`
	ChatNotify          string    `json:"chat_notify"`
	PhotoNotify         string    `json:"photo_notify"`
	FileNotify          string    `json:"file_notify"`
	WikiNotify          string    `json:"wiki_notify"`
	DatabaseNotify      string    `json:"database_notify"`
	Email               string    `json:"email"`
	UserStatus          string    `json:"user_status"`
	UserName            string    `json:"user_name"`
	Timezone            string    `json:"timezone"`
	FullName            string    `json:"full_name"`
	AboutMe             string    `json:"about_me"`
	Location            string    `json:"location"`
	Website             string    `json:"website"`
	ProfilePrivacy      string    `json:"profile_privacy"`
	DontMungeMessageID  bool      `json:"dont_munge_message_id"`
	UseSignature        bool      `json:"use_signature"`
	UseSignatureEmail   bool      `json:"use_signature_email"`
	Signature           string    `json:"signature"`
	Color               string    `json:"color"`
	CoverPhotoURL       string    `json:"cover_photo_url"`
	IconURL             string    `json:"icon_url"`
	NiceGroupName       string    `json:"nice_group_name"`
	SubsCount           int       `json:"subs_count"`
	MostRecentMessage   time.Time `json:"most_recent_message"`
	Perms               struct {
		Object                          string `json:"object"`
		ArchivesVisible                 bool   `json:"archives_visible"`
		PollsVisible                    bool   `json:"polls_visible"`
		MembersVisible                  bool   `json:"members_visible"`
		ChatVisible                     bool   `json:"chat_visible"`
		CalendarVisible                 bool   `json:"calendar_visible"`
		FilesVisible                    bool   `json:"files_visible"`
		DatabaseVisible                 bool   `json:"database_visible"`
		PhotosVisible                   bool   `json:"photos_visible"`
		WikiVisible                     bool   `json:"wiki_visible"`
		MemberDirectoryVisible          bool   `json:"member_directory_visible"`
		HashtagsVisible                 bool   `json:"hashtags_visible"`
		GuidelinesVisible               bool   `json:"guidelines_visible"`
		SubgroupsVisible                bool   `json:"subgroups_visible"`
		OpenDonationsVisible            bool   `json:"open_donations_visible"`
		SponsorVisible                  bool   `json:"sponsor_visible"`
		ManageSubgroups                 bool   `json:"manage_subgroups"`
		DeleteGroup                     bool   `json:"delete_group"`
		DownloadArchives                bool   `json:"download_archives"`
		DownloadEntireGroup             bool   `json:"download_entire_group"`
		DownloadMembers                 bool   `json:"download_members"`
		ViewActivity                    bool   `json:"view_activity"`
		CreateHashtags                  bool   `json:"create_hashtags"`
		ManageHashtags                  bool   `json:"manage_hashtags"`
		ManageIntegrations              bool   `json:"manage_integrations"`
		ManageGroupSettings             bool   `json:"manage_group_settings"`
		MakeModerator                   bool   `json:"make_moderator"`
		ManageMemberSubscriptionOptions bool   `json:"manage_member_subscription_options"`
		ManagePendingMembers            bool   `json:"manage_pending_members"`
		RemoveMembers                   bool   `json:"remove_members"`
		BanMembers                      bool   `json:"ban_members"`
		ManageGroupBilling              bool   `json:"manage_group_billing"`
		ManageGroupPayments             bool   `json:"manage_group_payments"`
		EditArchives                    bool   `json:"edit_archives"`
		ManagePendingMessages           bool   `json:"manage_pending_messages"`
		InviteMembers                   bool   `json:"invite_members"`
		CanPost                         bool   `json:"can_post"`
		CanVote                         bool   `json:"can_vote"`
		ManagePolls                     bool   `json:"manage_polls"`
		ManagePhotos                    bool   `json:"manage_photos"`
		ManageMembers                   bool   `json:"manage_members"`
		ManageCalendar                  bool   `json:"manage_calendar"`
		ManageChats                     bool   `json:"manage_chats"`
		ViewMemberDirectory             bool   `json:"view_member_directory"`
		ManageFiles                     bool   `json:"manage_files"`
		ManageWiki                      bool   `json:"manage_wiki"`
		ManageSubscription              bool   `json:"manage_subscription"`
		PublicPage                      bool   `json:"public_page"`
		SubPage                         bool   `json:"sub_page"`
		ModPage                         bool   `json:"mod_page"`
	} `json:"perms"`
	ExtraMemberData []struct {
		ColID          int       `json:"col_id"`
		ColType        string    `json:"col_type"`
		Text           string    `json:"text,omitempty"`
		Checked        bool      `json:"checked,omitempty"`
		Date           time.Time `json:"date,omitempty"`
		Time           time.Time `json:"time,omitempty"`
		StreetAddress1 string    `json:"street_address1,omitempty"`
		StreetAddress2 string    `json:"street_address2,omitempty"`
		City           string    `json:"city,omitempty"`
		State          string    `json:"state,omitempty"`
		Zip            string    `json:"zip,omitempty"`
		Country        string    `json:"country,omitempty"`
		Title          string    `json:"title,omitempty"`
		URL            string    `json:"url,omitempty"`
		Desc           string    `json:"desc,omitempty"`
		ImageName      string    `json:"image_name,omitempty"`
	} `json:"extra_member_data"`
}
type MemberInfoList struct {
	Object        string `json:"object"`
	TotalCount    int    `json:"total_count"`
	StartItem     int    `json:"start_item"`
	EndItem       int    `json:"end_item"`
	HasMore       bool   `json:"has_more"`
	NextPageToken int    `json:"next_page_token"`
	SortField     string `json:"sort_field"`
	SecondOrder   string `json:"second_order"`
	Query         string `json:"query"`
	SortDir       string `json:"sort_dir"`
	// groupData
	Data []MemberInfo `json:"data"`
}

type PendingMsg struct {
	Object         string `json:"object"`
	ID             int    `json:"id"`
	Created        string `json:"created"`
	Updated        string `json:"updated"`
	GroupID        int    `json:"group_id"`
	UserID         int    `json:"user_id"`
	Subject        string `json:"subject"`
	IsSpecial      bool   `json:"is_special"`
	IsWebPost      bool   `json:"is_web_post"`
	HasAttachments bool   `json:"has_attachments"`
	EditorMemberID int    `json:"editor_member_id"`
	RawMessage     string `json:"raw_message"`
	MessageBody    string `json:"message_body"`
	BodyFormat     string `json:"body_format"`
	Sender         struct {
		ProfilePhotoURL string `json:"profile_photo_url"`
		Name            string `json:"name"`
		CanViewProfile  bool   `json:"can_view_profile"`
	} `json:"sender"`
	SenderEmail  string `json:"sender_email"`
	SenderName   string `json:"sender_name"`
	ClaimingUser struct {
		ProfilePhotoURL string `json:"profile_photo_url"`
		Name            string `json:"name"`
		CanViewProfile  bool   `json:"can_view_profile"`
	} `json:"claiming_user"`
	ClaimedDate string `json:"claimed_date"`
	Editor      struct {
		ProfilePhotoURL string `json:"profile_photo_url"`
		Name            string `json:"name"`
		CanViewProfile  bool   `json:"can_view_profile"`
	} `json:"editor"`
	EditReason string `json:"edit_reason"`
	Type       string `json:"type"`
	VirusName  string `json:"virus_name"`
}

type PendingMsgList struct {
	Object        string       `json:"object"`
	TotalCount    int          `json:"total_count"`
	StartItem     int          `json:"start_item"`
	EndItem       int          `json:"end_item"`
	HasMore       bool         `json:"has_more"`
	NextPageToken int          `json:"next_page_token"`
	SortField     string       `json:"sort_field"`
	SecondOrder   string       `json:"second_order"`
	Query         string       `json:"query"`
	SortDir       string       `json:"sort_dir"`
	Data          []PendingMsg `json:"data"`
}

func (mi MemberInfo) String() string {
	return Sprintf("MemberInfo { %s <%s> - [UserId %d] [GroupId %d]}", mi.FullName, mi.Email, mi.UserID, mi.GroupID)
}
func checkClose(err error, msg string) {
	if err != nil {
		log.Printf("checkClose : %s: %s", msg, err)
	}
}

// NewGroupsClient function to initialize the GroupsClient
func NewGroupsClient(baseURL string) *GroupsClient {
	return &GroupsClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *GroupsClient) Authenticate(email, password string) error {
	formData := url.Values{
		"email":    {email},
		"password": {password},
		"token":    {"true"},
	}

	groupsApiLoginUrl := Sprintf("%s/api/v1/login", c.BaseURL)

	resp, err := c.Client.Post(groupsApiLoginUrl, "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return Errorf("failed to send authentication request: %w", err)
	}
	defer func() {
		_ = resp.Body.Close() // don't ignore close, but don't return close errors
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return Errorf("authentication failed: status %d\nResponse body:\n%s", resp.StatusCode, string(body))
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return Errorf("failed to unmarshal token response: %w\nRaw response:\n%s", err, string(body))
	}

	c.Token = tokenResponse.Token
	return nil
}

// Authenticate method to get and store the API token
func (c *GroupsClient) AuthenticateOld(email, password string) error {
	formData := url.Values{
		"email":    {email},
		"password": {password},
		"token":    {"true"},
	}

	groupsApiLoginUrl := Sprintf("%s/api/v1/login", c.BaseURL)

	resp, err := c.Client.Post(groupsApiLoginUrl, "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return Errorf("received non-200 response code: %d\n Resp Body\n%s\n", resp.StatusCode, resp.Body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer checkClose(resp.Body.Close(), "GroupsClient.Authenticate: Error closing resp.Body")

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return err
	}

	c.Token = tokenResponse.Token
	return nil
}

// doRequest method to make authenticated HTTP requests
func (c *GroupsClient) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	time.Sleep(1 * time.Second)
	req, err := http.NewRequest(method, Sprintf("%s%s", c.BaseURL, endpoint), body)
	if err != nil {
		return nil, err
	}

	// Add the token to the Authorization header using basic auth format
	req.SetBasicAuth(c.Token, "")
	// FIXME gate this setting for POST reqs only??
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.Client.Do(req)
}

// GetOrg gets the org object for domain that the client is authenticated against
// https://groups.io/api#get_org
// https://groups.io/api#the-org-object
func (c *GroupsClient) GetOrg() (*Org, error) {
	resp, err := c.doRequest("GET", "/api/v1/getorg", nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Errorf("GetOrg : received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer checkClose(resp.Body.Close(), "GroupsClient.GetOrg() Error closing resp.Body")

	var orgDetails Org
	if err := json.Unmarshal(body, &orgDetails); err != nil {
		return nil, err
	}
	return &orgDetails, nil
}

// GetMemberInfoList method to get member info list of the authenticated user with pagination
// https://groups.io/api#get-subscriptions
func (c *GroupsClient) GetMemberInfoList() ([]MemberInfo, int, error) {
	// First call should not include the page_token parameter
	objectLimit := 100
	subCount := 0
	resp, err := c.doRequest("GET", Sprintf("/api/v1/getsubs?limit=%d", objectLimit), nil)

	if err != nil {
		return nil, subCount, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, subCount, Errorf("GetMemberInfoList: first call to getsubs, received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, subCount, err
	}

	var memberInfoList MemberInfoList
	if err := json.Unmarshal(body, &memberInfoList); err != nil {
		return nil, 0, err
	}
	checkClose(resp.Body.Close(), "GroupsClient.GetMemberInfoList() Error closing resp.Body")
	subCount = memberInfoList.TotalCount
	allSubscriptions := make([]MemberInfo, 0, subCount)
	allSubscriptions = append(allSubscriptions, memberInfoList.Data[:]...)
	nextPageToken := memberInfoList.NextPageToken
	hasMore := memberInfoList.HasMore

	// ref https://groups.io/api#pagination
	for hasMore != false {
		endpoint := Sprintf("/api/v1/getsubs?limit=100&page_token=%d", nextPageToken)
		forLoopResponse, err := c.doRequest("GET", endpoint, nil)

		if err != nil {
			return nil, subCount, err
		}

		if forLoopResponse.StatusCode != http.StatusOK {
			return nil, subCount, Errorf("received non-200 response code: %d", forLoopResponse.StatusCode)
		}

		body, err := io.ReadAll(forLoopResponse.Body)
		if err != nil {
			return nil, subCount, err
		}

		var subscription MemberInfoList
		if err := json.Unmarshal(body, &subscription); err != nil {
			return nil, subCount, err
		}

		allSubscriptions = append(allSubscriptions, subscription.Data[:]...)
		nextPageToken = subscription.NextPageToken
		hasMore = subscription.HasMore
		checkClose(forLoopResponse.Body.Close(), "GroupsClient.GetOrg() Error closing forLoopResponse.Body")

	}

	return allSubscriptions, subCount, nil
}

// GetMemberId  returns the membership ID of userId if they are a member of groupId
// using https://groups.io/api#getmembers
func (c *GroupsClient) GetMemberId(groupId int, userId int) (int, error) {
	// First call should not include the page_token parameter
	objectLimit := 100
	memberCount := 0
	resp, err := c.doRequest("GET", Sprintf("/api/v1/getmembers?group_id=%d&limit=%d", groupId, objectLimit), nil)

	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GetMemberId,  getmembers, resp: %v, groupId %d, (userId %d)", resp, groupId, userId)
		return 0, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	checkClose(resp.Body.Close(), "GroupsClient.GetMemberId() Error closing resp.Body")
	var memberInfoList MemberInfoList
	if err := json.Unmarshal(body, &memberInfoList); err != nil {
		return 0, err
	}

	memberCount = memberInfoList.TotalCount
	allMembers := make([]MemberInfo, 0, memberCount)
	allMembers = append(allMembers, memberInfoList.Data[:]...)
	nextPageToken := memberInfoList.NextPageToken
	hasMore := memberInfoList.HasMore

	// ref https://groups.io/api#pagination
	for hasMore != false {
		endpoint := Sprintf("/api/v1/getmembers?limit=100&page_token=%d", nextPageToken)
		forLoopResponse, err := c.doRequest("GET", endpoint, nil)

		if err != nil {
			return 0, err
		}

		if forLoopResponse.StatusCode != http.StatusOK {
			return 0, Errorf("received non-200 response code: %d", forLoopResponse.StatusCode)
		}
		body, err := io.ReadAll(forLoopResponse.Body)
		if err != nil {
			return 0, err
		}

		var members MemberInfoList
		if err = json.Unmarshal(body, &members); err != nil {
			return 0, err
		}

		allMembers = append(allMembers, members.Data[:]...)
		nextPageToken = members.NextPageToken
		hasMore = members.HasMore
		checkClose(forLoopResponse.Body.Close(), "GroupsClient.GetOrg() Error closing forLoopResponse.Body")
	}
	for _, member := range allMembers {
		if member.UserID == userId {
			return member.ID, nil
		}
	}
	return 0, Errorf("GetMemberId, UserId : %d not found in groupId %d\n", userId, groupId)
}

// GetMembers returns an array of MemberInfo for the given @groupId using https://groups.io/api#getmembers
func (c *GroupsClient) GetMembers(groupId int) ([]MemberInfo, error) {
	// First call should not include the page_token parameter
	objectLimit := 100
	memberCount := 0
	resp, err := c.doRequest("GET", Sprintf("/api/v1/getmembers?group_id=%d&limit=%d", groupId, objectLimit), nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GetMembers, getmembers, resp: %v, groupId %d", resp, groupId)
		return nil, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	checkClose(resp.Body.Close(), "GroupsClient.GetMembers() Error closing resp.Body")
	var memberInfoList MemberInfoList
	if err := json.Unmarshal(body, &memberInfoList); err != nil {
		return nil, err
	}

	memberCount = memberInfoList.TotalCount
	allMembers := make([]MemberInfo, 0, memberCount)
	allMembers = append(allMembers, memberInfoList.Data[:]...)
	nextPageToken := memberInfoList.NextPageToken
	hasMore := memberInfoList.HasMore

	// ref https://groups.io/api#pagination
	for hasMore != false {
		endpoint := Sprintf("/api/v1/getmembers?limit=100&page_token=%d", nextPageToken)
		forLoopResponse, err := c.doRequest("GET", endpoint, nil)

		if err != nil {
			return nil, err
		}

		if forLoopResponse.StatusCode != http.StatusOK {
			return nil, Errorf("received non-200 response code: %d", forLoopResponse.StatusCode)
		}
		body, err := io.ReadAll(forLoopResponse.Body)
		if err != nil {
			return nil, err
		}

		var members MemberInfoList
		if err = json.Unmarshal(body, &members); err != nil {
			return nil, err
		}

		allMembers = append(allMembers, members.Data[:]...)
		nextPageToken = members.NextPageToken
		hasMore = members.HasMore
		checkClose(forLoopResponse.Body.Close(), "GroupsClient.GetOrg() Error closing forLoopResponse.Body")
	}

	return allMembers, nil
}

// SearchMemberDetails retrieves the User data associated with fullEmail from the Org's main group
// the underlying groups.io end point returns a list but this function will only return data about a user's membership
// of main when the fullEmail is associated with one, and only one record,  in that list.
// https://groups.io/api#search-members
func (c *GroupsClient) SearchMemberDetails(fullEmail string) (*MemberInfo, error) {

	org, err := c.GetOrg()
	if err != nil {
		return nil, err
	}
	// NOTES in endpoint:
	// 1. The Org's parent group is being searched (org.ParentGroupID)
	// 2. The q parameter of searchmembers is a query and allows you to search for members with a partial string
	//
	// If you wanted to get a list of all members from a specific domain you could do that. But that is not what we are
	// doing here. We want to get a memberInfo record for a single user in the main group. This record can then be used
	// for administrative purposes elsewhere.
	searchQuery := Sprintf("/api/v1/searchmembers?group_id=%d&q=%s", org.ParentGroupID, url.QueryEscape(fullEmail))
	resp, err := c.doRequest(
		"GET",
		searchQuery,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Errorf("SearchMemberDetails: received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var memberInfoList MemberInfoList
	if err := json.Unmarshal(body, &memberInfoList); err != nil {
		return nil, err
	}
	if memberInfoList.TotalCount == 1 {
		member := memberInfoList.Data[memberInfoList.StartItem-1]
		return &member, nil
	} else {
		return nil, Errorf("SearchMemberDetails: %s returned %d items in list", fullEmail, memberInfoList.TotalCount)
	}
}

// GetAuthenticatedUser method to get user details from the API
func (c *GroupsClient) GetAuthenticatedUser() (*User, error) {
	resp, err := c.doRequest("GET", "/api/v1/getuser", nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Errorf("GetAuthenticatedUser : received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer checkClose(resp.Body.Close(), "GroupsClient.GetAuthenticatedUser: Error closing resp.Body")

	var loggedInUserDetails User
	if err := json.Unmarshal(body, &loggedInUserDetails); err != nil {
		return nil, err
	}

	return &loggedInUserDetails, nil
}

// GrantOwnerPermsToGroupMember assigns the OwnerRole to newOwner for each group in targetGroups
// returns number of groups that were updated or error.
func (c *GroupsClient) GrantOwnerPermsToGroupMember(newOwner MemberInfo, targetGroups []MemberInfo) (int, error) {
	var groupsUpdated int = 0
	var err error = nil
	for _, group := range targetGroups {
		thisGroupsMemberId, gmiError := c.GetMemberId(group.GroupID, newOwner.UserID)
		if gmiError == nil {
			m, ugmError := c.UpdateGroupMember(group.GroupID, thisGroupsMemberId, "mod_status", "sub_modstatus_owner")
			if ugmError == nil {
				groupsUpdated++
				log.Printf("INFO Member %s should now be an owner on group %s", m.FullName, group.GroupName)
			} else {
				log.Printf("WARN Member %s was not updated to owner of group %s", newOwner.FullName, group.GroupName)
			}
		} else {
			log.Printf("WARN: Member %s is not a member of group %s. GetMemberId error: %v", newOwner.FullName, group.GroupName, gmiError)
		}
	}
	return groupsUpdated, err
}

// RemoveGroupMember removes the member, memberId from the group, groupID.
// Returns
// nil, Error if the parameters are invalid
// MemberInfo, nil if the member was remove from the group
// https://groups.io/api#remove-member
func (c *GroupsClient) RemoveGroupMember(groupId int, memberId int) (MemberInfo, error) {
	mbr := MemberInfo{}
	formData := url.Values{}
	formData.Set("group_id", strconv.Itoa(groupId))
	formData.Set("member_info_id", strconv.Itoa(memberId))
	formData.Set("extra", "true")
	reqBody := strings.NewReader(formData.Encode())
	resp, reqErr := c.doRequest("POST", "/api/v1/removemember", reqBody)
	if reqErr != nil {
		return mbr, Errorf("RemoveGroupMember: request failed endpoint %s, formData %v: %w", "/api/v1/removemember", formData, reqErr)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errRespBody, _ := io.ReadAll(resp.Body)
		return mbr, Errorf("RemoveGroupMember: non-200 status %d, endpoint %s, formData %v, responseBody: %s",
			resp.StatusCode, "/api/v1/removemember", formData, errRespBody)
	}

	response, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return mbr, readErr
	}

	if unmErr := json.Unmarshal(response, &mbr); unmErr != nil {
		return mbr, unmErr
	}

	return mbr, nil
}

// UpdateGroupMember updates field to value for memberId on groupID, returns an err if this fails to happen
func (c *GroupsClient) UpdateGroupMember(groupId int, memberId int, field string, value string) (MemberInfo, error) {
	mbr := MemberInfo{}
	formData := url.Values{}
	formData.Set("group_id", strconv.Itoa(groupId))
	formData.Set("member_info_id", strconv.Itoa(memberId))
	formData.Set("extra", "true")
	formData.Set(field, value)
	reqBody := strings.NewReader(formData.Encode())
	resp, reqErr := c.doRequest("POST", "/api/v1/updatemember", reqBody)
	if reqErr != nil {
		Printf(
			"UpdateGroupMember: ERROR, endpoint: %s, formData: %v, reqErr %v",
			"/api/v1/updatemember",
			formData,
			reqErr,
		)
		Printf("UpdateGroupMember: ERROR, response: %+v, err: %+v", resp, reqErr)
		os.Exit(1)
		return mbr, reqErr
	}

	if resp.StatusCode != http.StatusOK {
		errRespBody, _ := io.ReadAll(resp.Body)
		return mbr,
			Errorf("UpdateGroupMember: ERROR, called : %s, formData : %+v response: %+v, err: %+v responseBody: %s",
				"/api/v1/updatemember ", formData, resp, reqErr, errRespBody)
	}

	response, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return mbr, readErr
	}

	if unmErr := json.Unmarshal(response, &mbr); unmErr != nil {
		return mbr, unmErr
	}

	return mbr, nil
}

// GetPendingMsgList method to get pending msg info list accessible to the authenticated user with pagination
// FIRST PASS, see if we can get all the pending messages by passing in the parent group ID from the Org
// https://groups.io/api#get-
func (c *GroupsClient) GetPendingMsgList() ([]PendingMsg, int, error) {

	org, err := c.GetOrg()
	if err != nil {
		return nil, 0, err
	}
	log.Printf("GetPendingMsgList: org %+v\n", org)
	v1Endpoint := "getpendingmessages"
	objectLimit := 100
	count := 0
	resource := Sprintf("/api/v1/%s?limit=%d&group_id=%d", v1Endpoint, objectLimit, org.ParentGroupID)
	resp, err := c.doRequest("GET", resource, nil)

	if err != nil {
		return nil, count, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, count, Errorf("GetPendingMsgList: first call to %s, received non-200 response code: %d", v1Endpoint, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, count, err
	}

	var pendingMsgList PendingMsgList
	if err := json.Unmarshal(body, &pendingMsgList); err != nil {
		return nil, 0, err
	}
	checkClose(resp.Body.Close(), "GroupsClient.GetPendingMsgList() Error closing resp.Body")
	count = pendingMsgList.TotalCount
	allPendingMsgs := make([]PendingMsg, 0, count)
	allPendingMsgs = append(allPendingMsgs, pendingMsgList.Data[:]...)
	nextPageToken := pendingMsgList.NextPageToken
	hasMore := pendingMsgList.HasMore

	// ref https://groups.io/api#pagination
	for hasMore != false {
		endpoint := Sprintf("/api/v1/%s?limit=100&page_token=%d", v1Endpoint, nextPageToken)
		forLoopResponse, err := c.doRequest("GET", endpoint, nil)

		if err != nil {
			return nil, count, err
		}

		if forLoopResponse.StatusCode != http.StatusOK {
			return nil, count, Errorf("GroupsClient.GetPendingMsgList() for loop received non-200 response code: %d", forLoopResponse.StatusCode)
		}

		body, err := io.ReadAll(forLoopResponse.Body)
		if err != nil {
			return nil, count, err
		}

		var thisPagesPendingMsgList PendingMsgList
		if err := json.Unmarshal(body, &thisPagesPendingMsgList); err != nil {
			return nil, count, err
		}

		allPendingMsgs = append(allPendingMsgs, thisPagesPendingMsgList.Data[:]...)
		nextPageToken = thisPagesPendingMsgList.NextPageToken
		hasMore = thisPagesPendingMsgList.HasMore
		checkClose(forLoopResponse.Body.Close(), "GroupsClient.GetPendingMsgList() Error closing forLoopResponse.Body")

	}

	return allPendingMsgs, count, nil
}

// RemoveFromGroups removes member from the array of targetGroups
// Returns the number of groups successfully updated, the group names removed from, and an aggregated error (if any)
func (c *GroupsClient) RemoveFromGroups(member MemberInfo, targetGroups []MemberInfo) (int, []string, error) {
	var groupsUpdated int
	var removedGroups []string
	var errs []string

	for _, group := range targetGroups {
		thisGroupsMemberId, gmiError := c.GetMemberId(group.GroupID, member.UserID)
		if gmiError != nil {
			continue
		}

		m, rgmError := c.RemoveGroupMember(group.GroupID, thisGroupsMemberId)
		if rgmError != nil {
			msg := Sprintf("group %s (%d): remove failed: %v", group.GroupName, group.GroupID, rgmError)
			log.Printf("WARN %s", msg)
			errs = append(errs, msg)
			continue
		}

		groupsUpdated++
		removedGroups = append(removedGroups, group.GroupName)
		log.Printf("INFO Member %s removed from group %s", m.FullName, group.GroupName)
	}

	if len(errs) > 0 {
		return groupsUpdated, removedGroups, Errorf("RemoveFromGroups encountered issues: %s", strings.Join(errs, "; "))
	}
	return groupsUpdated, removedGroups, nil
}
