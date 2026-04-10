This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule AddHeaderAction

When included in a receipt rule, this action adds a header to the received
email.

For information about adding a header using a receipt rule, see the [Amazon\
SES Developer Guide](../../../ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderName" : String,
  "HeaderValue" : String
}

```

### YAML

```yaml

  HeaderName: String
  HeaderValue: String

```

## Properties

`HeaderName`

The name of the header to add to the incoming message. The name must contain at least
one character, and can contain up to 50 characters. It consists of alphanumeric
( `a–z, A–Z, 0–9`) characters and dashes.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderValue`

The content to include in the header. This value can contain up to 2048 characters. It
can't contain newline ( `\n`) or carriage return ( `\r`)
characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

BounceAction

All content copied from https://docs.aws.amazon.com/.
