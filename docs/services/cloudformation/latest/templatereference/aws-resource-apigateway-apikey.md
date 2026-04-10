This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::ApiKey

The `AWS::ApiGateway::ApiKey` resource creates a unique key that you can distribute to clients who are executing API Gateway `Method` resources that require an API key. To specify which API key clients must use, map the API key with the `RestApi` and `Stage` resources that include the methods that require a key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::ApiKey",
  "Properties" : {
      "CustomerId" : String,
      "Description" : String,
      "Enabled" : Boolean,
      "GenerateDistinctId" : Boolean,
      "Name" : String,
      "StageKeys" : [ StageKey, ... ],
      "Tags" : [ Tag, ... ],
      "Value" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::ApiKey
Properties:
  CustomerId: String
  Description: String
  Enabled: Boolean
  GenerateDistinctId: Boolean
  Name: String
  StageKeys:
    - StageKey
  Tags:
    - Tag
  Value: String

```

## Properties

`CustomerId`

An AWS Marketplace customer identifier, when integrating with the AWS SaaS Marketplace.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the ApiKey.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether the ApiKey can be used by callers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerateDistinctId`

Specifies whether ( `true`) or not ( `false`) the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A name for the API key. If you don't specify a name, CloudFormation generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](../userguide/aws-properties-name.md).

###### Important

If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StageKeys`

DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.

_Required_: No

_Type_: Array of [StageKey](aws-properties-apigateway-apikey-stagekey.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

_Required_: No

_Type_: Array of [Tag](aws-properties-apigateway-apikey-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Specifies a value of the API key.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the API key ID, such as `m2m1k7sybf`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`APIKeyId`

The ID for the API key. For example: `abc123`.

## Examples

### API Key

The following example creates an API key and associates it with the `Test` stage of the `TestAPIDeployment` deployment. To ensure that CloudFormation creates the stage and deployment (which are declared elsewhere in the same template) before the API key, the example adds an explicit dependency on the deployment and stage. Without this dependency, CloudFormation might create the API key first, which would cause the association to fail because the deployment and stage wouldn't exist.

#### JSON

```json

{
    "ApiKey": {
        "Type": "AWS::ApiGateway::ApiKey",
        "DependsOn": [
            "TestAPIDeployment",
            "Test"
        ],
        "Properties": {
            "Name": "TestApiKey",
            "Description": "CloudFormation API Key V1",
            "Enabled": true,
            "StageKeys": [
                {
                    "RestApiId": {
                        "Ref": "RestApi"
                    },
                    "StageName": "Test"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

ApiKey:
  Type: 'AWS::ApiGateway::ApiKey'
  DependsOn:
    - TestAPIDeployment
    - Test
  Properties:
    Name: TestApiKey
    Description: CloudFormation API Key V1
    Enabled: true
    StageKeys:
      - RestApiId: !Ref RestApi
        StageName: Test

```

## See also

- [apikey:create](../../../apigateway/latest/api/api-createapikey.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::Account

StageKey

All content copied from https://docs.aws.amazon.com/.
