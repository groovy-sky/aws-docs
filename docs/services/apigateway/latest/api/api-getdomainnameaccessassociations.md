# GetDomainNameAccessAssociations

Represents a collection on DomainNameAccessAssociations resources.

## Request Syntax

```nohighlight

GET /domainnameaccessassociations?limit=limit&position=position&resourceOwner=resourceOwner HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetDomainNameAccessAssociations_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetDomainNameAccessAssociations_RequestSyntax)**

The current pagination position in the paged result set.

**[resourceOwner](#API_GetDomainNameAccessAssociations_RequestSyntax)**

The owner of the domain name access association. Use `SELF` to only list the domain name access associations owned by your own account.
Use `OTHER_ACCOUNTS` to list the domain name access associations with your private custom domain names that are owned by other AWS
accounts.

Valid Values: `SELF | OTHER_ACCOUNTS`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
      {
         "accessAssociationSource": "string",
         "accessAssociationSourceType": "string",
         "domainNameAccessAssociationArn": "string",
         "domainNameArn": "string",
         "tags": {
            "string" : "string"
         }
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetDomainNameAccessAssociations_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [DomainNameAccessAssociation](api-domainnameaccessassociation.md) objects

**[position](#API_GetDomainNameAccessAssociations_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Get the domain name access associations for domain name access associations owned by your own account

This example illustrates one usage of GetDomainNameAccessAssociations.

```request

GET /domainnameaccessassociations?limit=10&resourceOwner=SELF
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160601T173728Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160601/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
    "_links": {
        "curies": {
            "href": "http://docs.aws.amazon.com/apigateway/latest/developerguide/domainnameaccessassociation-domainname-{rel}.html",
            "name": "domainnameaccessassociation",
            "templated": true
        },
        "self": {
            "href": "/domainnameaccessassociations?limit=10"
        },
        "domainnameaccessassociation:create": {
            "href": "/domainnameaccessassociations"
        },
        "item": {
            "href": "/domainnameaccessassociations/arn:aws:apigateway:us-east-1:012345678910:/domainnameaccessassociations/domainname/private.a.b.c.com+abcd1234/vpcesource/vpce-abcd1234"
        }
    },
    "_embedded": {
        "item": {
            "_links": {
                "self": {
                    "href": "/domainnameaccessassociations/arn:aws:apigateway:us-east-1:012345678910:/domainnameaccessassociations/domainname/private.a.b.c.com+abcd1234/vpcesource/vpce-abcd1234"
                }
            },
            "accessAssociationSource": "vpce-abcd1234",
            "accessAssociationSourceType": "VPCE",
            "domainNameAccessAssociationArn": "arn:aws:apigateway:us-east-1:012345678910:/domainnameaccessassociations/domainname/private.a.b.c.com+abcd1234/vpcesource/vpce-abcd1234",
            "domainNameArn": "arn:aws:apigateway:us-east-1:012345678910:/domainnames/private.a.b.c.com+abcd1234"
        }
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getdomainnameaccessassociations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getdomainnameaccessassociations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDomainName

GetDomainNames

All content copied from https://docs.aws.amazon.com/.
