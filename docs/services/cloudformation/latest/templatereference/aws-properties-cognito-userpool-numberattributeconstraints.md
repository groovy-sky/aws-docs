This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool NumberAttributeConstraints

The minimum and maximum values of an attribute that is of the number type, for example
`custom:age`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxValue" : String,
  "MinValue" : String
}

```

### YAML

```yaml

  MaxValue: String
  MinValue: String

```

## Properties

`MaxValue`

The maximum length of a number attribute value. Must be a number less than or equal to
`2^1023`, represented as a string with a length of 131072 characters or
fewer.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinValue`

The minimum value of an attribute that is of the number data type.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaConfig

PasswordPolicy

All content copied from https://docs.aws.amazon.com/.
