This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy LastLaunched

Defines criteria to exclude AMIs from lifecycle actions based on the last
time they were used to launch an instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Unit: String
  Value: Integer

```

## Properties

`Unit`

Defines the unit of time that the lifecycle policy uses to calculate elapsed time
since the last instance launched from the AMI. For example: days, weeks, months, or years.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The integer number of units for the time period. For example `6` (months).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `365`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncludeResources

PolicyDetail

All content copied from https://docs.aws.amazon.com/.
