module github.com/turbot/steampipe-plugin-zendesk

go 1.15

require (
	github.com/nukosuke/go-zendesk v0.9.4
	github.com/turbot/steampipe-plugin-sdk v0.2.4
)

replace github.com/nukosuke/go-zendesk => ../go-zendesk
