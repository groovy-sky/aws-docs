# ImageScanningConfiguration

The image scanning configuration for a repository.

## Contents

**scanOnPush**

The setting that determines whether images are scanned after being pushed to a
repository. If set to `true`, images will be scanned after being pushed. If
this parameter is not specified, it will default to `false` and images will
not be scanned unless a scan is manually started with the [API\_StartImageScan](api-startimagescan.md) API.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagescanningconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagescanningconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagescanningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageScanFindingsSummary

ImageScanStatus

All content copied from https://docs.aws.amazon.com/.
