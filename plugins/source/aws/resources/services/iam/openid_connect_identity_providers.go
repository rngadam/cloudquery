// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func OpenidConnectIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:      "aws_iam_openid_connect_identity_providers",
		Resolver:  fetchIamOpenidConnectIdentityProviders,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "client_id_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ClientIDList"),
			},
			{
				Name:     "create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateDate"),
			},
			{
				Name:     "thumbprint_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ThumbprintList"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
