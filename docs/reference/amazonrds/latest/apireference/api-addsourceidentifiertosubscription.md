# AddSourceIdentifierToSubscription

Adds a source identifier to an existing RDS event notification subscription.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceIdentifier**

The identifier of the event source to be added.

Constraints:

- If the source type is a DB instance, a `DBInstanceIdentifier` value must be supplied.

- If the source type is a DB cluster, a `DBClusterIdentifier` value must be supplied.

- If the source type is a DB parameter group, a `DBParameterGroupName` value must be supplied.

- If the source type is a DB security group, a `DBSecurityGroupName` value must be supplied.

- If the source type is a DB snapshot, a `DBSnapshotIdentifier` value must be supplied.

- If the source type is a DB cluster snapshot, a `DBClusterSnapshotIdentifier` value must be supplied.

- If the source type is an RDS Proxy, a `DBProxyName` value must be supplied.

Type: String

Required: Yes

**SubscriptionName**

The name of the RDS event notification subscription you want to add a source identifier to.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**EventSubscription**

Contains the results of a successful invocation of the `DescribeEventSubscriptions` action.

Type: [EventSubscription](api-eventsubscription.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**SourceNotFound**

The requested source could not be found.

HTTP Status Code: 404

**SubscriptionNotFound**

The subscription name does not exist.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of AddSourceIdentifierToSubscription.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=AddSourceIdentifierToSubscription
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SourceIdentifier=mysqldb
   &SubscriptionName=EventSubscription04
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140422/us-east-1/rds/aws4_request
   &X-Amz-Date=20140422T230442Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=347d5e788e809cd06c50214b12750a3c39716bf65b239bb6f7ee8ff5374e2df9

```

#### Sample Response

```

<AddSourceIdentifierToSubscriptionResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <AddSourceIdentifierToSubscriptionResult>
    <EventSubscription>
      <SourceType>db-instance</SourceType>
      <Enabled>true</Enabled>
      <CustomerAwsId>803#########</CustomerAwsId>
      <Status>modifying</Status>
      <SourceIdsList>
        <SourceId>mysqldb</SourceId>
      </SourceIdsList>
      <SubscriptionCreationTime>2014-04-22 23:03:19.776</SubscriptionCreationTime>
      <EventCategoriesList>
        <EventCategory>creation</EventCategory>
        <EventCategory>deletion</EventCategory>
      </EventCategoriesList>
      <CustSubscriptionId>EventSubscription04</CustSubscriptionId>
      <SnsTopicArn>arn:aws:sns:us-east-1:803#########:mytopic</SnsTopicArn>
    </EventSubscription>
  </AddSourceIdentifierToSubscriptionResult>
  <ResponseMetadata>
    <RequestId>6c05f0b0-bf71-11d3-f4c6-37db295f7674</RequestId>
  </ResponseMetadata>
</AddSourceIdentifierToSubscriptionResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/addsourceidentifiertosubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/addsourceidentifiertosubscription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddRoleToDBInstance

AddTagsToResource

All content copied from https://docs.aws.amazon.com/.
