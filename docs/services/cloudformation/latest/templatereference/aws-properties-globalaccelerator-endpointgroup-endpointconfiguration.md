This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::EndpointGroup EndpointConfiguration

A complex type for endpoints. A resource must be valid and active when you add it as an endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachmentArn" : String,
  "ClientIPPreservationEnabled" : Boolean,
  "EndpointId" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  AttachmentArn: String
  ClientIPPreservationEnabled: Boolean
  EndpointId: String
  Weight: Integer

```

## Properties

`AttachmentArn`

The Amazon Resource Name (ARN) of the cross-account attachment that specifies the endpoints (resources)
that can be added to accelerators and principals that have permission to add the endpoints.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientIPPreservationEnabled`

Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint.
The value is true or false. The default value is true for new accelerators.

If the value is set to true, the client's IP address is preserved in the `X-Forwarded-For` request header as
traffic travels to applications on the Application Load Balancer endpoint fronted by the accelerator.

For more information, see [Preserve Client IP Addresses](https://docs.aws.amazon.com/global-accelerator/latest/dg/preserve-client-ip-address.html) in the _AWS Global Accelerator Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointId`

An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon
Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address
allocation ID. For Amazon EC2 instances, this is the EC2 instance ID. A resource must be valid and active
when you add it as an endpoint.

For cross-account endpoints, this must be the ARN of the resource.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The weight associated with the endpoint. When you add weights to endpoints, you configure Global Accelerator to route traffic
based on proportions that you specify. For example, you might specify endpoint weights of 4, 5, 5, and 6 (sum=20). The
result is that 4/20 of your traffic, on average, is routed to the first endpoint, 5/20 is routed both to the second
and third endpoints, and 6/20 is routed to the last endpoint. For more information, see [Endpoint Weights](https://docs.aws.amazon.com/global-accelerator/latest/dg/about-endpoints-endpoint-weights.html) in the
_AWS Global Accelerator Developer Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GlobalAccelerator::EndpointGroup

PortOverride
