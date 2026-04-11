# Access

Contains information about actions and resources that define permissions to check
against a policy.

## Contents

**actions**

A list of actions for the access permissions. Any strings that can be used as an action
in an IAM policy can be used in the list of actions to check.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: No

**resources**

A list of resources for the access permissions. Any strings that can be used as an
Amazon Resource Name (ARN) in an IAM policy can be used in the list of resources to
check. You can only use a wildcard in the portion of the ARN that specifies the resource
ID.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Length Constraints: Minimum length of 0. Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/access.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/access.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/access.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Data Types

AccessPreview
