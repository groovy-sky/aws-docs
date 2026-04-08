# ModifyDBSubnetGroup

Modifies an existing DB subnet group. DB subnet groups must contain at least one subnet in at least two AZs in the AWS Region.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBSubnetGroupName**

The name for the DB subnet group. This value is stored as a lowercase string.
You can't modify the default subnet group.

Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.

Example: `mydbsubnetgroup`

Type: String

Required: Yes

**SubnetIds.SubnetIdentifier.N**

The EC2 subnet IDs for the DB subnet group.

Type: Array of strings

Required: Yes

**DBSubnetGroupDescription**

The description for the DB subnet group.

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**DBSubnetGroup**

Contains the details of an Amazon RDS DB subnet group.

This data type is used as a response element
in the `DescribeDBSubnetGroups` action.

Type: [DBSubnetGroup](api-dbsubnetgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBSubnetGroupDoesNotCoverEnoughAZs**

Subnets in the DB subnet group should cover at least two Availability Zones unless there is only one Availability Zone.

HTTP Status Code: 400

**DBSubnetGroupNotFoundFault**

`DBSubnetGroupName` doesn't refer to an existing DB subnet group.

HTTP Status Code: 404

**DBSubnetQuotaExceededFault**

The request would result in the user exceeding the allowed number of subnets in a
DB subnet groups.

HTTP Status Code: 400

**InvalidDBSubnetGroupStateFault**

The DB subnet group cannot be deleted because it's in use.

HTTP Status Code: 400

**InvalidSubnet**

The requested subnet is invalid, or multiple subnets were requested that are not all in a common VPC.

HTTP Status Code: 400

**SubnetAlreadyInUse**

The DB subnet is already in use in the Availability Zone.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of ModifyDBSubnetGroup.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=ModifyDBSubnetGroup
   &DBSubnetGroupDescription=A%20new%20Description
   &DBSubnetGroupName=myawsuser-sngrp
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SubnetIds.member.1=subnet-e4d398a1
   &SubnetIds.member.2=subnet-c2bdb6ba
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140425/us-east-1/rds/aws4_request
   &X-Amz-Date=20140425T200214Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=213c429d925cb1608fc13a1dde48715bcac3b0794536ee90beac34203265f9af

```

#### Sample Response

```

<ModifyDBSubnetGroupResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ModifyDBSubnetGroupResult>
    <DBSubnetGroup>
      <VpcId>vpc-33ac97ea</VpcId>
      <SubnetGroupStatus>Complete</SubnetGroupStatus>
      <DBSubnetGroupDescription>A new Description</DBSubnetGroupDescription>
      <DBSubnetGroupName>myawsuser-sngrp</DBSubnetGroupName>
      <Subnets>
        <Subnet>
          <SubnetStatus>Active</SubnetStatus>
          <SubnetIdentifier>subnet-e4d398a1</SubnetIdentifier>
          <SubnetAvailabilityZone>
            <Name>us-east-1b</Name>
            <ProvisionedIopsCapable>false</ProvisionedIopsCapable>
          </SubnetAvailabilityZone>
        </Subnet>
        <Subnet>
          <SubnetStatus>Active</SubnetStatus>
          <SubnetIdentifier>subnet-c2bdb6ba</SubnetIdentifier>
          <SubnetAvailabilityZone>
            <Name>us-east-1c</Name>
            <ProvisionedIopsCapable>false</ProvisionedIopsCapable>
          </SubnetAvailabilityZone>
        </Subnet>
      </Subnets>
    </DBSubnetGroup>
  </ModifyDBSubnetGroupResult>
  <ResponseMetadata>
    <RequestId>6dd028be-bba3-11d3-f4c6-37db295f7674</RequestId>
  </ResponseMetadata>
</ModifyDBSubnetGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbsubnetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbsubnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBSnapshotAttribute

ModifyEventSubscription

All content copied from https://docs.aws.amazon.com/.
