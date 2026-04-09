# DataSourceVpcConfiguration

Provides configuration information needed to connect to an Amazon VPC (Virtual
Private Cloud).

## Contents

**securityGroupIds**

A list of identifiers of security groups within your Amazon VPC. The security
groups should enable Amazon Q Business to connect to the data source.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `[-0-9a-zA-Z]+`

Required: Yes

**subnetIds**

A list of identifiers for subnets within your Amazon VPC. The subnets should
be able to connect to each other in the VPC, and they should have outgoing access to the
Internet through a NAT device.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `[-0-9a-zA-Z]+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/datasourcevpcconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/datasourcevpcconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/datasourcevpcconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceSyncJobMetrics

DateAttributeBoostingConfiguration

All content copied from https://docs.aws.amazon.com/.
