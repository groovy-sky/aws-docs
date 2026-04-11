This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume AutocommitPeriod

Sets the autocommit period of files in an FSx for ONTAP SnapLock volume, which determines
how long the files must
remain unmodified before they're automatically transitioned to the write once, read many (WORM) state.

For more information, see
[Autocommit](../../../fsx/latest/ontapguide/worm-state.md#worm-state-autocommit).

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

Defines the type of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
Setting this value to `NONE` disables autocommit. The default value is `NONE`.

_Required_: Yes

_Type_: String

_Allowed values_: `MINUTES | HOURS | DAYS | MONTHS | YEARS | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Defines the amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
The following ranges are valid:

- `Minutes`: 5 - 65,535

- `Hours`: 1 - 65,535

- `Days`: 1 - 3,650

- `Months`: 1 - 120

- `Years`: 1 - 10

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AggregateConfiguration

ClientConfigurations

All content copied from https://docs.aws.amazon.com/.
