This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DocumentationVersion

The `AWS::ApiGateway::DocumentationVersion` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](../../../apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.md) in the _API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::DocumentationVersion",
  "Properties" : {
      "Description" : String,
      "DocumentationVersion" : String,
      "RestApiId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::DocumentationVersion
Properties:
  Description: String
  DocumentationVersion: String
  RestApiId: String

```

## Properties

`Description`

A description about the new documentation snapshot.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentationVersion`

The version identifier of the to-be-updated documentation version.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Associate documentation version with stage

The following example associates a documentation version with an API stage.

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
        "property": {
            "Type": "String"
        },
        "stageName": {
            "Type": "String"
        },
        "type": {
            "Type": "String"
        },
        "version": {
            "Type": "String"
        }
    },
    "Resources": {
        "Deployment": {
            "Type": "AWS::ApiGateway::Deployment",
            "Properties": {
                "RestApiId": {
                    "Ref": "RestApi"
                }
            },
            "DependsOn": [
                "Method"
            ]
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
                "Property": {
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
        },
        "Method": {
            "Type": "AWS::ApiGateway::Method",
            "Properties": {
                "AuthorizationType": "NONE",
                "HttpMethod": "POST",
                "ResourceId": {
                    "Fn::GetAtt": [
                        "RestApi",
                        "RootResourceId"
                    ]
                },
                "RestApiId": {
                    "Ref": "RestApi"
                },
                "Integration": {
                    "Type": "MOCK"
                }
            }
        },
        "RestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Name": {
                    "Ref": "apiName"
                }
            }
        },
        "Stage": {
            "Type": "AWS::ApiGateway::Stage",
            "Properties": {
                "DeploymentId": {
                    "Ref": "Deployment"
                },
                "DocumentationVersion": {
                    "Ref": "version"
                },
                "RestApiId": {
                    "Ref": "RestApi"
                },
                "StageName": {
                    "Ref": "stageName"
                }
            },
            "DependsOn": "DocumentationVersion"
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
  property:
    Type: String
  stageName:
    Type: String
  type:
    Type: String
  version:
    Type: String
Resources:
  Deployment:
    Type: AWS::ApiGateway::Deployment
    Properties:
      RestApiId: !Ref RestApi
    DependsOn:
      - Method
  DocumentationPart:
    Type: AWS::ApiGateway::DocumentationPart
    Properties:
      Location:
        Type: !Ref type
      RestApiId: !Ref RestApi
      Property: !Ref property
  DocumentationVersion:
    Type: AWS::ApiGateway::DocumentationVersion
    Properties:
      Description: !Ref description
      DocumentationVersion: !Ref version
      RestApiId: !Ref RestApi
    DependsOn: DocumentationPart
  Method:
    Type: AWS::ApiGateway::Method
    Properties:
      AuthorizationType: NONE
      HttpMethod: POST
      ResourceId: !GetAtt
        - RestApi
        - RootResourceId
      RestApiId: !Ref RestApi
      Integration:
        Type: MOCK
  RestApi:
    Type: AWS::ApiGateway::RestApi
    Properties:
      Name: !Ref apiName
  Stage:
    Type: AWS::ApiGateway::Stage
    Properties:
      DeploymentId: !Ref Deployment
      DocumentationVersion: !Ref version
      RestApiId: !Ref RestApi
      StageName: !Ref stageName
    DependsOn: DocumentationVersion
```

## See also

- [documentationpart:create](../../../apigateway/latest/api/api-createdocumentationpart.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Location

AWS::ApiGateway::DomainName

All content copied from https://docs.aws.amazon.com/.
