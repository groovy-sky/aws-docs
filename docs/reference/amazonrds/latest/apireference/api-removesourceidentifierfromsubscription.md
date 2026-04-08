# RemoveSourceIdentifierFromSubscription

Removes a source identifier from an existing RDS event notification subscription.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceIdentifier**

The source identifier to be removed from the subscription, such as the **DB instance identifier**
for a DB instance or the name of a security group.

Type: String

Required: Yes

**SubscriptionName**

The name of the RDS event notification subscription you want to remove a source identifier from.

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

This example illustrates one usage of RemoveSourceIdentifierFromSubscription.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
   ?Action=RemoveSourceIdentifierFromSubscription
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &SourceIdentifier=si-sample
   &SubscriptionName=myawsuser-secgrp
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140428/us-east-1/rds/aws4_request
   &X-Amz-Date=20140428T222718Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=4419f0015657ee120d781849ffdc6642eeafeee42bf1d18c4b2ed8eb732f7bf8

```

#### Sample Response

```

<RemoveSourceIdentifierFromSubscriptionResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <RemoveSourceIdentifierFromSubscriptionResult>
    <EventSubscription>
      <Enabled>true</Enabled>
      <CustomerAwsId>802#########</CustomerAwsId>
      <SourceType>db-security-group</SourceType>
      <Status>active</Status>
      <SubscriptionCreationTime>2014-04-25 21:43:25.542</SubscriptionCreationTime>
      <EventCategoriesList>
        <EventCategory>configuration change</EventCategory>
        <EventCategory>failure</EventCategory>
      </EventCategoriesList>
      <CustSubscriptionId>myawsuser-secgrp</CustSubscriptionId>
      <SnsTopicArn>arn:aws:sns:us-east-1:802#########:myawsuser-RDS</SnsTopicArn>
    </EventSubscription>
  </RemoveSourceIdentifierFromSubscriptionResult>
  <ResponseMetadata>
    <RequestId>326cdeb9-be23-11d3-91a5-a90441261bc4</RequestId>
  </ResponseMetadata>
</RemoveSourceIdentifierFromSubscriptionResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/removesourceidentifierfromsubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/removesourceidentifierfromsubscription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoveRoleFromDBInstance

RemoveTagsFromResource

All content copied from https://docs.aws.amazon.com/.
