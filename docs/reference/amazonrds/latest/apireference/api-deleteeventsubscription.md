# DeleteEventSubscription

Deletes an RDS event notification subscription.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SubscriptionName**

The name of the RDS event notification subscription you want to delete.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**EventSubscription**

Contains the results of a successful invocation of the `DescribeEventSubscriptions` action.

Type: [EventSubscription](api-eventsubscription.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidEventSubscriptionState**

This error can occur if someone else is modifying a subscription. You should retry the action.

HTTP Status Code: 400

**SubscriptionNotFound**

The subscription name does not exist.

HTTP Status Code: 404

## Examples

### Example

This example illustrates one usage of DeleteEventSubscription.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=DeleteEventSubscription
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SubscriptionName=EventSubscription04
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140423/us-east-1/rds/aws4_request
   &X-Amz-Date=20140423T203337Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=05aa834e364a9e1a279d44cc955694518fc96fff638c74faa2be45783102e785

```

#### Sample Response

```

<DeleteEventSubscriptionResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DeleteEventSubscriptionResult>
    <EventSubscription>
      <Enabled>true</Enabled>
      <CustomerAwsId>803#########</CustomerAwsId>
      <SourceType>db-instance</SourceType>
      <Status>deleting</Status>
      <SourceIdsList>
        <SourceId>mysqldb</SourceId>
      </SourceIdsList>
      <SubscriptionCreationTime>2014-04-22 23:03:19.776</SubscriptionCreationTime>
      <CustSubscriptionId>EventSubscription04</CustSubscriptionId>
      <SnsTopicArn>arn:aws:sns:us-east-1:803#########:myawsuser-RDS</SnsTopicArn>
    </EventSubscription>
  </DeleteEventSubscriptionResult>
  <ResponseMetadata>
    <RequestId>7b4cf02a-ba25-11d3-a691-857dc0addcc9</RequestId>
  </ResponseMetadata>
</DeleteEventSubscriptionResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deleteeventsubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deleteeventsubscription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBSubnetGroup

DeleteGlobalCluster

All content copied from https://docs.aws.amazon.com/.
