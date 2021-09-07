package zendesk

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableZendeskJiraLink() *plugin.Table {
	return &plugin.Table{
		Name:        "zendesk_jira_link",
		Description: "Links between Zendesk tickets and Jira issues.",
		List: &plugin.ListConfig{
			Hydrate: listJiraLink,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getJiraLink,
		},
		Columns: []*plugin.Column{
			{
                                Name:        "id",
                                Description: "Automatically assigned when the link is created.",
                                Type:        proto.ColumnType_INT,
                        },
                        {
                                Name:        "created_at",
                                Description: "The time at which the link was created.",
                                Type:        proto.ColumnType_STRING,
                        },
                        {
                                Name:        "issue_id",
                                Description: "The id of the Jira issue.",
                                Type:        proto.ColumnType_INT,
                        },
                        {
                                Name:        "issue_key",
                                Description: "The key for the Jira issue.",
                                Type:        proto.ColumnType_STRING,
                        },
                        {
                                Name:        "ticket_id",
                                Description: "The id of the Zendesk ticket.",
                                Type:        proto.ColumnType_STRING,
                        },
                        {
                                Name:        "updated_at",
                                Description: "The time at which the link was updated.",
                                Type:        proto.ColumnType_STRING,
                        },
		},
	}
}

func listJiraLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := &zendesk.OrganizationListOptions{
		PageOptions: zendesk.PageOptions{
			Page:    1,
			PerPage: 100,
		},
	}	
	jira_links, page, err := conn.GetJiraLinks(ctx)
	for true {	
		if err != nil {
			return nil, err
		}
		for _, t := range jira_links {
			d.StreamListItem(ctx, t)
		}
		if !page.HasNext() {
			break
		}
		opts.Page++
	}
	return nil, nil
}

func getJiraLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	quals := d.KeyColumnQuals

	i := quals["id"].GetInt64Value()
	if err != nil {
		return nil, err
	}
	result, err := conn.GetJiraLink(ctx, i)
	if err != nil {
		return nil, err
	}
	return result, nil	
}
