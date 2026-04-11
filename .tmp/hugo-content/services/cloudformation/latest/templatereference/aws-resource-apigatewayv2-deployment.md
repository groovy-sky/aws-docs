This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Deployment

The `AWS::ApiGatewayV2::Deployment` resource creates a deployment for
an API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::Deployment",
  "Properties" : {
      "ApiId" : String,
      "Description" : String,
      "StageName" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::Deployment
Properties:
  ApiId: String
  Description: String
  StageName: String

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description for the deployment resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageName`

The name of an existing stage to associate with the deployment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the deployment ID, such as
`123abc`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DeploymentId`

The deployment ID.

## Examples

### Deployment creation example

The following example creates a `deployment` resource for
the `MyApi` API, which has the `MyRoute` route
defined.

#### JSON

```json

{
    "Deployment": {
        "Type": "AWS::ApiGatewayV2::Deployment",
        "DependsOn": [
            "MyRoute"
        ],
        "Properties": {
            "Description": "My deployment",
            "ApiId": {
                "Ref": "MyApi"
            },
            "StageName": "Beta"
        }
    }
}
```

#### YAML

```yaml

Deployment:
  Type: 'AWS::ApiGatewayV2::Deployment'
  DependsOn:
    - MyRoute
  Properties:
    Description: My deployment
    ApiId: !Ref MyApi
    StageName: Beta
```

## See also

- [CreateDeployment](../../../apigatewayv2/latest/api-reference/apis-apiid-deployments.md#CreateDeployment) in the _Amazon API_
_Gateway Version 2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JWTConfiguration

AWS::ApiGatewayV2::DomainName

All content copied from https://docs.aws.amazon.com/.
