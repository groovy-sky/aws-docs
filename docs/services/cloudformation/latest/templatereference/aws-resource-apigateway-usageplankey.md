This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::UsagePlanKey

The `AWS::ApiGateway::UsagePlanKey` resource associates an API key with a usage plan. This association determines which users the usage plan is applied to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::UsagePlanKey",
  "Properties" : {
      "KeyId" : String,
      "KeyType" : String,
      "UsagePlanId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::UsagePlanKey
Properties:
  KeyId: String
  KeyType: String
  UsagePlanId: String

```

## Properties

`KeyId`

The Id of the UsagePlanKey resource.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyType`

The type of a UsagePlanKey resource for a plan customer.

_Required_: Yes

_Type_: String

_Allowed values_: `API_KEY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UsagePlanId`

The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the key and ID of the usage plan combined with a ":", such as `123abcdef:abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID for the usage plan key. For example: `abc123`.

## Examples

### Create usage plan key

#### JSON

```json

{
    "usagePlanKey": {
        "Type": "AWS::ApiGateway::UsagePlanKey",
        "Properties": {
            "KeyId": {
                "Ref": "myApiKey"
            },
            "KeyType": "API_KEY",
            "UsagePlanId": {
                "Ref": "myUsagePlan"
            }
        }
    }
}
```

#### YAML

```yaml

usagePlanKey:
  Type: 'AWS::ApiGateway::UsagePlanKey'
  Properties:
    KeyId: !Ref myApiKey
    KeyType: API_KEY
    UsagePlanId: !Ref myUsagePlan

```

## See also

- [usageplankey:create](../../../apigateway/latest/api/api-createusageplankey.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThrottleSettings

AWS::ApiGateway::VpcLink

All content copied from https://docs.aws.amazon.com/.
