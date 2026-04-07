This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::Firewall AvailabilityZoneMapping

Defines the mapping between an Availability Zone and a firewall endpoint for a transit gateway-attached firewall. Each mapping represents where the firewall can process traffic. You use these mappings when calling `CreateFirewall`, `AssociateAvailabilityZones`, and `DisassociateAvailabilityZones`.

To retrieve the current Availability Zone mappings for a firewall, use `DescribeFirewall`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String
}

```

### YAML

```yaml

  AvailabilityZone: String

```

## Properties

`AvailabilityZone`

The ID of the Availability Zone where the firewall endpoint is located. For example, `us-east-2a`. The Availability Zone must be in the same Region as the transit gateway.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::NetworkFirewall::Firewall

SubnetMapping
