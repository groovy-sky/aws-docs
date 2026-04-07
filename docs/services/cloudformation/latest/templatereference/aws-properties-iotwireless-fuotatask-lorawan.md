This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::FuotaTask LoRaWAN

The LoRaWAN information used with a FUOTA task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RfRegion" : String,
  "StartTime" : String
}

```

### YAML

```yaml

  RfRegion: String
  StartTime: String

```

## Properties

`RfRegion`

The frequency band (RFRegion) value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

Start time of a FUOTA task.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTWireless::FuotaTask

Tag
