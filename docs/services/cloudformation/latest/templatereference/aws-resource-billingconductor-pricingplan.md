This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::PricingPlan

Creates a pricing plan that is used for computing AWS charges for billing groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BillingConductor::PricingPlan",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "PricingRuleArns" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BillingConductor::PricingPlan
Properties:
  Description: String
  Name: String
  PricingRuleArns:
    - String
  Tags:
    - Tag

```

## Properties

`Description`

The pricing plan description.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a pricing plan.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\+=\.\-@]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PricingRuleArns`

The `PricingRuleArns` that are associated with the Pricing Plan.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A map that contains tag keys and tag values that are attached to a pricing plan.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-pricingplan-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the created pricing plan.

`CreationTime`

The time the pricing plan was created.

`LastModifiedTime`

The most recent time the pricing plan was modified.

`Size`

The pricing rules count currently associated with this pricing plan list element.

## Examples

### Pricing plan with a pricing rule attached

The following example creates a pricing plan with a 10% global mark up pricing rule attached.

#### JSON

```json

{
  "Resources": {
      "TestPricingRule": {
          "Type": "AWS::BillingConductor::PricingRule",
          "Properties": {
              "Name": "TestPricingRule",
              "Description": "Test pricing rule created through CloudFormation. Mark everything by 10%.",
              "Type": "MARKUP",
              "Scope": "GLOBAL",
              "ModifierPercentage": 10
          }
      },
      "TestPricingPlan": {
          "Type": "AWS::BillingConductor::PricingPlan",
          "Properties": {
              "Name": "TestPricingPlan",
              "Description": "Test pricing plan created through CloudFormation.",
              "PricingRuleArns": [
                  {"Fn::GetAtt": ["TestPricingRule", "Arn"]}
              ]
          }
      }
  }
}
```

#### YAML

```yaml

Resources:
  TestPricingRule:
      Type: 'AWS::BillingConductor::PricingRule'
      Properties:
        Name: 'TestPricingRule'
        Description: 'Test pricing rule created through CloudFormation. Mark everything by 10%.'
        Type: 'MARKUP'
        Scope: 'GLOBAL'
        ModifierPercentage: 10
  TestPricingPlan:
    Type: 'AWS::BillingConductor::PricingPlan'
    Properties:
      Name: 'TestPricingPlan'
      Description: 'Test pricing plan created through CloudFormation.'
      PricingRuleArns:
        - !GetAtt TestPricingRule.Arn

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
