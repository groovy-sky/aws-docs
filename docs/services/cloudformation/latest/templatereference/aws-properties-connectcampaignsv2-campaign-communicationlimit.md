This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign CommunicationLimit

Contains information about a communication limit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Frequency" : Integer,
  "MaxCountPerRecipient" : Integer,
  "Unit" : String
}

```

### YAML

```yaml

  Frequency: Integer
  MaxCountPerRecipient: Integer
  Unit: String

```

## Properties

`Frequency`

The frequency of communication limit evaluation.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCountPerRecipient`

The maximum outreaching count for each recipient.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of communication limit evaluation.

_Required_: Yes

_Type_: String

_Allowed values_: `DAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChannelSubtypeConfig

CommunicationLimits

All content copied from https://docs.aws.amazon.com/.
