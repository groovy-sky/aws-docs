This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::Channel MultitrackInputConfiguration

A complex type that specifies multitrack input configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "MaximumResolution" : String,
  "Policy" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  MaximumResolution: String
  Policy: String

```

## Properties

`Enabled`

Indicates whether multitrack input is enabled. Can be set to `true` only if channel type is `STANDARD`. Setting `enabled` to `true` with any other channel type will cause an exception. If `true`, then `policy`, `maximumResolution`, and `containerFormat` are required, and `containerFormat` must be set to `FRAGMENTED_MP4`. Default: `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumResolution`

Maximum resolution for multitrack input. Required if `enabled` is `true`.

_Required_: No

_Type_: String

_Allowed values_: `SD | HD | FULL_HD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policy`

Indicates whether multitrack input is allowed or required. Required if `enabled` is `true`.

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | REQUIRE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IVS::Channel

Tag

All content copied from https://docs.aws.amazon.com/.
