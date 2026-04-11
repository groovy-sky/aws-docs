This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SNS::Topic Subscription

`Subscription` is an embedded property that describes the subscription endpoints
of an Amazon SNS topic.

###### Note

For full control over subscription behavior (for example, delivery policy, filtering,
raw message delivery, and cross-region subscriptions), use the [AWS::SNS::Subscription](../userguide/aws-resource-sns-subscription.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Endpoint" : String,
  "Protocol" : String
}

```

### YAML

```yaml

  Endpoint: String
  Protocol: String

```

## Properties

`Endpoint`

The endpoint that receives notifications from the Amazon SNS topic. The endpoint
value depends on the protocol that you specify. For more information, see the
`Endpoint` parameter of the `
                            Subscribe
                        ` action in the _Amazon SNS API Reference_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The subscription's protocol. For more information, see the `Protocol`
parameter of the `
                            Subscribe
                        ` action in the _Amazon SNS API Reference_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfig

Tag

All content copied from https://docs.aws.amazon.com/.
