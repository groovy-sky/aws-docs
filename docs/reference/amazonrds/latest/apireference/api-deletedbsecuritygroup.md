# DeleteDBSecurityGroup

Deletes a DB security group.

The specified DB security group must not be associated with any DB instances.

###### Note

EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that
you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](../../../../services/ec2/latest/userguide/vpc-migrate.md) in the
_Amazon EC2 User Guide_, the blog [EC2-Classic Networking is Retiring – \
Here’s How to Prepare](http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare), and [Moving a DB instance not in a VPC \
into a VPC](../../../../services/amazonrds/latest/userguide/user-vpc-non-vpc2vpc.md) in the _Amazon RDS User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSecurityGroupName**

The name of the DB security group to delete.

###### Note

You can't delete the default DB security group.

Constraints:

- Must be 1 to 255 letters, numbers, or hyphens.

- First character must be a letter

- Can't end with a hyphen or contain two consecutive hyphens

- Must not be "Default"

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBSecurityGroupNotFound**

`DBSecurityGroupName` doesn't refer to an existing DB security group.

HTTP Status Code: 404

**InvalidDBSecurityGroupState**

The state of the DB security group doesn't allow deletion.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of DeleteDBSecurityGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DeleteDBSecurityGroup
   &DBSecurityGroupName=mydbsecuritygroup
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140423/us-east-1/rds/aws4_request
   &X-Amz-Date=20140423T203336Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=873c15061fe60b9db8ea63137e5af82b157019696fc3e9764ef2abd9d71c640a

```

#### Sample Response

```

<DeleteDBSecurityGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ResponseMetadata>
    <RequestId>7aec7454-ba25-11d3-855b-576787000e19</RequestId>
  </ResponseMetadata>
</DeleteDBSecurityGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deletedbsecuritygroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deletedbsecuritygroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBProxyEndpoint

DeleteDBShardGroup

All content copied from https://docs.aws.amazon.com/.
