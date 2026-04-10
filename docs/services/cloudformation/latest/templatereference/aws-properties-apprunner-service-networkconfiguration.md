This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service NetworkConfiguration

Describes configuration settings related to network traffic of an AWS App Runner service. Consists of embedded objects for each configurable network
feature.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EgressConfiguration" : EgressConfiguration,
  "IngressConfiguration" : IngressConfiguration,
  "IpAddressType" : String
}

```

### YAML

```yaml

  EgressConfiguration:
    EgressConfiguration
  IngressConfiguration:
    IngressConfiguration
  IpAddressType: String

```

## Properties

`EgressConfiguration`

Network configuration settings for outbound message traffic.

_Required_: No

_Type_: [EgressConfiguration](aws-properties-apprunner-service-egressconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngressConfiguration`

Network configuration settings for inbound message traffic.

_Required_: No

_Type_: [IngressConfiguration](aws-properties-apprunner-service-ingressconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

App Runner provides you with the option to choose between _IPv4_ and _dual stack_ (IPv4 and IPv6)
for your incoming public network configuration. This is an optional parameter.
If you do not specify an `IpAddressType`, it defaults to select IPv4.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL_STACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyValuePair

ServiceObservabilityConfiguration

All content copied from https://docs.aws.amazon.com/.
