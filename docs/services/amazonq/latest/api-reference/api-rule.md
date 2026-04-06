# Rule

Guardrail rules for an Amazon Q Business application. Amazon Q Business supports only one rule
at a time.

## Contents

**ruleType**

The type of rule.

Type: String

Valid Values: `CONTENT_BLOCKER_RULE | CONTENT_RETRIEVAL_RULE`

Required: Yes

**excludedUsersAndGroups**

Users and groups to be excluded from a rule.

Type: [UsersAndGroups](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_UsersAndGroups.html) object

Required: No

**includedUsersAndGroups**

Users and groups to be included in a rule.

Type: [UsersAndGroups](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_UsersAndGroups.html) object

Required: No

**ruleConfiguration**

The configuration information for a rule.

Type: [RuleConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_RuleConfiguration.html) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/Rule)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/Rule)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/Rule)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RetrieverContentSource

RuleConfiguration
