// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lex_bot_alias", botAliasDataSourceType)
}

// botAliasDataSourceType returns the Terraform awscc_lex_bot_alias data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lex::BotAlias resource type.
func botAliasDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
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
		"bot_alias_id": {
			// Property: BotAliasId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique ID of resource",
			//   "maxLength": 10,
			//   "minLength": 10,
			//   "pattern": "^[0-9a-zA-Z]+$",
			//   "type": "string"
			// }
			Description: "Unique ID of resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"bot_alias_locale_settings": {
			// Property: BotAliasLocaleSettings
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of bot alias locale settings to add to the bot alias.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A locale setting in alias",
			//     "properties": {
			//       "BotAliasLocaleSetting": {
			//         "additionalProperties": false,
			//         "description": "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
			//         "properties": {
			//           "CodeHookSpecification": {
			//             "additionalProperties": false,
			//             "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
			//             "properties": {
			//               "LambdaCodeHook": {
			//                 "additionalProperties": false,
			//                 "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
			//                 "properties": {
			//                   "CodeHookInterfaceVersion": {
			//                     "description": "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
			//                     "maxLength": 5,
			//                     "minLength": 1,
			//                     "type": "string"
			//                   },
			//                   "LambdaArn": {
			//                     "description": "The Amazon Resource Name (ARN) of the Lambda function.",
			//                     "maxLength": 2048,
			//                     "minLength": 20,
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "CodeHookInterfaceVersion",
			//                   "LambdaArn"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "required": [
			//               "LambdaCodeHook"
			//             ],
			//             "type": "object"
			//           },
			//           "Enabled": {
			//             "description": "Whether the Lambda code hook is enabled",
			//             "type": "boolean"
			//           }
			//         },
			//         "required": [
			//           "Enabled"
			//         ],
			//         "type": "object"
			//       },
			//       "LocaleId": {
			//         "description": "A string used to identify the locale",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "LocaleId",
			//       "BotAliasLocaleSetting"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of bot alias locale settings to add to the bot alias.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"bot_alias_locale_setting": {
						// Property: BotAliasLocaleSetting
						Description: "You can use this parameter to specify a specific Lambda function to run different functions in different locales.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"code_hook_specification": {
									// Property: CodeHookSpecification
									Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"lambda_code_hook": {
												// Property: LambdaCodeHook
												Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"code_hook_interface_version": {
															// Property: CodeHookInterfaceVersion
															Description: "The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.",
															Type:        types.StringType,
															Computed:    true,
														},
														"lambda_arn": {
															// Property: LambdaArn
															Description: "The Amazon Resource Name (ARN) of the Lambda function.",
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
								"enabled": {
									// Property: Enabled
									Description: "Whether the Lambda code hook is enabled",
									Type:        types.BoolType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"locale_id": {
						// Property: LocaleId
						Description: "A string used to identify the locale",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"bot_alias_name": {
			// Property: BotAliasName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for a resource.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "^([0-9a-zA-Z][_-]?)+$",
			//   "type": "string"
			// }
			Description: "A unique identifier for a resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bot_alias_status": {
			// Property: BotAliasStatus
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "Creating",
			//     "Available",
			//     "Deleting",
			//     "Failed"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"bot_alias_tags": {
			// Property: BotAliasTags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to add to the bot alias.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Lex resources",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of tags to add to the bot alias.",
			Attributes: tfsdk.SetNestedAttributes(
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
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"bot_id": {
			// Property: BotId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique ID of resource",
			//   "maxLength": 10,
			//   "minLength": 10,
			//   "pattern": "^[0-9a-zA-Z]+$",
			//   "type": "string"
			// }
			Description: "Unique ID of resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"bot_version": {
			// Property: BotVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of a bot.",
			//   "maxLength": 5,
			//   "minLength": 1,
			//   "pattern": "^(DRAFT|[0-9]+)$",
			//   "type": "string"
			// }
			Description: "The version of a bot.",
			Type:        types.StringType,
			Computed:    true,
		},
		"conversation_log_settings": {
			// Property: ConversationLogSettings
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
			//   "properties": {
			//     "AudioLogSettings": {
			//       "description": "List of audio log settings",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Settings for logging audio of conversations between Amazon Lex and a user. You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.",
			//         "properties": {
			//           "Destination": {
			//             "additionalProperties": false,
			//             "description": "The location of audio log files collected when conversation logging is enabled for a bot.",
			//             "properties": {
			//               "S3Bucket": {
			//                 "additionalProperties": false,
			//                 "description": "Specifies an Amazon S3 bucket for logging audio conversations",
			//                 "properties": {
			//                   "KmsKeyArn": {
			//                     "description": "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
			//                     "maxLength": 2048,
			//                     "minLength": 20,
			//                     "pattern": "^arn:[\\w\\-]+:kms:[\\w\\-]+:[\\d]{12}:(?:key\\/[\\w\\-]+|alias\\/[a-zA-Z0-9:\\/_\\-]{1,256})$",
			//                     "type": "string"
			//                   },
			//                   "LogPrefix": {
			//                     "description": "The Amazon S3 key of the deployment package.",
			//                     "maxLength": 1024,
			//                     "minLength": 0,
			//                     "type": "string"
			//                   },
			//                   "S3BucketArn": {
			//                     "description": "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
			//                     "maxLength": 2048,
			//                     "minLength": 1,
			//                     "pattern": "^arn:[\\w\\-]+:s3:::[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "LogPrefix",
			//                   "S3BucketArn"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Enabled": {
			//             "description": "",
			//             "type": "boolean"
			//           }
			//         },
			//         "required": [
			//           "Destination",
			//           "Enabled"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "TextLogSettings": {
			//       "description": "List of text log settings",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Contains information about code hooks that Amazon Lex calls during a conversation.",
			//         "properties": {
			//           "Destination": {
			//             "additionalProperties": false,
			//             "description": "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
			//             "properties": {
			//               "CloudWatch": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "CloudWatchLogGroupArn": {
			//                     "description": "A string used to identify this tag",
			//                     "maxLength": 2048,
			//                     "minLength": 1,
			//                     "type": "string"
			//                   },
			//                   "LogPrefix": {
			//                     "description": "A string containing the value for the tag",
			//                     "maxLength": 1024,
			//                     "minLength": 0,
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "CloudWatchLogGroupArn",
			//                   "LogPrefix"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Enabled": {
			//             "description": "",
			//             "type": "boolean"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Contains information about code hooks that Amazon Lex calls during a conversation.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"audio_log_settings": {
						// Property: AudioLogSettings
						Description: "List of audio log settings",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"destination": {
									// Property: Destination
									Description: "The location of audio log files collected when conversation logging is enabled for a bot.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"s3_bucket": {
												// Property: S3Bucket
												Description: "Specifies an Amazon S3 bucket for logging audio conversations",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"kms_key_arn": {
															// Property: KmsKeyArn
															Description: "The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.",
															Type:        types.StringType,
															Computed:    true,
														},
														"log_prefix": {
															// Property: LogPrefix
															Description: "The Amazon S3 key of the deployment package.",
															Type:        types.StringType,
															Computed:    true,
														},
														"s3_bucket_arn": {
															// Property: S3BucketArn
															Description: "The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.",
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
								"enabled": {
									// Property: Enabled
									Description: "",
									Type:        types.BoolType,
									Computed:    true,
								},
							},
							tfsdk.SetNestedAttributesOptions{},
						),
						Computed: true,
					},
					"text_log_settings": {
						// Property: TextLogSettings
						Description: "List of text log settings",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"destination": {
									// Property: Destination
									Description: "Defines the Amazon CloudWatch Logs destination log group for conversation text logs.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cloudwatch": {
												// Property: CloudWatch
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"cloudwatch_log_group_arn": {
															// Property: CloudWatchLogGroupArn
															Description: "A string used to identify this tag",
															Type:        types.StringType,
															Computed:    true,
														},
														"log_prefix": {
															// Property: LogPrefix
															Description: "A string containing the value for the tag",
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
								"enabled": {
									// Property: Enabled
									Description: "",
									Type:        types.BoolType,
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
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the bot alias. Use the description to help identify the bot alias in lists.",
			//   "maxLength": 200,
			//   "type": "string"
			// }
			Description: "A description of the bot alias. Use the description to help identify the bot alias in lists.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sentiment_analysis_settings": {
			// Property: SentimentAnalysisSettings
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
			//   "properties": {
			//     "DetectSentiment": {
			//       "description": "Enable to call Amazon Comprehend for Sentiment natively within Lex",
			//       "type": "boolean"
			//     }
			//   },
			//   "required": [
			//     "DetectSentiment"
			//   ],
			//   "type": "object"
			// }
			Description: "Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"detect_sentiment": {
						// Property: DetectSentiment
						Description: "Enable to call Amazon Comprehend for Sentiment natively within Lex",
						Type:        types.BoolType,
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
		Description: "Data Source schema for AWS::Lex::BotAlias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::BotAlias").WithTerraformTypeName("awscc_lex_bot_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"audio_log_settings":          "AudioLogSettings",
		"bot_alias_id":                "BotAliasId",
		"bot_alias_locale_setting":    "BotAliasLocaleSetting",
		"bot_alias_locale_settings":   "BotAliasLocaleSettings",
		"bot_alias_name":              "BotAliasName",
		"bot_alias_status":            "BotAliasStatus",
		"bot_alias_tags":              "BotAliasTags",
		"bot_id":                      "BotId",
		"bot_version":                 "BotVersion",
		"cloudwatch":                  "CloudWatch",
		"cloudwatch_log_group_arn":    "CloudWatchLogGroupArn",
		"code_hook_interface_version": "CodeHookInterfaceVersion",
		"code_hook_specification":     "CodeHookSpecification",
		"conversation_log_settings":   "ConversationLogSettings",
		"description":                 "Description",
		"destination":                 "Destination",
		"detect_sentiment":            "DetectSentiment",
		"enabled":                     "Enabled",
		"key":                         "Key",
		"kms_key_arn":                 "KmsKeyArn",
		"lambda_arn":                  "LambdaArn",
		"lambda_code_hook":            "LambdaCodeHook",
		"locale_id":                   "LocaleId",
		"log_prefix":                  "LogPrefix",
		"s3_bucket":                   "S3Bucket",
		"s3_bucket_arn":               "S3BucketArn",
		"sentiment_analysis_settings": "SentimentAnalysisSettings",
		"text_log_settings":           "TextLogSettings",
		"value":                       "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
