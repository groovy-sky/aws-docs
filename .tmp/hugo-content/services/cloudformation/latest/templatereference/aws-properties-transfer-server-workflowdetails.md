This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Server WorkflowDetails

Container for the `WorkflowDetail` data type. It is used by actions that
trigger a workflow to begin execution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnPartialUpload" : [ WorkflowDetail, ... ],
  "OnUpload" : [ WorkflowDetail, ... ]
}

```

### YAML

```yaml

  OnPartialUpload:
    - WorkflowDetail
  OnUpload:
    - WorkflowDetail

```

## Properties

`OnPartialUpload`

A trigger that starts a workflow if a file is only partially uploaded. You can attach
a workflow to a server that executes whenever there is a partial upload.

A _partial upload_ occurs when a file is open when the session
disconnects.

###### Note

`OnPartialUpload` can contain a maximum of one
`WorkflowDetail` object.

_Required_: No

_Type_: Array of [WorkflowDetail](aws-properties-transfer-server-workflowdetail.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnUpload`

A trigger that starts a workflow: the workflow begins to execute after a file is
uploaded.

To remove an associated workflow from a server, you can provide an empty
`OnUpload` object, as in the following example.

`aws transfer update-server --server-id s-01234567890abcdef --workflow-details '{"OnUpload":[]}'`

###### Note

`OnUpload` can contain a maximum of one `WorkflowDetail`
object.

_Required_: No

_Type_: Array of [WorkflowDetail](aws-properties-transfer-server-workflowdetail.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowDetail

AWS::Transfer::User

All content copied from https://docs.aws.amazon.com/.
