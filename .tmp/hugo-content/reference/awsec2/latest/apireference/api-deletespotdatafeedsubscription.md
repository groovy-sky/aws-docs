# DeleteSpotDatafeedSubscription

Deletes the data feed for Spot Instances.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request deletes the data feed for the AWS account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteSpotDatafeedSubscription
&AUTHPARAMS
```

#### Sample Response

```

<DeleteSpotDatafeedSubscriptionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</DeleteSpotDatafeedSubscriptionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletespotdatafeedsubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletespotdatafeedsubscription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteSnapshot

DeleteSubnet
