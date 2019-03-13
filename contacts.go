package gohubspot

import (
	"fmt"
	"strings"
)

type AllContactsResp struct {
	Contacts []struct {
		AddedAt          int64              `json:"addedAt"`
		Vid              int                `json:"vid"`
		CanonicalVid     int                `json:"canonical-vid"`
		MergedVids       []interface{}      `json:"merged-vids"`
		PortalID         int                `json:"portal-id"`
		IsContact        bool               `json:"is-contact"`
		ProfileToken     string             `json:"profile-token"`
		ProfileURL       string             `json:"profile-url"`
		Properties       *ContactProperties `json:"properties"`
		FormSubmissions  []interface{}      `json:"form-submissions"`
		ListMemberships  []interface{}      `json:"list-memberships"`
		IdentityProfiles []struct {
			Vid                     int   `json:"vid"`
			SavedAtTimestamp        int64 `json:"saved-at-timestamp"`
			DeletedChangedTimestamp int   `json:"deleted-changed-timestamp"`
			Identities              []struct {
				Type      string `json:"type"`
				Value     string `json:"value"`
				Timestamp int64  `json:"timestamp"`
				IsPrimary bool   `json:"is-primary,omitempty"`
			} `json:"identities"`
		} `json:"identity-profiles"`
		MergeAudits []interface{} `json:"merge-audits"`
	} `json:"contacts"`
	HasMore   bool `json:"has-more"`
	VidOffset int  `json:"vid-offset"`
}

type ContactProperties map[string]struct {
	Value interface{} `json:"value"`
}

func (cp *ContactProperties) GetString(key string) string {
	v, ex := (*cp)[key]
	if !ex {
		return ""
	}
	return fmt.Sprintf("%s", v.Value)
}
func (cp *ContactProperties) GetStrings(key string) []string {
	v, ex := (*cp)[key]
	if !ex {
		return []string{}
	}
	return strings.Split(fmt.Sprintf("%s", v.Value), ";")
}
func (cp *ContactProperties) MultiGetString(key ...string) string {
	vs := make([]string, 0)
	for _, k := range key {
		v := cp.GetString(k)
		if v != "" {
			vs = append(vs, v)
		}
	}
	return strings.Join(vs, " ")
}

func (cp *ContactProperties) PriorityGetString(key ...string) string {
	for _, k := range key {
		v := cp.GetString(k)
		if v != "" {
			return v
		}
	}
	return ""
}
