This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FinSpace::Environment SuperuserParameters

Configuration information for the superuser.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailAddress" : String,
  "FirstName" : String,
  "LastName" : String
}

```

### YAML

```yaml

  EmailAddress: String
  FirstName: String
  LastName: String

```

## Properties

`EmailAddress`

The email address of the superuser.

_Required_: No

_Type_: String

_Pattern_: `[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+[.]+[A-Za-z]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirstName`

The first name of the superuser.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]{1,50}$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LastName`

The last name of the superuser.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]{1,50}$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FederationParameters

Tag
