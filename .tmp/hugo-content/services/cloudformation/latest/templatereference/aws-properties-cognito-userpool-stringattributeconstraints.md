This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool StringAttributeConstraints

The minimum and maximum length values of an attribute that is of the string type, for
example `custom:department`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxLength" : String,
  "MinLength" : String
}

```

### YAML

```yaml

  MaxLength: String
  MinLength: String

```

## Properties

`MaxLength`

The maximum length of a string attribute value. Must be a number less than or equal to
`2^1023`, represented as a string with a length of 131072 characters or
fewer.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinLength`

The minimum length of a string attribute value.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SmsConfiguration

UserAttributeUpdateSettings

All content copied from https://docs.aws.amazon.com/.
