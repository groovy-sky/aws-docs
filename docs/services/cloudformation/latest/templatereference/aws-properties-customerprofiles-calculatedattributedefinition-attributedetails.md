This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::CalculatedAttributeDefinition AttributeDetails

Mathematical expression and a list of attribute items specified in that
expression.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ AttributeItem, ... ],
  "Expression" : String
}

```

### YAML

```yaml

  Attributes:
    - AttributeItem
  Expression: String

```

## Properties

`Attributes`

Mathematical expression and a list of attribute items specified in that
expression.

_Required_: Yes

_Type_: Array of [AttributeItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-calculatedattributedefinition-attributeitem.html)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

Mathematical expression that is performed on attribute items provided in the attribute
list. Each element in the expression should follow the structure of
\\"{ObjectTypeName.AttributeName}\\".

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CustomerProfiles::CalculatedAttributeDefinition

AttributeItem
