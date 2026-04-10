This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule WorkflowUpdate

Used to update information about the investigation into the finding.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Status" : String
}

```

### YAML

```yaml

  Status: String

```

## Properties

`Status`

The status of the investigation into the finding. The workflow status is specific to an individual finding. It does not affect the generation of new findings. For example, setting the workflow status to `SUPPRESSED` or `RESOLVED` does not prevent a new finding for the same issue.

The allowed values are the following.

- `NEW` \- The initial state of a finding, before it is reviewed.

Security Hub CSPM also resets `WorkFlowStatus` from `NOTIFIED` or
`RESOLVED` to `NEW` in the following cases:

- The record state changes from `ARCHIVED` to
`ACTIVE`.

- The compliance status changes from `PASSED` to either
`WARNING`, `FAILED`, or
`NOT_AVAILABLE`.

- `NOTIFIED` \- Indicates that you notified the resource owner about the
security issue. Used when the initial reviewer is not the resource owner, and needs
intervention from the resource owner.

- `RESOLVED` \- The finding was reviewed and remediated and is now
considered resolved.

- `SUPPRESSED` \- Indicates that you reviewed the finding and don't believe that any action is needed. The finding is no longer updated.

_Required_: Yes

_Type_: String

_Allowed values_: `NEW | NOTIFIED | RESOLVED | SUPPRESSED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringFilter

AWS::SecurityHub::AutomationRuleV2

All content copied from https://docs.aws.amazon.com/.
