# CreateCidrCollection

Creates a CIDR collection in the current AWS account.

## Request Syntax

```nohighlight

POST /2013-04-01/cidrcollection HTTP/1.1
<?xml version="1.0" encoding="UTF-8"?>
<CreateCidrCollectionRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <CallerReference>string</CallerReference>
   <Name>string</Name>
</CreateCidrCollectionRequest>
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in XML format.

**[CreateCidrCollectionRequest](#API_CreateCidrCollection_RequestSyntax)**

Root level tag for the CreateCidrCollectionRequest parameters.

Required: Yes

**[CallerReference](#API_CreateCidrCollection_RequestSyntax)**

A client-specific token that allows requests to be securely retried so that the
intended outcome will only occur once, retries receive a similar response, and there are
no additional edge cases to handle.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `\p{ASCII}+`

Required: Yes

**[Name](#API_CreateCidrCollection_RequestSyntax)**

A unique identifier for the account that can be used to reference the collection from
other API calls.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[0-9A-Za-z_\-]+`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 201
Location: Location
<?xml version="1.0" encoding="UTF-8"?>
<CreateCidrCollectionResponse>
   <Collection>
      <Arn>string</Arn>
      <Id>string</Id>
      <Name>string</Name>
      <Version>long</Version>
   </Collection>
</CreateCidrCollectionResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The response returns the following HTTP headers.

**[Location](#API_CreateCidrCollection_ResponseSyntax)**

A unique URL that represents the location for the CIDR collection.

Length Constraints: Maximum length of 1024.

The following data is returned in XML format by the service.

**[CreateCidrCollectionResponse](#API_CreateCidrCollection_ResponseSyntax)**

Root level tag for the CreateCidrCollectionResponse parameters.

Required: Yes

**[Collection](#API_CreateCidrCollection_ResponseSyntax)**

A complex type that contains information about the CIDR collection.

Type: [CidrCollection](api-cidrcollection.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CidrCollectionAlreadyExistsException**

A CIDR collection with this name and a different caller reference already exists in this account.

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

**LimitsExceeded**

This operation can't be completed because the current account has reached the
limit on the resource you are trying to create. To request a higher limit, [create a case](http://aws.amazon.com/route53-request) with the AWS Support
Center.

**message**

HTTP Status Code: 400

## Examples

### Example request

This example illustrates one usage of CreateCidrCollection.

```

POST /2013-04-01/cidrcollection
<?xml version="1.0" encoding="UTF-8"?>
<CreateCidrCollectionRequest xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <Name>isp-city-cidrs</Name>
    <CallerReference>ref1</CallerReference>
</CreateCidrCollectionRequest>
```

### Example response

This example illustrates one usage of CreateCidrCollection.

```

HTTP/1.1 201 Created
<?xml version="1.0"?>
<CreateCidrCollectionResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
 <Collection>
   <Arn>arn:aws:route53:::cidrcollection/c8c02a84-aaaa-bbbb-e0d2-d833a2f80106</Arn>
   <Id>c8c02a84-aaaa-bbbb-e0d2-d833a2f80106</Id>
   <Name>isp-city-cidrs</Name>
   <Version>1</Version>
  </Collection>
</CreateCidrCollectionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/createcidrcollection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/createcidrcollection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ChangeTagsForResource

CreateHealthCheck
