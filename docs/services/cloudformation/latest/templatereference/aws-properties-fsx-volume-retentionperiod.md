This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume RetentionPeriod

Specifies the retention period of an FSx for ONTAP SnapLock volume. After it is set, it can't be changed.
Files can't be
deleted or modified during the retention period.

For more information, see
[Working with the retention \
period in SnapLock](../../../fsx/latest/ontapguide/snaplock-retention.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Type: String
  Value: Integer

```

## Properties

`Type`

Defines the type of time for the retention period of an FSx for ONTAP SnapLock volume.
Set it to
one of the valid types. If you set it to `INFINITE`, the files are retained forever. If you set it to
`UNSPECIFIED`, the files are retained until you set an explicit retention period.

_Required_: Yes

_Type_: String

_Allowed values_: `SECONDS | MINUTES | HOURS | DAYS | MONTHS | YEARS | INFINITE | UNSPECIFIED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Defines the amount of time for the retention period of an FSx for ONTAP SnapLock volume.
You can't set a value for `INFINITE` or `UNSPECIFIED`. For all other options, the
following ranges are valid:

- `Seconds`: 0 - 65,535

- `Minutes`: 0 - 65,535

- `Hours`: 0 - 24

- `Days`: 0 - 365

- `Months`: 0 - 12

- `Years`: 0 - 100

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginSnapshot

SnaplockConfiguration

All content copied from https://docs.aws.amazon.com/.
