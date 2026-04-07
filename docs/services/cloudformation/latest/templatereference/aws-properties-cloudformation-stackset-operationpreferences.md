This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet OperationPreferences

The user-specified preferences for how CloudFormation performs a StackSet
operation. For more information on maximum concurrent accounts and failure tolerance,
see [StackSet operation options](../userguide/stacksets-concepts.md#stackset-ops-options) in the _CloudFormation User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConcurrencyMode" : String,
  "FailureToleranceCount" : Integer,
  "FailureTolerancePercentage" : Integer,
  "MaxConcurrentCount" : Integer,
  "MaxConcurrentPercentage" : Integer,
  "RegionConcurrencyType" : String,
  "RegionOrder" : [ String, ... ]
}

```

### YAML

```yaml

  ConcurrencyMode: String
  FailureToleranceCount: Integer
  FailureTolerancePercentage: Integer
  MaxConcurrentCount: Integer
  MaxConcurrentPercentage: Integer
  RegionConcurrencyType: String
  RegionOrder:
    - String

```

## Properties

`ConcurrencyMode`

Specifies how the concurrency level behaves during the operation execution.

- `STRICT_FAILURE_TOLERANCE`: This option dynamically lowers the concurrency
level to ensure the number of failed accounts never exceeds the value of
`FailureToleranceCount` +1\. The initial actual concurrency is set to the lower of
either the value of the `MaxConcurrentCount`, or the value of
`FailureToleranceCount` +1\. The actual concurrency is then reduced proportionally
by the number of failures. This is the default behavior.

If failure tolerance or Maximum concurrent accounts are set to percentages, the behavior
is similar.

- `SOFT_FAILURE_TOLERANCE`: This option decouples
`FailureToleranceCount` from the actual concurrency. This allows StackSet
operations to run at the concurrency level set by the `MaxConcurrentCount` value, or
`MaxConcurrentPercentage`, regardless of the number of failures.

_Required_: No

_Type_: String

_Allowed values_: `STRICT_FAILURE_TOLERANCE | SOFT_FAILURE_TOLERANCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureToleranceCount`

The number of accounts per Region this operation can fail in before CloudFormation stops the operation in that Region. If the operation is stopped in a
Region, CloudFormation doesn't attempt the operation in any subsequent
Regions.

Conditional: You must specify either `FailureToleranceCount` or
`FailureTolerancePercentage` (but not both).

_Required_: Conditional

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureTolerancePercentage`

The percentage of accounts per Region this stack operation can fail in before CloudFormation stops the operation in that Region. If the operation is stopped in a
Region, CloudFormation doesn't attempt the operation in any subsequent
Regions.

When calculating the number of accounts based on the specified percentage, CloudFormation rounds _down_ to the next whole number.

Conditional: You must specify either `FailureToleranceCount` or
`FailureTolerancePercentage`, but not both.

_Required_: Conditional

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentCount`

The maximum number of accounts in which to perform this operation at one time. This is
dependent on the value of `FailureToleranceCount`.
`MaxConcurrentCount` is at most one more than the
`FailureToleranceCount`.

Note that this setting lets you specify the _maximum_ for
operations. For large deployments, under certain circumstances the actual number of
accounts acted upon concurrently may be lower due to service throttling.

Conditional: You must specify either `MaxConcurrentCount` or
`MaxConcurrentPercentage`, but not both.

_Required_: Conditional

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentPercentage`

The maximum percentage of accounts in which to perform this operation at one
time.

When calculating the number of accounts based on the specified percentage, CloudFormation rounds down to the next whole number. This is true except in cases
where rounding down would result is zero. In this case, CloudFormation sets the
number as one instead.

Note that this setting lets you specify the _maximum_ for
operations. For large deployments, under certain circumstances the actual number of
accounts acted upon concurrently may be lower due to service throttling.

Conditional: You must specify either `MaxConcurrentCount` or
`MaxConcurrentPercentage`, but not both.

_Required_: Conditional

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionConcurrencyType`

The concurrency type of deploying StackSets operations in Regions, could be in parallel or
one Region at a time.

_Required_: No

_Type_: String

_Allowed values_: `SEQUENTIAL | PARALLEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionOrder`

The order of the Regions where you want to perform the stack operation.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ManagedExecution

Parameter
