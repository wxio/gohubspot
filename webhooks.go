package gohubspot

import (
	"fmt"
)

type WebhooksService service

type WebhookSettings struct {
	WebhookURL            string `json:"WebhookUrl"`
	MaxConcurrentRequests int    `json:"MaxConcurrentRequests"`
}

type SubscriptionDetails struct {
	SubscriptionType string `json:"subscriptionType"`
	PropertyName     string `json:"propertyName"`
}

type Webhook struct {
	SubscriptionDetails SubscriptionDetails `json:"subscriptionDetails"`
	Enabled             bool                `json:"enabled"`
}

type WebhookPayload struct {
	ObjectID         int    `json:"objectId"`
	PropertyName     string `json:"propertyName,omitempty"`
	PropertyValue    string `json:"propertyValue,omitempty"`
	ChangeSource     string `json:"changeSource"`
	EventID          int64  `json:"eventId"`
	SubscriptionID   int    `json:"subscriptionId"`
	PortalID         int    `json:"portalId"`
	AppID            int    `json:"appId"`
	OccurredAt       int64  `json:"occurredAt"`
	SubscriptionType string `json:"subscriptionType"`
	AttemptNumber    int    `json:"attemptNumber"`
}

// GET https://api.hubapi.com/webhooks/v1/{appId}/settings
// { "webhookUrl": "https://testing.com/webhook", "maxConcurrentRequests": 20}
// PUT https://api.hubapi.com/webhooks/v1/{appId}/settings
// { "webhookUrl": "https://testing.com/webhook-modified", "maxConcurrentRequests": 25}
// GET https://api.hubapi.com/webhooks/v1/{appId}/subscriptions

// PUT https://api.hubapi.com/webhooks/v1/{appId}/subscriptions/{subscriptionId}
// { "enabled" : false }

// DELETE https://api.hubapi.com/webhooks/v1/{appId}/subscriptions/{subscriptionId}

func (s *WebhooksService) GetSettings(appId string) (res *WebhookSettings, err error) {
	res = &WebhookSettings{}
	url := "/webhooks/v1/" + appId + "/settings"
	err = s.client.RunGet(url, res)
	return res, err
}
func (s *WebhooksService) PostSettings(appId string, set *WebhookSettings) (err error) {
	url := "/webhooks/v1/" + appId + "/settings"
	var res interface{}
	err = s.client.RunPost(url, set, &res)
	fmt.Printf("%v\n", res)
	return err
}

func (s *WebhooksService) Get(appId string) ([]Webhook, error) {
	res := make([]Webhook, 0)
	url := "/webhooks/v1/" + appId + "/subscriptions"
	err := s.client.RunGet(url, &res)
	return res, err
}

// POST https://api.hubapi.com/webhooks/v1/{appId}/subscriptions
// { "subscriptionDetails" : { "subscriptionType" : "company.propertyChange", "propertyName" : "companyname" }, "enabled" : false}
// contact.creation - To get notified if any contact is created in a customer's portal.
// contact.deletion - To get notified if any contact is deleted in a customer's portal.
// contact.propertyChange - To get notified if a specified property is changed for any contact in a customer's portal.
// company.creation - To get notified if any company is created in a customer's portal.
// company.deletion - To get notified if any company is deleted in a customer's portal.
// company.propertyChange - To get notified if a specified property is changed for any company in a customer's portal.
// deal.creation - To get notified if any deal is created in a customer's portal.
// deal.deletion - To get notified if any deal is deleted in a customer's portal.
// deal.propertyChange - To get notified if a specified property is changed for any deal in a customer's portal.
func (s *WebhooksService) PostSubscription(appId string, wh *Webhook) error {
	url := "/webhooks/v1/" + appId + "/subscriptions"
	var res interface{}
	err := s.client.RunPost(url, wh, &res)
	fmt.Printf("'%v'\n", res)
	return err
}

func (s *WebhooksService) GetByToken(token string) (*Contact, error) {
	url := fmt.Sprintf("/Webhooks/v1/contact/utk/%s/profile", token)
	res := new(Contact)
	err := s.client.RunGet(url, res)
	return res, err
}
