# Subnet

This data type is used as a response element for the `DescribeDBSubnetGroups` operation.

## Contents

###### Note

In the following list, the required parameters are described first.

**SubnetAvailabilityZone**

Contains Availability Zone information.

This data type is used as an element in the `OrderableDBInstanceOption`
data type.

Type: [AvailabilityZone](api-availabilityzone.md) object

Required: No

**SubnetIdentifier**

The identifier of the subnet.

Type: String

Required: No

**SubnetOutpost**

If the subnet is associated with an Outpost, this value specifies the Outpost.

For more information about RDS on Outposts, see [Amazon RDS on AWS Outposts](../../../../services/amazonrds/latest/userguide/rds-on-outposts.md)
in the _Amazon RDS User Guide._

Type: [Outpost](api-outpost.md) object

Required: No

**SubnetStatus**

The status of the subnet.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/subnet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/subnet.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/subnet.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceRegion

SupportedEngineLifecycle

All content copied from https://docs.aws.amazon.com/.
