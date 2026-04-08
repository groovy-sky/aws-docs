# DeleteQueryLoggingConfig

Deletes a configuration for DNS query logging. If you delete a configuration, Amazon
Route 53 stops sending query logs to CloudWatch Logs. Route 53 doesn't delete any logs
that are already in CloudWatch Logs.

For more information about DNS query logs, see [CreateQueryLoggingConfig](api-createqueryloggingconfig.md).

## Request Syntax

```nohighlight

DELETE /2013-04-01/queryloggingconfig/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_DeleteQueryLoggingConfig_RequestSyntax)**

The ID of the configuration that you want to delete.

Length Constraints: Minimum length of 1. Maximum length of 36.

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

**ConcurrentModification**

Another user submitted a request to create, update, or delete the object at the same
time that you did. Retry the request.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchQueryLoggingConfig**

There is no DNS query logging configuration with the specified ID.

HTTP Status Code: 404

## Examples

### Example Request

The following request deletes the configuration with the ID
`87654321-dcba-1234-abcd-1a2b3c4d5e6f`.

```

DELETE /2013-04-01/queryloggingconfig HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<DeleteQueryLoggingConfigRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Id>87654321-dcba-1234-abcd-1a2b3c4d5e6f</Id>
</DeleteQueryLoggingConfigRequest>
```

### Example Response

This example illustrates one usage of DeleteQueryLoggingConfig.

```

HTTP/1.1 200 OK
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletequeryloggingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletequeryloggingconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteKeySigningKey

DeleteReusableDelegationSet
