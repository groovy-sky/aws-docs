This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource DocumentAttributeTarget

The target document attribute or metadata field you want to alter when ingesting
documents into Amazon Kendra.

For example, you can delete customer identification numbers associated with the
documents, stored in the document metadata field called 'Customer\_ID'. You set the
target key as 'Customer\_ID' and the deletion flag to `TRUE`. This removes all
customer ID values in the field 'Customer\_ID'. This would scrub personally identifiable
information from each document's metadata.

Amazon Kendra cannot create a target field if it has not already been created as
an index field. After you create your index field, you can create a document metadata
field using `DocumentAttributeTarget`. Amazon Kendra then will map your
newly created metadata field to your index field.

You can also use this with [DocumentAttributeCondition](../../../kendra/latest/dg/api-documentattributecondition.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetDocumentAttributeKey" : String,
  "TargetDocumentAttributeValue" : DocumentAttributeValue,
  "TargetDocumentAttributeValueDeletion" : Boolean
}

```

### YAML

```yaml

  TargetDocumentAttributeKey: String
  TargetDocumentAttributeValue:
    DocumentAttributeValue
  TargetDocumentAttributeValueDeletion: Boolean

```

## Properties

`TargetDocumentAttributeKey`

The identifier of the target document attribute or metadata field.

For example, 'Department' could be an identifier for the target attribute or metadata
field that includes the department names associated with the documents.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDocumentAttributeValue`

The target value you want to create for the target attribute.

For example, 'Finance' could be the target value for the target attribute key
'Department'.

_Required_: No

_Type_: [DocumentAttributeValue](aws-properties-kendra-datasource-documentattributevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDocumentAttributeValueDeletion`

`TRUE` to delete the existing target value for your specified target
attribute key. You cannot create a target value and set this to `TRUE`. To
create a target value ( `TargetDocumentAttributeValue`), set this to
`FALSE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeCondition

DocumentAttributeValue

All content copied from https://docs.aws.amazon.com/.
