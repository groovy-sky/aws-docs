This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource DocumentAttributeCondition

The condition used for the target document attribute or metadata field when ingesting
documents into Amazon Kendra. You use this with [DocumentAttributeTarget to\
apply the condition](../../../kendra/latest/dg/api-documentattributetarget.md).

For example, you can create the 'Department' target field and have it prefill
department names associated with the documents based on information in the 'Source\_URI'
field. Set the condition that if the 'Source\_URI' field contains 'financial' in its URI
value, then prefill the target field 'Department' with the target value 'Finance' for
the document.

Amazon Kendra cannot create a target field if it has not already been created as
an index field. After you create your index field, you can create a document metadata
field using `DocumentAttributeTarget`. Amazon Kendra then will map your
newly created metadata field to your index field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionDocumentAttributeKey" : String,
  "ConditionOnValue" : DocumentAttributeValue,
  "Operator" : String
}

```

### YAML

```yaml

  ConditionDocumentAttributeKey: String
  ConditionOnValue:
    DocumentAttributeValue
  Operator: String

```

## Properties

`ConditionDocumentAttributeKey`

The identifier of the document attribute used for the condition.

For example, 'Source\_URI' could be an identifier for the attribute or metadata field
that contains source URIs associated with the documents.

Amazon Kendra currently does not support `_document_body` as an
attribute key used for the condition.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConditionOnValue`

The value used by the operator.

For example, you can specify the value 'financial' for strings in the 'Source\_URI'
field that partially match or contain this value.

_Required_: No

_Type_: [DocumentAttributeValue](aws-properties-kendra-datasource-documentattributevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The condition operator.

For example, you can use 'Contains' to partially match a string.

_Required_: Yes

_Type_: String

_Allowed values_: `GreaterThan | GreaterThanOrEquals | LessThan | LessThanOrEquals | Equals | NotEquals | Contains | NotContains | Exists | NotExists | BeginsWith`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceVpcConfiguration

DocumentAttributeTarget

All content copied from https://docs.aws.amazon.com/.
