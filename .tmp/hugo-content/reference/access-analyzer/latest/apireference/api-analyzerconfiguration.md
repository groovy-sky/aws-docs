# AnalyzerConfiguration

Contains information about the configuration of an analyzer for an AWS organization or
account.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**internalAccess**

Specifies the configuration of an internal access analyzer for an AWS organization or
account. This configuration determines how the analyzer evaluates access within your AWS
environment.

Type: [InternalAccessConfiguration](api-internalaccessconfiguration.md) object

Required: No

**unusedAccess**

Specifies the configuration of an unused access analyzer for an AWS organization or
account.

Type: [UnusedAccessConfiguration](api-unusedaccessconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/analyzerconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/analyzerconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/analyzerconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnalyzedResourceSummary

AnalyzerSummary
