This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration MemberSpecification

Basic metadata used to construct a new member.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "DisplayName" : String,
  "MemberAbilities" : [ String, ... ],
  "MLMemberAbilities" : MLMemberAbilities,
  "PaymentConfiguration" : PaymentConfiguration
}

```

### YAML

```yaml

  AccountId: String
  DisplayName: String
  MemberAbilities:
    - String
  MLMemberAbilities:
    MLMemberAbilities
  PaymentConfiguration:
    PaymentConfiguration

```

## Properties

`AccountId`

The identifier used to reference members of the collaboration. Currently only supports
AWS account ID.

_Required_: Yes

_Type_: String

_Pattern_: `^\d+$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The member's display name.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemberAbilities`

The abilities granted to the collaboration member.

_Allowed Values_: `CAN_QUERY` \| `CAN_RECEIVE_RESULTS`

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MLMemberAbilities`

The ML abilities granted to the collaboration member.

_Required_: No

_Type_: [MLMemberAbilities](aws-properties-cleanrooms-collaboration-mlmemberabilities.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PaymentConfiguration`

The collaboration member's payment responsibilities set by the collaboration creator.

If the collaboration creator hasn't speciﬁed anyone as the member paying for query
compute costs, then the member who can query is the default payer.

_Required_: No

_Type_: [PaymentConfiguration](aws-properties-cleanrooms-collaboration-paymentconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobComputePaymentConfig

MLMemberAbilities
