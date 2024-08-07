// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_log_stream", logStreamDataSource)
}

// logStreamDataSource returns the Terraform awscc_logs_log_stream data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::LogStream resource.
func logStreamDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LogGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the log group where the log stream is created.",
		//	  "type": "string"
		//	}
		"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the log group where the log stream is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogStreamName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the log stream. The name must be unique wihtin the log group.",
		//	  "type": "string"
		//	}
		"log_stream_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the log stream. The name must be unique wihtin the log group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Logs::LogStream",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogStream").WithTerraformTypeName("awscc_logs_log_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"log_group_name":  "LogGroupName",
		"log_stream_name": "LogStreamName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
