This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::QuickConnect PhoneNumberQuickConnectConfig

Contains information about a phone number for a quick connect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PhoneNumber" : String
}

```

### YAML

```yaml

  PhoneNumber: String

```

## Properties

`PhoneNumber`

The phone number in E.164 format.

_Required_: Yes

_Type_: String

_Pattern_: `^\+[1-9]\d{1,14}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::QuickConnect

QueueQuickConnectConfig

All content copied from https://docs.aws.amazon.com/.
