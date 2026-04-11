This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel NielsenCBET

Complete these fields only if you want to insert watermarks of type Nielsen CBET

The parent of this entity is NielsenWatermarksSettings

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CbetCheckDigitString" : String,
  "CbetStepaside" : String,
  "Csid" : String
}

```

### YAML

```yaml

  CbetCheckDigitString:
    String
  CbetStepaside: String
  Csid: String

```

## Properties

`CbetCheckDigitString`

Enter the CBET check digits to use in the watermark.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CbetStepaside`

Determines the method of CBET insertion mode when prior encoding is detected on the same layer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Csid`

Enter the CBET Source ID (CSID) to use in the watermark

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkInputSettings

NielsenConfiguration

All content copied from https://docs.aws.amazon.com/.
