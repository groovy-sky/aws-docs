# Image

Describes an image.

## Contents

**Name**

The name of the image.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**Applications**

The applications associated with the image.

Type: Array of [Application](api-application.md) objects

Required: No

**AppstreamAgentVersion**

The version of the WorkSpaces Applications agent to use for instances that are launched from this image.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**Arn**

The ARN of the image.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**BaseImageArn**

The ARN of the image from which this image was created.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**CreatedTime**

The time the image was created.

Type: Timestamp

Required: No

**Description**

The description to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**DisplayName**

The image name to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**DynamicAppProvidersEnabled**

Indicates whether dynamic app providers are enabled within an WorkSpaces Applications image or not.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**ImageBuilderName**

The name of the image builder that was used to create the private image. If the image is shared, copied, or updated by using Managed Image Updates, this value is null.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**ImageBuilderSupported**

Indicates whether an image builder can be launched from this image.

Type: Boolean

Required: No

**ImageErrors**

Describes the errors that are returned when a new image can't be created.

Type: Array of [ResourceError](api-resourceerror.md) objects

Required: No

**ImagePermissions**

The permissions to provide to the destination AWS account for the specified image.

Type: [ImagePermissions](api-imagepermissions.md) object

Required: No

**ImageSharedWithOthers**

Indicates whether the image is shared with another account ID.

Type: String

Valid Values: `TRUE | FALSE`

Required: No

**ImageType**

The type of the image. Images created through AMI import have type "custom", while WorkSpaces Applications provided images have type "native". Custom images support additional instance types including GeneralPurpose, MemoryOptimized, ComputeOptimized, and Accelerated instance families.

Type: String

Valid Values: `CUSTOM | NATIVE`

Required: No

**LatestAppstreamAgentVersion**

Indicates whether the image is using the latest WorkSpaces Applications agent version or not.

Type: String

Valid Values: `TRUE | FALSE`

Required: No

**ManagedSoftwareIncluded**

Indicates whether the image includes license-included applications.

Type: Boolean

Required: No

**Platform**

The operating system platform of the image.

Type: String

Valid Values: `WINDOWS | WINDOWS_SERVER_2016 | WINDOWS_SERVER_2019 | WINDOWS_SERVER_2022 | WINDOWS_SERVER_2025 | AMAZON_LINUX2 | RHEL8 | ROCKY_LINUX8 | UBUNTU_PRO_2404`

Required: No

**PublicBaseImageReleasedDate**

The release date of the public base image.
For private images, this date is the release date of the base image from which the image was created.

Type: Timestamp

Required: No

**State**

The image starts in the `PENDING` state. If image creation succeeds, the
state is `AVAILABLE`. If image creation fails, the state is `FAILED`.

Type: String

Valid Values: `PENDING | AVAILABLE | FAILED | COPYING | DELETING | CREATING | IMPORTING | VALIDATING`

Required: No

**StateChangeReason**

The reason why the last state change occurred.

Type: [ImageStateChangeReason](api-imagestatechangereason.md) object

Required: No

**SupportedInstanceFamilies**

The supported instances families that determine which image a customer can use when the customer launches a fleet or image builder. The following instances families are supported:

- General Purpose

- Compute Optimized

- Memory Optimized

- Graphics G4

- Graphics G5

- Graphics G6

Type: Array of strings

Length Constraints: Minimum length of 1.

Required: No

**Visibility**

Indicates whether the image is public or private.

Type: String

Valid Values: `PUBLIC | PRIVATE | SHARED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/image.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/image.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/image.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FleetError

ImageBuilder

All content copied from https://docs.aws.amazon.com/.
