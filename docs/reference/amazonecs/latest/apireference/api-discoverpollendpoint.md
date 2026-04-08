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

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

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

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecs-2014-11-13/discoverpollendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/discoverpollendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeTaskSets

ExecuteCommand
