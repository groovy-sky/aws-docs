---
title: "ImageScanFindings"
---

# ImageScanFindings

The details of an image scan.

## Contents

**enhancedFindings**

Details about the enhanced scan findings from Amazon Inspector.

Type: Array of [EnhancedImageScanFinding](api-enhancedimagescanfinding.md) objects

Required: No

**findings**

The findings from the image scan.

Type: Array of [ImageScanFinding](api-imagescanfinding.md) objects

Required: No

**findingSeverityCounts**

The image vulnerability counts, sorted by severity.

Type: String to integer map

Valid Keys: `INFORMATIONAL | LOW | MEDIUM | HIGH | CRITICAL | UNDEFINED`

Valid Range: Minimum value of 0.

Required: No

**imageScanCompletedAt**

The time of the last completed image scan.

Type: Timestamp

Required: No

**vulnerabilitySourceUpdatedAt**

The time when the vulnerability data was last scanned.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/ImageScanFindings)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/ImageScanFindings)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/ImageScanFindings)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageScanFinding

ImageScanFindingsSummary

All content copied from https://docs.aws.amazon.com/.
