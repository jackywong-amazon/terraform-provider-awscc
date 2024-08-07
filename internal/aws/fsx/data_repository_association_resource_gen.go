// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package fsx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_fsx_data_repository_association", dataRepositoryAssociationResource)
}

// dataRepositoryAssociationResource returns the Terraform awscc_fsx_data_repository_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::FSx::DataRepositoryAssociation resource.
func dataRepositoryAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The system-generated, unique ID of the data repository association.",
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The system-generated, unique ID of the data repository association.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BatchImportMetaDataOnCreate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true.",
		//	  "type": "boolean"
		//	}
		"batch_import_meta_data_on_create": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to true.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataRepositoryPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to.",
		//	  "type": "string"
		//	}
		"data_repository_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository files will be imported from or exported to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The globally unique ID of the file system, assigned by Amazon FSx.",
		//	  "type": "string"
		//	}
		"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The globally unique ID of the file system, assigned by Amazon FSx.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.",
		//	  "type": "string"
		//	}
		"file_system_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImportedFileChunkSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.",
		//	  "type": "integer"
		//	}
		"imported_file_chunk_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify Amazon Web Services resources. We require an ARN when you need to specify a resource unambiguously across all of Amazon Web Services. For more information, see Amazon Resource Names (ARNs) in the Amazon Web Services General Reference.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
		//	  "properties": {
		//	    "AutoExportPolicy": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.",
		//	      "properties": {
		//	        "Events": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "enum": [
		//	              "NEW",
		//	              "CHANGED",
		//	              "DELETED"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 3,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "Events"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "AutoImportPolicy": {
		//	      "additionalProperties": false,
		//	      "description": "Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.",
		//	      "properties": {
		//	        "Events": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "enum": [
		//	              "NEW",
		//	              "CHANGED",
		//	              "DELETED"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 3,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "Events"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoExportPolicy
				"auto_export_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Events
						"events": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeAtMost(3),
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"NEW",
										"CHANGED",
										"DELETED",
									),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: AutoImportPolicy
				"auto_import_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Events
						"events": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeAtMost(3),
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"NEW",
										"CHANGED",
										"DELETED",
									),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of Tag values, with a maximum of 50 elements.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of Tag values, with a maximum of 50 elements.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::FSx::DataRepositoryAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FSx::DataRepositoryAssociation").WithTerraformTypeName("awscc_fsx_data_repository_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id":                   "AssociationId",
		"auto_export_policy":               "AutoExportPolicy",
		"auto_import_policy":               "AutoImportPolicy",
		"batch_import_meta_data_on_create": "BatchImportMetaDataOnCreate",
		"data_repository_path":             "DataRepositoryPath",
		"events":                           "Events",
		"file_system_id":                   "FileSystemId",
		"file_system_path":                 "FileSystemPath",
		"imported_file_chunk_size":         "ImportedFileChunkSize",
		"key":                              "Key",
		"resource_arn":                     "ResourceARN",
		"s3":                               "S3",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(120).WithDeleteTimeoutInMinutes(180)

	opts = opts.WithUpdateTimeoutInMinutes(180)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
