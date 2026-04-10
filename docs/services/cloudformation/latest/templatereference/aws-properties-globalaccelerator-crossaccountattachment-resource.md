This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::CrossAccountAttachment Resource

A resource is one of the following: the ARN for an AWS resource that is supported by
AWS Global Accelerator to be added as an endpoint, or a CIDR range that specifies a bring your own IP (BYOIP) address pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String,
  "EndpointId" : String,
  "Region" : String
}

```

### YAML

```yaml

  Cidr: String
  EndpointId: String
  Region: String

```

## Properties

`Cidr`

An IP address range, in CIDR format, that is specified as resource. The address must
be provisioned and advertised in AWS Global Accelerator by following the bring your own IP address (BYOIP) process
for Global Accelerator

For more information, see
[Bring your own IP addresses (BYOIP)](../../../global-accelerator/latest/dg/using-byoip.md) in
the AWS Global Accelerator Developer Guide.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointId`

The endpoint ID for the endpoint that is specified as a AWS resource.

An endpoint ID for the cross-account feature is the ARN of an AWS resource, such as a
Network Load Balancer, that Global Accelerator supports as an endpoint for an accelerator.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The AWS Region where a shared endpoint resource is located.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GlobalAccelerator::CrossAccountAttachment

Tag

All content copied from https://docs.aws.amazon.com/.
