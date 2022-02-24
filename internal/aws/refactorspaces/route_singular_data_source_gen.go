// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package refactorspaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_refactorspaces_route", routeDataSourceType)
}

// routeDataSourceType returns the Terraform awscc_refactorspaces_route data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::RefactorSpaces::Route resource type.
func routeDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_identifier": {
			// Property: ApplicationIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^app-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:(aws[a-zA-Z-]*)?:refactor-spaces:[a-zA-Z0-9\\-]+:\\w{12}:[a-zA-Z_0-9+=,.@\\-_/]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"environment_identifier": {
			// Property: EnvironmentIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^env-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"path_resource_to_id": {
			// Property: PathResourceToId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"route_identifier": {
			// Property: RouteIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^rte-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"route_type": {
			// Property: RouteType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEFAULT",
			//     "URI_PATH"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"service_identifier": {
			// Property: ServiceIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "^svc-([0-9A-Za-z]{10}$)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Environment resource",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
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
			//   "type": "array"
			// }
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
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
		"uri_path_route": {
			// Property: UriPathRoute
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ActivationState": {
			//       "enum": [
			//         "ACTIVE"
			//       ],
			//       "type": "string"
			//     },
			//     "IncludeChildPaths": {
			//       "type": "boolean"
			//     },
			//     "Methods": {
			//       "insertionOrder": false,
			//       "items": {
			//         "enum": [
			//           "DELETE",
			//           "GET",
			//           "HEAD",
			//           "OPTIONS",
			//           "PATCH",
			//           "POST",
			//           "PUT"
			//         ],
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "SourcePath": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "pattern": "^(/[a-zA-Z0-9._-]+)+$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ActivationState"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"activation_state": {
						// Property: ActivationState
						Type:     types.StringType,
						Computed: true,
					},
					"include_child_paths": {
						// Property: IncludeChildPaths
						Type:     types.BoolType,
						Computed: true,
					},
					"methods": {
						// Property: Methods
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"source_path": {
						// Property: SourcePath
						Type:     types.StringType,
						Computed: true,
					},
				},
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
		Description: "Data Source schema for AWS::RefactorSpaces::Route",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RefactorSpaces::Route").WithTerraformTypeName("awscc_refactorspaces_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activation_state":       "ActivationState",
		"application_identifier": "ApplicationIdentifier",
		"arn":                    "Arn",
		"environment_identifier": "EnvironmentIdentifier",
		"include_child_paths":    "IncludeChildPaths",
		"key":                    "Key",
		"methods":                "Methods",
		"path_resource_to_id":    "PathResourceToId",
		"route_identifier":       "RouteIdentifier",
		"route_type":             "RouteType",
		"service_identifier":     "ServiceIdentifier",
		"source_path":            "SourcePath",
		"tags":                   "Tags",
		"uri_path_route":         "UriPathRoute",
		"value":                  "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
