resource "awscc_apigateway_rest_api" "example" {
  name = "exampleAPI"
  description = "Rest API Gateway"
  endpoint_configuration = {
    types = ["REGIONAL"]
  }

  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}


resource "awscc_apigateway_deployment" "example" {
  rest_api_id = awscc_apigateway_rest_api.example.id
}

resource "aws_api_gateway_stage" "example" {
  deployment_id = awscc_apigateway_deployment.example.deployment_id
  rest_api_id   = awscc_apigateway_rest_api.example.id
  stage_name    = "example_stage"
}

resource "aws_wafv2_web_acl" "example" {
  name        = "managed-rule-example"
  description = "example of a managed rule"
  scope       = "REGIONAL"

  default_action {
    allow {}
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name                = "example-metric-name"
    sampled_requests_enabled   = false
  }
}

resource "awscc_wafv2_web_acl_association" "example" {
  resource_arn = aws_api_gateway_stage.example.arn
  web_acl_arn  = aws_wafv2_web_acl.example.arn
}
