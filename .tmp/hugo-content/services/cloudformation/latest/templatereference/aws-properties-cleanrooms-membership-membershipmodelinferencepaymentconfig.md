This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipModelInferencePaymentConfig

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

Indicates whether the collaboration member has accepted to pay for model inference costs
( `TRUE`) or has not accepted to pay for model inference costs
( `FALSE`).

If the collaboration creator has not specified anyone to pay for model inference costs,
then the member who can query is the default payer.

An error message is returned for the following reasons:

- If you set the value to `FALSE` but you are responsible to pay for
model inference costs.

- If you set the value to `TRUE` but you are not responsible to pay for
model inference costs.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MembershipMLPaymentConfig

MembershipModelTrainingPaymentConfig

All content copied from https://docs.aws.amazon.com/.
