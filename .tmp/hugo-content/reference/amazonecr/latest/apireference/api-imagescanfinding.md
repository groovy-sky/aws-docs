# ImageScanFinding

Contains information about an image scan finding.

## Contents

**attributes**

A collection of attributes of the host from which the finding is generated.

Type: Array of [Attribute](api-attribute.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**description**

The description of the finding.

Type: String

Required: No

**name**

The name associated with the finding, usually a CVE number.

Type: String

Required: No

**severity**

The finding severity.

Type: String

Valid Values: `INFORMATIONAL | LOW | MEDIUM | HIGH | CRITICAL | UNDEFINED`

Required: No

**uri**

A link containing additional details about the security vulnerability.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagescanfinding.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagescanfinding.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagescanfinding.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageReplicationStatus

ImageScanFindings

All content copied from https://docs.aws.amazon.com/.
