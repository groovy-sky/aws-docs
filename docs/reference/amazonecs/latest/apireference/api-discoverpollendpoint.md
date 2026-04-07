# DiscoverPollEndpoint

###### Note

This action is only used by the Amazon ECS agent, and it is not intended for use
outside of the agent.

Returns an endpoint for the Amazon ECS agent to poll for updates.

## Request Syntax

```nohighlight

{
   "cluster": "string",
   "containerInstance": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[cluster](#API_DiscoverPollEndpoint_RequestSyntax)**

The short name or full Amazon Resource Name (ARN) of the cluster that the container
instance belongs to.

Type: String

Required: No

**[containerInstance](#API_DiscoverPollEndpoint_RequestSyntax)**

The container instance ID or full ARN of the container instance. For more information
about the ARN format, see [Amazon Resource Name (ARN)](../../../../services/amazonecs/latest/developerguide/ecs-account-settings.md#ecs-resource-ids) in the _Amazon ECS Developer_
_Guide_.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "endpoint": "string",
   "serviceConnectEndpoint": "string",
   "telemetryEndpoint": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[endpoint](#API_DiscoverPollEndpoint_ResponseSyntax)**

The endpoint for the Amazon ECS agent to poll.

Type: String

**[serviceConnectEndpoint](#API_DiscoverPollEndpoint_ResponseSyntax)**

The endpoint for the Amazon ECS agent to poll for Service Connect configuration. For
more information, see [Service Connect](../../../../services/amazonecs/latest/developerguide/service-connect.md)
in the _Amazon Elastic Container Service Developer Guide_.

Type: String

**[telemetryEndpoint](#API_DiscoverPollEndpoint_ResponseSyntax)**

The telemetry endpoint for the Amazon ECS agent.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/CommonErrors.html).

**ClientException**

These errors are usually caused by a client action. This client action might be using
an action or resource on behalf of a user that doesn't have permissions to use the
action or resource. Or, it might be specifying an identifier that isn't valid.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server issue.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ecs-2014-11-13/DiscoverPollEndpoint)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/DiscoverPollEndpoint)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeTaskSets

ExecuteCommand
