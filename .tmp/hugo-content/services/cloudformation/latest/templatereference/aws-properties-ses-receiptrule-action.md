This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule Action

An action that Amazon SES can take when it receives an email on behalf of one or more email
addresses or domains that you own. An instance of this data type can represent only one
action.

For information about setting up receipt rules, see the [Amazon SES\
Developer Guide](../../../ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddHeaderAction" : AddHeaderAction,
  "BounceAction" : BounceAction,
  "ConnectAction" : ConnectAction,
  "LambdaAction" : LambdaAction,
  "S3Action" : S3Action,
  "SNSAction" : SNSAction,
  "StopAction" : StopAction,
  "WorkmailAction" : WorkmailAction
}

```

### YAML

```yaml

  AddHeaderAction:
    AddHeaderAction
  BounceAction:
    BounceAction
  ConnectAction:
    ConnectAction
  LambdaAction:
    LambdaAction
  S3Action:
    S3Action
  SNSAction:
    SNSAction
  StopAction:
    StopAction
  WorkmailAction:
    WorkmailAction

```

## Properties

`AddHeaderAction`

Adds a header to the received email.

_Required_: No

_Type_: [AddHeaderAction](aws-properties-ses-receiptrule-addheaderaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BounceAction`

Rejects the received email by returning a bounce response to the sender and,
optionally, publishes a notification to Amazon Simple Notification Service (Amazon SNS).

_Required_: No

_Type_: [BounceAction](aws-properties-ses-receiptrule-bounceaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectAction`

The action that informs a traffic policy resource to either allow or block the email
if it matches a condition in the policy statement.

_Required_: No

_Type_: [ConnectAction](aws-properties-ses-receiptrule-connectaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaAction`

Calls an AWS Lambda function, and optionally, publishes a notification to Amazon SNS.

_Required_: No

_Type_: [LambdaAction](aws-properties-ses-receiptrule-lambdaaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Action`

Saves the received message to an Amazon Simple Storage Service (Amazon S3) bucket and, optionally, publishes a
notification to Amazon SNS.

_Required_: No

_Type_: [S3Action](aws-properties-ses-receiptrule-s3action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SNSAction`

Publishes the email content within a notification to Amazon SNS.

_Required_: No

_Type_: [SNSAction](aws-properties-ses-receiptrule-snsaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StopAction`

Terminates the evaluation of the receipt rule set and optionally publishes a
notification to Amazon SNS.

_Required_: No

_Type_: [StopAction](aws-properties-ses-receiptrule-stopaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkmailAction`

Calls Amazon WorkMail and, optionally, publishes a notification to Amazon SNS.

_Required_: No

_Type_: [WorkmailAction](aws-properties-ses-receiptrule-workmailaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::ReceiptRule

AddHeaderAction

All content copied from https://docs.aws.amazon.com/.
