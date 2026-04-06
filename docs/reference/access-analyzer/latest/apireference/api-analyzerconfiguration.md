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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/AnalyzerConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/AnalyzerConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/AnalyzerConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalyzedResourceSummary

AnalyzerSummary
