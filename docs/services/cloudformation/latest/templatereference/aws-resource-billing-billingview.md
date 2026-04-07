This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Billing::BillingView

Creates a billing view with the specified billing view attributes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Billing::BillingView",
  "Properties" : {
      "DataFilterExpression" : DataFilterExpression,
      "Description" : String,
      "Name" : String,
      "SourceViews" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Billing::BillingView
Properties:
  DataFilterExpression:
    DataFilterExpression
  Description: String
  Name: String
  SourceViews:
    - String
  Tags:
    - Tag

```

## Properties

`DataFilterExpression`

See [Expression](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_billing_Expression.html). Billing view only supports `LINKED_ACCOUNT`, `Tags`, and `CostCategories`.

_Required_: No

_Type_: [DataFilterExpression](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billing-billingview-datafilterexpression.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the billing view.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the billing view.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\+=\.\-@]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceViews`

A list of billing views used as the data source for the custom billing view.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key value map specifying tags associated to the billing view being created.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billing-billingview-tags.html) of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-billing-billingview-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) that can be used to uniquely identify the billing view.

`BillingViewType`

The type of billing group.

`CreatedAt`

The time when the billing view was created.

`OwnerAccountId`

The account owner of the billing view.

`UpdatedAt`

The time when the billing view was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Billing

DataFilterExpression
