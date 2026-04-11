This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Invoicing::InvoiceUnit

An invoice unit is a set of mutually exclusive account that correspond to your business entity. Invoice units allow you separate AWS account costs and configures your invoice for each business entity going forward.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Invoicing::InvoiceUnit",
  "Properties" : {
      "Description" : String,
      "InvoiceReceiver" : String,
      "Name" : String,
      "ResourceTags" : [ ResourceTag, ... ],
      "Rule" : Rule,
      "TaxInheritanceDisabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Invoicing::InvoiceUnit
Properties:
  Description: String
  InvoiceReceiver: String
  Name: String
  ResourceTags:
    - ResourceTag
  Rule:
    Rule
  TaxInheritanceDisabled: Boolean

```

## Properties

`Description`

The assigned description for an invoice unit. This information can't be modified or deleted.

_Required_: No

_Type_: String

_Pattern_: `^[\S\s]*$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvoiceReceiver`

The account that receives invoices related to the invoice unit.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A unique name that is distinctive within your AWS.

_Required_: Yes

_Type_: String

_Pattern_: `^(?! )[\p{L}\p{N}\p{Z}-_]*(?<! )$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceTags`

The tag structure that contains a tag key and value.

_Required_: No

_Type_: Array of [ResourceTag](aws-properties-invoicing-invoiceunit-resourcetag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rule`

An `InvoiceUnitRule` object used the categorize invoice units.

_Required_: Yes

_Type_: [Rule](aws-properties-invoicing-invoiceunit-rule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaxInheritanceDisabled`

Whether the invoice unit based tax inheritance is/ should be enabled or disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`InvoiceUnitArn`

The ARN to identify an invoice unit. This information can't be modified or deleted.

`LastModified`

The last time the invoice unit was updated. This is important to determine the version of invoice unit configuration used to create the invoices. Any invoice created after this modified time will use this invoice unit configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Invoicing

ResourceTag

All content copied from https://docs.aws.amazon.com/.
