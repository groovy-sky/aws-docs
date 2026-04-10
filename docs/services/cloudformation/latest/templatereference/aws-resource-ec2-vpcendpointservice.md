This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCEndpointService

Creates a VPC endpoint service configuration to which service consumers (AWS accounts,
users, and IAM roles) can connect.

To create an endpoint service configuration, you must first create one of the following
for your service:

- A [Network Load\
Balancer](../../../elasticloadbalancing/latest/network/introduction.md). Service consumers connect to your service using an interface
endpoint.

- A [Gateway Load\
Balancer](../../../elasticloadbalancing/latest/gateway/introduction.md). Service consumers connect to your service using a Gateway Load
Balancer endpoint.

For more information, see the [AWS PrivateLink User Guide](../../../vpc/latest/privatelink.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCEndpointService",
  "Properties" : {
      "AcceptanceRequired" : Boolean,
      "ContributorInsightsEnabled" : Boolean,
      "GatewayLoadBalancerArns" : [ String, ... ],
      "NetworkLoadBalancerArns" : [ String, ... ],
      "PayerResponsibility" : String,
      "SupportedIpAddressTypes" : [ String, ... ],
      "SupportedRegions" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCEndpointService
Properties:
  AcceptanceRequired: Boolean
  ContributorInsightsEnabled: Boolean
  GatewayLoadBalancerArns:
    - String
  NetworkLoadBalancerArns:
    - String
  PayerResponsibility: String
  SupportedIpAddressTypes:
    - String
  SupportedRegions:
    - String
  Tags:
    - Tag

```

## Properties

`AcceptanceRequired`

Indicates whether requests from service consumers to create an endpoint to your service
must be accepted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContributorInsightsEnabled`

Indicates whether to enable the built-in Contributor Insights rules provided by
AWS PrivateLink.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GatewayLoadBalancerArns`

The Amazon Resource Names (ARNs) of the Gateway Load Balancers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkLoadBalancerArns`

The Amazon Resource Names (ARNs) of the Network Load Balancers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayerResponsibility`

The entity that is responsible for the endpoint costs. The default is the endpoint owner.
If you set the payer responsibility to the service owner, you cannot set it back to the
endpoint owner.

_Required_: No

_Type_: String

_Allowed values_: `ServiceOwner`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedIpAddressTypes`

The supported IP address types. The possible values are `ipv4` and `ipv6`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedRegions`

The Regions from which service consumers can access the service.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to associate with the service.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-vpcendpointservice-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the VPC endpoint service configuration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ServiceId`

The ID of the endpoint service.

## See also

- [CreateVpcEndpointServiceConfiguration](../../../../reference/awsec2/latest/apireference/apireference-query-createvpcendpointserviceconfiguration.md) in the _Amazon EC2 API_
_Reference_

- [VPC endpoint services](../../../vpc/latest/privatelink/endpoint-service.md) in _AWS_
_PrivateLink_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPCEndpointConnectionNotification

Tag

All content copied from https://docs.aws.amazon.com/.
