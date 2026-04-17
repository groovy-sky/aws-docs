---
title: "AWS::Deadline::Queue SchedulingConfiguration"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Queue SchedulingConfiguration

The `SchedulingConfiguration` property type specifies Property description not available. for an [AWS::Deadline::Queue](aws-resource-deadline-queue.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PriorityBalanced" : PriorityBalancedSchedulingConfiguration,
  "PriorityFifo" : Json,
  "WeightedBalanced" : WeightedBalancedSchedulingConfiguration
}

```

### YAML

```yaml

  PriorityBalanced:
    PriorityBalancedSchedulingConfiguration
  PriorityFifo: Json
  WeightedBalanced:
    WeightedBalancedSchedulingConfiguration

```

## Properties

`PriorityBalanced`

Property description not available.

_Required_: No

_Type_: [PriorityBalancedSchedulingConfiguration](aws-properties-deadline-queue-prioritybalancedschedulingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PriorityFifo`

Property description not available.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeightedBalanced`

Property description not available.

_Required_: No

_Type_: [WeightedBalancedSchedulingConfiguration](aws-properties-deadline-queue-weightedbalancedschedulingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PriorityBalancedSchedulingConfiguration

SchedulingMaxPriorityOverride

All content copied from https://docs.aws.amazon.com/.
