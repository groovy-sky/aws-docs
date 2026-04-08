# ModifyDBRecommendation

Updates the recommendation status and recommended action status for the specified recommendation.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**RecommendationId**

The identifier of the recommendation to update.

Type: String

Required: Yes

**Locale**

The language of the modified recommendation.

Type: String

Required: No

**RecommendedActionUpdates.member.N**

The list of recommended action status to update. You can update multiple recommended actions at one time.

Type: Array of [RecommendedActionUpdate](api-recommendedactionupdate.md) objects

Required: No

**Status**

The recommendation status to update.

Valid values:

- active

- dismissed

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**DBRecommendation**

The recommendation for your DB instances, DB clusters, and DB parameter groups.

Type: [DBRecommendation](api-dbrecommendation.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Modifying the recommended action status for a recommendation

This example illustrates one usage of ModifyDBRecommendation.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=ModifyDBRecommendation
    &RecommendationId=15e811d7-ec23-4d94-8d28-74cd2e7729ad
    &RecommendedActionUpdates.member.1.ActionId=806effbdc8853c4bf0e794c0c240ee8e
    &RecommendedActionUpdates.member.1.Status=applied
    &Locale=es
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20230222/us-east-1/rds/aws4_request
    &X-Amz-Date=20230222T200807Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=2d4f2b9e8abc31122b5546f94c0499bba47de813cb875f9b9c78e8e19c9afe1b

```

#### Sample Response

```

<ModifyDBRecommendationResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ModifyDBRecommendationResult>
    <DBRecommendation>
      <RecommendationId>15e811d7-ec23-4d94-8d28-74cd2e7729ad</RecommendationId>
      <TypeId>config_recommendation::multi_az_instance</TypeId>
      <Severity>low</Severity>
      <ResourceArn>arn:aws:rds:us-west-2:636812126935:db:mariadb-instance</ResourceArn>
      <Status>pending</Status>
      <CreatedTime>2023-10-05T18:04:04.017000+00:00</CreatedTime>
      <UpdatedTime>2023-10-20T19:17:18+00:00</UpdatedTime>
      <Detection>**1 resource** is not a Multi-AZ instance</Detection>
      <Recommendation>Set up Multi-AZ for the impacted DB instances</Recommendation>
      <Description>We recommend that you use Multi-AZ deployment. The Multi-AZ deployments enhance the availability and durability of the DB instance. Click Info for more details about Multi-AZ deployment and pricing.</Description>
      <RecommendedActions>
        <member>
          <ActionId>806effbdc8853c4bf0e794c0c240ee8e</ActionId>
          <Operation>modifyDbInstance</Operation>
          <Parameters>
            <member>
              <Key>MultiAZ</Key>
              <Value>true</Value>
            </member>
            <member>
              <Key>DBInstanceIdentifier</Key>
              <Value>mariadb-instance</Value>
            </member>
          </Parameters>
          <ApplyModes>
            <member>immediately</member>
            <member>next-maintenance-window</member>
          </ApplyModes>
          <Status>applied</Status>
          <ContextAttributes>
            <member>
              <Key>resourceArn</Key>
              <Value>arn:aws:rds:us-west-2:636812126935:db:mariadb-instance</Value>
            </member>
            <member>
              <Key>engineName</Key>
              <Value>mariadb</Value>
            </member>
          </ContextAttributes>
        </member>
      </RecommendedActions>
      <Category>reliability</Category>
      <Source>RDS</Source>
      <TypeDetection>**[resource-count] resources** are not Multi-AZ instances</TypeDetection>
      <TypeRecommendation>Set up Multi-AZ for the impacted DB instances</TypeRecommendation>
      <Impact>Data availability at risk</Impact>
      <AdditionalInfo>In an Amazon RDS Multi-AZ deployment, Amazon RDS automatically creates a primary database instance and replicates the data to an instance in a different availability zone. When it detects a failure, Amazon RDS automatically fails over to a standby instance without manual intervention.</AdditionalInfo>
      <Links>
        <member>
          <Text>Pricing for Amazon RDS Multi-AZ</Text>
          <Url>https://aws.amazon.com/rds/features/multi-az/#Pricing</Url>
        </member>
      </Links>
    </DBRecommendation>
  </ModifyDBRecommendationResult>
</ModifyDBRecommendationResponse>

```

### Modifying the recommendation status for the specified recommendation ID

This example illustrates one usage of ModifyDBRecommendation.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=ModifyDBRecommendation
    &RecommendationId=8c9132b0-267d-4493-b3c4-aedd0920809d
    &Status=dismissed
    &Locale=es
    &SignatureMethod=HmacSHA256
    &SignatureVersion=4
    &Version=2014-10-31
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20230222/us-east-1/rds/aws4_request
    &X-Amz-Date=20230222T200807Z
    &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
    &X-Amz-Signature=2d4f2b9e8abc31122b5546f94c0499bba47de813cb875f9b9c78e8e19c9afe1b

```

#### Sample Response

```

<ModifyDBRecommendationResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <ModifyDBRecommendationResult>
    <DBRecommendation>
      <RecommendationId>8c9132b0-267d-4493-b3c4-aedd0920809d</RecommendationId>
      <TypeId>config_recommendation::enhanced_monitoring_off</TypeId>
      <Severity>low</Severity>
      <ResourceArn>arn:aws:rds:us-west-2:636812126935:db:mariadb-instance</ResourceArn>
      <Status>dismissed</Status>
      <CreatedTime>2023-10-05T18:04:03.957000+00:00</CreatedTime>
      <UpdatedTime>2023-10-20T19:20:22+00:00</UpdatedTime>
      <Detection>**1 resource** doesn't have Enhanced Monitoring enabled</Detection>
      <Recommendation>Turn on Enhanced Monitoring</Recommendation>
      <Description>Your database resources don't have Enhanced Monitoring turned on. Enhanced Monitoring provides real-time operating system metrics for monitoring and troubleshooting.</Description>
      <RecommendedActions>
        <member>
          <ActionId>a2e5e55f28854f9ec12f45c227d85f48</ActionId>
          <Operation>modifyDbInstance</Operation>
          <Parameters>
            <member>
              <Key>MonitoringInterval</Key>
              <Value>60</Value>
            </member>
            <member>
              <Key>DBInstanceIdentifier</Key>
              <Value>mariadb-instance</Value>
            </member>
          </Parameters>
          <ApplyModes>
            <mmeber>immediately</mmeber>
          </ApplyModes>
          <Status>ready</Status>
          <ContextAttributes>
            <member>
              <Key>resourceArn</Key>
              <Value>arn:aws:rds:us-west-2:636812126935:db:mariadb-instance</Value>
            </member>
            <member>
              <Key>engineName</Key>
              <Value>mariadb</Value>
            </member>
            <member>
              <Key>recommendedValue</Key>
              <Value>60</Value>
            </member>
          </ContextAttributes>
        </member>
      </RecommendedActions>
      <Category>reliability</Category>
      <Source>RDS</Source>
      <TypeDetection>**[resource-count] resources** don't have Enhanced Monitoring enabled</TypeDetection>
      <TypeRecommendation>Turn on Enhanced Monitoring</TypeRecommendation>
      <Impact>Reduced operational visibility</Impact>
      <AdditionalInfo>Enhanced Monitoring for Amazon RDS provides additional visibility on the health of your DB instances. We recommend that you turn on Enhanced Monitoring. When the Enhanced Monitoring option is turned on for your DB instance, it collects vital operating system metrics and process information.</AdditionalInfo>
      <Links>
        <member>
          <Text>Turning Enhanced Monitoring on and off</Text>
          <Url>https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html</Url>
        </member>
      </Links>
    </DBRecommendation>
  </ModifyDBRecommendationResult>
</ModifyDBRecommendationResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifydbrecommendation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifydbrecommendation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyDBProxyTargetGroup

ModifyDBShardGroup

All content copied from https://docs.aws.amazon.com/.
