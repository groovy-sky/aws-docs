This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate SystemEndpointAttributes

The system endpoint attributes that are used with the message template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : String
}

```

### YAML

```yaml

  Address: String

```

## Properties

`Address`

The customer's phone number if used with `customerEndpoint`, or the number the customer dialed to
call your contact center if used with `systemEndpoint`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SystemAttributes

Tag

All content copied from https://docs.aws.amazon.com/.
