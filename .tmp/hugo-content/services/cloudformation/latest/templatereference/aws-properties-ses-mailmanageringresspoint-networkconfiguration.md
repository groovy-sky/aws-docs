This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerIngressPoint NetworkConfiguration

The network type (IPv4-only, Dual-Stack, PrivateLink) of the ingress endpoint resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrivateNetworkConfiguration" : PrivateNetworkConfiguration,
  "PublicNetworkConfiguration" : PublicNetworkConfiguration
}

```

### YAML

```yaml

  PrivateNetworkConfiguration:
    PrivateNetworkConfiguration
  PublicNetworkConfiguration:
    PublicNetworkConfiguration

```

## Properties

`PrivateNetworkConfiguration`

Specifies the network configuration for the private ingress point.

_Required_: No

_Type_: [PrivateNetworkConfiguration](aws-properties-ses-mailmanageringresspoint-privatenetworkconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicNetworkConfiguration`

Specifies the network configuration for the public ingress point.

_Required_: No

_Type_: [PublicNetworkConfiguration](aws-properties-ses-mailmanageringresspoint-publicnetworkconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressPointConfiguration

PrivateNetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
