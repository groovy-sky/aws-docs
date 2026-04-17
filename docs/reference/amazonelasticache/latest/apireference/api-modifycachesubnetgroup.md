---
title: "ModifyCacheSubnetGroup"
---

# ModifyCacheSubnetGroup

Modifies an existing cache subnet group.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheSubnetGroupName**

The name for the cache subnet group. This value is stored as a lowercase
string.

Constraints: Must contain no more than 255 alphanumeric characters or hyphens.

Example: `mysubnetgroup`

Type: String

Required: Yes

**CacheSubnetGroupDescription**

A description of the cache subnet group.

Type: String

Required: No

**SubnetIds.SubnetIdentifier.N**

The EC2 subnet IDs for the cache subnet group.

Type: Array of strings

Required: No

## Response Elements

The following element is returned by the service.

**CacheSubnetGroup**

Represents the output of one of the following operations:

- `CreateCacheSubnetGroup`

- `ModifyCacheSubnetGroup`

Type: [CacheSubnetGroup](api-cachesubnetgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheSubnetGroupNotFoundFault**

The requested cache subnet group name does not refer to an existing cache subnet
group.

HTTP Status Code: 400

**CacheSubnetQuotaExceededFault**

The request cannot be processed because it would exceed the allowed number of subnets
in a cache subnet group.

HTTP Status Code: 400

**InvalidSubnet**

An invalid subnet identifier was specified.

HTTP Status Code: 400

**SubnetInUse**

The requested subnet is being used by another cache subnet group.

HTTP Status Code: 400

**SubnetNotAllowedFault**

At least one subnet ID does not match the other subnet IDs. This mismatch typically
occurs when a user sets one subnet ID to a regional Availability Zone and a different
one to an outpost. Or when a user sets the subnet ID to an Outpost when not subscribed
on this service.

HTTP Status Code: 400

## Examples

### ModifyCacheSubnetGroup

This example illustrates one usage of ModifyCacheSubnetGroup.

#### Sample Request

```

https://elasticache.amazonaws.com/
   ?Action=ModifyCacheSubnetGroup
   &CacheSubnetGroupName=myCachesubnetgroup
   &CacheSubnetGroupDescription=My%20modified%20CacheSubnetGroup
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<ModifyCacheSubnetGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
    <ModifyCacheSubnetGroupResult>
            <CacheSubnetGroup>
                <VpcId>990524496922</VpcId>
                <CacheSubnetGroupDescription>My modified CacheSubnetGroup</CacheSubnetGroupDescription>
                <CacheSubnetGroupName>myCachesubnetgroup</CacheSubnetGroupName>
                <Subnets>
                    <Subnet>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetIdentifier>subnet-7c5b4115</SubnetIdentifier>
                        <SubnetAvailabilityZone>
                            <Name>us-west-2c</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                    <Subnet>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetIdentifier>subnet-7b5b4112</SubnetIdentifier>
                        <SubnetAvailabilityZone>
                            <Name>us-west-2b</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                    <Subnet>
                        <SubnetStatus>Active</SubnetStatus>
                        <SubnetIdentifier>subnet-3ea6bd57</SubnetIdentifier>
                        <SubnetAvailabilityZone>
                            <Name>us-west-2c</Name>
                        </SubnetAvailabilityZone>
                    </Subnet>
                </Subnets>
            </CacheSubnetGroup>
        </ModifyCacheSubnetGroupResult>
        <ResponseMetadata>
            <RequestId>ed662948-a57b-11df-9e38-7ffab86c801f</RequestId>
        </ResponseMetadata>
    </ModifyCacheSubnetGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/elasticache-2015-02-02/ModifyCacheSubnetGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/ModifyCacheSubnetGroup)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyCacheParameterGroup

ModifyGlobalReplicationGroup

All content copied from https://docs.aws.amazon.com/.
