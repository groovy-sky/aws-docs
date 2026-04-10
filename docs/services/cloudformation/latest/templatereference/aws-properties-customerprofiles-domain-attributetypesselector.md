This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain AttributeTypesSelector

Configures information about the `AttributeTypesSelector` which rule-based
identity resolution uses to match profiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : [ String, ... ],
  "AttributeMatchingModel" : String,
  "EmailAddress" : [ String, ... ],
  "PhoneNumber" : [ String, ... ]
}

```

### YAML

```yaml

  Address:
    - String
  AttributeMatchingModel: String
  EmailAddress:
    - String
  PhoneNumber:
    - String

```

## Properties

`Address`

The `Address` type. You can choose from `Address`,
`BusinessAddress`, `MaillingAddress`, and
`ShippingAddress`. You only can use the `Address` type in the
`MatchingRule`. For example, if you want to match a profile based on
`BusinessAddress.City` or `MaillingAddress.City`, you can
choose the `BusinessAddress` and the `MaillingAddress` to
represent the `Address` type and specify the `Address.City` on the
matching rule.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeMatchingModel`

Configures the `AttributeMatchingModel`, you can either choose
`ONE_TO_ONE` or `MANY_TO_MANY`.

_Required_: Yes

_Type_: String

_Allowed values_: `ONE_TO_ONE | MANY_TO_MANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailAddress`

The Email type. You can choose from `EmailAddress`,
`BusinessEmailAddress` and `PersonalEmailAddress`. You only
can use the `EmailAddress` type in the `MatchingRule`. For
example, if you want to match profile based on `PersonalEmailAddress` or
`BusinessEmailAddress`, you can choose the
`PersonalEmailAddress` and the `BusinessEmailAddress` to
represent the `EmailAddress` type and only specify the
`EmailAddress` on the matching rule.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneNumber`

The `PhoneNumber` type. You can choose from `PhoneNumber`,
`HomePhoneNumber`, and `MobilePhoneNumber`. You only can use
the `PhoneNumber` type in the `MatchingRule`. For example, if you
want to match a profile based on `Phone` or `HomePhone`, you can
choose the `Phone` and the `HomePhone` to represent the
`PhoneNumber` type and only specify the `PhoneNumber` on the
matching rule.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CustomerProfiles::Domain

AutoMerging

All content copied from https://docs.aws.amazon.com/.
