# DescribeBlueGreenDeployments

Describes one or more blue/green deployments.

For more information, see [Using Amazon RDS Blue/Green Deployments \
for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User Guide_ and
[Using Amazon RDS Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**BlueGreenDeploymentIdentifier**

The blue/green deployment identifier. If you specify this parameter, the response only
includes information about the specific blue/green deployment. This parameter isn't
case-sensitive.

Constraints:

- Must match an existing blue/green deployment identifier.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][0-9A-Za-z-:._]*`

Required: No

**Filters.Filter.N**

A filter that specifies one or more blue/green deployments to describe.

Valid Values:

- `blue-green-deployment-identifier` \- Accepts system-generated
identifiers for blue/green deployments. The results list only includes
information about the blue/green deployments with the specified
identifiers.

- `blue-green-deployment-name` \- Accepts user-supplied names for blue/green deployments.
The results list only includes information about the blue/green deployments with the
specified names.

- `source` \- Accepts source databases for a blue/green deployment.
The results list only includes information about the blue/green deployments with
the specified source databases.

- `target` \- Accepts target databases for a blue/green deployment.
The results list only includes information about the blue/green deployments with
the specified target databases.

Type: Array of [Filter](api-filter.md) objects

Required: No

**Marker**

An optional pagination token provided by a previous
`DescribeBlueGreenDeployments` request. If you specify this parameter,
the response only includes records beyond the marker, up to the value specified by
`MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response.
If more records exist than the specified `MaxRecords` value,
a pagination token called a marker is included in the response so you can retrieve the remaining results.

Default: 100

Constraints:

- Must be a minimum of 20.

- Can't exceed 100.

Type: Integer

Valid Range: Minimum value of 20. Maximum value of 100.

Required: No

## Response Elements

The following elements are returned by the service.

**BlueGreenDeployments.member.N**

A list of blue/green deployments in the current account and AWS Region.

Type: Array of [BlueGreenDeployment](api-bluegreendeployment.md) objects

**Marker**

A pagination token that can be used in a later
`DescribeBlueGreenDeployments` request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BlueGreenDeploymentNotFoundFault**

`BlueGreenDeploymentIdentifier` doesn't refer to an existing blue/green deployment.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DescribeBlueGreenDeployments.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeBlueGreenDeployments
   &BlueGreenDeploymentIdentifier=bgd-clyvb1zv1geqensv
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20141031/us-west-2/rds/aws4_request
   &X-Amz-Date=20230110T005253Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=8a684aebe6d5219bb3572316a341963324d6ef339bd0dcfa5854f1a01d401214
```

#### Sample Response

```

<DescribeBlueGreenDeploymentsResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeBlueGreenDeploymentsResult>
    <BlueGreenDeployments>
      <member>
        <TagList/>
        <BlueGreenDeploymentName>my-blue-green-deployment</BlueGreenDeploymentName>
        <CreateTime>2023-01-10T20:08:48.940Z</CreateTime>
        <SwitchoverDetails>
          <member>
            <SourceMember>arn:aws:rds:us-west-2:123456789012:db:database-1</SourceMember>
            <TargetMember>arn:aws:rds:us-west-2:123456789012:db:database-1-green-mhv83d</TargetMember>
            <Status>PROVISIONING</Status>
          </member>
        </SwitchoverDetails>
        <Source>arn:aws:rds:us-west-2:123456789012:db:database-1</Source>
        <BlueGreenDeploymentIdentifier>bgd-clyvb1zv1geqensv</BlueGreenDeploymentIdentifier>
        <Tasks>
          <member>
            <Name>CREATING_READ_REPLICA_OF_SOURCE</Name>
            <Status>IN_PROGRESS</Status>
          </member>
          <member>
            <Name>DB_ENGINE_VERSION_UPGRADE</Name>
            <Status>PENDING</Status>
          </member>
          <member>
            <Name>CONFIGURE_BACKUPS</Name>
            <Status>PENDING</Status>
          </member>
        </Tasks>
        <Target>arn:aws:rds:us-west-2:123456789012:db:database-1-green-mhv83d</Target>
        <Status>PROVISIONING</Status>
      </member>
    </BlueGreenDeployments>
  </DescribeBlueGreenDeploymentsResult>
  <ResponseMetadata>
    <RequestId>a534de7b-dc20-4b16-863a-24f456385d3a</RequestId>
  </ResponseMetadata>
</DescribeBlueGreenDeploymentsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describebluegreendeployments.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describebluegreendeployments.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeAccountAttributes

DescribeCertificates

All content copied from https://docs.aws.amazon.com/.
