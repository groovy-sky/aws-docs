# ExportFilter

Used to select which agent's data is to be exported. A single agent ID may be selected
for export using the [StartExportTask](api-startexporttask.md) action.

## Contents

**condition**

Supported condition: `EQUALS`

Type: String

Length Constraints: Maximum length of 200.

Pattern: `\S+`

Required: Yes

**name**

A single `ExportFilter` name. Supported filters:
`agentIds`.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `[\s\S]*\S[\s\S]*`

Required: Yes

**values**

A single agent ID for a Discovery Agent. An agent ID can be found using the [DescribeAgents](api-describeagents.md) action. Typically an ADS agent ID is in the form
`o-0123456789abcdef0`.

Type: Array of strings

Length Constraints: Maximum length of 1000.

Pattern: `(^$|[\s\S]*\S[\s\S]*)`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/exportfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/exportfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/exportfilter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Ec2RecommendationsExportPreferences

ExportInfo
