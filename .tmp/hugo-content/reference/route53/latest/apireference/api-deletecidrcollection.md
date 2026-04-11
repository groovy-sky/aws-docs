# DeleteCidrCollection

Deletes a CIDR collection in the current AWS account. The collection
must be empty before it can be deleted.

## Request Syntax

```nohighlight

DELETE /2013-04-01/cidrcollection/CidrCollectionId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[CidrCollectionId](#API_DeleteCidrCollection_RequestSyntax)**

The UUID of the collection to delete.

Pattern: `[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}`

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

**CidrCollectionInUseException**

This CIDR collection is in use, and isn't empty.

HTTP Status Code: 400

**ConcurrentModification**

Another user submitted a request to create, update, or delete the object at the same
time that you did. Retry the request.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchCidrCollectionException**

The CIDR collection you specified, doesn't exist.

HTTP Status Code: 404

## Examples

### Example request

This example illustrates one usage of DeleteCidrCollection.

```

DELETE /2013-04-01/cidrcollection/c8c02a84-aaaa-bbbb-e0d2-d833a2f80106
```

### Example response

This example illustrates one usage of DeleteCidrCollection.

```

HTTP/1.1 200 OK
<?xml version="1.0"?>
<DeleteCidrCollectionResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/"/>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/deletecidrcollection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/deletecidrcollection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeactivateKeySigningKey

DeleteHealthCheck
