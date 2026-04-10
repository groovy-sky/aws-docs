# LookupAttribute

Specifies an attribute and value that filter the events returned.

## Contents

**AttributeKey**

Specifies an attribute on which to filter the events returned.

Type: String

Valid Values: `EventId | EventName | ReadOnly | Username | ResourceType | ResourceName | EventSource | AccessKeyId`

Required: Yes

**AttributeValue**

Specifies a value for the specified `AttributeKey`.

The maximum length for the `AttributeValue` is 2000 characters. The
following characters (' `_`', ' ` `', ' `,`',
' `\\n`') count as two characters towards the 2000 character limit.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2000.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/lookupattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/lookupattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/lookupattribute.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InsightSelector

PartitionKey

All content copied from https://docs.aws.amazon.com/.
