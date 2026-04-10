This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel DvbTdtSettings

The DVB Time and Date Table (TDT).

The parent of this entity is M2tsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RepInterval" : Integer
}

```

### YAML

```yaml

  RepInterval: Integer

```

## Properties

`RepInterval`

The number of milliseconds between instances of this table in the output transport stream.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DvbSubSourceSettings

Eac3AtmosSettings

All content copied from https://docs.aws.amazon.com/.
