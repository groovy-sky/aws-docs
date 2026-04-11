# ActiveInstance

Describes a running instance in a Spot Fleet.

## Contents

**instanceHealth**

The health status of the instance. If the status of either the instance status check
or the system status check is `impaired`, the health status of the instance
is `unhealthy`. Otherwise, the health status is `healthy`.

Type: String

Valid Values: `healthy | unhealthy`

Required: No

**instanceId**

The ID of the instance.

Type: String

Required: No

**instanceType**

The instance type.

Type: String

Required: No

**spotInstanceRequestId**

The ID of the Spot Instance request.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/activeinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/activeinstance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/activeinstance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AccountAttributeValue

ActiveVpnTunnelStatus
