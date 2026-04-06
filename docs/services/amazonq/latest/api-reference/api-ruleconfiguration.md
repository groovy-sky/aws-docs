# RuleConfiguration

Provides configuration information about a rule.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**contentBlockerRule**

A rule for configuring how Amazon Q Business responds when it encounters a a blocked
topic.

Type: [ContentBlockerRule](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_ContentBlockerRule.html) object

Required: No

**contentRetrievalRule**

Rules for retrieving content from data sources connected to a Amazon Q Business
application for a specific topic control configuration.

Type: [ContentRetrievalRule](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_ContentRetrievalRule.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/RuleConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/RuleConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/RuleConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Rule

S3
