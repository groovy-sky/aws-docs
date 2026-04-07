# AttributeFilter

The filter criteria used on responses based on document attributes or metadata
fields.

## Contents

**andAllFilters**

Performs a logical `AND` operation on all supplied filters.

Type: Array of [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_AttributeFilter.html) objects

Required: No

**containsAll**

Returns `true` when a document contains all the specified document attributes
or metadata fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `stringListValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**containsAny**

Returns `true` when a document contains any of the specified document
attributes or metadata fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `stringListValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**equalsTo**

Performs an _equals_ operation on two document attributes or metadata
fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue`, `longValue`,
`stringListValue` and `stringValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**greaterThan**

Performs a _greater than_ operation on two document attributes or
metadata fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue` and
`longValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**greaterThanOrEquals**

Performs a _greater than or equals_ operation on two document
attributes or metadata fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue` and `longValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**lessThan**

Performs a _less than_ operation on two document attributes or metadata
fields. Supported for the following [document attribute value types](api-documentattributevalue.md): `dateValue` and
`longValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**lessThanOrEquals**

Performs a _less than or equals_ operation on two document attributes
or metadata fields.Supported for the following [document attribute value type](api-documentattributevalue.md): `dateValue` and `longValue`.

Type: [DocumentAttribute](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_DocumentAttribute.html) object

Required: No

**notFilter**

Performs a logical `NOT` operation on all supplied filters.

Type: [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_AttributeFilter.html) object

Required: No

**orAllFilters**

Performs a logical `OR` operation on all supplied filters.

Type: Array of [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_AttributeFilter.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/AttributeFilter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/AttributeFilter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/AttributeFilter)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AppDefinitionInput

BatchCreateCategoryInputCategory
