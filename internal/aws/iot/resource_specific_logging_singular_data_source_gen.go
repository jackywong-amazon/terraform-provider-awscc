// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iot_resource_specific_logging", resourceSpecificLoggingDataSourceType)
}

// resourceSpecificLoggingDataSourceType returns the Terraform awscc_iot_resource_specific_logging data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoT::ResourceSpecificLogging resource type.
func resourceSpecificLoggingDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"log_level": {
			// Property: LogLevel
			// CloudFormation resource type schema:
			// {
			//   "description": "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			//   "enum": [
			//     "ERROR",
			//     "WARN",
			//     "INFO",
			//     "DEBUG",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_id": {
			// Property: TargetId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
			//   "maxLength": 140,
			//   "minLength": 13,
			//   "pattern": "[a-zA-Z0-9:_-]+",
			//   "type": "string"
			// }
			Description: "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_name": {
			// Property: TargetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The target name.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9:_-]+",
			//   "type": "string"
			// }
			Description: "The target name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			// {
			//   "description": "The target type. Value must be THING_GROUP.",
			//   "enum": [
			//     "THING_GROUP"
			//   ],
			//   "type": "string"
			// }
			Description: "The target type. Value must be THING_GROUP.",
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
		Description: "Data Source schema for AWS::IoT::ResourceSpecificLogging",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ResourceSpecificLogging").WithTerraformTypeName("awscc_iot_resource_specific_logging")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"log_level":   "LogLevel",
		"target_id":   "TargetId",
		"target_name": "TargetName",
		"target_type": "TargetType",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
