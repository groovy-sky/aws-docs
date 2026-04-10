This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration ModelTrainingPaymentConfig

An object representing the collaboration member's model training payment responsibilities set by the
collaboration creator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsResponsible" : Boolean
}

```

### YAML

```yaml

  IsResponsible: Boolean

```

## Properties

`IsResponsible`

Indicates whether the collaboration creator has configured the collaboration member to
pay for model training costs ( `TRUE`) or has not configured the collaboration
member to pay for model training costs ( `FALSE`).

Exactly one member can be configured to pay for model training costs. An error is
returned if the collaboration creator sets a `TRUE` value for more than one
member in the collaboration.

If the collaboration creator hasn't specified anyone as the member paying for model training costs, then the member who can query is the default payer. An error is returned if
the collaboration creator sets a `FALSE` value for the member who can
query.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelInferencePaymentConfig

PaymentConfiguration

All content copied from https://docs.aws.amazon.com/.
