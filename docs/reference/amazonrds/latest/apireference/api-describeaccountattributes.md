# DescribeAccountAttributes

Lists all of the attributes for a customer account. The attributes include Amazon RDS quotas for the account, such as the number of DB instances allowed. The description for a quota includes the quota name, current usage toward that quota, and the quota's maximum value.

This command doesn't take any parameters.

## Response Elements

The following element is returned by the service.

**AccountQuotas.AccountQuota.N**

A list of `AccountQuota` objects. Within this list, each quota has a name,
a count of usage toward the quota maximum, and a maximum value for the quota.

Type: Array of [AccountQuota](api-accountquota.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Example

This example illustrates one usage of DescribeAccountAttributes.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
        ?Action=DescribeAccountAttributes
        &SignatureMethod=HmacSHA256
        &SignatureVersion=4
        &Version=2014-10-31
        &X-Amz-Algorithm=AWS4-HMAC-SHA256
        &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20141216/us-west-2/rds/aws4_request
        &X-Amz-Date=20141216T192233Z
        &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
        &X-Amz-Signature=b49545dd3c933bdded80655d433d84bf743261ea1bebb33a7922c5c2c5240cd8

```

#### Sample Response

```

<DescribeAccountAttributesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeAccountAttributesResult>
    <AccountQuotaList>
      <AccountQuota>
        <AccountQuotaName>DBInstances</AccountQuotaName>
        <Used>22</Used>
        <Max>40</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>ReservedDBInstances</AccountQuotaName>
        <Used>6</Used>
        <Max>40</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>AllocatedStorage</AccountQuotaName>
        <Used>27459</Used>
        <Max>100000</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>DBSecurityGroupsPerVPC</AccountQuotaName>
        <Used>11</Used>
        <Max>25</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>AuthorizationsPerDBSecurityGroup</AccountQuotaName>
        <Used>10</Used>
        <Max>20</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>DBParameterGroups</AccountQuotaName>
        <Used>40</Used>
        <Max>50</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>ManualSnapshots</AccountQuotaName>
        <Used>32</Used>
        <Max>50</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>EventSubscriptions</AccountQuotaName>
        <Used>3</Used>
        <Max>20</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>DBSubnetGroups</AccountQuotaName>
        <Used>19</Used>
        <Max>20</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>OptionGroups</AccountQuotaName>
        <Used>14</Used>
        <Max>20</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>SubnetsPerDBSubnetGroup</AccountQuotaName>
        <Used>6</Used>
        <Max>20</Max
      </AccountQuota>
      <AccountQuota>
        <AccountQuotaName>ReadReplicasPerMaster</AccountQuotaName>
        <Used>2</Used>
        <Max>5</Max
      </AccountQuota>
      <AccountQuota>
        <Used>1</Used>
        <AccountQuotaName>DBClusterRoles</AccountQuotaName>
        <Max>5</Max>
      </AccountQuota>
    </AccountQuotaList>
  </DescribeAccountAttributesResult>
  <ResponseMetadata>
    <RequestId>0ce48079-68e4-11de-8c8e-eb648410240d</RequestId>
  </ResponseMetadata>
</DescribeAccountAttributesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describeaccountattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describeaccountattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeregisterDBProxyTargets

DescribeBlueGreenDeployments

All content copied from https://docs.aws.amazon.com/.
