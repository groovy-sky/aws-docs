# TargetNetwork

Describes a target network associated with a Client VPN endpoint.

## Contents

**associationId**

The ID of the association.

Type: String

Required: No

**clientVpnEndpointId**

The ID of the Client VPN endpoint with which the target network is associated.

Type: String

Required: No

**SecurityGroups.N**

The IDs of the security groups applied to the target network association.

Type: Array of strings

Required: No

**status**

The current state of the target network association.

Type: [AssociationStatus](api-associationstatus.md) object

Required: No

**targetNetworkId**

The ID of the subnet specified as the target network.

Type: String

Required: No

**vpcId**

The ID of the VPC in which the target network (subnet) is located.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TargetNetwork)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TargetNetwork)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TargetNetwork)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetGroupsConfig

TargetReservationValue
