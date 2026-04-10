This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction Subscriber

The subscriber to a budget notification. The subscriber consists of a subscription
type and either an Amazon SNS topic or an email address.

For example, an email subscriber has the following parameters:

- A `subscriptionType` of `EMAIL`

- An `address` of `example@example.com`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : String,
  "Type" : String
}

```

### YAML

```yaml

  Address: String
  Type: String

```

## Properties

`Address`

The address that AWS sends budget notifications to, either an SNS topic
or an email.

When you create a subscriber, the value of `Address` can't contain line
breaks.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of notification that AWS sends to a subscriber.

_Required_: Yes

_Type_: String

_Allowed values_: `SNS | EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SsmActionDefinition

Next

All content copied from https://docs.aws.amazon.com/.
