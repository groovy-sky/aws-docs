---
title: "AWS::CloudWatch::Alarm EvaluationCriteria"
---

This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm EvaluationCriteria

The evaluation criteria for an alarm. This is a union type that currently
supports `PromQLCriteria`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PromQLCriteria" : AlarmPromQLCriteria
}

```

### YAML

```yaml

  PromQLCriteria:
    AlarmPromQLCriteria

```

## Properties

`PromQLCriteria`

The PromQL criteria for the alarm evaluation.

_Required_: No

_Type_: [AlarmPromQLCriteria](aws-properties-cloudwatch-alarm-alarmpromqlcriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dimension

Metric

All content copied from https://docs.aws.amazon.com/.
