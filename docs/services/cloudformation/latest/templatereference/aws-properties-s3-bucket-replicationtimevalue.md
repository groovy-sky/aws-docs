This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ReplicationTimeValue

A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics
`EventThreshold`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Minutes" : Integer
}

```

### YAML

```yaml

  Minutes: Integer

```

## Properties

`Minutes`

Contains an integer specifying time in minutes.

Valid value: 15

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationTime

RoutingRule

All content copied from https://docs.aws.amazon.com/.
