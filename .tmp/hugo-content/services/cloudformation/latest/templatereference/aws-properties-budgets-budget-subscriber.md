This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::Budget Subscriber

The `Subscriber` property type specifies who to notify for a Billing and Cost Management budget notification.
The subscriber consists of a subscription type, and either an Amazon SNS topic or an email address.

For example, an email subscriber would have the following parameters:

- A `subscriptionType` of `EMAIL`

- An `address` of `example@example.com`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : String,
  "SubscriptionType" : String
}

```

### YAML

```yaml

  Address: String
  SubscriptionType: String

```

## Properties

`Address`

The address that AWS sends budget notifications to, either an SNS topic
or an email.

When you create a subscriber, the value of `Address` can't contain line
breaks.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubscriptionType`

The type of notification that AWS sends to a subscriber.

_Required_: Yes

_Type_: String

_Allowed values_: `SNS | EMAIL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Subscriber](../../../../reference/aws-cost-management/latest/apireference/api-budgets-subscriber.md)
in the _AWS Cost Explorer Service Cost Management APIs_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Spend

TagValues

All content copied from https://docs.aws.amazon.com/.
