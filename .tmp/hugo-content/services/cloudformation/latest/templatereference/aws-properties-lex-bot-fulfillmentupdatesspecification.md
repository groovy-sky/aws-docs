This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot FulfillmentUpdatesSpecification

Provides information for updating the user on the progress of
fulfilling an intent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Active" : Boolean,
  "StartResponse" : FulfillmentStartResponseSpecification,
  "TimeoutInSeconds" : Integer,
  "UpdateResponse" : FulfillmentUpdateResponseSpecification
}

```

### YAML

```yaml

  Active: Boolean
  StartResponse:
    FulfillmentStartResponseSpecification
  TimeoutInSeconds: Integer
  UpdateResponse:
    FulfillmentUpdateResponseSpecification

```

## Properties

`Active`

Determines whether fulfillment updates are sent to the user. When
this field is true, updates are sent.

If the `active` field is set to true, the
`startResponse`, `updateResponse`, and
`timeoutInSeconds` fields are required.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartResponse`

Provides configuration information for the message sent to users
when the fulfillment Lambda functions starts running.

_Required_: No

_Type_: [FulfillmentStartResponseSpecification](aws-properties-lex-bot-fulfillmentstartresponsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInSeconds`

The length of time that the fulfillment Lambda function should run
before it times out.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateResponse`

Provides configuration information for messages sent periodically to
the user while the fulfillment Lambda function is running.

_Required_: No

_Type_: [FulfillmentUpdateResponseSpecification](aws-properties-lex-bot-fulfillmentupdateresponsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FulfillmentUpdateResponseSpecification

GenerativeAISettings

All content copied from https://docs.aws.amazon.com/.
