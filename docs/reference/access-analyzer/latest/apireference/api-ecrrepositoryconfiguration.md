# EcrRepositoryConfiguration

The proposed access control configuration for an Amazon ECR repository. You can propose a
configuration for a new Amazon ECR repository or an existing Amazon ECR repository that you own by
specifying the Amazon ECR policy. For more information, see [Repository](../../../amazonecr/latest/apireference/api-repository.md).

- If the configuration is for an existing Amazon ECR repository and you do not specify
the Amazon ECR policy, then the access preview uses the existing Amazon ECR policy for the
repository.

- If the access preview is for a new resource and you do not specify the policy,
then the access preview assumes an Amazon ECR repository without a policy.

- To propose deletion of an existing Amazon ECR repository policy, you can specify an
empty string for the Amazon ECR policy.

## Contents

**repositoryPolicy**

The JSON repository policy text to apply to the Amazon ECR repository. For more information,
see [Private repository\
policy examples](../../../../services/amazonecr/latest/userguide/repository-policy-examples.md) in the _Amazon ECR User Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/ecrrepositoryconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/ecrrepositoryconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/ecrrepositoryconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EbsSnapshotConfiguration

EfsFileSystemConfiguration
