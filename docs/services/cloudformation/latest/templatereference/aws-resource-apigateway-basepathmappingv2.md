This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::BasePathMappingV2

The `AWS::ApiGateway::BasePathMappingV2` resource creates a base path that clients who call your
API must use in the invocation URL. Supported only for private custom domain names.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::BasePathMappingV2",
  "Properties" : {
      "BasePath" : String,
      "DomainNameArn" : String,
      "RestApiId" : String,
      "Stage" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::BasePathMappingV2
Properties:
  BasePath: String
  DomainNameArn: String
  RestApiId: String
  Stage: String

```

## Properties

`BasePath`

The base path name that callers of the private API must provide as part of the URL after the domain name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainNameArn`

The ARN of the domain name for the BasePathMappingV2 resource to be described.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The private API's identifier. This identifier is unique across all of your APIs in API Gateway.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

Represents a unique identifier for a version of a deployed private RestApi that is callable by users. The Stage must depend on the `RestApi`'s stage. To create a dependency, add a DependsOn attribute to the BasePathMappingV2 resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Base path mapping creation example

The following example creates a `BasePathMappingV2` resource called `MyBasePathMappingV2`.

#### JSON

```json

{
    "MyBasePathMappingV2": {
        "Type": "AWS::ApiGateway::BasePathMappingV2",
        "DependsOn": [
            "PrivateApiStage"
         ],
        "Properties": {
            "BasePath": "prod",
            "DomainNameArn": {
                "Fn::GetAtt": [
                    "MyDomainName",
                    "DomainNameArn"
                ]
            },
            "RestApiId": "a1b2c3",
            "Stage": "users"
        }
    }
}
```

#### YAML

```yaml

MyBasePathMappingV2:
  Type: AWS::ApiGateway::BasePathMappingV2
  DependsOn:
   - PrivateApiStage
  Properties:
    BasePath: prod
    DomainNameArn: !GetAtt MyDomainName.DomainNameArn
    RestApiId: a1b2c3
    Stage: users
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::BasePathMapping

AWS::ApiGateway::ClientCertificate
