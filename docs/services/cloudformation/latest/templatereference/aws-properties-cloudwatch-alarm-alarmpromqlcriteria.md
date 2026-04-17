---
title: "AWS::CloudWatch::Alarm AlarmPromQLCriteria"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm AlarmPromQLCriteria

Contains the configuration that determines how a PromQL alarm evaluates its
contributors, including the query to run and the durations that define when contributors
transition between states.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PendingPeriod" : Integer,
  "Query" : String,
  "RecoveryPeriod" : Integer
}

```

### YAML

```yaml

  PendingPeriod: Integer
  Query: String
  RecoveryPeriod: Integer

```

## Properties

`PendingPeriod`

The duration, in seconds, that a contributor must be continuously breaching before
it transitions to the `ALARM` state.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Query`

The PromQL query that the alarm evaluates. The query must return a result of vector
type. Each entry in the vector result represents an alarm contributor.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecoveryPeriod`

The duration, in seconds, that a contributor must continuously not be breaching
before it transitions back to the `OK` state.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudWatch::Alarm

Dimension

All content copied from https://docs.aws.amazon.com/.
