This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayMulticastDomain Options

The options for the transit gateway multicast domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoAcceptSharedAssociations" : String,
  "Igmpv2Support" : String,
  "StaticSourcesSupport" : String
}

```

### YAML

```yaml

  AutoAcceptSharedAssociations: String
  Igmpv2Support: String
  StaticSourcesSupport: String

```

## Properties

`AutoAcceptSharedAssociations`

Indicates whether to automatically accept cross-account subnet associations that are associated with the transit gateway multicast domain.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Igmpv2Support`

Specify whether to enable Internet Group Management Protocol (IGMP) version 2 for the transit gateway multicast domain.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticSourcesSupport`

Specify whether to enable support for statically configuring multicast group sources for a domain.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::TransitGatewayMulticastDomain

Tag
