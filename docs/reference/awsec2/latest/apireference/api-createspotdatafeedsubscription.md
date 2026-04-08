# CreateSpotDatafeedSubscription

Creates a data feed for Spot Instances, enabling you to view Spot Instance usage logs.
You can create one data feed per AWS account. For more information, see
[Spot Instance data feed](../../../../services/ec2/latest/userguide/spot-data-feeds.md)
in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Bucket**

The name of the Amazon S3 bucket in which to store the Spot Instance data feed. For
more information about bucket names, see [Bucket naming rules](../../../../services/s3/latest/userguide/bucketnamingrules.md)
in the _Amazon S3 User Guide_.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Prefix**

The prefix for the data feed file names.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**spotDatafeedSubscription**

The Spot Instance data feed subscription.

Type: [SpotDatafeedSubscription](api-spotdatafeedsubscription.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a Spot Instance data feed for the account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateSpotDatafeedSubscription
&Bucket=amzn-s3-demo-bucket
&AUTHPARAMS
```

#### Sample Response

```

<CreateSpotDatafeedSubscriptionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <spotDatafeedSubscription>
    <ownerId>123456789012</ownerId>
    <bucket>amzn-s3-demo-bucket</bucket>
    <prefix>spotdata_</prefix>
    <state>Active</state>
  </spotDatafeedSubscription>
</CreateSpotDatafeedSubscriptionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createspotdatafeedsubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createspotdatafeedsubscription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateSnapshots

CreateStoreImageTask
