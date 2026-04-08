# DescribeDBSecurityGroups

Returns a list of `DBSecurityGroup` descriptions. If a `DBSecurityGroupName` is specified,
the list will contain only the descriptions of the specified DB security group.

###### Note

EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that
you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](../../../../services/ec2/latest/userguide/vpc-migrate.md) in the
_Amazon EC2 User Guide_, the blog [EC2-Classic Networking is Retiring – \
Here’s How to Prepare](http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare), and [Moving a DB instance not in a VPC \
into a VPC](../../../../services/amazonrds/latest/userguide/user-vpc-non-vpc2vpc.md) in the _Amazon RDS User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSecurityGroupName**

The name of the DB security group to return details for.

Type: String

Required: No

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeDBSecurityGroups` request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so that
you can retrieve the remaining results.

Default: 100

Constraints: Minimum 20, maximum 100.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**DBSecurityGroups.DBSecurityGroup.N**

A list of `DBSecurityGroup` instances.

Type: Array of [DBSecurityGroup](api-dbsecuritygroup.md) objects

**Marker**

An optional pagination token provided by a previous request.
If this parameter is specified, the response includes
only records beyond the marker,
up to the value specified by `MaxRecords`.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBSecurityGroupNotFound**

`DBSecurityGroupName` doesn't refer to an existing DB security group.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeDBSecurityGroups.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeDBSecurityGroups
   &MaxRecords=100
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194732Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=b14bcddedcf2fd7ffbbcc45ed2caa99cd848ee309a19070f946ad2a54f5331fe

```

#### Sample Response

```

<DescribeDBSecurityGroupsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeDBSecurityGroupsResult>
    <DBSecurityGroups>
      <DBSecurityGroup>
        <EC2SecurityGroups/>
        <DBSecurityGroupDescription>My security group</DBSecurityGroupDescription>
        <IPRanges>
          <IPRange>
            <CIDRIP>192.0.0.0/24</CIDRIP>
            <Status>authorized</Status>
          </IPRange>
          <IPRange>
            <CIDRIP>190.0.1.0/29</CIDRIP>
            <Status>authorized</Status>
          </IPRange>
          <IPRange>
            <CIDRIP>190.0.2.0/29</CIDRIP>
            <Status>authorized</Status>
          </IPRange>
          <IPRange>
            <CIDRIP>10.0.0.0/8</CIDRIP>
            <Status>authorized</Status>
          </IPRange>
        </IPRanges>
        <OwnerId>803#########</OwnerId>
        <DBSecurityGroupName>my-secgrp</DBSecurityGroupName>
      </DBSecurityGroup>
      <DBSecurityGroup>
        <EC2SecurityGroups/>
        <DBSecurityGroupDescription>default</DBSecurityGroupDescription>
        <IPRanges/>
        <OwnerId>803#########</OwnerId>
        <DBSecurityGroupName>default</DBSecurityGroupName>
      </DBSecurityGroup>
   </DBSecurityGroups>
  </DescribeDBSecurityGroupsResult>
  <ResponseMetadata>
    <RequestId>b76e692c-b98c-11d3-a907-5a2c468b9cb0</RequestId>
  </ResponseMetadata>
</DescribeDBSecurityGroupsResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describedbsecuritygroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describedbsecuritygroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDBRecommendations

DescribeDBShardGroups

All content copied from https://docs.aws.amazon.com/.
