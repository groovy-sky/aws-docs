# DeleteServiceLevelObjective

Deletes the specified service level objective.

## Request Syntax

```nohighlight

DELETE /slo/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DeleteServiceLevelObjective_RequestSyntax)**

The ARN or name of the service level objective to delete.

Pattern: `[0-9A-Za-z][-._0-9A-Za-z ]{0,126}[0-9A-Za-z]$|^arn:(aws|aws-us-gov):application-signals:[^:]*:[^:]*:slo/[0-9A-Za-z][-._0-9A-Za-z ]{0,126}[0-9A-Za-z]`

Required: Yes

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

**ResourceNotFoundException**

Resource not found.

**ResourceId**

Can't find the resource id.

**ResourceType**

The resource type is not valid.

HTTP Status Code: 404

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/deleteservicelevelobjective.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/deleteservicelevelobjective.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteGroupingConfiguration

GetService
