This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Stage

The `AWS::ApiGatewayV2::Stage` resource specifies a stage for an API. Each stage is a named
reference to a deployment of the API and is made available for client applications to call. To learn more, see
[Working with stages for\
HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-stages.html) and [Deploy a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-set-up-websocket-deployment.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::Stage",
  "Properties" : {
      "AccessLogSettings" : AccessLogSettings,
      "ApiId" : String,
      "AutoDeploy" : Boolean,
      "ClientCertificateId" : String,
      "DefaultRouteSettings" : RouteSettings,
      "DeploymentId" : String,
      "Description" : String,
      "RouteSettings" : Json,
      "StageName" : String,
      "StageVariables" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::Stage
Properties:
  AccessLogSettings:
    AccessLogSettings
  ApiId: String
  AutoDeploy: Boolean
  ClientCertificateId: String
  DefaultRouteSettings:
    RouteSettings
  DeploymentId: String
  Description: String
  RouteSettings: Json
  StageName: String
  StageVariables: Json
  Tags:
    - Tag

```

## Properties

`AccessLogSettings`

Settings for logging access in this stage.

_Required_: No

_Type_: [AccessLogSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-stage-accesslogsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoDeploy`

Specifies whether updates to an API automatically trigger a new deployment. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCertificateId`

The identifier of a client certificate for a `Stage`. Supported only for WebSocket APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRouteSettings`

The default route settings for the stage.

_Required_: No

_Type_: [RouteSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-stage-routesettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentId`

The deployment identifier for the API stage. Can't be updated if `autoDeploy` is enabled.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the API stage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouteSettings`

Route settings for the stage.

_Required_: No

_Type_: [Json](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-stage-routesettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageName`

The stage name. Stage names can contain only alphanumeric characters, hyphens, and underscores, or be `$default`. Maximum length is 128 characters.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageVariables`

A map that defines the stage variables for a `Stage`. Variable names can have alphanumeric and underscore characters, and the values must match \[A-Za-z0-9-.\_~:/?#&=,\]+.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the stage name, such as
`MyTestStage`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Stage creation example

The following example creates a `stage` resource called
`MyStage` and associates it with an existing
`deployment` called `MyDeployment`.

#### JSON

```json

{
    "MyStage": {
        "Type": "AWS::ApiGatewayV2::Stage",
        "Properties": {
            "StageName": "Prod",
            "Description": "Prod Stage",
            "DeploymentId": {
                "Ref": "MyDeployment"
            },
            "ApiId": {
                "Ref": "CFNWebSocket"
            },
            "DefaultRouteSettings": {
                "DetailedMetricsEnabled": true,
                "LoggingLevel": "INFO",
                "DataTraceEnabled": false,
                "ThrottlingBurstLimit": 10,
                "ThrottlingRateLimit": 10
            },
            "AccessLogSettings": {
                "DestinationArn": "arn:aws:logs:us-east-1:123456789:log-group:my-log-group",
                "Format": "{\"requestId\":\"$context.requestId\", \"ip\": \"$context.identity.sourceIp\", \"caller\":\"$context.identity.caller\", \"user\":\"$context.identity.user\",\"requestTime\":\"$context.requestTime\", \"eventType\":\"$context.eventType\",\"routeKey\":\"$context.routeKey\", \"status\":\"$context.status\",\"connectionId\":\"$context.connectionId\"}"
            }
        }
    }
}
```

#### YAML

```yaml

MyStage:
  Type: 'AWS::ApiGatewayV2::Stage'
  Properties:
    StageName: Prod
    Description: Prod Stage
    DeploymentId: !Ref MyDeployment
    ApiId: !Ref CFNWebSocket
    DefaultRouteSettings:
      DetailedMetricsEnabled: true
      LoggingLevel: INFO
      DataTraceEnabled: false
      ThrottlingBurstLimit: 10
      ThrottlingRateLimit: 10
    AccessLogSettings:
      DestinationArn: 'arn:aws:logs:us-east-1:123456789:log-group:my-log-group'
      Format: >-
        {"requestId":"$context.requestId", "ip": "$context.identity.sourceIp",
        "caller":"$context.identity.caller",
        "user":"$context.identity.user","requestTime":"$context.requestTime",
        "eventType":"$context.eventType","routeKey":"$context.routeKey",
        "status":"$context.status","connectionId":"$context.connectionId"}

```

## See also

- [CreateStage](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-stages.html#CreateStage) in the _Amazon API Gateway_
_Version 2 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MatchHeaderValue

AccessLogSettings
