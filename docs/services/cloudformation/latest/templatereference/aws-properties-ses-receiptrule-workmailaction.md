This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule WorkmailAction

When included in a receipt rule, this action calls Amazon WorkMail and, optionally,
publishes a notification to Amazon Simple Notification Service (Amazon SNS). It usually isn't necessary to set this up
manually, because Amazon WorkMail adds the rule automatically during its setup
procedure.

For information using a receipt rule to call Amazon WorkMail, see the [Amazon SES\
Developer Guide](../../../ses/latest/dg/receiving-email-action-workmail.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OrganizationArn" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  OrganizationArn: String
  TopicArn: String

```

## Properties

`OrganizationArn`

The Amazon Resource Name (ARN) of the Amazon WorkMail organization. Amazon WorkMail
ARNs use the following format:

`arn:aws:workmail:<region>:<awsAccountId>:organization/<workmailOrganizationId>`

You can find the ID of your organization by using the [ListOrganizations](../../../../reference/workmail/latest/apireference/api-listorganizations.md) operation in Amazon WorkMail. Amazon WorkMail
organization IDs begin with " `m-`", followed by a string of alphanumeric
characters.

For information about Amazon WorkMail organizations, see the [Amazon WorkMail Administrator Guide](../../../workmail/latest/adminguide/organizations-overview.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS topic to notify when the WorkMail action
is called. You can find the ARN of a topic by using the [ListTopics](../../../sns/latest/api/api-listtopics.md) operation in
Amazon SNS.

For more information about Amazon SNS topics, see the [Amazon SNS Developer Guide](../../../sns/latest/dg/createtopic.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopAction

AWS::SES::ReceiptRuleSet

All content copied from https://docs.aws.amazon.com/.
