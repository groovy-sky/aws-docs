This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Job ConnectionsList

Specifies the connections used by a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Connections" : [ String, ... ]
}

```

### YAML

```yaml

  Connections:
    - String

```

## Properties

`Connections`

A list of connections used by the job.

_Required_: No

_Type_: Array of String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::Job

ExecutionProperty

All content copied from https://docs.aws.amazon.com/.
