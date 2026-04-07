This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Resource

The `AWS::ApiGateway::Resource` resource creates a resource in an API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::Resource",
  "Properties" : {
      "ParentId" : String,
      "PathPart" : String,
      "RestApiId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::Resource
Properties:
  ParentId: String
  PathPart: String
  RestApiId: String

```

## Properties

`ParentId`

The parent resource's identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PathPart`

The last path segment for this resource.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ID, such as `abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceId`

The ID for the resource. For example: `abc123`.

## Examples

### Create resource

The following example creates a `stack` resource for the `MyApi` API.

#### JSON

```json

{
    "Stack": {
        "Type": "AWS::ApiGateway::Resource",
        "Properties": {
            "RestApiId": {
                "Ref": "MyApi"
            },
            "ParentId": {
                "Fn::GetAtt": [
                    "MyApi",
                    "RootResourceId"
                ]
            },
            "PathPart": "stack"
        }
    }
}
```

#### YAML

```yaml

Stack:
  Type: 'AWS::ApiGateway::Resource'
  Properties:
    RestApiId: !Ref MyApi
    ParentId: !GetAtt
      - MyApi
      - RootResourceId
    PathPart: stack

```

## See also

- [resource:create](../../../apigateway/latest/api/api-createresource.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::RequestValidator

AWS::ApiGateway::RestApi
