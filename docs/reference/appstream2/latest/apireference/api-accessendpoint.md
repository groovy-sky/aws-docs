# AccessEndpoint

Describes an interface VPC endpoint (interface endpoint) that lets you create a private connection between the virtual private cloud (VPC) that you specify and WorkSpaces Applications. When you specify an interface endpoint for a stack, users of the stack can connect to WorkSpaces Applications only through that endpoint. When you specify an interface endpoint for an image builder, administrators can connect to the image builder only through that endpoint.

## Contents

**EndpointType**

The type of interface endpoint.

Type: String

Valid Values: `STREAMING`

Required: Yes

**VpceId**

The identifier (ID) of the VPC in which the interface endpoint is used.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/accessendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/accessendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/accessendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

AdminAppLicenseUsageRecord

All content copied from https://docs.aws.amazon.com/.
