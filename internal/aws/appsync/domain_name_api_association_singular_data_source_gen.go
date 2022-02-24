// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_appsync_domain_name_api_association", domainNameApiAssociationDataSourceType)
}

// domainNameApiAssociationDataSourceType returns the Terraform awscc_appsync_domain_name_api_association data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::AppSync::DomainNameApiAssociation resource type.
func domainNameApiAssociationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_association_identifier": {
			// Property: ApiAssociationIdentifier
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"api_id": {
			// Property: ApiId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 253,
			//   "minLength": 1,
			//   "pattern": "^(\\*[a-z\\d-]*\\.)?([a-z\\d-]+\\.)+[a-z\\d-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::AppSync::DomainNameApiAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppSync::DomainNameApiAssociation").WithTerraformTypeName("awscc_appsync_domain_name_api_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_association_identifier": "ApiAssociationIdentifier",
		"api_id":                     "ApiId",
		"domain_name":                "DomainName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
