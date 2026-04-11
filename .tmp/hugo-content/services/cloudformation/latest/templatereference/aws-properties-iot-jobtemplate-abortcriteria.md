This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate AbortCriteria

The criteria that determine when and how a job abort takes place.

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

The type of job action to take to initiate the job abort.

_Required_: Yes

_Type_: String

_Allowed values_: `CANCEL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FailureType`

The type of job execution failures that can initiate a job abort.

_Required_: Yes

_Type_: String

_Allowed values_: `FAILED | REJECTED | TIMED_OUT | ALL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinNumberOfExecutedThings`

The minimum number of things which must receive job execution notifications before
the job can be aborted.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThresholdPercentage`

The minimum percentage of job execution failures that must occur to initiate the
job abort.

AWS IoT Core supports up to two digits after the decimal (for example, 10.9 and 10.99,
but not 10.999).

_Required_: Yes

_Type_: Number

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AbortConfig

ExponentialRolloutRate

All content copied from https://docs.aws.amazon.com/.
