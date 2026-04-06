# AttributeFilter

Enables filtering of responses based on document attributes or metadata fields.

## Contents

**andAllFilters**

Performs a logical `AND` operation on all supplied filters.

Type: Array of [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AttributeFilter.html) objects

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

Type: [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AttributeFilter.html) object

Required: No

**orAllFilters**

Performs a logical `OR` operation on all supplied filters.

Type: Array of [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AttributeFilter.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/AttributeFilter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/AttributeFilter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/AttributeFilter)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttachmentsConfiguration

AudioExtractionConfiguration
