This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link ModuleParameters

Describes the parameters of a module.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NoBid" : NoBidModuleParameters,
  "OpenRtbAttribute" : OpenRtbAttributeModuleParameters
}

```

### YAML

```yaml

  NoBid:
    NoBidModuleParameters
  OpenRtbAttribute:
    OpenRtbAttributeModuleParameters

```

## Properties

`NoBid`

Describes the parameters of a no bid module.

_Required_: No

_Type_: [NoBidModuleParameters](aws-properties-rtbfabric-link-nobidmoduleparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenRtbAttribute`

Describes the parameters of an open RTB attribute module.

_Required_: No

_Type_: [OpenRtbAttributeModuleParameters](aws-properties-rtbfabric-link-openrtbattributemoduleparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModuleConfiguration

NoBidAction

All content copied from https://docs.aws.amazon.com/.
