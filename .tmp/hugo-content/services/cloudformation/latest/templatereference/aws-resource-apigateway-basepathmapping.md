This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::BasePathMapping

The `AWS::ApiGateway::BasePathMapping` resource creates a base path that clients who call your API
must use in the invocation URL. Supported only for public custom domain names.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::BasePathMapping",
  "Properties" : {
      "BasePath" : String,
      "DomainName" : String,
      "RestApiId" : String,
      "Stage" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::BasePathMapping
Properties:
  BasePath: String
  DomainName: String
  RestApiId: String
  Stage: String

```

## Properties

`BasePath`

The base path name that callers of the API must provide as part of the URL after the domain name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainName`

The domain name of the BasePathMapping resource to be described.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

The name of the associated stage.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [basepathmapping:create](../../../apigateway/latest/api/api-createbasepathmapping.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::Authorizer

AWS::ApiGateway::BasePathMappingV2

All content copied from https://docs.aws.amazon.com/.
