// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ses_configuration_set", configurationSetDataSourceType)
}

// configurationSetDataSourceType returns the Terraform awscc_ses_configuration_set data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SES::ConfigurationSet resource type.
func configurationSetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the configuration set.",
			//   "pattern": "^[a-zA-Z0-9_-]{1,64}$",
			//   "type": "string"
			// }
			Description: "The name of the configuration set.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SES::ConfigurationSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::ConfigurationSet").WithTerraformTypeName("awscc_ses_configuration_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"name": "Name",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
