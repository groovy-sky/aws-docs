# EC2SecurityGroup

This data type is used as a response element in the following actions:

- `AuthorizeDBSecurityGroupIngress`

- `DescribeDBSecurityGroups`

- `RevokeDBSecurityGroupIngress`

## Contents

###### Note

In the following list, the required parameters are described first.

**EC2SecurityGroupId**

Specifies the id of the EC2 security group.

Type: String

Required: No

**EC2SecurityGroupName**

Specifies the name of the EC2 security group.

Type: String

Required: No

**EC2SecurityGroupOwnerId**

Specifies the AWS ID of the owner of the EC2 security group
specified in the `EC2SecurityGroupName` field.

Type: String

Required: No

**Status**

Provides the status of the EC2 security group. Status can be "authorizing", "authorized", "revoking", and "revoked".

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/ec2securitygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/ec2securitygroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/ec2securitygroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DoubleRange

Endpoint

All content copied from https://docs.aws.amazon.com/.
