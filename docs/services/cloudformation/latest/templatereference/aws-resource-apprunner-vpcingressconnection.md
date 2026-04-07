This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::VpcIngressConnection

The `AWS::AppRunner::VpcIngressConnection` resource is an AWS App Runner resource type that specifies an App Runner VPC
Ingress Connection.

App Runner requires this resource when you want to associate your App Runner service to an Amazon VPC endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppRunner::VpcIngressConnection",
  "Properties" : {
      "IngressVpcConfiguration" : IngressVpcConfiguration,
      "ServiceArn" : String,
      "Tags" : [ Tag, ... ],
      "VpcIngressConnectionName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppRunner::VpcIngressConnection
Properties:
  IngressVpcConfiguration:
    IngressVpcConfiguration
  ServiceArn: String
  Tags:
    - Tag
  VpcIngressConnectionName: String

```

## Properties

`IngressVpcConfiguration`

Specifications for the customer’s Amazon VPC and the related AWS PrivateLink VPC endpoint that are used to create the VPC Ingress Connection
resource.

_Required_: Yes

_Type_: [IngressVpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apprunner-vpcingressconnection-ingressvpcconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceArn`

The Amazon Resource Name (ARN) for this App Runner service that is used to create the VPC Ingress Connection resource.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[0-9]{12}:(\w|/|-){1,1011}`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional list of metadata items that you can associate with the VPC Ingress Connection resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apprunner-vpcingressconnection-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcIngressConnectionName`

The customer-provided VPC Ingress Connection name.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9][A-Za-z0-9\-_]{3,39}`

_Minimum_: `4`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainName`

The domain name associated with the VPC Ingress Connection resource.

`Status`

The current status of the VPC Ingress Connection.
The VPC Ingress Connection displays one of the following statuses: `AVAILABLE`, `PENDING_CREATION`, `PENDING_UPDATE`, `PENDING_DELETION`, `FAILED_CREATION`, `FAILED_UPDATE`, `FAILED_DELETION`, and `DELETED`..

`VpcIngressConnectionArn`

The Amazon Resource Name (ARN) of the VPC Ingress Connection.

## Examples

### VPC Ingress Connection

This example illustrates creating a VPC Ingress Connection.

#### JSON

```json

{
  "Type": "AWS::AppRunner::VpcIngressConnection",
  "Properties": {
    "IngressVpcConfiguration": {
      "VpcEndpointId": "vpc-1a2b3c4d",
      "VpcId": "vpc-4a6b7c8d"
    },
    "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/my-service",
    "VpcIngressConnectionName": "my-vpc-ingress-connection"
  }
}
```

#### YAML

```yaml

Type: AWS::AppRunner::VpcIngressConnection
Properties:
  IngressVpcConfiguration:
    VpcEndpointId: vpce-1a2b3c4d
    VpcId: vpc-4a6b7c8d
  ServiceArn: arn:aws:apprunner:us-east-1:123456789012:service/my-service
  VpcIngressConnectionName: my-vpc-ingress-connection
```

## See also

- [Configure network settings for incoming traffic](https://docs.aws.amazon.com/apprunner/latest/dg/network-pl.html) in the
_AWS App Runner Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

IngressVpcConfiguration
