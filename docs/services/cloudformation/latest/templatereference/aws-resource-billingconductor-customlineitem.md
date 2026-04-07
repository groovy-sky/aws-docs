This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem

Creates a custom line item that can be used to create a one-time or recurring, fixed or percentage-based charge that you can apply to a single billing group. You can apply custom line items to the current or previous billing period. You can create either a fee or a discount custom line item.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BillingConductor::CustomLineItem",
  "Properties" : {
      "AccountId" : String,
      "BillingGroupArn" : String,
      "BillingPeriodRange" : BillingPeriodRange,
      "ComputationRule" : String,
      "CustomLineItemChargeDetails" : CustomLineItemChargeDetails,
      "Description" : String,
      "Name" : String,
      "PresentationDetails" : PresentationDetails,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BillingConductor::CustomLineItem
Properties:
  AccountId: String
  BillingGroupArn: String
  BillingPeriodRange:
    BillingPeriodRange
  ComputationRule: String
  CustomLineItemChargeDetails:
    CustomLineItemChargeDetails
  Description: String
  Name: String
  PresentationDetails:
    PresentationDetails
  Tags:
    - Tag

```

## Properties

`AccountId`

The AWS account in which this custom line item will be applied to.

_Required_: No

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BillingGroupArn`

The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-cn)?:billingconductor::[0-9]{12}:billinggroup/?[a-zA-Z0-9]{10,12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BillingPeriodRange`

A time range for which the custom line item is effective.

_Required_: No

_Type_: [BillingPeriodRange](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-billingperiodrange.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationRule`

The computation rule that determines how the custom line item charges are computed and reflected in the bill.

_Required_: No

_Type_: String

_Allowed values_: `CONSOLIDATED | ITEMIZED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomLineItemChargeDetails`

The charge details of a custom line item. It should contain only one of `Flat` or `Percentage`.

_Required_: No

_Type_: [CustomLineItemChargeDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-customlineitemchargedetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The custom line item's description. This is shown on the Bills page in association with the charge value.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The custom line item's name.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\+=\.\-@]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PresentationDetails`

Configuration details specifying how the custom line item charges are presented, including which service the charges are shown under.

_Required_: No

_Type_: [PresentationDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-presentationdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A map that contains tag keys and tag values that are attached to a custom line item.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-customlineitem-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) that references the billing group where the custom line item applies to.

`AssociationSize`

The number of resources that are associated to the custom line item.

`CreationTime`

The time created.

`CurrencyCode`

The custom line item's charge value currency. Only one of the valid values can be used.

`LastModifiedTime`

The most recent time the custom line item was modified.

`ProductCode`

The product code associated with the custom line item.

## Examples

- [Flat custom line item](#aws-resource-billingconductor-customlineitem--examples--Flat_custom_line_item)

- [Percentage custom line item](#aws-resource-billingconductor-customlineitem--examples--Percentage_custom_line_item)

### Flat custom line item

The following example shows a flat charge custom line item of $10 attached to a billing group.

#### JSON

```json

{
  "Resources": {
      "TestFlatCustomLineItem": {
          "Type": "AWS::BillingConductor::CustomLineItem",
          "Properties": {
              "Name": "TestFlatCustomLineItem",
              "Description": "Test flat custom line item created through CloudFormation for a $10 charge.",
              "BillingGroupArn": {
                  "Fn::GetAtt": ["TestBillingGroup", "Arn"]
              },
              "CustomLineItemChargeDetails": {
                  "Flat": {
                      "ChargeValue": 10
                  },
                  "Type": "FEE"
              },
              "BillingPeriodRange": {
              "InclusiveStartBillingPeriod": "2022-11"
              }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  TestFlatCustomLineItem:
      Type: 'AWS::BillingConductor::CustomLineItem'
      Properties:
          Name: 'TestFlatCustomLineItem'
          Description: 'Test flat custom line item created through CloudFormation for a $10 charge.'
          BillingGroupArn: !GetAtt TestBillingGroup.Arn
          CustomLineItemChargeDetails:
              Flat:
                  ChargeValue: 10
              Type: 'FEE'
          BillingPeriodRange:
            InclusiveStartBillingPeriod: 2022-11

```

### Percentage custom line item

The following example shows a percentage charge custom line item that is 10% of the billing group total cost.

#### JSON

```json

{
  "Resources": {
      "TestPercentageCustomLineItem": {
          "Type": "AWS::BillingConductor::CustomLineItem",
          "Properties": {
              "Name": "TestPercentageCustomLineItem",
              "Description": "Test percentage custom line item created through CloudFormation for a %10 additional charge on the overall total bill of the billing group.",
              "BillingGroupArn": {
                  "Fn::GetAtt": ["TestBillingGroup", "Arn"]
              },
              "CustomLineItemChargeDetails": {
                  "Percentage": {
                      "PercentageValue": 10,
                      "ChildAssociatedResources": [
                          {"Fn::GetAtt": ["TestBillingGroup", "Arn"]}
                      ]
                  },
                  "Type": "FEE"
              }
          }
      }
  }
}
```

#### YAML

```yaml

Resources:
        TestPercentageCustomLineItem:
        Type: 'AWS::BillingConductor::CustomLineItem'
        Properties:
        Name: 'TestPercentageCustomLineItem'
        Description: 'Test percentage custom line item created through CloudFormation for a %10 additional charge on the overall total bill of the billing group.'
        BillingGroupArn: !GetAtt TestBillingGroup.Arn
        CustomLineItemChargeDetails:
        Percentage:
        PercentageValue: 10
        ChildAssociatedResources:
        - !GetAtt TestBillingGroup.Arn
        Type: 'FEE'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

BillingPeriodRange
