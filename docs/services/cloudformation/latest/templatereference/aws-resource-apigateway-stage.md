This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Stage

The `AWS::ApiGateway::Stage` resource creates a stage for a deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::Stage",
  "Properties" : {
      "AccessLogSetting" : AccessLogSetting,
      "CacheClusterEnabled" : Boolean,
      "CacheClusterSize" : String,
      "CanarySetting" : CanarySetting,
      "ClientCertificateId" : String,
      "DeploymentId" : String,
      "Description" : String,
      "DocumentationVersion" : String,
      "MethodSettings" : [ MethodSetting, ... ],
      "RestApiId" : String,
      "StageName" : String,
      "Tags" : [ Tag, ... ],
      "TracingEnabled" : Boolean,
      "Variables" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::Stage
Properties:
  AccessLogSetting:
    AccessLogSetting
  CacheClusterEnabled: Boolean
  CacheClusterSize: String
  CanarySetting:
    CanarySetting
  ClientCertificateId: String
  DeploymentId: String
  Description: String
  DocumentationVersion: String
  MethodSettings:
    - MethodSetting
  RestApiId: String
  StageName: String
  Tags:
    - Tag
  TracingEnabled: Boolean
  Variables:
    Key: Value

```

## Properties

`AccessLogSetting`

Access log settings, including the access log format and access log destination ARN.

_Required_: No

_Type_: [AccessLogSetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-stage-accesslogsetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheClusterEnabled`

Specifies whether a cache cluster is enabled for the stage. To activate a method-level cache, set `CachingEnabled` to `true` for a method.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheClusterSize`

The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](../../../apigateway/latest/developerguide/api-gateway-caching.md).

_Required_: No

_Type_: String

_Allowed values_: `0.5 | 1.6 | 6.1 | 13.5 | 28.4 | 58.2 | 118 | 237`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CanarySetting`

Settings for the canary deployment in this stage.

_Required_: No

_Type_: [CanarySetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-stage-canarysetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCertificateId`

The identifier of a client certificate for an API stage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentId`

The identifier of the Deployment that the stage points to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The stage's description.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentationVersion`

The version of the associated API documentation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MethodSettings`

A map that defines the method settings for a Stage resource. Keys (designated as `/{method_setting_key` below) are method paths defined as `{resource_path}/{http_method}` for an individual method override, or `/\*/\*` for overriding all methods in the stage.

_Required_: No

_Type_: Array of [MethodSetting](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-stage-methodsetting.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageName`

The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-stage-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TracingEnabled`

Specifies whether active tracing with X-ray is enabled for the Stage.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+`.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the stage, such as `MyTestStage`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create stage

The following example creates a stage for the `TestDeployment` deployment. The stage also specifies method settings for the `MyRestApi` API.

#### JSON

```json

{
    "Resources": {
        "Prod": {
            "Type": "AWS::ApiGateway::Stage",
            "Properties": {
                "StageName": "Prod",
                "Description": "Prod Stage",
                "RestApiId": {
                    "Ref": "MyRestApi"
                },
                "DeploymentId": {
                    "Ref": "TestDeployment"
                },
                "DocumentationVersion": {
                    "Ref": "MyDocumentationVersion"
                },
                "ClientCertificateId": {
                    "Ref": "ClientCertificate"
                },
                "Variables": {
                    "Stack": "Prod"
                },
                "MethodSettings": [
                    {
                        "ResourcePath": "/",
                        "HttpMethod": "GET",
                        "MetricsEnabled": "true",
                        "DataTraceEnabled": "false"
                    },
                    {
                        "ResourcePath": "/stack",
                        "HttpMethod": "POST",
                        "MetricsEnabled": "true",
                        "DataTraceEnabled": "false",
                        "ThrottlingBurstLimit": "999"
                    },
                    {
                        "ResourcePath": "/stack",
                        "HttpMethod": "GET",
                        "MetricsEnabled": "true",
                        "DataTraceEnabled": "false",
                        "ThrottlingBurstLimit": "555"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  Prod:
    Type: AWS::ApiGateway::Stage
    Properties:
      StageName: Prod
      Description: Prod Stage
      RestApiId: !Ref MyRestApi
      DeploymentId: !Ref TestDeployment
      DocumentationVersion: !Ref MyDocumentationVersion
      ClientCertificateId: !Ref ClientCertificate
      Variables:
        Stack: Prod
      MethodSettings:
        - ResourcePath: /
          HttpMethod: GET
          MetricsEnabled: 'true'
          DataTraceEnabled: 'false'
        - ResourcePath: /stack
          HttpMethod: POST
          MetricsEnabled: 'true'
          DataTraceEnabled: 'false'
          ThrottlingBurstLimit: '999'
        - ResourcePath: /stack
          HttpMethod: GET
          MetricsEnabled: 'true'
          DataTraceEnabled: 'false'
          ThrottlingBurstLimit: '555'
```

## See also

- [stage:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateStage.html) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AccessLogSetting
