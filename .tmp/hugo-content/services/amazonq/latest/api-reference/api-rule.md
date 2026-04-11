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

Type: [UsersAndGroups](api-usersandgroups.md) object

Required: No

**includedUsersAndGroups**

Users and groups to be included in a rule.

Type: [UsersAndGroups](api-usersandgroups.md) object

Required: No

**ruleConfiguration**

The configuration information for a rule.

Type: [RuleConfiguration](api-ruleconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/rule.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/rule.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/rule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetrieverContentSource

RuleConfiguration

All content copied from https://docs.aws.amazon.com/.
