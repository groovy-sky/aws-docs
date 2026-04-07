This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::PricingRule

Creates a pricing rule which can be associated with a pricing plan, or a set of pricing plans.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BillingConductor::PricingRule",
  "Properties" : {
      "BillingEntity" : String,
      "Description" : String,
      "ModifierPercentage" : Number,
      "Name" : String,
      "Operation" : String,
      "Scope" : String,
      "Service" : String,
      "Tags" : [ Tag, ... ],
      "Tiering" : Tiering,
      "Type" : String,
      "UsageType" : String
    }
}

```

### YAML

```yaml

Type: AWS::BillingConductor::PricingRule
Properties:
  BillingEntity: String
  Description: String
  ModifierPercentage: Number
  Name: String
  Operation: String
  Scope: String
  Service: String
  Tags:
    - Tag
  Tiering:
    Tiering
  Type: String
  UsageType: String

```

## Properties

`BillingEntity`

The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplace.

_Required_: No

_Type_: String

_Allowed values_: `AWS | AWS Marketplace | AISPL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The pricing rule description.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModifierPercentage`

A percentage modifier applied on the public pricing rates.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a pricing rule.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\+=\.\-@]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operation`

Operation is the specific AWS action covered by this line item. This describes the specific usage of the line item.

If the `Scope` attribute is set to `SKU`, this attribute indicates which operation the `PricingRule` is modifying. For example, a value of `RunInstances:0202` indicates the operation of running an Amazon EC2 instance.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scope`

The scope of pricing rule that indicates if it's globally applicable or service-specific.

_Required_: Yes

_Type_: String

_Allowed values_: `GLOBAL | SERVICE | BILLING_ENTITY | SKU`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Service`

If the `Scope` attribute is `SERVICE`, this attribute indicates which service the `PricingRule` is applicable for.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9\.\-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A map that contains tag keys and tag values that are attached to a pricing rule.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-pricingrule-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tiering`

The set of tiering configurations for the pricing rule.

_Required_: No

_Type_: [Tiering](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billingconductor-pricingrule-tiering.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of pricing rule.

_Required_: Yes

_Type_: String

_Allowed values_: `MARKUP | DISCOUNT | TIERING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsageType`

Usage Type is the unit that each service uses to measure the usage of a specific type of resource.

_Required_: No

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) used to uniquely identify a pricing rule.

`AssociatedPricingPlanCount`

The pricing plans count that this pricing rule is associated with.

`CreationTime`

The time the pricing rule was created.

`LastModifiedTime`

The most recent time the pricing rule was modified.

## Examples

- [Simple pricing rule](#aws-resource-billingconductor-pricingrule--examples--Simple_pricing_rule)

- [AWS Marketplace pricing rule](#aws-resource-billingconductor-pricingrule--examples--pricing_rule)

### Simple pricing rule

The following example shows a pricing rule that applies a 10% global markup.

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

```

### AWS Marketplace pricing rule

The following example shows a pricing rule that targets only AWS Marketplace charges.

#### JSON

```json

{
  "Resources": {
      "TestPricingRule": {
          "Type": "AWS::BillingConductor::PricingRule",
          "Properties": {
              "Name": "TestPricingRule",
              "Description": "Test pricing rule created through CloudFormation. Keep all MP charges at public on demand rate. "
              "Type": "MARKUP",
              "Scope": "BILLING_ENTITY",
              "BillingEntity": "AWS Marketplace",
              "ModifierPercentage": 0
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
        Description: 'Test pricing rule created through CloudFormation. Keep all MP charges at public on demand rate.'
        Type: 'MARKUP'
        Scope: 'BILLING_ENTITY'
        BillingEntity: 'AWS Marketplace'
        ModifierPercentage: 0

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

FreeTier
