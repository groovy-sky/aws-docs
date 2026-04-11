This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Server WorkflowDetail

Specifies the workflow ID for the workflow to assign and the execution role that's
used for executing the workflow.

In addition to a workflow to execute when a file is uploaded completely,
`WorkflowDetails` can also contain a workflow ID (and execution role) for
a workflow to execute on partial upload. A partial upload occurs when a file is open
when the session disconnects.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionRole" : String,
  "WorkflowId" : String
}

```

### YAML

```yaml

  ExecutionRole: String
  WorkflowId: String

```

## Properties

`ExecutionRole`

Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer
can assume, so that all workflow steps can operate on the required resources

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*role/\S+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowId`

A unique identifier for the workflow.

_Required_: Yes

_Type_: String

_Pattern_: `^w-([a-z0-9]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WorkflowDetails

All content copied from https://docs.aws.amazon.com/.
