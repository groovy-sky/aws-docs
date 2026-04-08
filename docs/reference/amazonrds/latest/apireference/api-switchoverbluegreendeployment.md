# SwitchoverBlueGreenDeployment

Switches over a blue/green deployment.

Before you switch over, production traffic is routed to the databases in the blue environment.
After you switch over, production traffic is routed to the databases in the green environment.

For more information, see [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User_
_Guide_ and [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**BlueGreenDeploymentIdentifier**

The resource ID of the blue/green deployment.

Constraints:

- Must match an existing blue/green deployment resource ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][0-9A-Za-z-:._]*`

Required: Yes

**SwitchoverTimeout**

The amount of time, in seconds, for the switchover to complete.

Default: 300

If the switchover takes longer than the specified duration, then any changes are rolled back,
and no changes are made to the environments.

Type: Integer

Valid Range: Minimum value of 30.

Required: No

## Response Elements

The following element is returned by the service.

**BlueGreenDeployment**

Details about a blue/green deployment.

For more information, see [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User_
_Guide_ and [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

Type: [BlueGreenDeployment](api-bluegreendeployment.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BlueGreenDeploymentNotFoundFault**

`BlueGreenDeploymentIdentifier` doesn't refer to an existing blue/green deployment.

HTTP Status Code: 404

**InvalidBlueGreenDeploymentStateFault**

The blue/green deployment can't be switched over or deleted because there is an invalid configuration in
the green environment.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of SwitchoverBlueGreenDeployment.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=SwitchoverBlueGreenDeployment
   &BlueGreenDeploymentIdentifier=bgd-mdoyy2mn7vbkhhgg
   &SwitchoverTimeout=400
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20141031/us-west-2/rds/aws4_request
   &X-Amz-Date=20230110T190520Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=8a684aebe6d5219bb3572316a341963324d6ef339bd0dcfa5854f1a01d401214

```

#### Sample Response

```

<SwitchoverBlueGreenDeploymentResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <SwitchoverBlueGreenDeploymentResult>
    <BlueGreenDeployment>
      <TagList/>
      <BlueGreenDeploymentName>my-blue-green-deployment</BlueGreenDeploymentName>
      <CreateTime>2023-01-10T18:42:09.330Z</CreateTime>
      <SwitchoverDetails>
        <member>
          <SourceMember>arn:aws:rds:us-west-2:123456789012:db:database-1</SourceMember>
          <TargetMember>arn:aws:rds:us-west-2:123456789012:db:database-1-green-7jtrw5</TargetMember>
          <Status>AVAILABLE</Status>
        </member>
      </SwitchoverDetails>
      <Source>arn:aws:rds:us-west-2:123456789012:db:database-1</Source>
      <BlueGreenDeploymentIdentifier>bgd-mdoyy2mn7vbkhhgg</BlueGreenDeploymentIdentifier>
      <Tasks>
        <member>
          <Name>CREATING_READ_REPLICA_OF_SOURCE</Name>
          <Status>COMPLETED</Status>
        </member>
        <member>
          <Name>CONFIGURE_BACKUPS</Name>
          <Status>COMPLETED</Status>
        </member>
      </Tasks>
      <Target>arn:aws:rds:us-west-2:123456789012:db:database-1-green-7jtrw5</Target>
      <Status>SWITCHOVER_IN_PROGRESS</Status>
    </BlueGreenDeployment>
  </SwitchoverBlueGreenDeploymentResult>
  <ResponseMetadata>
    <RequestId>c4f69d85-87e5-4fbb-b6b8-ccdb17404af6</RequestId>
  </ResponseMetadata>
</SwitchoverBlueGreenDeploymentResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/switchoverbluegreendeployment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/switchoverbluegreendeployment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopDBInstanceAutomatedBackupsReplication

SwitchoverGlobalCluster

All content copied from https://docs.aws.amazon.com/.
