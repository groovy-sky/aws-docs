# RevokeCacheSecurityGroupIngress

Revokes ingress from a cache security group. Use this operation to disallow access
from an Amazon EC2 security group that had been previously authorized.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheSecurityGroupName**

The name of the cache security group to revoke ingress from.

Type: String

Required: Yes

**EC2SecurityGroupName**

The name of the Amazon EC2 security group to revoke access from.

Type: String

Required: Yes

**EC2SecurityGroupOwnerId**

The Amazon account number of the Amazon EC2 security group owner. Note that this is
not the same thing as an Amazon access key ID - you must provide a valid Amazon account
number for this parameter.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**CacheSecurityGroup**

Represents the output of one of the following operations:

- `AuthorizeCacheSecurityGroupIngress`

- `CreateCacheSecurityGroup`

- `RevokeCacheSecurityGroupIngress`

Type: [CacheSecurityGroup](api-cachesecuritygroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AuthorizationNotFound**

The specified Amazon EC2 security group is not authorized for the specified cache
security group.

HTTP Status Code: 404

**CacheSecurityGroupNotFound**

The requested cache security group name does not refer to an existing cache security
group.

HTTP Status Code: 404

**InvalidCacheSecurityGroupState**

The current state of the cache security group does not allow deletion.

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

## Examples

### RevokeCacheSecurityGroupIngress

This example illustrates one usage of RevokeCacheSecurityGroupIngress.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=RevokeCacheSecurityGroupIngress
   &EC2SecurityGroupName=default
   &CacheSecurityGroupName=mygroup
   &EC2SecurityGroupOwnerId=1234-5678-1234
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<RevokeCacheSecurityGroupIngressResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
    <RevokeCacheSecurityGroupIngressResult>
        <CacheSecurityGroup>
            <EC2SecurityGroups>
                <EC2SecurityGroup>
                    <Status>revoking</Status>
                    <EC2SecurityGroupName>default</EC2SecurityGroupName>
                    <EC2SecurityGroupOwnerId>123456781234</EC2SecurityGroupOwnerId>
                </EC2SecurityGroup>
            </EC2SecurityGroups>
            <CacheSecurityGroupName>mygroup</CacheSecurityGroupName>
            <OwnerId>123456789012</OwnerId>
            <Description>My security group</Description>
        </CacheSecurityGroup>
    </RevokeCacheSecurityGroupIngressResult>
    <ResponseMetadata>
        <RequestId>02ae3699-3650-11e0-a564-8f11342c56b0</RequestId>
    </ResponseMetadata>
</RevokeCacheSecurityGroupIngressResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/revokecachesecuritygroupingress.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResetCacheParameterGroup

StartMigration

All content copied from https://docs.aws.amazon.com/.
