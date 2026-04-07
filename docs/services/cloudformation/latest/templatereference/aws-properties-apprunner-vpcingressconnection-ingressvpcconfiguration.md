This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::VpcIngressConnection IngressVpcConfiguration

Specifications for the customer’s VPC and related PrivateLink VPC endpoint that are used to associate with the VPC Ingress Connection resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcEndpointId" : String,
  "VpcId" : String
}

```

### YAML

```yaml

  VpcEndpointId: String
  VpcId: String

```

## Properties

`VpcEndpointId`

The ID of the VPC endpoint that your App Runner service connects to.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC that is used for the VPC endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppRunner::VpcIngressConnection

Tag
