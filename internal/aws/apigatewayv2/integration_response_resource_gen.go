// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigatewayv2_integration_response", integrationResponseResource)
}

// integrationResponseResource returns the Terraform awscc_apigatewayv2_integration_response resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGatewayV2::IntegrationResponse resource.
func integrationResponseResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API identifier",
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API identifier",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContentHandlingStrategy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " Specifies how to handle response payload content type conversions",
		//	  "type": "string"
		//	}
		"content_handling_strategy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: " Specifies how to handle response payload content type conversions",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IntegrationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The integration ID",
		//	  "type": "string"
		//	}
		"integration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The integration ID",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IntegrationResponseId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Integration Response ID generated by service",
		//	  "type": "string"
		//	}
		"integration_response_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Integration Response ID generated by service",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IntegrationResponseKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The integration response key",
		//	  "type": "string"
		//	}
		"integration_response_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The integration response key",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResponseParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A key-value map specifying response parameters that are passed to the method response from the backend",
		//	  "type": "object"
		//	}
		"response_parameters": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value map specifying response parameters that are passed to the method response from the backend",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResponseTemplates
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The collection of response templates for the integration response as a string-to-string map of key-value pairs",
		//	  "type": "object"
		//	}
		"response_templates": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The collection of response templates for the integration response as a string-to-string map of key-value pairs",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateSelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The template selection expression for the integration response. Supported only for WebSocket APIs",
		//	  "type": "string"
		//	}
		"template_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The template selection expression for the integration response. Supported only for WebSocket APIs",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "Schema for ApiGatewayV2 Integration Response",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::IntegrationResponse").WithTerraformTypeName("awscc_apigatewayv2_integration_response")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":                        "ApiId",
		"content_handling_strategy":     "ContentHandlingStrategy",
		"integration_id":                "IntegrationId",
		"integration_response_id":       "IntegrationResponseId",
		"integration_response_key":      "IntegrationResponseKey",
		"response_parameters":           "ResponseParameters",
		"response_templates":            "ResponseTemplates",
		"template_selection_expression": "TemplateSelectionExpression",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
