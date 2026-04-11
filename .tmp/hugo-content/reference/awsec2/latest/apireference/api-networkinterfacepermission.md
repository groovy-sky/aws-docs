# NetworkInterfacePermission

Describes a permission for a network interface.

## Contents

**awsAccountId**

The AWS account ID.

Type: String

Required: No

**awsService**

The AWS service.

Type: String

Required: No

**networkInterfaceId**

The ID of the network interface.

Type: String

Required: No

**networkInterfacePermissionId**

The ID of the network interface permission.

Type: String

Required: No

**permission**

The type of permission.

Type: String

Valid Values: `INSTANCE-ATTACH | EIP-ASSOCIATE`

Required: No

**permissionState**

Information about the state of the permission.

Type: [NetworkInterfacePermissionState](api-networkinterfacepermissionstate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinterfacepermission.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinterfacepermission.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinterfacepermission.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInterfaceIpv6Address

NetworkInterfacePermissionState
