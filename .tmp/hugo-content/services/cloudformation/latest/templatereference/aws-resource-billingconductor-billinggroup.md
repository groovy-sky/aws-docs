This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::BillingGroup

Creates a billing group that resembles a consolidated billing family that AWS charges, based off of the predefined pricing plan computation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BillingConductor::BillingGroup",
  "Properties" : {
      "AccountGrouping" : AccountGrouping,
      "ComputationPreference" : ComputationPreference,
      "Description" : String,
      "Name" : String,
      "PrimaryAccountId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BillingConductor::BillingGroup
Properties:
  AccountGrouping:
    AccountGrouping
  ComputationPreference:
    ComputationPreference
  Description: String
  Name: String
  PrimaryAccountId: String
  Tags:
    - Tag

```

## Properties

`AccountGrouping`

The set of accounts that will be under the billing group. The set of accounts resemble the linked accounts in a consolidated billing family.

_Required_: Yes

_Type_: [AccountGrouping](aws-properties-billingconductor-billinggroup-accountgrouping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationPreference`

The preferences and settings that will be used to compute the AWScharges for a billing group.

_Required_: Yes

_Type_: [ComputationPreference](aws-properties-billingconductor-billinggroup-computationpreference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the billing group.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The billing group's name.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\+=\.\-@]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryAccountId`

The account ID that serves as the main account in a billing group.

_Required_: No

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A map that contains tag keys and tag values that are attached to a billing group.

_Required_: No

_Type_: Array of [Tag](aws-properties-billingconductor-billinggroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the created billing group.

`CreationTime`

The time the billing group was created.

`LastModifiedTime`

The most recent time the billing group was modified.

`Size`

The number of accounts in the particular billing group.

`Status`

The billing group status. Only one of the valid values can be used.

`StatusReason`

The reason why the billing group is in its current status.

## Examples

### Create a billing group with a pricing plan

The following example is a billing group that takes a list of linked account IDs and the primary account as input. This example creates a billing group with a pricing plan that has no pricing rules.

#### JSON

```json

{
  "Parameters": {
      "LinkedAccountIds": {
      "Type": "ListNumber"
      },
      "PrimaryAccountId": {
          "Type": "Number"
      }
  },
  "Resources": {
      "TestBillingGroup": {
          "Type": "AWS::BillingConductor::BillingGroup",
          "Properties": {
              "Name": "TestBillingGroup",
              "Description": "Test billing group created through CloudFormation with 1 linked account. The linked account is also the primary account.",
              "PrimaryAccountId": {
                  "Ref": "PrimaryAccountId"
              },
              "AccountGrouping": {
                  "LinkedAccountIds":{
                      "Ref": "LinkedAccountIds"
                  }
              },
              "ComputationPreference": {
                  "PricingPlanArn": {
                      "Fn::GetAtt": ["TestPricingPlan", "Arn"]
                  }
              }
          }
      }
  }
}
```

#### YAML

```yaml

Parameters:
  LinkedAccountIds:
  Type: ListNumber
  PrimaryAccountId:
      Type: Number
Resources:
  TestBillingGroup:
      Type: 'AWS::BillingConductor::BillingGroup'
      Properties:
        Name: 'TestBillingGroup'
        Description: 'Test billing group created through CloudFormation with 1 linked account. The linked account is also the primary account.'
        PrimaryAccountId: !Ref PrimaryAccountId
        AccountGrouping:
            LinkedAccountIds: !Ref LinkedAccountIds
        ComputationPreference:
            PricingPlanArn: !GetAtt TestPricingPlan.Arn

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Billing Conductor

AccountGrouping

All content copied from https://docs.aws.amazon.com/.
