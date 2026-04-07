# Filter

A filter that can use conditional operators.

For more information about filters, see [Querying Discovered\
Configuration Items](../../../../services/application-discovery/latest/userguide/discovery-api-queries.md) in the _AWS Application Discovery Service User_
_Guide_.

## Contents

**condition**

A conditional operator. The following operators are valid: EQUALS, NOT\_EQUALS,
CONTAINS, NOT\_CONTAINS. If you specify multiple filters, the system utilizes all filters as
though concatenated by _AND_. If you specify multiple values for a
particular filter, the system differentiates the values using _OR_. Calling
either _DescribeConfigurations_ or _ListConfigurations_
returns attributes of matching configuration items.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `\S+`

Required: Yes

**name**

The name of the filter.

Type: String

Length Constraints: Maximum length of 10000.

Pattern: `[\s\S]*`

Required: Yes

**values**

A string value on which to filter. For example, if you choose the
`destinationServer.osVersion` filter name, you could specify `Ubuntu`
for the value.

Type: Array of strings

Length Constraints: Maximum length of 1000.

Pattern: `(^$|[\s\S]*\S[\s\S]*)`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/discovery-2015-11-01/Filter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/discovery-2015-11-01/Filter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/discovery-2015-11-01/Filter)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FailedConfiguration

ImportTask
