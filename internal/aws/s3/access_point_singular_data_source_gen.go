// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_s3_access_point", accessPointDataSourceType)
}

// accessPointDataSourceType returns the Terraform awscc_s3_access_point data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::S3::AccessPoint resource type.
func accessPointDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"alias": {
			// Property: Alias
			// CloudFormation resource type schema:
			// {
			//   "description": "The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.",
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "the Amazon Resource Name (ARN) of the specified accesspoint.",
			//   "type": "string"
			// }
			Description: "the Amazon Resource Name (ARN) of the specified accesspoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the bucket that you want to associate this Access Point with.",
			//   "maxLength": 255,
			//   "minLength": 3,
			//   "type": "string"
			// }
			Description: "The name of the bucket that you want to associate this Access Point with.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.",
			//   "maxLength": 50,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"network_origin": {
			// Property: NetworkOrigin
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
			//   "enum": [
			//     "Internet",
			//     "VPC"
			//   ],
			//   "type": "string"
			// }
			Description: "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "The Access Point Policy you want to apply to this access point.",
			//   "type": "object"
			// }
			Description: "The Access Point Policy you want to apply to this access point.",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"policy_status": {
			// Property: PolicyStatus
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "IsPublic": {
			//       "description": "Specifies whether the policy is public or not.",
			//       "enum": [
			//         "true",
			//         "false"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"is_public": {
						// Property: IsPublic
						Description: "Specifies whether the policy is public or not.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"public_access_block_configuration": {
			// Property: PublicAccessBlockConfiguration
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "BlockPublicAcls": {
			//       "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
			//       "type": "boolean"
			//     },
			//     "BlockPublicPolicy": {
			//       "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
			//       "type": "boolean"
			//     },
			//     "IgnorePublicAcls": {
			//       "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
			//       "type": "boolean"
			//     },
			//     "RestrictPublicBuckets": {
			//       "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"block_public_acls": {
						// Property: BlockPublicAcls
						Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"block_public_policy": {
						// Property: BlockPublicPolicy
						Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"ignore_public_acls": {
						// Property: IgnorePublicAcls
						Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
						Type:        types.BoolType,
						Computed:    true,
					},
					"restrict_public_buckets": {
						// Property: RestrictPublicBuckets
						Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"vpc_configuration": {
			// Property: VpcConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "The Virtual Private Cloud (VPC) configuration for a bucket access point.",
			//   "properties": {
			//     "VpcId": {
			//       "description": "If this field is specified, this access point will only allow connections from the specified VPC ID.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Virtual Private Cloud (VPC) configuration for a bucket access point.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"vpc_id": {
						// Property: VpcId
						Description: "If this field is specified, this access point will only allow connections from the specified VPC ID.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::S3::AccessPoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::AccessPoint").WithTerraformTypeName("awscc_s3_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":                             "Alias",
		"arn":                               "Arn",
		"block_public_acls":                 "BlockPublicAcls",
		"block_public_policy":               "BlockPublicPolicy",
		"bucket":                            "Bucket",
		"ignore_public_acls":                "IgnorePublicAcls",
		"is_public":                         "IsPublic",
		"name":                              "Name",
		"network_origin":                    "NetworkOrigin",
		"policy":                            "Policy",
		"policy_status":                     "PolicyStatus",
		"public_access_block_configuration": "PublicAccessBlockConfiguration",
		"restrict_public_buckets":           "RestrictPublicBuckets",
		"vpc_configuration":                 "VpcConfiguration",
		"vpc_id":                            "VpcId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_s3_access_point", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}