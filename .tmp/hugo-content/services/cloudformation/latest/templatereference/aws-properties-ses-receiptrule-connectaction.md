This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule ConnectAction

When included in a receipt rule, this action parses the received message and
starts an email contact in Amazon Connect on your behalf.

###### Note

When you receive emails, the maximum email size (including headers) is 40 MB.
Additionally, emails may only have up to 10 attachments.
Emails larger than 40 MB or with more than 10 attachments will be bounced.

We recommend that you configure this action via Amazon Connect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IAMRoleARN" : String,
  "InstanceARN" : String
}

```

### YAML

```yaml

  IAMRoleARN: String
  InstanceARN: String

```

## Properties

`IAMRoleARN`

The Amazon Resource Name (ARN) of the IAM role to be used by Amazon Simple Email Service while starting email contacts
to the Amazon Connect instance. This role should have permission to invoke `connect:StartEmailContact`
for the given Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[\w-]+:iam::[0-9]+:role/[\w-]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceARN`

The Amazon Resource Name (ARN) for the Amazon Connect instance that Amazon SES integrates with for starting
email contacts.

For more information about Amazon Connect instances, see the [Amazon Connect Administrator Guide](../../../connect/latest/adminguide/amazon-connect-instances.md)

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov):connect:[a-z]{2}-[a-z]+-[0-9-]{1}:[0-9]{1,20}:instance/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BounceAction

LambdaAction

All content copied from https://docs.aws.amazon.com/.
