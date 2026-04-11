This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashUtcTiming

Determines the type of UTC timing included in the DASH Media Presentation Description (MPD).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimingMode" : String,
  "TimingSource" : String
}

```

### YAML

```yaml

  TimingMode: String
  TimingSource: String

```

## Properties

`TimingMode`

The UTC timing mode.

_Required_: No

_Type_: String

_Allowed values_: `HTTP_HEAD | HTTP_ISO | HTTP_XSDATE | UTC_DIRECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimingSource`

The the method that the player uses to synchronize to coordinated universal time (UTC) wall clock time.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashTtmlConfiguration

Encryption

All content copied from https://docs.aws.amazon.com/.
