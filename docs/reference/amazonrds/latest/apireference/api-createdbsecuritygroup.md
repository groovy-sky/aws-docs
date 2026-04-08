# CreateDBSecurityGroup

Creates a new DB security group. DB security groups control access to a DB instance.

A DB security group controls access to EC2-Classic DB instances that are not in a VPC.

###### Note

EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that
you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](../../../../services/ec2/latest/userguide/vpc-migrate.md) in the
_Amazon EC2 User Guide_, the blog [EC2-Classic Networking is Retiring – \
Here’s How to Prepare](http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare), and [Moving a DB instance not in a VPC \
into a VPC](../../../../services/amazonrds/latest/userguide/user-vpc-non-vpc2vpc.md) in the _Amazon RDS User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSecurityGroupDescription**

The description for the DB security group.

Type: String

Required: Yes

**DBSecurityGroupName**

The name for the DB security group. This value is stored as a lowercase string.

Constraints:

- Must be 1 to 255 letters, numbers, or hyphens.

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

- Must not be "Default"

Example: `mysecuritygroup`

Type: String

Required: Yes

**Tags.Tag.N**

Tags to assign to the DB security group.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**DBSecurityGroup**

Contains the details for an Amazon RDS DB security group.

This data type is used as a response element
in the `DescribeDBSecurityGroups` action.

Type: [DBSecurityGroup](api-dbsecuritygroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBSecurityGroupAlreadyExists**

A DB security group with the name specified in
`DBSecurityGroupName` already exists.

HTTP Status Code: 400

**DBSecurityGroupNotSupported**

A DB security group isn't allowed for this action.

HTTP Status Code: 400

**QuotaExceeded.DBSecurityGroup**

The request would result in the user exceeding the allowed number of DB security
groups.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of CreateDBSecurityGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=CreateDBSecurityGroup
   &DBSecurityGroupDescription=My%20new%20DB%20Security%20Group
   &DBSecurityGroupName=mydbsecuritygroup00
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140424/us-east-1/rds/aws4_request
   &X-Amz-Date=20140424T190716Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=c2f180a3f0f5d73b47f9c229937a78f3569bf14392db8093d9b2e6785609ab45

```

#### Sample Response

```

<CreateDBSecurityGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <CreateDBSecurityGroupResult>
    <DBSecurityGroup>
      <EC2SecurityGroups/>
      <DBSecurityGroupDescription>My new DB Security Group</DBSecurityGroupDescription>
      <IPRanges/>
      <OwnerId>803#########</OwnerId>
      <DBSecurityGroupName>mydbsecuritygroup00</DBSecurityGroupName>
    </DBSecurityGroup>
  </CreateDBSecurityGroupResult>
  <ResponseMetadata>
    <RequestId>e68ef6fa-afc1-11c3-845a-476777009d19</RequestId>
  </ResponseMetadata>
</CreateDBSecurityGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/createdbsecuritygroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/createdbsecuritygroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDBProxyEndpoint

CreateDBShardGroup

All content copied from https://docs.aws.amazon.com/.
