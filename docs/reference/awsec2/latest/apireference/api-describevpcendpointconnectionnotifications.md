# DescribeVpcEndpointConnectionNotifications

Describes the connection notifications for VPC endpoints and VPC endpoint
services.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ConnectionNotificationId**

The ID of the notification.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `connection-notification-arn` \- The ARN of the SNS topic for the
notification.

- `connection-notification-id` \- The ID of the
notification.

- `connection-notification-state` \- The state of the notification
( `Enabled` \| `Disabled`).

- `connection-notification-type` \- The type of notification
( `Topic`).

- `service-id` \- The ID of the endpoint service.

- `vpc-endpoint-id` \- The ID of the VPC endpoint.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return in a single call. To retrieve the remaining
results, make another request with the returned `NextToken` value.

Type: Integer

Required: No

**NextToken**

The token to request the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**connectionNotificationSet**

The notifications.

Type: Array of [ConnectionNotification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ConnectionNotification.html) objects

**nextToken**

The token to use to retrieve the next page of results. This value is
`null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all of your connection notifications.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcEndpointConnectionNotifications
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVpcEndpointConnectionNotificationsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>48541e40-9b6f-488e-8da7-a52a7example</requestId>
    <connectionNotificationSet>
        <item>
            <connectionNotificationArn>arn:aws:sns:us-east-1:123456789012:EndpointNotification</connectionNotificationArn>
            <connectionEvents>
                <item>Accept</item>
                <item>Connect</item>
                <item>Delete</item>
                <item>Reject</item>
            </connectionEvents>
            <connectionNotificationType>Topic</connectionNotificationType>
            <connectionNotificationState>Enabled</connectionNotificationState>
            <connectionNotificationId>vpce-nfn-123cb952bc8af7123</connectionNotificationId>
            <vpcEndpointId>vpce-1234151a02f327123</vpcEndpointId>
        </item>
    </connectionNotificationSet>
</DescribeVpcEndpointConnectionNotificationsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeVpcEndpointConnectionNotifications)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeVpcEndpointAssociations

DescribeVpcEndpointConnections
