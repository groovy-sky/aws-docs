# FieldTypeDetails

Contains details regarding the supported field type and the operators that can be applied
for filtering.

## Contents

**fieldType**

The type of field, such as string, integer, date, and so on.

Type: String

Required: Yes

**filterOperators**

The list of operators supported by a field.

Type: Array of strings

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | CONTAINS | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: Yes

**fieldLengthRange**

This is the allowable length range for this field's value.

Type: [Range](api-range.md) object

Required: No

**fieldValueRange**

The range of values this field can hold.

Type: [Range](api-range.md) object

Required: No

**supportedDateFormat**

The date format that the field supports.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: No

**supportedValues**

The list of values that a field can contain. For example, a Boolean
`fieldType` can have two values: "true" and "false".

Type: Array of strings

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: No

**valueRegexPattern**

The regular expression pattern for the field name.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/fieldtypedetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/fieldtypedetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/fieldtypedetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecutionResult

FlowDefinition

All content copied from https://docs.aws.amazon.com/.
