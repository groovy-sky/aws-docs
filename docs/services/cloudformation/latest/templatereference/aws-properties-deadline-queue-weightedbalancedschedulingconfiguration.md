---
title: "AWS::Deadline::Queue WeightedBalancedSchedulingConfiguration"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Queue WeightedBalancedSchedulingConfiguration

The `WeightedBalancedSchedulingConfiguration` property type specifies Property description not available. for an [AWS::Deadline::Queue](aws-resource-deadline-queue.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorWeight" : Number,
  "MaxPriorityOverride" : SchedulingMaxPriorityOverride,
  "MinPriorityOverride" : SchedulingMinPriorityOverride,
  "PriorityWeight" : Number,
  "RenderingTaskBuffer" : Integer,
  "RenderingTaskWeight" : Number,
  "SubmissionTimeWeight" : Number
}

```

### YAML

```yaml

  ErrorWeight: Number
  MaxPriorityOverride:
    SchedulingMaxPriorityOverride
  MinPriorityOverride:
    SchedulingMinPriorityOverride
  PriorityWeight: Number
  RenderingTaskBuffer: Integer
  RenderingTaskWeight: Number
  SubmissionTimeWeight: Number

```

## Properties

`ErrorWeight`

Property description not available.

_Required_: No

_Type_: Number

_Minimum_: `-10000`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxPriorityOverride`

Property description not available.

_Required_: No

_Type_: [SchedulingMaxPriorityOverride](aws-properties-deadline-queue-schedulingmaxpriorityoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinPriorityOverride`

Property description not available.

_Required_: No

_Type_: [SchedulingMinPriorityOverride](aws-properties-deadline-queue-schedulingminpriorityoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PriorityWeight`

Property description not available.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenderingTaskBuffer`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RenderingTaskWeight`

Property description not available.

_Required_: No

_Type_: Number

_Minimum_: `-10000`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubmissionTimeWeight`

Property description not available.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WindowsUser

All content copied from https://docs.aws.amazon.com/.
