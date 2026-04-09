# SigningConfiguration

The signing configuration for a registry, which specifies rules
for automatically signing images when pushed.

## Contents

**rules**

A list of signing rules. Each rule defines a signing profile and optional repository
filters that determine which images are automatically signed. Maximum of 10 rules.

Type: Array of [SigningRule](api-signingrule.md) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/signingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/signingconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/signingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScoreDetails

SigningRepositoryFilter

All content copied from https://docs.aws.amazon.com/.
