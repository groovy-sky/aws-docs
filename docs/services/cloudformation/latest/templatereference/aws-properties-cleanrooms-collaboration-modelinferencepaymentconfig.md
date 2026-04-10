This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration ModelInferencePaymentConfig

An object representing the collaboration member's model inference payment responsibilities set by the
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
pay for model inference costs ( `TRUE`) or has not configured the collaboration
member to pay for model inference costs ( `FALSE`).

Exactly one member can be configured to pay for model inference costs. An error is
returned if the collaboration creator sets a `TRUE` value for more than one
member in the collaboration.

If the collaboration creator hasn't specified anyone as the member paying for model inference costs, then the member who can query is the default payer. An error is returned if
the collaboration creator sets a `FALSE` value for the member who can
query.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MLPaymentConfig

ModelTrainingPaymentConfig

All content copied from https://docs.aws.amazon.com/.
