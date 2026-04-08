# SnsTopicConfiguration

The proposed access control configuration for an Amazon SNS topic. You can propose a
configuration for a new Amazon SNS topic or an existing Amazon SNS topic that you own by specifying
the policy. If the configuration is for an existing Amazon SNS topic and you do not specify the
Amazon SNS policy, then the access preview uses the existing Amazon SNS policy for the topic. If the
access preview is for a new resource and you do not specify the policy, then the access
preview assumes an Amazon SNS topic without a policy. To propose deletion of an existing Amazon SNS
topic policy, you can specify an empty string for the Amazon SNS policy. For more information,
see [Topic](../../../../services/sns/latest/api/api-topic.md).

## Contents

**topicPolicy**

The JSON policy text that defines who can access an Amazon SNS topic. For more information,
see [Example cases for Amazon SNS access control](../../../../services/sns/latest/dg/sns-access-policy-use-cases.md) in the _Amazon SNS Developer_
_Guide_.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 30720.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/snstopicconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/snstopicconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/snstopicconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecretsManagerSecretConfiguration

SortCriteria
