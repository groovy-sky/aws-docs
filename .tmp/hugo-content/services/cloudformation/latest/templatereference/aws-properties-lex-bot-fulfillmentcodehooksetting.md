This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot FulfillmentCodeHookSetting

Determines if a Lambda function should be invoked for a specific
intent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "FulfillmentUpdatesSpecification" : FulfillmentUpdatesSpecification,
  "IsActive" : Boolean,
  "PostFulfillmentStatusSpecification" : PostFulfillmentStatusSpecification
}

```

### YAML

```yaml

  Enabled: Boolean
  FulfillmentUpdatesSpecification:
    FulfillmentUpdatesSpecification
  IsActive: Boolean
  PostFulfillmentStatusSpecification:
    PostFulfillmentStatusSpecification

```

## Properties

`Enabled`

Indicates whether a Lambda function should be invoked to fulfill a
specific intent.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FulfillmentUpdatesSpecification`

Provides settings for update messages sent to the user for
long-running Lambda fulfillment functions. Fulfillment updates can be
used only with streaming conversations.

_Required_: No

_Type_: [FulfillmentUpdatesSpecification](aws-properties-lex-bot-fulfillmentupdatesspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsActive`

Determines whether the fulfillment code hook is used. When
`active` is false, the code hook doesn't run.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostFulfillmentStatusSpecification`

Provides settings for messages sent to the user for after the Lambda
fulfillment function completes. Post-fulfillment messages can be sent
for both streaming and non-streaming conversations.

_Required_: No

_Type_: [PostFulfillmentStatusSpecification](aws-properties-lex-bot-postfulfillmentstatusspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExternalSourceSetting

FulfillmentStartResponseSpecification

All content copied from https://docs.aws.amazon.com/.
