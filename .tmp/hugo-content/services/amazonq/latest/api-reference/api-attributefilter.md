# AttributeFilter

Enables filtering of responses based on document attributes or metadata fields.

## Contents

**andAllFilters**

Performs a logical `AND` operation on all supplied filters.

Type: Array of [AttributeFilter](api-attributefilter.md) objects

Required: No

**containsAll**

Returns `true` when a document contains all the specified document
attributes or metadata fields. Supported for the following [document attribute value types](api-documentattributevalue.md):
`stringListValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**containsAny**

Returns `true` when a document contains any of the specified document
attributes or metadata fields. Supported for the following [document attribute value types](api-documentattributevalue.md):
`stringListValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**equalsTo**

Performs an equals operation on two document attributes or metadata fields. Supported
for the following [document attribute value types](api-documentattributevalue.md): `dateValue`,
`longValue`, `stringListValue` and
`stringValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**greaterThan**

Performs a greater than operation on two document attributes or metadata fields.
Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue`
and `longValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**greaterThanOrEquals**

Performs a greater or equals than operation on two document attributes or metadata
fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue`
and `longValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**lessThan**

Performs a less than operation on two document attributes or metadata fields.
Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue`
and `longValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**lessThanOrEquals**

Performs a less than or equals operation on two document attributes or metadata
fields.Supported for the following [document attribute value type](api-documentattributevalue.md): `dateValue`
and `longValue`.

Type: [DocumentAttribute](api-documentattribute.md) object

Required: No

**notFilter**

Performs a logical `NOT` operation on all supplied filters.

Type: [AttributeFilter](api-attributefilter.md) object

Required: No

**orAllFilters**

Performs a logical `OR` operation on all supplied filters.

Type: Array of [AttributeFilter](api-attributefilter.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/attributefilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/attributefilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/attributefilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttachmentsConfiguration

AudioExtractionConfiguration

All content copied from https://docs.aws.amazon.com/.
