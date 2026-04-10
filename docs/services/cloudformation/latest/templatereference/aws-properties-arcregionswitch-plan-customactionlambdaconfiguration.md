This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan CustomActionLambdaConfiguration

Configuration for AWS Lambda functions that perform custom actions during a Region switch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Lambdas" : [ Lambdas, ... ],
  "RegionToRun" : String,
  "RetryIntervalMinutes" : Number,
  "TimeoutMinutes" : Number,
  "Ungraceful" : LambdaUngraceful
}

```

### YAML

```yaml

  Lambdas:
    - Lambdas
  RegionToRun: String
  RetryIntervalMinutes: Number
  TimeoutMinutes: Number
  Ungraceful:
    LambdaUngraceful

```

## Properties

`Lambdas`

The AWS Lambda functions for the execution block.

_Required_: Yes

_Type_: [Array](aws-properties-arcregionswitch-plan-lambdas.md) of [Lambdas](aws-properties-arcregionswitch-plan-lambdas.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionToRun`

The AWS Region for the function to run in. For recovery workflows use `activatingRegion` or `deactivatingRegion`. For post-recovery workflows, use `activeRegion` (the Region with customer traffic) or `inactiveRegion` (the Region with no customer traffic).

_Required_: Yes

_Type_: String

_Allowed values_: `activatingRegion | deactivatingRegion | activeRegion | inactiveRegion`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryIntervalMinutes`

The retry interval specified.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The timeout value specified for the configuration.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ungraceful`

The settings for ungraceful execution.

_Required_: No

_Type_: [LambdaUngraceful](aws-properties-arcregionswitch-plan-lambdaungraceful.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociatedAlarm

DocumentDbConfiguration

All content copied from https://docs.aws.amazon.com/.
