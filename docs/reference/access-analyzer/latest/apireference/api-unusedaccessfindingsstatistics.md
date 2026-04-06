# UnusedAccessFindingsStatistics

Provides aggregate statistics about the findings for the specified unused access
analyzer.

## Contents

**topAccounts**

A list of one to ten AWS accounts that have the most active findings for the unused
access analyzer.

Type: Array of [FindingAggregationAccountDetails](api-findingaggregationaccountdetails.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**totalActiveFindings**

The total number of active findings for the unused access analyzer.

Type: Integer

Required: No

**totalArchivedFindings**

The total number of archived findings for the unused access analyzer.

Type: Integer

Required: No

**totalResolvedFindings**

The total number of resolved findings for the unused access analyzer.

Type: Integer

Required: No

**unusedAccessTypeStatistics**

A list of details about the total number of findings for each type of unused access for
the analyzer.

Type: Array of [UnusedAccessTypeStatistics](api-unusedaccesstypestatistics.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/unusedaccessfindingsstatistics.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/unusedaccessfindingsstatistics.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/unusedaccessfindingsstatistics.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnusedAccessConfiguration

UnusedAccessTypeStatistics
