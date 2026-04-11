This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::ApiMapping

The `AWS::ApiGatewayV2::ApiMapping` resource contains an API mapping.
An API mapping relates a path of your custom domain name to a stage of your API. A
custom domain name can have multiple API mappings, but the paths can't overlap. A
custom domain can map only to APIs of the same protocol type. For more
information, see [CreateApiMapping](../../../apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.md#CreateApiMapping) in the _Amazon API Gateway V2 API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::ApiMapping",
  "Properties" : {
      "ApiId" : String,
      "ApiMappingKey" : String,
      "DomainName" : String,
      "Stage" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::ApiMapping
Properties:
  ApiId: String
  ApiMappingKey: String
  DomainName: String
  Stage: String

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiMappingKey`

The API mapping key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Stage`

The API stage.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the API mapping resource ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApiMappingId`

The API mapping resource ID.

## Examples

### API mapping creation example

The following example creates an `ApiMapping` resource
called `MyApiMapping`.

#### JSON

```json

{
    "MyApiMapping": {
        "Type": "AWS::ApiGatewayV2::ApiMapping",
        "Properties": {
            "DomainName": "mydomainame.us-east-1.com",
            "ApiId": {
                "Ref": "MyApi"
            },
            "Stage": {
                "Ref": "MyStage"
            }
        }
    }
}
```

#### YAML

```yaml

MyApiMapping:
  Type: 'AWS::ApiGatewayV2::ApiMapping'
  Properties:
    DomainName: mydomainame.us-east-1.com
    ApiId: !Ref MyApi
    Stage: !Ref MyStage

```

## See also

- [CreateApiMapping](../../../apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.md#CreateApiMapping) in the _Amazon API_
_Gateway Version 2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StageOverrides

AWS::ApiGatewayV2::Authorizer

All content copied from https://docs.aws.amazon.com/.
