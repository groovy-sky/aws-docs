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

In environments created with a [CreateEnvironment:NetworkFabricType](https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType) of `NONE` you need to configure
[VPC to VPC connectivity](https://docs.aws.amazon.com/whitepapers/latest/aws-vpc-connectivity-options/amazon-vpc-to-amazon-vpc-connectivity-options.html) between your service VPC and the application proxy VPC to
route traffic through the application proxy to a service with a private URL endpoint. For more
information, see [Create an application](https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/userguide/getting-started-create-application.html) in the _Refactor Spaces User Guide_.

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

_Type_: [ApiGatewayProxyInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-refactorspaces-application-apigatewayproxyinput.html)

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

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-refactorspaces-application-tag.html)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Migration Hub Refactor Spaces

ApiGatewayProxyInput
