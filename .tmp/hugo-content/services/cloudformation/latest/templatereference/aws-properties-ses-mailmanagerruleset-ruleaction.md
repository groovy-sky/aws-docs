This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleAction

The action for a rule to take. Only one of the contained actions can be set.

###### Important

This data type is a UNION, so only one of the following members can be specified
when used or returned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddHeader" : AddHeaderAction,
  "Archive" : ArchiveAction,
  "DeliverToMailbox" : DeliverToMailboxAction,
  "DeliverToQBusiness" : DeliverToQBusinessAction,
  "Drop" : Json,
  "PublishToSns" : SnsAction,
  "Relay" : RelayAction,
  "ReplaceRecipient" : ReplaceRecipientAction,
  "Send" : SendAction,
  "WriteToS3" : S3Action
}

```

### YAML

```yaml

  AddHeader:
    AddHeaderAction
  Archive:
    ArchiveAction
  DeliverToMailbox:
    DeliverToMailboxAction
  DeliverToQBusiness:
    DeliverToQBusinessAction
  Drop: Json
  PublishToSns:
    SnsAction
  Relay:
    RelayAction
  ReplaceRecipient:
    ReplaceRecipientAction
  Send:
    SendAction
  WriteToS3:
    S3Action

```

## Properties

`AddHeader`

This action adds a header. This can be used to add arbitrary email headers.

_Required_: No

_Type_: [AddHeaderAction](aws-properties-ses-mailmanagerruleset-addheaderaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Archive`

This action archives the email. This can be used to deliver an email to an
archive.

_Required_: No

_Type_: [ArchiveAction](aws-properties-ses-mailmanagerruleset-archiveaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliverToMailbox`

This action delivers an email to a WorkMail mailbox.

_Required_: No

_Type_: [DeliverToMailboxAction](aws-properties-ses-mailmanagerruleset-delivertomailboxaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliverToQBusiness`

This action delivers an email to an Amazon Q Business application for ingestion into
its knowledge base.

_Required_: No

_Type_: [DeliverToQBusinessAction](aws-properties-ses-mailmanagerruleset-delivertoqbusinessaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Drop`

This action terminates the evaluation of rules in the rule set.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishToSns`

This action publishes the email content to an Amazon SNS topic.

_Required_: No

_Type_: [SnsAction](aws-properties-ses-mailmanagerruleset-snsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Relay`

This action relays the email to another SMTP server.

_Required_: No

_Type_: [RelayAction](aws-properties-ses-mailmanagerruleset-relayaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceRecipient`

The action replaces certain or all recipients with a different set of
recipients.

_Required_: No

_Type_: [ReplaceRecipientAction](aws-properties-ses-mailmanagerruleset-replacerecipientaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Send`

This action sends the email to the internet.

_Required_: No

_Type_: [SendAction](aws-properties-ses-mailmanagerruleset-sendaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteToS3`

This action writes the MIME content of the email to an S3 bucket.

_Required_: No

_Type_: [S3Action](aws-properties-ses-mailmanagerruleset-s3action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

RuleBooleanExpression

All content copied from https://docs.aws.amazon.com/.
