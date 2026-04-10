This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem BillingPeriodRange

The billing period range in which the custom line item request will be applied.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExclusiveEndBillingPeriod" : String,
  "InclusiveStartBillingPeriod" : String
}

```

### YAML

```yaml

  ExclusiveEndBillingPeriod: String
  InclusiveStartBillingPeriod: String

```

## Properties

`ExclusiveEndBillingPeriod`

The exclusive end billing period that defines a billing period range where a custom line is applied.

_Required_: No

_Type_: String

_Pattern_: `\d{4}-(0?[1-9]|1[012])`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InclusiveStartBillingPeriod`

The inclusive start billing period that defines a billing period range where a custom line is applied.

_Required_: No

_Type_: String

_Pattern_: `\d{4}-(0?[1-9]|1[012])`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BillingConductor::CustomLineItem

CustomLineItemChargeDetails

All content copied from https://docs.aws.amazon.com/.
