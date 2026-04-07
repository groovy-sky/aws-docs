This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::BillingGroup AccountGrouping

The set of accounts that will be under the billing group. The set of accounts resemble the linked accounts in a consolidated billing family.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoAssociate" : Boolean,
  "LinkedAccountIds" : [ String, ... ],
  "ResponsibilityTransferArn" : String
}

```

### YAML

```yaml

  AutoAssociate: Boolean
  LinkedAccountIds:
    - String
  ResponsibilityTransferArn: String

```

## Properties

`AutoAssociate`

Specifies if this billing group will automatically associate newly added AWS accounts
that join your consolidated billing family.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkedAccountIds`

The account IDs that make up the billing group. Account IDs must be a part of the consolidated billing family, and not associated with another billing group.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponsibilityTransferArn`

The Amazon Resource Name (ARN) that identifies the transfer relationship owned by the Bill Transfer account (caller account). When specified, the PrimaryAccountId is no longer required.

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z0-9][a-z0-9-.]{0,62}:organizations::[0-9]{12}:transfer/o-[a-z0-9]{10,32}/(billing)/(inbound|outbound)/rt-[0-9a-z]{8,32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::BillingConductor::BillingGroup

ComputationPreference
