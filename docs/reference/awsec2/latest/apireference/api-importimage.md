# ImportImage

###### Note

To import your virtual machines (VMs) with a console-based experience, you can use the
_Import virtual machine images to AWS_ template in the [Migration Hub Orchestrator console](https://console.aws.amazon.com/migrationhub/orchestrator). For more
information, see the [_AWS Migration Hub Orchestrator User Guide_](https://docs.aws.amazon.com/migrationhub-orchestrator/latest/userguide/import-vm-images.html).

Import single or multi-volume disk images or EBS snapshots into an Amazon Machine Image (AMI).

###### Important

AWS VM Import/Export strongly recommends specifying a value for either the
`--license-type` or `--usage-operation` parameter when you create a new
VM Import task. This ensures your operating system is licensed appropriately and your billing is
optimized.

For more information, see [Importing a \
VM as an image using VM Import/Export](https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html) in the _VM Import/Export User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Architecture**

The architecture of the virtual machine.

Valid values: `i386` \| `x86_64`

Type: String

Required: No

**BootMode**

The boot mode of the virtual machine.

###### Note

The `uefi-preferred` boot mode isn't supported for importing images. For more
information, see [Boot modes](https://docs.aws.amazon.com/vm-import/latest/userguide/prerequisites.html#vmimport-boot-modes) in
the _VM Import/Export User Guide_.

Type: String

Valid Values: `legacy-bios | uefi | uefi-preferred`

Required: No

**ClientData**

The client-specific data.

Type: [ClientData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ClientData.html) object

Required: No

**ClientToken**

The token to enable idempotency for VM import requests.

Type: String

Required: No

**Description**

A description string for the import image task.

Type: String

Required: No

**DiskContainer.N**

Information about the disk containers.

Type: Array of [ImageDiskContainer](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImageDiskContainer.html) objects

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Encrypted**

Specifies whether the destination AMI of the imported image should be encrypted. The default KMS key for EBS is used
unless you specify a non-default KMS key using `KmsKeyId`. For more information, see [Amazon EBS Encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
_Amazon Elastic Compute Cloud User Guide_.

Type: Boolean

Required: No

**Hypervisor**

The target hypervisor platform.

Valid values: `xen`

Type: String

Required: No

**KmsKeyId**

An identifier for the symmetric KMS key to use when creating the
encrypted AMI. This parameter is only required if you want to use a non-default KMS key; if this
parameter is not specified, the default KMS key for EBS is used. If a `KmsKeyId` is
specified, the `Encrypted` flag must also be set.

The KMS key identifier may be provided in any of the following formats:

- Key ID

- Key alias

- ARN using key ID. The ID ARN contains the `arn:aws:kms` namespace, followed by the Region of the key, the AWS account ID of the key owner, the `key` namespace, and then the key ID. For example, arn:aws:kms: _us-east-1_: _012345678910_:key/ _abcd1234-a123-456a-a12b-a123b4cd56ef_.

- ARN using key alias. The alias ARN contains the `arn:aws:kms` namespace, followed by the Region of the key, the AWS account ID of the key owner, the `alias` namespace, and then the key alias. For example, arn:aws:kms: _us-east-1_: _012345678910_:alias/ _ExampleAlias_.

AWS parses `KmsKeyId` asynchronously, meaning that the action you call may appear to complete even
though you provided an invalid identifier. This action will eventually report failure.

The specified KMS key must exist in the Region that the AMI is being copied to.

Amazon EBS does not support asymmetric KMS keys.

Type: String

Required: No

**LicenseSpecifications.N**

The ARNs of the license configurations.

Type: Array of [ImportImageLicenseConfigurationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportImageLicenseConfigurationRequest.html) objects

Required: No

**LicenseType**

The license type to be used for the Amazon Machine Image (AMI) after importing.

Specify `AWS` to replace the source-system license with an AWS
license or `BYOL` to retain the source-system license. Leaving this parameter
undefined is the same as choosing `AWS` when importing a Windows Server operating
system, and the same as choosing `BYOL` when importing a Windows client operating
system (such as Windows 10) or a Linux operating system.

To use `BYOL`, you must have existing licenses with rights to use these licenses in a third party
cloud, such as AWS. For more information, see [Prerequisites](https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html#prerequisites-image) in the
VM Import/Export User Guide.

Type: String

Required: No

**Platform**

The operating system of the virtual machine. If you import a VM that is compatible with
Unified Extensible Firmware Interface (UEFI) using an EBS snapshot, you must specify a value for
the platform.

Valid values: `Windows` \| `Linux`

Type: String

Required: No

**RoleName**

The name of the role to use when not using the default role, 'vmimport'.

Type: String

Required: No

**TagSpecification.N**

The tags to apply to the import image task during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**UsageOperation**

The usage operation value. For more information, see [Licensing options](https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#prerequisites) in the _VM Import/Export User Guide_.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**architecture**

The architecture of the virtual machine.

Type: String

**description**

A description of the import task.

Type: String

**encrypted**

Indicates whether the AMI is encrypted.

Type: Boolean

**hypervisor**

The target hypervisor of the import task.

Type: String

**imageId**

The ID of the Amazon Machine Image (AMI) created by the import task.

Type: String

**importTaskId**

The task ID of the import image task.

Type: String

**kmsKeyId**

The identifier for the symmetric KMS key that was used to create the encrypted AMI.

Type: String

**licenseSpecifications**

The ARNs of the license configurations.

Type: Array of [ImportImageLicenseConfigurationResponse](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportImageLicenseConfigurationResponse.html) objects

**licenseType**

The license type of the virtual machine.

Type: String

**platform**

The operating system of the virtual machine.

Type: String

**progress**

The progress of the task.

Type: String

**requestId**

The ID of the request.

Type: String

**snapshotDetailSet**

Information about the snapshots.

Type: Array of [SnapshotDetail](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SnapshotDetail.html) objects

**status**

A brief status of the task.

Type: String

**statusMessage**

A detailed status message of the import task.

Type: String

**tagSet**

Any tags assigned to the import image task.

Type: Array of [Tag](api-tag.md) objects

**usageOperation**

The usage operation value.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ImportImage)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ImportImage)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ImportImage)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ImportImage)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ImportImage)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ImportImage)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ImportImage)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ImportImage)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ImportImage)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ImportImage)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImportClientVpnClientCertificateRevocationList

ImportInstance
