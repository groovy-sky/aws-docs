# PullThroughCacheRule

The details of a pull through cache rule.

## Contents

**createdAt**

The date and time the pull through cache was created.

Type: Timestamp

Required: No

**credentialArn**

The ARN of the Secrets Manager secret associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 612.

Pattern: `^arn:aws(-\w+)*:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\-pullthroughcache\/[a-zA-Z0-9\/_+=.@-]+$`

Required: No

**customRoleArn**

The ARN of the IAM role associated with the pull through cache rule.

Type: String

Length Constraints: Maximum length of 2048.

Required: No

**ecrRepositoryPrefix**

The Amazon ECR repository prefix associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: No

**registryId**

The AWS account ID associated with the registry the pull through cache rule is
associated with.

Type: String

Pattern: `[0-9]{12}`

Required: No

**updatedAt**

The date and time, in JavaScript date format, when the pull through cache rule was
last updated.

Type: Timestamp

Required: No

**upstreamRegistry**

The name of the upstream source registry associated with the pull through cache
rule.

Type: String

Valid Values: `ecr | ecr-public | quay | k8s | docker-hub | github-container-registry | azure-container-registry | gitlab-container-registry`

Required: No

**upstreamRegistryUrl**

The upstream registry URL associated with the pull through cache rule.

Type: String

Required: No

**upstreamRepositoryPrefix**

The upstream repository prefix associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/pullthroughcacherule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/pullthroughcacherule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/pullthroughcacherule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PackageVulnerabilityDetails

Recommendation

All content copied from https://docs.aws.amazon.com/.
