This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource DocumentAttributeCondition

The condition used for the target document attribute or metadata field when ingesting
documents into Amazon Q Business. You use this with [`DocumentAttributeTarget`](../../../amazonq/latest/api-reference/api-documentattributetarget.md) to apply the condition.

For example, you can create the 'Department' target field and have it prefill
department names associated with the documents based on information in the 'Source\_URI'
field. Set the condition that if the 'Source\_URI' field contains 'financial' in its URI
value, then prefill the target field 'Department' with the target value 'Finance' for
the document.

Amazon Q Business can't create a target field if it has not already been created as an
index field. After you create your index field, you can create a document metadata field
using `DocumentAttributeTarget`. Amazon Q Business then will map your newly
created metadata field to your index field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Operator" : String,
  "Value" : DocumentAttributeValue
}

```

### YAML

```yaml

  Key: String
  Operator: String
  Value:
    DocumentAttributeValue

```

## Properties

`Key`

The identifier of the document attribute used for the condition.

For example, 'Source\_URI' could be an identifier for the attribute or metadata field
that contains source URIs associated with the documents.

Amazon Q Business currently doesn't support `_document_body` as an attribute
key used for the condition.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The identifier of the document attribute used for the condition.

For example, 'Source\_URI' could be an identifier for the attribute or metadata field
that contains source URIs associated with the documents.

Amazon Q Business currently does not support `_document_body` as an attribute
key used for the condition.

_Required_: Yes

_Type_: String

_Allowed values_: `GREATER_THAN | GREATER_THAN_OR_EQUALS | LESS_THAN | LESS_THAN_OR_EQUALS | EQUALS | NOT_EQUALS | CONTAINS | NOT_CONTAINS | EXISTS | NOT_EXISTS | BEGINS_WITH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a document attribute. You can only provide one value for a document
attribute.

_Required_: No

_Type_: [DocumentAttributeValue](aws-properties-qbusiness-datasource-documentattributevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceVpcConfiguration

DocumentAttributeTarget

All content copied from https://docs.aws.amazon.com/.
