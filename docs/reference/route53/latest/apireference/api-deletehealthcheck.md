# DeleteHealthCheck

Deletes a health check.

###### Important

Amazon Route 53 does not prevent you from deleting a health check even if the
health check is associated with one or more resource record sets. If you delete a
health check and you don't update the associated resource record sets, the future
status of the health check can't be predicted and may change. This will affect the
routing of DNS queries for your DNS failover configuration. For more information,
see [Replacing and Deleting Health Checks](../../../../services/route53/latest/developerguide/health-checks-creating-deleting.md#health-checks-deleting.html) in the _Amazon Route 53_
_Developer Guide_.

If you're using AWS Cloud Map and you configured Cloud Map to create a Route 53
health check when you register an instance, you can't use the Route 53
`DeleteHealthCheck` command to delete the health check. The health check
is deleted automatically when you deregister the instance; there can be a delay of
several hours before the health check is deleted from Route 53.

## Request Syntax

```nohighlight

DELETE /2013-04-01/healthcheck/HealthCheckId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[HealthCheckId](#API_DeleteHealthCheck_RequestSyntax)**

The ID of the health check that you want to delete.

Length Constraints: Maximum length of 64.

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

**HealthCheckInUse**

_This error has been deprecated._

This error code is not in use.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchHealthCheck**

No health check exists with the specified ID.

**message**

HTTP Status Code: 404

## Examples

### Example Request

This example illustrates one usage of DeleteHealthCheck.

```

DELETE /2013-04-01/healthcheck/abcdef11-2222-3333-4444-555555fedcba
```

### Example Response

This example illustrates one usage of DeleteHealthCheck.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<DeleteHealthCheckResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
</DeleteHealthCheckResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletehealthcheck.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletehealthcheck.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteCidrCollection

DeleteHostedZone
