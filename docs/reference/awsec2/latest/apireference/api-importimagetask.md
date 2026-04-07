# ImportImageTask

Describes an import image task.

## Contents

**architecture**

The architecture of the virtual machine.

Valid values: `i386` \| `x86_64` \| `arm64`

Type: String

Required: No

**bootMode**

The boot mode of the virtual machine.

Type: String

Valid Values: `legacy-bios | uefi | uefi-preferred`

Required: No

**description**

A description of the import task.

Type: String

Required: No

**encrypted**

Indicates whether the image is encrypted.

Type: Boolean

Required: No

**hypervisor**

The target hypervisor for the import task.

Valid values: `xen`

Type: String

Required: No

**imageId**

The ID of the Amazon Machine Image (AMI) of the imported virtual machine.

Type: String

Required: No

**importTaskId**

The ID of the import image task.

Type: String

Required: No

**kmsKeyId**

The identifier for the KMS key that was used to create the encrypted image.

Type: String

Required: No

**LicenseSpecifications.N**

The ARNs of the license configurations that are associated with the import image task.

Type: Array of [ImportImageLicenseConfigurationResponse](api-importimagelicenseconfigurationresponse.md) objects

Required: No

**licenseType**

The license type of the virtual machine.

Type: String

Required: No

**platform**

The description string for the import image task.

Type: String

Required: No

**progress**

The percentage of progress of the import image task.

Type: String

Required: No

**SnapshotDetailSet.N**

Information about the snapshots.

Type: Array of [SnapshotDetail](api-snapshotdetail.md) objects

Required: No

**status**

A brief status for the import image task.

Type: String

Required: No

**statusMessage**

A descriptive status message for the import image task.

Type: String

Required: No

**TagSet.N**

The tags for the import image task.

Type: Array of [Tag](api-tag.md) objects

Required: No

**usageOperation**

The usage operation value.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/importimagetask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/importimagetask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/importimagetask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImportImageLicenseConfigurationResponse

ImportInstanceLaunchSpecification
