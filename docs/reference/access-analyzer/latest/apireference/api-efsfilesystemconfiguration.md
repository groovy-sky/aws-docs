# EfsFileSystemConfiguration

The proposed access control configuration for an Amazon EFS file system. You can propose a
configuration for a new Amazon EFS file system or an existing Amazon EFS file system that you own by
specifying the Amazon EFS policy. For more information, see [Using file systems in Amazon EFS](../../../../services/efs/latest/ug/using-fs.md).

- If the configuration is for an existing Amazon EFS file system and you do not specify
the Amazon EFS policy, then the access preview uses the existing Amazon EFS policy for the file
system.

- If the access preview is for a new resource and you do not specify the policy,
then the access preview assumes an Amazon EFS file system without a policy.

- To propose deletion of an existing Amazon EFS file system policy, you can specify an
empty string for the Amazon EFS policy.

## Contents

**fileSystemPolicy**

The JSON policy definition to apply to the Amazon EFS file system. For more information on
the elements that make up a file system policy, see [Amazon EFS Resource-based policies](../../../../services/efs/latest/ug/access-control-overview-access-control-manage-access-intro-resource-policies.md).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/efsfilesystemconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/efsfilesystemconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/efsfilesystemconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EcrRepositoryConfiguration

ExternalAccessDetails
