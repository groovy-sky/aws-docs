---
title: "DBSecurityGroup"
---

# DBSecurityGroup

Contains the details for an Amazon RDS DB security group.

This data type is used as a response element
in the `DescribeDBSecurityGroups` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**DBSecurityGroupArn**

The Amazon Resource Name (ARN) for the DB security group.

Type: String

Required: No

**DBSecurityGroupDescription**

Provides the description of the DB security group.

Type: String

Required: No

**DBSecurityGroupName**

Specifies the name of the DB security group.

Type: String

Required: No

**EC2SecurityGroups.EC2SecurityGroup.N**

Contains a list of `EC2SecurityGroup` elements.

Type: Array of [EC2SecurityGroup](api-ec2securitygroup.md) objects

Required: No

**IPRanges.IPRange.N**

Contains a list of `IPRange` elements.

Type: Array of [IPRange](api-iprange.md) objects

Required: No

**OwnerId**

Provides the AWS ID of the owner of a specific DB security group.

Type: String

Required: No

**VpcId**

Provides the VpcId of the DB security group.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/DBSecurityGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/DBSecurityGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/DBSecurityGroup)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBRecommendation

DBSecurityGroupMembership

All content copied from https://docs.aws.amazon.com/.
