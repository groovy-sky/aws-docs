# RuleConfiguration

Provides configuration information about a rule.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**contentBlockerRule**

A rule for configuring how Amazon Q Business responds when it encounters a a blocked
topic.

Type: [ContentBlockerRule](api-contentblockerrule.md) object

Required: No

**contentRetrievalRule**

Rules for retrieving content from data sources connected to a Amazon Q Business
application for a specific topic control configuration.

Type: [ContentRetrievalRule](api-contentretrievalrule.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/ruleconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/ruleconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/ruleconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

S3

All content copied from https://docs.aws.amazon.com/.
