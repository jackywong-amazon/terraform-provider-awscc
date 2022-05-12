// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_networkmanager_connect_attachments", connectAttachmentsDataSourceType)
}

// connectAttachmentsDataSourceType returns the Terraform awscc_networkmanager_connect_attachments data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::NetworkManager::ConnectAttachment resource type.
func connectAttachmentsDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::NetworkManager::ConnectAttachment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::ConnectAttachment").WithTerraformTypeName("awscc_networkmanager_connect_attachments")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
