This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Rotation CoverageTime

Information about when an on-call shift begins and ends.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "StartTime" : String
}

```

### YAML

```yaml

  EndTime: String
  StartTime: String

```

## Properties

`EndTime`

Information about when an on-call rotation shift ends.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

Information about when an on-call rotation shift begins.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSMContacts::Rotation

MonthlySetting

All content copied from https://docs.aws.amazon.com/.
