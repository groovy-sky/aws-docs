# PrivateDnsNameConfiguration

Information about the private DNS name for the service endpoint.

## Contents

**name**

The name of the record subdomain the service provider needs to create. The service provider adds the `value` text to the `name`.

Type: String

Required: No

**state**

The verification state of the VPC endpoint service.

Consumers
of the endpoint service can use the private name only when the state is
`verified`.

Type: String

Valid Values: `pendingVerification | verified | failed`

Required: No

**type**

The endpoint service verification type, for example TXT.

Type: String

Required: No

**value**

The value the service provider adds to the private DNS name domain record before verification.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/privatednsnameconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/privatednsnameconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/privatednsnameconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PrivateDnsDetails

PrivateDnsNameOptionsOnLaunch
