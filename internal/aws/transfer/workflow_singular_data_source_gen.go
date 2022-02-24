// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_transfer_workflow", workflowDataSourceType)
}

// workflowDataSourceType returns the Terraform awscc_transfer_workflow data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Transfer::Workflow resource type.
func workflowDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			//   "maxLength": 1600,
			//   "minLength": 20,
			//   "pattern": "arn:.*",
			//   "type": "string"
			// }
			Description: "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A textual description for the workflow.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "^[\\w\\- ]*$",
			//   "type": "string"
			// }
			Description: "A textual description for the workflow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"on_exception_steps": {
			// Property: OnExceptionSteps
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the steps (actions) to take if any errors are encountered during execution of the workflow.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The basic building block of a workflow.",
			//     "properties": {
			//       "CopyStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that performs a file copy.",
			//         "properties": {
			//           "DestinationFileLocation": {
			//             "additionalProperties": false,
			//             "description": "Specifies the location for the file being copied. Only applicable for the Copy type of workflow steps.",
			//             "properties": {
			//               "S3FileLocation": {
			//                 "additionalProperties": false,
			//                 "description": "Specifies the details for the S3 file being copied.",
			//                 "properties": {
			//                   "Bucket": {
			//                     "description": "Specifies the S3 bucket that contains the file being copied.",
			//                     "maxLength": 63,
			//                     "minLength": 3,
			//                     "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
			//                     "type": "string"
			//                   },
			//                   "Key": {
			//                     "description": "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
			//                     "maxLength": 1024,
			//                     "minLength": 0,
			//                     "pattern": ".*",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           },
			//           "OverwriteExisting": {
			//             "description": "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
			//             "enum": [
			//               "TRUE",
			//               "FALSE"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "CustomStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that invokes a lambda function.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           },
			//           "Target": {
			//             "description": "The ARN for the lambda function that is being called.",
			//             "maxLength": 170,
			//             "minLength": 0,
			//             "pattern": "arn:[a-z-]+:lambda:.*$",
			//             "type": "string"
			//           },
			//           "TimeoutSeconds": {
			//             "description": "Timeout, in seconds, for the step.",
			//             "maximum": 1800,
			//             "minimum": 1,
			//             "type": "integer"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "DeleteStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that deletes the file.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "TagStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that creates one or more tags.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           },
			//           "Tags": {
			//             "description": "Array that contains from 1 to 10 key/value pairs.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "Specifies the key-value pair that are assigned to a file during the execution of a Tagging step.",
			//               "properties": {
			//                 "Key": {
			//                   "description": "The name assigned to the tag that you create.",
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "description": "The value that corresponds to the key.",
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Key",
			//                 "Value"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 10,
			//             "type": "array",
			//             "uniqueItems": true
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Type": {
			//         "enum": [
			//           "COPY",
			//           "CUSTOM",
			//           "DELETE",
			//           "TAG"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 8,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Specifies the steps (actions) to take if any errors are encountered during execution of the workflow.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"copy_step_details": {
						// Property: CopyStepDetails
						Description: "Details for a step that performs a file copy.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"destination_file_location": {
									// Property: DestinationFileLocation
									Description: "Specifies the location for the file being copied. Only applicable for the Copy type of workflow steps.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"s3_file_location": {
												// Property: S3FileLocation
												Description: "Specifies the details for the S3 file being copied.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"bucket": {
															// Property: Bucket
															Description: "Specifies the S3 bucket that contains the file being copied.",
															Type:        types.StringType,
															Computed:    true,
														},
														"key": {
															// Property: Key
															Description: "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"overwrite_existing": {
									// Property: OverwriteExisting
									Description: "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"custom_step_details": {
						// Property: CustomStepDetails
						Description: "Details for a step that invokes a lambda function.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"target": {
									// Property: Target
									Description: "The ARN for the lambda function that is being called.",
									Type:        types.StringType,
									Computed:    true,
								},
								"timeout_seconds": {
									// Property: TimeoutSeconds
									Description: "Timeout, in seconds, for the step.",
									Type:        types.Int64Type,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"delete_step_details": {
						// Property: DeleteStepDetails
						Description: "Details for a step that deletes the file.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"tag_step_details": {
						// Property: TagStepDetails
						Description: "Details for a step that creates one or more tags.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"tags": {
									// Property: Tags
									Description: "Array that contains from 1 to 10 key/value pairs.",
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Description: "The name assigned to the tag that you create.",
												Type:        types.StringType,
												Computed:    true,
											},
											"value": {
												// Property: Value
												Description: "The value that corresponds to the key.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
										tfsdk.SetNestedAttributesOptions{},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"steps": {
			// Property: Steps
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the details for the steps that are in the specified workflow.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The basic building block of a workflow.",
			//     "properties": {
			//       "CopyStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that performs a file copy.",
			//         "properties": {
			//           "DestinationFileLocation": {
			//             "additionalProperties": false,
			//             "description": "Specifies the location for the file being copied. Only applicable for the Copy type of workflow steps.",
			//             "properties": {
			//               "S3FileLocation": {
			//                 "additionalProperties": false,
			//                 "description": "Specifies the details for the S3 file being copied.",
			//                 "properties": {
			//                   "Bucket": {
			//                     "description": "Specifies the S3 bucket that contains the file being copied.",
			//                     "maxLength": 63,
			//                     "minLength": 3,
			//                     "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
			//                     "type": "string"
			//                   },
			//                   "Key": {
			//                     "description": "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
			//                     "maxLength": 1024,
			//                     "minLength": 0,
			//                     "pattern": ".*",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           },
			//           "OverwriteExisting": {
			//             "description": "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
			//             "enum": [
			//               "TRUE",
			//               "FALSE"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "CustomStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that invokes a lambda function.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           },
			//           "Target": {
			//             "description": "The ARN for the lambda function that is being called.",
			//             "maxLength": 170,
			//             "minLength": 0,
			//             "pattern": "arn:[a-z-]+:lambda:.*$",
			//             "type": "string"
			//           },
			//           "TimeoutSeconds": {
			//             "description": "Timeout, in seconds, for the step.",
			//             "maximum": 1800,
			//             "minimum": 1,
			//             "type": "integer"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "DeleteStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that deletes the file.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "TagStepDetails": {
			//         "additionalProperties": false,
			//         "description": "Details for a step that creates one or more tags.",
			//         "properties": {
			//           "Name": {
			//             "description": "The name of the step, used as an identifier.",
			//             "maxLength": 30,
			//             "minLength": 0,
			//             "pattern": "^[\\w-]*$",
			//             "type": "string"
			//           },
			//           "Tags": {
			//             "description": "Array that contains from 1 to 10 key/value pairs.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "Specifies the key-value pair that are assigned to a file during the execution of a Tagging step.",
			//               "properties": {
			//                 "Key": {
			//                   "description": "The name assigned to the tag that you create.",
			//                   "maxLength": 128,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "description": "The value that corresponds to the key.",
			//                   "maxLength": 256,
			//                   "minLength": 0,
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Key",
			//                 "Value"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 10,
			//             "type": "array",
			//             "uniqueItems": true
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Type": {
			//         "enum": [
			//           "COPY",
			//           "CUSTOM",
			//           "DELETE",
			//           "TAG"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 8,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Specifies the details for the steps that are in the specified workflow.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"copy_step_details": {
						// Property: CopyStepDetails
						Description: "Details for a step that performs a file copy.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"destination_file_location": {
									// Property: DestinationFileLocation
									Description: "Specifies the location for the file being copied. Only applicable for the Copy type of workflow steps.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"s3_file_location": {
												// Property: S3FileLocation
												Description: "Specifies the details for the S3 file being copied.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"bucket": {
															// Property: Bucket
															Description: "Specifies the S3 bucket that contains the file being copied.",
															Type:        types.StringType,
															Computed:    true,
														},
														"key": {
															// Property: Key
															Description: "The name assigned to the file when it was created in S3. You use the object key to retrieve the object.",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"overwrite_existing": {
									// Property: OverwriteExisting
									Description: "A flag that indicates whether or not to overwrite an existing file of the same name. The default is FALSE.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"custom_step_details": {
						// Property: CustomStepDetails
						Description: "Details for a step that invokes a lambda function.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"target": {
									// Property: Target
									Description: "The ARN for the lambda function that is being called.",
									Type:        types.StringType,
									Computed:    true,
								},
								"timeout_seconds": {
									// Property: TimeoutSeconds
									Description: "Timeout, in seconds, for the step.",
									Type:        types.Int64Type,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"delete_step_details": {
						// Property: DeleteStepDetails
						Description: "Details for a step that deletes the file.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"tag_step_details": {
						// Property: TagStepDetails
						Description: "Details for a step that creates one or more tags.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the step, used as an identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"tags": {
									// Property: Tags
									Description: "Array that contains from 1 to 10 key/value pairs.",
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Description: "The name assigned to the tag that you create.",
												Type:        types.StringType,
												Computed:    true,
											},
											"value": {
												// Property: Value
												Description: "The value that corresponds to the key.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
										tfsdk.SetNestedAttributesOptions{},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Creates a key-value pair for a specific resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The name assigned to the tag that you create.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Contains one or more values that you assigned to the key name you create.",
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
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name assigned to the tag that you create.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "Contains one or more values that you assigned to the key name you create.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"workflow_id": {
			// Property: WorkflowId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the workflow.",
			//   "maxLength": 19,
			//   "minLength": 19,
			//   "pattern": "^w-([a-z0-9]{17})$",
			//   "type": "string"
			// }
			Description: "A unique identifier for the workflow.",
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
		Description: "Data Source schema for AWS::Transfer::Workflow",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Workflow").WithTerraformTypeName("awscc_transfer_workflow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"bucket":                    "Bucket",
		"copy_step_details":         "CopyStepDetails",
		"custom_step_details":       "CustomStepDetails",
		"delete_step_details":       "DeleteStepDetails",
		"description":               "Description",
		"destination_file_location": "DestinationFileLocation",
		"key":                       "Key",
		"name":                      "Name",
		"on_exception_steps":        "OnExceptionSteps",
		"overwrite_existing":        "OverwriteExisting",
		"s3_file_location":          "S3FileLocation",
		"steps":                     "Steps",
		"tag_step_details":          "TagStepDetails",
		"tags":                      "Tags",
		"target":                    "Target",
		"timeout_seconds":           "TimeoutSeconds",
		"type":                      "Type",
		"value":                     "Value",
		"workflow_id":               "WorkflowId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
