// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kendra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_kendra_faq", faqDataSourceType)
}

// faqDataSourceType returns the Terraform awscc_kendra_faq data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Kendra::Faq resource type.
func faqDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "FAQ description",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "FAQ description",
			Type:        types.StringType,
			Computed:    true,
		},
		"file_format": {
			// Property: FileFormat
			// CloudFormation resource type schema:
			// {
			//   "description": "FAQ file format",
			//   "enum": [
			//     "CSV",
			//     "CSV_WITH_HEADER",
			//     "JSON"
			//   ],
			//   "type": "string"
			// }
			Description: "FAQ file format",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique ID of the FAQ",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Unique ID of the FAQ",
			Type:        types.StringType,
			Computed:    true,
		},
		"index_id": {
			// Property: IndexId
			// CloudFormation resource type schema:
			// {
			//   "description": "Index ID",
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "type": "string"
			// }
			Description: "Index ID",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "FAQ name",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "FAQ name",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "FAQ role ARN",
			//   "maxLength": 1284,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "FAQ role ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"s3_path": {
			// Property: S3Path
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "FAQ S3 path",
			//   "properties": {
			//     "Bucket": {
			//       "maxLength": 63,
			//       "minLength": 3,
			//       "pattern": "[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]",
			//       "type": "string"
			//     },
			//     "Key": {
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Bucket",
			//     "Key"
			//   ],
			//   "type": "object"
			// }
			Description: "FAQ S3 path",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket": {
						// Property: Bucket
						Type:     types.StringType,
						Computed: true,
					},
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags for labeling the FAQ",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Kendra resources",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array"
			// }
			Description: "Tags for labeling the FAQ",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Kendra::Faq",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Kendra::Faq").WithTerraformTypeName("awscc_kendra_faq")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":         "Arn",
		"bucket":      "Bucket",
		"description": "Description",
		"file_format": "FileFormat",
		"id":          "Id",
		"index_id":    "IndexId",
		"key":         "Key",
		"name":        "Name",
		"role_arn":    "RoleArn",
		"s3_path":     "S3Path",
		"tags":        "Tags",
		"value":       "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
