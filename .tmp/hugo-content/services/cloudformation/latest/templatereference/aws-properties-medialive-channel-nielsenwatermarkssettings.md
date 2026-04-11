This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel NielsenWatermarksSettings

Settings to configure Nielsen Watermarks in the audio encode.

The parent of this entity is AudioWatermarkSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NielsenCbetSettings" : NielsenCBET,
  "NielsenDistributionType" : String,
  "NielsenNaesIiNwSettings" : NielsenNaesIiNw
}

```

### YAML

```yaml

  NielsenCbetSettings:
    NielsenCBET
  NielsenDistributionType: String
  NielsenNaesIiNwSettings:
    NielsenNaesIiNw

```

## Properties

`NielsenCbetSettings`

Complete these fields only if you want to insert watermarks of type Nielsen CBET

_Required_: No

_Type_: [NielsenCBET](aws-properties-medialive-channel-nielsencbet.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenDistributionType`

Choose the distribution types that you want to assign to the watermarks:
\- PROGRAM\_CONTENT
\- FINAL\_DISTRIBUTOR

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenNaesIiNwSettings`

Complete these fields only if you want to insert watermarks of type Nielsen NAES II (N2) and Nielsen NAES VI (NW).

_Required_: No

_Type_: [NielsenNaesIiNw](aws-properties-medialive-channel-nielsennaesiinw.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NielsenNaesIiNw

Output

All content copied from https://docs.aws.amazon.com/.
