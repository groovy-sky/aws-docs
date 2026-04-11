This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource DocumentAttributeTarget

The target document attribute or metadata field you want to alter when ingesting
documents into Amazon Q Business.

For example, you can delete all customer identification numbers associated with the
documents, stored in the document metadata field called 'Customer\_ID' by setting the
target key as 'Customer\_ID' and the deletion flag to `TRUE`. This removes all
customer ID values in the field 'Customer\_ID'. This would scrub personally identifiable
information from each document's metadata.

Amazon Q Business can't create a target field if it has not already been created as an
index field. After you create your index field, you can create a document metadata field
using [`DocumentAttributeTarget`](../../../amazonq/latest/api-reference/api-documentattributetarget.md). Amazon Q Business
will then map your newly created document attribute to your index field.

You can also use this with [`DocumentAttributeCondition`](../../../amazonq/latest/api-reference/api-documentattributecondition.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeValueOperator" : String,
  "Key" : String,
  "Value" : DocumentAttributeValue
}

```

### YAML

```yaml

  AttributeValueOperator: String
  Key: String
  Value:
    DocumentAttributeValue

```

## Properties

`AttributeValueOperator`

`TRUE` to delete the existing target value for your specified target
attribute key. You cannot create a target value and set this to
`TRUE`.

_Required_: No

_Type_: String

_Allowed values_: `DELETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The identifier of the target document attribute or metadata field. For example,
'Department' could be an identifier for the target attribute or metadata field that
includes the department names associated with the documents.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a document attribute. You can only provide one value for a document
attribute.

_Required_: No

_Type_: [DocumentAttributeValue](aws-properties-qbusiness-datasource-documentattributevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeCondition

DocumentAttributeValue

All content copied from https://docs.aws.amazon.com/.
