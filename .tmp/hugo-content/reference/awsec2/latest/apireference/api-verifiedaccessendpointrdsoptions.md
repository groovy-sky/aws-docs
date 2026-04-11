# VerifiedAccessEndpointRdsOptions

Describes the RDS options for a Verified Access endpoint.

## Contents

**port**

The port.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 65535.

Required: No

**protocol**

The protocol.

Type: String

Valid Values: `http | https | tcp`

Required: No

**rdsDbClusterArn**

The ARN of the DB cluster.

Type: String

Required: No

**rdsDbInstanceArn**

The ARN of the RDS instance.

Type: String

Required: No

**rdsDbProxyArn**

The ARN of the RDS proxy.

Type: String

Required: No

**rdsEndpoint**

The RDS endpoint.

Type: String

Required: No

**SubnetIdSet.N**

The IDs of the subnets.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccessendpointrdsoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccessendpointrdsoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccessendpointrdsoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessEndpointPortRange

VerifiedAccessEndpointStatus
