This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipJobComputePaymentConfig

An object representing the payment responsibilities accepted by the
collaboration member for query and job compute costs.

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

Indicates whether the collaboration member has accepted to pay for job
compute costs ( `TRUE`) or has not accepted to pay for query and job compute costs
( `FALSE`).

There is only one member who pays for queries and jobs.

An error message is returned for the following reasons:

- If you set the value to `FALSE` but you are responsible to
pay for query and job compute costs.

- If you set the value to `TRUE` but you are not responsible to
pay for query and job compute costs.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRooms::Membership

MembershipMLPaymentConfig

All content copied from https://docs.aws.amazon.com/.
