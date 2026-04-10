This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel NielsenNaesIiNw

Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI
(NW).

The parent of this entity is NielsenWatermarksSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CheckDigitString" : String,
  "Sid" : Number,
  "Timezone" : String
}

```

### YAML

```yaml

  CheckDigitString:
    String
  Sid: Number
  Timezone: String

```

## Properties

`CheckDigitString`

Enter the check digit string for the watermark

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sid`

Enter the Nielsen Source ID (SID) to include in the watermark

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timezone`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NielsenConfiguration

NielsenWatermarksSettings

All content copied from https://docs.aws.amazon.com/.
