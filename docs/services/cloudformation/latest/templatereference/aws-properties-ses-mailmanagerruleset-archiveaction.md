This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet ArchiveAction

The action to archive the email by delivering the email to an Amazon SES archive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "TargetArchive" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  TargetArchive: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there
are configuration errors. For example, the specified archive has been deleted.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArchive`

The identifier of the archive to send the email to.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Analysis

DeliverToMailboxAction

All content copied from https://docs.aws.amazon.com/.
