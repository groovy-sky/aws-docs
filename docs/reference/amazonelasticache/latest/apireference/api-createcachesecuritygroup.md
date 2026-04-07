# CreateCacheSecurityGroup

Creates a new cache security group. Use a cache security group to control access to
one or more clusters.

Cache security groups are only used when you are creating a cluster outside of an
Amazon Virtual Private Cloud (Amazon VPC). If you are creating a cluster inside of a
VPC, use a cache subnet group instead. For more information, see [CreateCacheSubnetGroup](api-createcachesubnetgroup.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/CommonParameters.html).

**CacheSecurityGroupName**

A name for the cache security group. This value is stored as a lowercase
string.

Constraints: Must contain no more than 255 alphanumeric characters. Cannot be the word
"Default".

Example: `mysecuritygroup`

Type: String

Required: Yes

**Description**

A description for the cache security group.

Type: String

Required: Yes

**Tags.Tag.N**

A list of tags to be added to this resource. A tag is a key-value pair. A tag key must
be accompanied by a tag value, although null is accepted.

Type: Array of [Tag](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Tag.html) objects

Required: No

## Response Elements

The following element is returned by the service.

**CacheSecurityGroup**

Represents the output of one of the following operations:

- `AuthorizeCacheSecurityGroupIngress`

- `CreateCacheSecurityGroup`

- `RevokeCacheSecurityGroupIngress`

Type: [CacheSecurityGroup](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheSecurityGroup.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/CommonErrors.html).

**CacheSecurityGroupAlreadyExists**

A cache security group with the specified name already exists.

HTTP Status Code: 400

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**QuotaExceeded.CacheSecurityGroup**

The request cannot be processed because it would exceed the allowed number of cache
security groups.

HTTP Status Code: 400

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

## Examples

### CreateCacheSecurityGroup

This example illustrates one usage of CreateCacheSecurityGroup.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=CreateCacheSecurityGroup
   &CacheSecurityGroupName=mycachesecuritygroup
   &Description=My%20cache%20security%20group
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<CreateCacheSecurityGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <CreateCacheSecurityGroupResult>
      <CacheSecurityGroup>
         <EC2SecurityGroups/>
         <CacheSecurityGroupName>mycachesecuritygroup</CacheSecurityGroupName>
         <OwnerId>123456789012</OwnerId>
         <Description>My cache security group</Description>
      </CacheSecurityGroup>
   </CreateCacheSecurityGroupResult>
   <ResponseMetadata>
      <RequestId>2b1c8035-b7fa-11e0-9326-b7275b9d4a6c</RequestId>
   </ResponseMetadata>
</CreateCacheSecurityGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/elasticache-2015-02-02/CreateCacheSecurityGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/CreateCacheSecurityGroup)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateCacheParameterGroup

CreateCacheSubnetGroup
