This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerIngressPoint PrivateNetworkConfiguration

Specifies the network configuration for the private ingress point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcEndpointId" : String
}

```

### YAML

```yaml

  VpcEndpointId: String

```

## Properties

`VpcEndpointId`

The identifier of the VPC endpoint to associate with this private ingress point.

_Required_: Yes

_Type_: String

_Pattern_: `^vpce-[a-zA-Z0-9]{17}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

PublicNetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
