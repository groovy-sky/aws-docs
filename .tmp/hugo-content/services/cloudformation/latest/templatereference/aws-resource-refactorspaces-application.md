This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Application

###### Note

AWS Migration Hub is no longer open to new customers as of November 7, 2025. For capabilities similar to AWS Migration Hub, explore [AWS Migration Hub](https://aws.amazon.com/transform).

Creates an AWS Migration Hub Refactor Spaces application. The account that owns the environment also owns the
applications created inside the environment, regardless of the account that creates the
application. Refactor Spaces provisions an Amazon API Gateway, API Gateway VPC link, and
Network Load Balancer for the application proxy inside your account.

In environments created with a [CreateEnvironment:NetworkFabricType](../../../../reference/migrationhub-refactor-spaces/latest/apireference/api-createenvironment.md#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType) of `NONE` you need to configure
[VPC to VPC connectivity](../../../whitepapers/latest/aws-vpc-connectivity-options/amazon-vpc-to-amazon-vpc-connectivity-options.md) between your service VPC and the application proxy VPC to
route traffic through the application proxy to a service with a private URL endpoint. For more
information, see [Create an application](../../../migrationhub-refactor-spaces/latest/userguide/getting-started-create-application.md) in the _Refactor Spaces User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RefactorSpaces::Application",
  "Properties" : {
      "ApiGatewayProxy" : ApiGatewayProxyInput,
      "EnvironmentIdentifier" : String,
      "Name" : String,
      "ProxyType" : String,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::RefactorSpaces::Application
Properties:
  ApiGatewayProxy:
    ApiGatewayProxyInput
  EnvironmentIdentifier: String
  Name: String
  ProxyType: String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`ApiGatewayProxy`

The endpoint URL of the Amazon API Gateway proxy.

_Required_: No

_Type_: [ApiGatewayProxyInput](aws-properties-refactorspaces-application-apigatewayproxyinput.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentIdentifier`

The unique identifier of the environment.

_Required_: Yes

_Type_: String

_Pattern_: `^env-([0-9A-Za-z]{10}$)`

_Minimum_: `14`

_Maximum_: `14`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!app-)[a-zA-Z0-9]+[a-zA-Z0-9-_ ]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProxyType`

The proxy type of the proxy created within the application.

_Required_: Yes

_Type_: String

_Allowed values_: `API_GATEWAY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the application.

_Required_: No

_Type_: Array of [Tag](aws-properties-refactorspaces-application-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the virtual private cloud (VPC).

_Required_: Yes

_Type_: String

_Pattern_: `^vpc-[-a-f0-9]{8}([-a-f0-9]{9})?$`

_Minimum_: `12`

_Maximum_: `21`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a composite ID following this format:
`<EnvironmentId>|<ApplicationId>`, for example,
`env-1234654123|app-1234654123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApiGatewayId`

The resource ID of the API Gateway for the proxy.

`ApplicationIdentifier`

The unique identifier of the application.

`Arn`

The Amazon Resource Name (ARN) of the application.

`NlbArn`

The Amazon Resource Name (ARN) of the Network Load Balancer.

`NlbName`

The name of the Network Load Balancer configured by the API Gateway proxy.

`ProxyUrl`

The endpoint URL of the Amazon API Gateway proxy.

`StageName`

The name of the API Gateway stage. The name defaults to `prod`.

`VpcLinkId`

The `VpcLink` ID of the API Gateway proxy.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Migration Hub Refactor Spaces

ApiGatewayProxyInput

All content copied from https://docs.aws.amazon.com/.
