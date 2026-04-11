# DocumentAttributeCondition

The condition used for the target document attribute or metadata field when ingesting
documents into Amazon Q Business. You use this with [`DocumentAttributeTarget`](api-documentattributetarget.md) to apply the condition.

For example, you can create the 'Department' target field and have it prefill
department names associated with the documents based on information in the 'Source\_URI'
field. Set the condition that if the 'Source\_URI' field contains 'financial' in its URI
value, then prefill the target field 'Department' with the target value 'Finance' for
the document.

Amazon Q Business can't create a target field if it has not already been created as an
index field. After you create your index field, you can create a document metadata field
using `DocumentAttributeTarget`. Amazon Q Business then will map your newly
created metadata field to your index field.

## Contents

**key**

The identifier of the document attribute used for the condition.

For example, 'Source\_URI' could be an identifier for the attribute or metadata field
that contains source URIs associated with the documents.

Amazon Q Business currently doesn't support `_document_body` as an attribute
key used for the condition.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

Required: Yes

**operator**

The identifier of the document attribute used for the condition.

For example, 'Source\_URI' could be an identifier for the attribute or metadata field
that contains source URIs associated with the documents.

Amazon Q Business currently does not support `_document_body` as an attribute
key used for the condition.

Type: String

Valid Values: `GREATER_THAN | GREATER_THAN_OR_EQUALS | LESS_THAN | LESS_THAN_OR_EQUALS | EQUALS | NOT_EQUALS | CONTAINS | NOT_CONTAINS | EXISTS | NOT_EXISTS | BEGINS_WITH`

Required: Yes

**value**

The value of a document attribute. You can only provide one value for a document
attribute.

Type: [DocumentAttributeValue](api-documentattributevalue.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentattributecondition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentattributecondition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentattributecondition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeBoostingConfiguration

DocumentAttributeConfiguration

All content copied from https://docs.aws.amazon.com/.
