# StartDiscovery

Enables this AWS account to be able to use CloudWatch Application Signals
by creating the _AWSServiceRoleForCloudWatchApplicationSignals_ service-linked role. This service-
linked role has the following permissions:

- `xray:GetServiceGraph`

- `logs:StartQuery`

- `logs:GetQueryResults`

- `cloudwatch:GetMetricData`

- `cloudwatch:ListMetrics`

- `tag:GetResources`

- `autoscaling:DescribeAutoScalingGroups`

A service-linked CloudTrail event channel is created to process CloudTrail events and return change event information. This includes last deployment time, userName, eventName, and other event metadata.

After completing this step, you still need to instrument your Java and Python applications to send data
to Application Signals. For more information, see
[Enabling Application Signals](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-application-signals-enable.md).

## Request Syntax

```

POST /start-discovery HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 403

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/startdiscovery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/startdiscovery.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutGroupingConfiguration

TagResource
