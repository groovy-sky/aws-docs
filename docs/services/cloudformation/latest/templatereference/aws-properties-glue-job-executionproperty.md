This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Job ExecutionProperty

An execution property of a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConcurrentRuns" : Number
}

```

### YAML

```yaml

  MaxConcurrentRuns: Number

```

## Properties

`MaxConcurrentRuns`

The maximum number of concurrent runs allowed for the job. The default is 1. An error
is returned when this threshold is reached. The maximum value you can specify is
controlled by a service limit.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectionsList

JobCommand
