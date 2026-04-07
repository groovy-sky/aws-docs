# DeleteQueryLoggingConfig

Deletes a configuration for DNS query logging. If you delete a configuration, Amazon
Route 53 stops sending query logs to CloudWatch Logs. Route 53 doesn't delete any logs
that are already in CloudWatch Logs.

For more information about DNS query logs, see [CreateQueryLoggingConfig](https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html).

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53-2013-04-01/DeleteQueryLoggingConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/DeleteQueryLoggingConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteKeySigningKey

DeleteReusableDelegationSet
