This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rbin::Rule RetentionPeriod

Information about the retention period for which the retention rule is to
retain resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RetentionPeriodUnit" : String,
  "RetentionPeriodValue" : Integer
}

```

### YAML

```yaml

  RetentionPeriodUnit: String
  RetentionPeriodValue: Integer

```

## Properties

`RetentionPeriodUnit`

The unit of time in which the retention period is measured. Currently, only `DAYS`
is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionPeriodValue`

The period value for which the retention rule is to retain resources, measured in days.
The supported retention periods are:

- EBS volumes: 1 - 7 days

- EBS snapshots and EBS-backed AMIs: 1 - 365 days

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `3650`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

Tag

All content copied from https://docs.aws.amazon.com/.
