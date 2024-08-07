// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediaconnect_bridge_output", bridgeOutputDataSource)
}

// bridgeOutputDataSource returns the Terraform awscc_mediaconnect_bridge_output data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaConnect::BridgeOutput resource.
func bridgeOutputDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BridgeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) of the bridge.",
		//	  "type": "string"
		//	}
		"bridge_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) of the bridge.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The network output name.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The network output name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkOutput
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The output of the bridge.",
		//	  "properties": {
		//	    "IpAddress": {
		//	      "description": "The network output IP Address.",
		//	      "type": "string"
		//	    },
		//	    "NetworkName": {
		//	      "description": "The network output's gateway network name.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "The network output port.",
		//	      "type": "integer"
		//	    },
		//	    "Protocol": {
		//	      "description": "The network output protocol.",
		//	      "enum": [
		//	        "rtp-fec",
		//	        "rtp",
		//	        "udp"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Ttl": {
		//	      "description": "The network output TTL.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "Protocol",
		//	    "IpAddress",
		//	    "Port",
		//	    "NetworkName",
		//	    "Ttl"
		//	  ],
		//	  "type": "object"
		//	}
		"network_output": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IpAddress
				"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The network output IP Address.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: NetworkName
				"network_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The network output's gateway network name.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The network output port.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Protocol
				"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The network output protocol.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Ttl
				"ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The network output TTL.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The output of the bridge.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaConnect::BridgeOutput",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::BridgeOutput").WithTerraformTypeName("awscc_mediaconnect_bridge_output")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bridge_arn":     "BridgeArn",
		"ip_address":     "IpAddress",
		"name":           "Name",
		"network_name":   "NetworkName",
		"network_output": "NetworkOutput",
		"port":           "Port",
		"protocol":       "Protocol",
		"ttl":            "Ttl",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
