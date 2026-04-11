# StringCondition

This contains the value of the string and can contain
one or more operators.

## Contents

**Value**

The value of the string.

Type: String

Required: Yes

**Operator**

A string that defines what values will be
returned.

If this is included, avoid combinations of
operators that will return all possible values.
For example, including both `EQUALS_TO`
and `NOT_EQUALS_TO` with a value of `4`
will return all values.

Type: String

Valid Values: `EQUALS_TO | NOT_EQUALS_TO | CONTAINS | DOES_NOT_CONTAIN | BEGINS_WITH | ENDS_WITH | DOES_NOT_BEGIN_WITH | DOES_NOT_END_WITH`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/stringcondition.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/stringcondition.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/stringcondition.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SearchScopeSummary

TimeCondition

All content copied from https://docs.aws.amazon.com/.
