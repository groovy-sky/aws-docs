# DocumentAttributeTarget

The target document attribute or metadata field you want to alter when ingesting
documents into Amazon Q Business.

For example, you can delete all customer identification numbers associated with the
documents, stored in the document metadata field called 'Customer\_ID' by setting the
target key as 'Customer\_ID' and the deletion flag to `TRUE`. This removes all
customer ID values in the field 'Customer\_ID'. This would scrub personally identifiable
information from each document's metadata.

Amazon Q Business can't create a target field if it has not already been created as an
index field. After you create your index field, you can create a document metadata field
using [`DocumentAttributeTarget`](api-documentattributetarget.md). Amazon Q Business
will then map your newly created document attribute to your index field.

You can also use this with [`DocumentAttributeCondition`](api-documentattributecondition.md).

## Contents

**key**

The identifier of the target document attribute or metadata field. For example,
'Department' could be an identifier for the target attribute or metadata field that
includes the department names associated with the documents.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

Required: Yes

**attributeValueOperator**

`TRUE` to delete the existing target value for your specified target
attribute key. You cannot create a target value and set this to
`TRUE`.

Type: String

Valid Values: `DELETE`

Required: No

**value**

The value of a document attribute. You can only provide one value for a document
attribute.

Type: [DocumentAttributeValue](api-documentattributevalue.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentattributetarget.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentattributetarget.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentattributetarget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeConfiguration

DocumentAttributeValue

All content copied from https://docs.aws.amazon.com/.
