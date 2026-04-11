This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject ModalityRoutingConfiguration

This element allows you to set up where JPEG, PNG, MOV, and MP4 files
get routed to for processing. JPEG routing applies to both "JPEG" and "JPG"
file extensions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "jpeg" : String,
  "mov" : String,
  "mp4" : String,
  "png" : String
}

```

### YAML

```yaml

  jpeg: String
  mov: String
  mp4: String
  png: String

```

## Properties

`jpeg`

Sets whether JPEG files are routed to document or image processing.

_Required_: No

_Type_: String

_Allowed values_: `DOCUMENT | IMAGE | VIDEO | AUDIO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`mov`

Sets whether MOV files are routed to audio or video processing.

_Required_: No

_Type_: String

_Allowed values_: `DOCUMENT | IMAGE | VIDEO | AUDIO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`mp4`

Sets whether MP4 files are routed to audio or video processing.

_Required_: No

_Type_: String

_Allowed values_: `DOCUMENT | IMAGE | VIDEO | AUDIO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`png`

Sets whether PNG files are routed to document or image processing.

_Required_: No

_Type_: String

_Allowed values_: `DOCUMENT | IMAGE | VIDEO | AUDIO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModalityProcessingConfiguration

OverrideConfiguration

All content copied from https://docs.aws.amazon.com/.
