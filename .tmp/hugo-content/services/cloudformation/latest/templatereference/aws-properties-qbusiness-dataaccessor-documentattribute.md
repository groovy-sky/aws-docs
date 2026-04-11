This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor DocumentAttribute

A document attribute or metadata field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : DocumentAttributeValue
}

```

### YAML

```yaml

  Name: String
  Value:
    DocumentAttributeValue

```

## Properties

`Name`

The identifier for the attribute.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the attribute.

_Required_: Yes

_Type_: [DocumentAttributeValue](aws-properties-qbusiness-dataaccessor-documentattributevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataAccessorIdcTrustedTokenIssuerConfiguration

DocumentAttributeValue

All content copied from https://docs.aws.amazon.com/.
