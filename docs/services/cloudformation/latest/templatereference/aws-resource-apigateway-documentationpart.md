This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DocumentationPart

The `AWS::ApiGateway::DocumentationPart` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the _API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::DocumentationPart",
  "Properties" : {
      "Location" : Location,
      "Properties" : String,
      "RestApiId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::DocumentationPart
Properties:
  Location:
    Location
  Properties: String
  RestApiId: String

```

## Properties

`Location`

The location of the targeted API entity of the to-be-created documentation part.

_Required_: Yes

_Type_: [Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-documentationpart-location.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Properties`

The new documentation content map of the targeted API entity. Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the documentation part, such as `abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`DocumentationPartId`

The ID for the documentation part.

## Examples

### Associate documentation part with documentation version

The following example associates a documentation part for an API entity with a documentation version.

#### JSON

```json

{
    "Parameters": {
        "apiName": {
            "Type": "String"
        },
        "description": {
            "Type": "String"
        },
        "version": {
            "Type": "String"
        },
        "type": {
            "Type": "String"
        },
        "property": {
            "Type": "String"
        }
    },
    "Resources": {
        "RestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Name": {
                    "Ref": "apiName"
                }
            }
        },
        "DocumentationPart": {
            "Type": "AWS::ApiGateway::DocumentationPart",
            "Properties": {
                "Location": {
                    "Type": {
                        "Ref": "type"
                    }
                },
                "RestApiId": {
                    "Ref": "RestApi"
                },
                "Properties": {
                    "Ref": "property"
                }
            }
        },
        "DocumentationVersion": {
            "Type": "AWS::ApiGateway::DocumentationVersion",
            "Properties": {
                "Description": {
                    "Ref": "description"
                },
                "DocumentationVersion": {
                    "Ref": "version"
                },
                "RestApiId": {
                    "Ref": "RestApi"
                }
            },
            "DependsOn": "DocumentationPart"
        }
    }
}

```

#### YAML

```yaml

Parameters:
  apiName:
    Type: String
  description:
    Type: String
  version:
    Type: String
  type:
    Type: String
  property:
    Type: String
Resources:
  RestApi:
    Type: AWS::ApiGateway::RestApi
    Properties:
      Name: !Ref apiName
  DocumentationPart:
    Type: AWS::ApiGateway::DocumentationPart
    Properties:
      Location:
        Type: !Ref type
      RestApiId: !Ref RestApi
      Properties: !Ref property
  DocumentationVersion:
    Type: AWS::ApiGateway::DocumentationVersion
    Properties:
      Description: !Ref description
      DocumentationVersion: !Ref version
      RestApiId: !Ref RestApi
    DependsOn: DocumentationPart

```

## See also

- [documentationpart:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateDocumentationPart.html) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Location
