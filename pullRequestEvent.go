package main

// PullRequestEvent is a struct into which we will be decoding the JSON data
// sent to us by GitHub.
type PullRequestEvent struct {
	Action      string `json:"action"`
	Number      int64  `json:"number"`
	PullRequest struct {
		Links struct {
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		Additions         int64         `json:"additions"`
		Assignee          interface{}   `json:"assignee"`
		Assignees         []interface{} `json:"assignees"`
		AuthorAssociation string        `json:"author_association"`
		Base              struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Repo  struct {
				ArchiveURL       string      `json:"archive_url"`
				Archived         bool        `json:"archived"`
				AssigneesURL     string      `json:"assignees_url"`
				BlobsURL         string      `json:"blobs_url"`
				BranchesURL      string      `json:"branches_url"`
				CloneURL         string      `json:"clone_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				CommentsURL      string      `json:"comments_url"`
				CommitsURL       string      `json:"commits_url"`
				CompareURL       string      `json:"compare_url"`
				ContentsURL      string      `json:"contents_url"`
				ContributorsURL  string      `json:"contributors_url"`
				CreatedAt        string      `json:"created_at"`
				DefaultBranch    string      `json:"default_branch"`
				DeploymentsURL   string      `json:"deployments_url"`
				Description      interface{} `json:"description"`
				Disabled         bool        `json:"disabled"`
				DownloadsURL     string      `json:"downloads_url"`
				EventsURL        string      `json:"events_url"`
				Fork             bool        `json:"fork"`
				Forks            int64       `json:"forks"`
				ForksCount       int64       `json:"forks_count"`
				ForksURL         string      `json:"forks_url"`
				FullName         string      `json:"full_name"`
				GitCommitsURL    string      `json:"git_commits_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitURL           string      `json:"git_url"`
				HasDownloads     bool        `json:"has_downloads"`
				HasIssues        bool        `json:"has_issues"`
				HasPages         bool        `json:"has_pages"`
				HasProjects      bool        `json:"has_projects"`
				HasWiki          bool        `json:"has_wiki"`
				Homepage         interface{} `json:"homepage"`
				HooksURL         string      `json:"hooks_url"`
				HTMLURL          string      `json:"html_url"`
				ID               int64       `json:"id"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				IssuesURL        string      `json:"issues_url"`
				KeysURL          string      `json:"keys_url"`
				LabelsURL        string      `json:"labels_url"`
				Language         interface{} `json:"language"`
				LanguagesURL     string      `json:"languages_url"`
				License          interface{} `json:"license"`
				MergesURL        string      `json:"merges_url"`
				MilestonesURL    string      `json:"milestones_url"`
				MirrorURL        interface{} `json:"mirror_url"`
				Name             string      `json:"name"`
				NodeID           string      `json:"node_id"`
				NotificationsURL string      `json:"notifications_url"`
				OpenIssues       int64       `json:"open_issues"`
				OpenIssuesCount  int64       `json:"open_issues_count"`
				Owner            struct {
					AvatarURL         string `json:"avatar_url"`
					EventsURL         string `json:"events_url"`
					FollowersURL      string `json:"followers_url"`
					FollowingURL      string `json:"following_url"`
					GistsURL          string `json:"gists_url"`
					GravatarID        string `json:"gravatar_id"`
					HTMLURL           string `json:"html_url"`
					ID                int64  `json:"id"`
					Login             string `json:"login"`
					NodeID            string `json:"node_id"`
					OrganizationsURL  string `json:"organizations_url"`
					ReceivedEventsURL string `json:"received_events_url"`
					ReposURL          string `json:"repos_url"`
					SiteAdmin         bool   `json:"site_admin"`
					StarredURL        string `json:"starred_url"`
					SubscriptionsURL  string `json:"subscriptions_url"`
					Type              string `json:"type"`
					URL               string `json:"url"`
				} `json:"owner"`
				Private         bool   `json:"private"`
				PullsURL        string `json:"pulls_url"`
				PushedAt        string `json:"pushed_at"`
				ReleasesURL     string `json:"releases_url"`
				Size            int64  `json:"size"`
				SSHURL          string `json:"ssh_url"`
				StargazersCount int64  `json:"stargazers_count"`
				StargazersURL   string `json:"stargazers_url"`
				StatusesURL     string `json:"statuses_url"`
				SubscribersURL  string `json:"subscribers_url"`
				SubscriptionURL string `json:"subscription_url"`
				SvnURL          string `json:"svn_url"`
				TagsURL         string `json:"tags_url"`
				TeamsURL        string `json:"teams_url"`
				TreesURL        string `json:"trees_url"`
				UpdatedAt       string `json:"updated_at"`
				URL             string `json:"url"`
				Watchers        int64  `json:"watchers"`
				WatchersCount   int64  `json:"watchers_count"`
			} `json:"repo"`
			Sha  string `json:"sha"`
			User struct {
				AvatarURL         string `json:"avatar_url"`
				EventsURL         string `json:"events_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				GravatarID        string `json:"gravatar_id"`
				HTMLURL           string `json:"html_url"`
				ID                int64  `json:"id"`
				Login             string `json:"login"`
				NodeID            string `json:"node_id"`
				OrganizationsURL  string `json:"organizations_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				ReposURL          string `json:"repos_url"`
				SiteAdmin         bool   `json:"site_admin"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				Type              string `json:"type"`
				URL               string `json:"url"`
			} `json:"user"`
		} `json:"base"`
		Body         string      `json:"body"`
		ChangedFiles int64       `json:"changed_files"`
		ClosedAt     interface{} `json:"closed_at"`
		Comments     int64       `json:"comments"`
		CommentsURL  string      `json:"comments_url"`
		Commits      int64       `json:"commits"`
		CommitsURL   string      `json:"commits_url"`
		CreatedAt    string      `json:"created_at"`
		Deletions    int64       `json:"deletions"`
		DiffURL      string      `json:"diff_url"`
		Draft        bool        `json:"draft"`
		Head         struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Repo  struct {
				ArchiveURL       string      `json:"archive_url"`
				Archived         bool        `json:"archived"`
				AssigneesURL     string      `json:"assignees_url"`
				BlobsURL         string      `json:"blobs_url"`
				BranchesURL      string      `json:"branches_url"`
				CloneURL         string      `json:"clone_url"`
				CollaboratorsURL string      `json:"collaborators_url"`
				CommentsURL      string      `json:"comments_url"`
				CommitsURL       string      `json:"commits_url"`
				CompareURL       string      `json:"compare_url"`
				ContentsURL      string      `json:"contents_url"`
				ContributorsURL  string      `json:"contributors_url"`
				CreatedAt        string      `json:"created_at"`
				DefaultBranch    string      `json:"default_branch"`
				DeploymentsURL   string      `json:"deployments_url"`
				Description      interface{} `json:"description"`
				Disabled         bool        `json:"disabled"`
				DownloadsURL     string      `json:"downloads_url"`
				EventsURL        string      `json:"events_url"`
				Fork             bool        `json:"fork"`
				Forks            int64       `json:"forks"`
				ForksCount       int64       `json:"forks_count"`
				ForksURL         string      `json:"forks_url"`
				FullName         string      `json:"full_name"`
				GitCommitsURL    string      `json:"git_commits_url"`
				GitRefsURL       string      `json:"git_refs_url"`
				GitTagsURL       string      `json:"git_tags_url"`
				GitURL           string      `json:"git_url"`
				HasDownloads     bool        `json:"has_downloads"`
				HasIssues        bool        `json:"has_issues"`
				HasPages         bool        `json:"has_pages"`
				HasProjects      bool        `json:"has_projects"`
				HasWiki          bool        `json:"has_wiki"`
				Homepage         interface{} `json:"homepage"`
				HooksURL         string      `json:"hooks_url"`
				HTMLURL          string      `json:"html_url"`
				ID               int64       `json:"id"`
				IssueCommentURL  string      `json:"issue_comment_url"`
				IssueEventsURL   string      `json:"issue_events_url"`
				IssuesURL        string      `json:"issues_url"`
				KeysURL          string      `json:"keys_url"`
				LabelsURL        string      `json:"labels_url"`
				Language         interface{} `json:"language"`
				LanguagesURL     string      `json:"languages_url"`
				License          interface{} `json:"license"`
				MergesURL        string      `json:"merges_url"`
				MilestonesURL    string      `json:"milestones_url"`
				MirrorURL        interface{} `json:"mirror_url"`
				Name             string      `json:"name"`
				NodeID           string      `json:"node_id"`
				NotificationsURL string      `json:"notifications_url"`
				OpenIssues       int64       `json:"open_issues"`
				OpenIssuesCount  int64       `json:"open_issues_count"`
				Owner            struct {
					AvatarURL         string `json:"avatar_url"`
					EventsURL         string `json:"events_url"`
					FollowersURL      string `json:"followers_url"`
					FollowingURL      string `json:"following_url"`
					GistsURL          string `json:"gists_url"`
					GravatarID        string `json:"gravatar_id"`
					HTMLURL           string `json:"html_url"`
					ID                int64  `json:"id"`
					Login             string `json:"login"`
					NodeID            string `json:"node_id"`
					OrganizationsURL  string `json:"organizations_url"`
					ReceivedEventsURL string `json:"received_events_url"`
					ReposURL          string `json:"repos_url"`
					SiteAdmin         bool   `json:"site_admin"`
					StarredURL        string `json:"starred_url"`
					SubscriptionsURL  string `json:"subscriptions_url"`
					Type              string `json:"type"`
					URL               string `json:"url"`
				} `json:"owner"`
				Private         bool   `json:"private"`
				PullsURL        string `json:"pulls_url"`
				PushedAt        string `json:"pushed_at"`
				ReleasesURL     string `json:"releases_url"`
				Size            int64  `json:"size"`
				SSHURL          string `json:"ssh_url"`
				StargazersCount int64  `json:"stargazers_count"`
				StargazersURL   string `json:"stargazers_url"`
				StatusesURL     string `json:"statuses_url"`
				SubscribersURL  string `json:"subscribers_url"`
				SubscriptionURL string `json:"subscription_url"`
				SvnURL          string `json:"svn_url"`
				TagsURL         string `json:"tags_url"`
				TeamsURL        string `json:"teams_url"`
				TreesURL        string `json:"trees_url"`
				UpdatedAt       string `json:"updated_at"`
				URL             string `json:"url"`
				Watchers        int64  `json:"watchers"`
				WatchersCount   int64  `json:"watchers_count"`
			} `json:"repo"`
			Sha  string `json:"sha"`
			User struct {
				AvatarURL         string `json:"avatar_url"`
				EventsURL         string `json:"events_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				GravatarID        string `json:"gravatar_id"`
				HTMLURL           string `json:"html_url"`
				ID                int64  `json:"id"`
				Login             string `json:"login"`
				NodeID            string `json:"node_id"`
				OrganizationsURL  string `json:"organizations_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				ReposURL          string `json:"repos_url"`
				SiteAdmin         bool   `json:"site_admin"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				Type              string `json:"type"`
				URL               string `json:"url"`
			} `json:"user"`
		} `json:"head"`
		HTMLURL             string        `json:"html_url"`
		ID                  int64         `json:"id"`
		IssueURL            string        `json:"issue_url"`
		Labels              []interface{} `json:"labels"`
		Locked              bool          `json:"locked"`
		MaintainerCanModify bool          `json:"maintainer_can_modify"`
		MergeCommitSha      interface{}   `json:"merge_commit_sha"`
		Mergeable           interface{}   `json:"mergeable"`
		MergeableState      string        `json:"mergeable_state"`
		Merged              bool          `json:"merged"`
		MergedAt            interface{}   `json:"merged_at"`
		MergedBy            interface{}   `json:"merged_by"`
		Milestone           interface{}   `json:"milestone"`
		NodeID              string        `json:"node_id"`
		Number              int64         `json:"number"`
		PatchURL            string        `json:"patch_url"`
		Rebaseable          interface{}   `json:"rebaseable"`
		RequestedReviewers  []interface{} `json:"requested_reviewers"`
		RequestedTeams      []interface{} `json:"requested_teams"`
		ReviewCommentURL    string        `json:"review_comment_url"`
		ReviewComments      int64         `json:"review_comments"`
		ReviewCommentsURL   string        `json:"review_comments_url"`
		State               string        `json:"state"`
		StatusesURL         string        `json:"statuses_url"`
		Title               string        `json:"title"`
		UpdatedAt           string        `json:"updated_at"`
		URL                 string        `json:"url"`
		User                struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"user"`
	} `json:"pull_request"`
	Repository struct {
		ArchiveURL       string      `json:"archive_url"`
		Archived         bool        `json:"archived"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        string      `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		DeploymentsURL   string      `json:"deployments_url"`
		Description      interface{} `json:"description"`
		Disabled         bool        `json:"disabled"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            int64       `json:"forks"`
		ForksCount       int64       `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasProjects      bool        `json:"has_projects"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         interface{} `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HTMLURL          string      `json:"html_url"`
		ID               int64       `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         interface{} `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		License          interface{} `json:"license"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NodeID           string      `json:"node_id"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       int64       `json:"open_issues"`
		OpenIssuesCount  int64       `json:"open_issues_count"`
		Owner            struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"owner"`
		Private         bool   `json:"private"`
		PullsURL        string `json:"pulls_url"`
		PushedAt        string `json:"pushed_at"`
		ReleasesURL     string `json:"releases_url"`
		Size            int64  `json:"size"`
		SSHURL          string `json:"ssh_url"`
		StargazersCount int64  `json:"stargazers_count"`
		StargazersURL   string `json:"stargazers_url"`
		StatusesURL     string `json:"statuses_url"`
		SubscribersURL  string `json:"subscribers_url"`
		SubscriptionURL string `json:"subscription_url"`
		SvnURL          string `json:"svn_url"`
		TagsURL         string `json:"tags_url"`
		TeamsURL        string `json:"teams_url"`
		TreesURL        string `json:"trees_url"`
		UpdatedAt       string `json:"updated_at"`
		URL             string `json:"url"`
		Watchers        int64  `json:"watchers"`
		WatchersCount   int64  `json:"watchers_count"`
	} `json:"repository"`
	Sender struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		NodeID            string `json:"node_id"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"sender"`
}
