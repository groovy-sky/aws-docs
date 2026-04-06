# Task

A class for modeling different type of tasks. Task implementation varies based on the
`TaskType`.

## Contents

**sourceFields**

The source fields to which a particular task is applied.

Type: Array of strings

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: Yes

**taskType**

Specifies the particular task implementation that Amazon AppFlow performs.

Type: String

Valid Values: `Arithmetic | Filter | Map | Map_all | Mask | Merge | Passthrough | Truncate | Validate | Partition`

Required: Yes

**connectorOperator**

The operation to be performed on the provided source fields.

Type: [ConnectorOperator](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ConnectorOperator.html) object

Required: No

**destinationField**

A field in a destination connector, or a field value against which Amazon AppFlow
validates a source field.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `.*`

Required: No

**taskProperties**

A map used to store task-related information. The execution service looks for particular
information based on the `TaskType`.

Type: String to string map

Valid Keys: `VALUE | VALUES | DATA_TYPE | UPPER_BOUND | LOWER_BOUND | SOURCE_DATA_TYPE | DESTINATION_DATA_TYPE | VALIDATION_ACTION | MASK_VALUE | MASK_LENGTH | TRUNCATE_LENGTH | MATH_OPERATION_FIELDS_ORDER | CONCAT_FORMAT | SUBFIELD_CATEGORY_MAP | EXCLUDE_SOURCE_FIELDS_LIST | INCLUDE_NEW_FIELDS | ORDERED_PARTITION_KEYS_LIST`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `.+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/Task)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/Task)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/Task)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SupportedFieldTypeDetails

TrendmicroConnectorProfileCredentials
