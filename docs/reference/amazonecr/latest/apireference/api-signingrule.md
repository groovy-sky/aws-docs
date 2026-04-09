# SigningRule

A signing rule that specifies a signing profile and optional
repository filters. When an image is pushed to a matching repository, a
signing job is created using the specified profile.

## Contents

**signingProfileArn**

The ARN of the AWS Signer signing profile to use for signing images that match this
rule. For more information about signing profiles, see [Signing profiles](../../../../services/signer/latest/developerguide/signing-profiles.md) in
the _AWS Signer Developer Guide_.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `^arn:aws(-[a-z]+)*:signer:[a-z0-9-]+:[0-9]{12}:\/signing-profiles\/[a-zA-Z0-9_]{2,}$`

Required: Yes

**repositoryFilters**

A list of repository filters that determine which repositories
have their images signed on push. If no filters are specified, all
images pushed to the registry are signed using the rule's signing
profile. Maximum of 100 filters per rule.

Type: Array of [SigningRepositoryFilter](api-signingrepositoryfilter.md) objects

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/signingrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/signingrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/signingrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SigningRepositoryFilter

SubjectIdentifier

All content copied from https://docs.aws.amazon.com/.
