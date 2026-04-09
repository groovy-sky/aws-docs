# NetworkAccessConfiguration

Describes the network details of the fleet or image builder instance.

## Contents

**EniId**

The resource identifier of the elastic network interface that is attached to instances in your VPC. All network interfaces have the eni-xxxxxxxx resource identifier.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**EniIpv6Addresses**

The IPv6 addresses assigned to the elastic network interface. This field supports IPv6 connectivity for WorkSpaces Applications instances.

Type: Array of strings

Length Constraints: Minimum length of 1.

Required: No

**EniPrivateIpAddress**

The private IP address of the elastic network interface that is attached to instances in your VPC.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/networkaccessconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/networkaccessconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/networkaccessconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LastReportGenerationExecutionError

ResourceError

All content copied from https://docs.aws.amazon.com/.
