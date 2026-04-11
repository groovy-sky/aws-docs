This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::SchedulingPolicy QuotaSharePolicy

The quota share scheduling policy details for a job queue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleResourceAssignmentStrategy" : String
}

```

### YAML

```yaml

  IdleResourceAssignmentStrategy: String

```

## Properties

`IdleResourceAssignmentStrategy`

The strategy that determines how idle resources are assigned to quota shares
that are borrowing capacity. Currently, only `FIFO` is supported.

_Required_: No

_Type_: String

_Allowed values_: `FIFO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FairsharePolicy

ShareAttributes

All content copied from https://docs.aws.amazon.com/.
