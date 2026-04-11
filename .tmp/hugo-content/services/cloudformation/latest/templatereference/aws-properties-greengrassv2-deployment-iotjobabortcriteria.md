This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment IoTJobAbortCriteria

Contains criteria that define when and how to cancel a job.

The deployment stops if the following conditions are true:

1. The number of things that receive the deployment exceeds the
    `minNumberOfExecutedThings`.

2. The percentage of failures with type `failureType` exceeds the
    `thresholdPercentage`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "FailureType" : String,
  "MinNumberOfExecutedThings" : Integer,
  "ThresholdPercentage" : Number
}

```

### YAML

```yaml

  Action: String
  FailureType: String
  MinNumberOfExecutedThings: Integer
  ThresholdPercentage: Number

```

## Properties

`Action`

The action to perform when the criteria are met.

_Required_: Yes

_Type_: String

_Allowed values_: `CANCEL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FailureType`

The type of job deployment failure that can cancel a job.

_Required_: Yes

_Type_: String

_Allowed values_: `FAILED | REJECTED | TIMED_OUT | ALL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinNumberOfExecutedThings`

The minimum number of things that receive the configuration before the job can
cancel.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThresholdPercentage`

The minimum percentage of `failureType` failures that occur before the job can
cancel.

This parameter supports up to two digits after the decimal (for example, you can specify
`10.9` or `10.99`, but not `10.999`).

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IoTJobAbortConfig

IoTJobExecutionsRolloutConfig

All content copied from https://docs.aws.amazon.com/.
