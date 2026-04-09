# DescribeCacheSubnetGroups

Returns a list of cache subnet group descriptions. If a subnet group name is
specified, the list contains only the description of that group. This is applicable only
when you have ElastiCache in VPC setup. All ElastiCache clusters now launch in VPC by
default.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheSubnetGroupName**

The name of the cache subnet group to return details for.

Type: String

Required: No

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than
the specified `MaxRecords` value, a marker is included in the response so
that the remaining results can be retrieved.

Default: 100

Constraints: minimum 20; maximum 100.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**CacheSubnetGroups.CacheSubnetGroup.N**

A list of cache subnet groups. Each element in the list contains detailed information
about one group.

Type: Array of [CacheSubnetGroup](api-cachesubnetgroup.md) objects

**Marker**

Provides an identifier to allow retrieval of paginated results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheSubnetGroupNotFoundFault**

The requested cache subnet group name does not refer to an existing cache subnet
group.

HTTP Status Code: 400

## Examples

### DescribeCacheSubnetGroups

Some of the output has been omitted for brevity.

#### Sample Request

```

https://elasticache.amazonaws.com/
   ?Action=DescribeCacheSubnetGroups
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<DescribeCacheSubnetGroupsResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
        <DescribeCacheSubnetGroupsResult>
            <CacheSubnetGroups>
                <CacheSubnetGroup>
                    <VpcId>990524496922</VpcId>
                    <CacheSubnetGroupDescription>description</CacheSubnetGroupDescription>
                    <CacheSubnetGroupName>subnet_grp1</CacheSubnetGroupName>
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

 (...output omitted...)

            </CacheSubnetGroups>
        </DescribeCacheSubnetGroupsResult>
        <ResponseMetadata>
            <RequestId>31d0faee-229b-11e1-81f1-df3a2a803dad</RequestId>
        </ResponseMetadata>
    </DescribeCacheSubnetGroupsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describecachesubnetgroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describecachesubnetgroups.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCacheSecurityGroups

DescribeEngineDefaultParameters

All content copied from https://docs.aws.amazon.com/.
