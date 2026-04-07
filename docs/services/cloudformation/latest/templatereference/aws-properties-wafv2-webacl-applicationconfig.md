This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ApplicationConfig

A list of `ApplicationAttribute` s that contains information about the application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ ApplicationAttribute, ... ]
}

```

### YAML

```yaml

  Attributes:
    - ApplicationAttribute

```

## Properties

`Attributes`

Contains the attribute name and a list of values for that attribute.

_Required_: Yes

_Type_: Array of [ApplicationAttribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-applicationattribute.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApplicationAttribute

AsnMatchStatement
