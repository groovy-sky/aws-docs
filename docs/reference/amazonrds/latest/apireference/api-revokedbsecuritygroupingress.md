# RevokeDBSecurityGroupIngress

Revokes ingress from a DBSecurityGroup for previously authorized IP ranges or EC2 or VPC security groups. Required
parameters for this API are one of CIDRIP, EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
EC2SecurityGroupName or EC2SecurityGroupId).

###### Note

EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that
you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](../../../../services/ec2/latest/userguide/vpc-migrate.md) in the
_Amazon EC2 User Guide_, the blog [EC2-Classic Networking is Retiring – \
Here’s How to Prepare](http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare), and [Moving a DB instance not in a VPC \
into a VPC](../../../../services/amazonrds/latest/userguide/user-vpc-non-vpc2vpc.md) in the _Amazon RDS User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSecurityGroupName**

The name of the DB security group to revoke ingress from.

Type: String

Required: Yes

**CIDRIP**

The IP range to revoke access from.
Must be a valid CIDR range. If `CIDRIP` is specified,
`EC2SecurityGroupName`, `EC2SecurityGroupId` and `EC2SecurityGroupOwnerId`
can't be provided.

Type: String

Required: No

**EC2SecurityGroupId**

The id of the EC2 security group to revoke access from.
For VPC DB security groups, `EC2SecurityGroupId` must be provided.
Otherwise, EC2SecurityGroupOwnerId and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.

Type: String

Required: No

**EC2SecurityGroupName**

The name of the EC2 security group to revoke access from.
For VPC DB security groups, `EC2SecurityGroupId` must be provided.
Otherwise, EC2SecurityGroupOwnerId and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.

Type: String

Required: No

**EC2SecurityGroupOwnerId**

The AWS account number of the owner of the EC2 security group
specified in the `EC2SecurityGroupName` parameter.
The AWS access key ID isn't an acceptable value.
For VPC DB security groups, `EC2SecurityGroupId` must be provided.
Otherwise, EC2SecurityGroupOwnerId and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.

Type: String

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

**AuthorizationNotFound**

The specified CIDR IP range or Amazon EC2 security group might not be authorized
for the specified DB security group.

Or, RDS might not be authorized to perform necessary actions using IAM on your
behalf.

HTTP Status Code: 404

**DBSecurityGroupNotFound**

`DBSecurityGroupName` doesn't refer to an existing DB security group.

HTTP Status Code: 404

**InvalidDBSecurityGroupState**

The state of the DB security group doesn't allow deletion.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of RevokeDBSecurityGroupIngress.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=RevokeDBSecurityGroupIngress
   &CIDRIP=192.0.0.1%2F32
   &DBSecurityGroupName=mydbsecuritygroup01
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140428/us-east-1/rds/aws4_request
   &X-Amz-Date=20140428T233956Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=d9edabccacae36138704fb2b3cf6755ef08123862191b19d74582497b75e544a

```

#### Sample Response

```

<RevokeDBSecurityGroupIngressResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <RevokeDBSecurityGroupIngressResult>
    <DBSecurityGroup>
      <EC2SecurityGroups/>
      <DBSecurityGroupDescription>My new DBSecurityGroup</DBSecurityGroupDescription>
      <IPRanges>
        <IPRange>
          <CIDRIP>192.0.0.1/32</CIDRIP>
          <Status>revoking</Status>
        </IPRange>
      </IPRanges>
      <OwnerId>803#########</OwnerId>
      <DBSecurityGroupName>mydbsecuritygroup01</DBSecurityGroupName>
    </DBSecurityGroup>
  </RevokeDBSecurityGroupIngressResult>
  <ResponseMetadata>
    <RequestId>579d8ba0-be2d-11d3-ae4f-eec568ed6b36</RequestId>
  </ResponseMetadata>
</RevokeDBSecurityGroupIngressResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/revokedbsecuritygroupingress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/revokedbsecuritygroupingress.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreDBInstanceToPointInTime

StartActivityStream

All content copied from https://docs.aws.amazon.com/.
