# GetDomainNames

Represents a collection of DomainName resources.

## Request Syntax

```nohighlight

GET /domainnames?limit=limit&position=position&resourceOwner=resourceOwner HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetDomainNames_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetDomainNames_RequestSyntax)**

The current pagination position in the paged result set.

**[resourceOwner](#API_GetDomainNames_RequestSyntax)**

The owner of the domain name.

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
         "certificateArn": "string",
         "certificateName": "string",
         "certificateUploadDate": number,
         "distributionDomainName": "string",
         "distributionHostedZoneId": "string",
         "domainName": "string",
         "domainNameArn": "string",
         "domainNameId": "string",
         "domainNameStatus": "string",
         "domainNameStatusMessage": "string",
         "endpointAccessMode": "string",
         "endpointConfiguration": {
            "ipAddressType": "string",
            "types": [ "string" ],
            "vpcEndpointIds": [ "string" ]
         },
         "managementPolicy": "string",
         "mutualTlsAuthentication": {
            "truststoreUri": "string",
            "truststoreVersion": "string",
            "truststoreWarnings": [ "string" ]
         },
         "ownershipVerificationCertificateArn": "string",
         "policy": "string",
         "regionalCertificateArn": "string",
         "regionalCertificateName": "string",
         "regionalDomainName": "string",
         "regionalHostedZoneId": "string",
         "routingMode": "string",
         "securityPolicy": "string",
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

**[item](#API_GetDomainNames_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [DomainName](api-domainname.md) objects

**[position](#API_GetDomainNames_ResponseSyntax)**

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

### Retrieve domain names

The following example request retrieves up to 10 custom DomainName resources for APIs under the caller's AWS account.

A successful response returns the requested `DomainName`
resources that can be navigated to by following the linked item or examining the embedded item resource.

#### Sample Request

```

GET /domainnames?limit=10 HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20160601T173728Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20160601/us-east-1/apigateway/aws4_request, SignedHeaders=content-type;host;x-amz-date, Signature={sig4_hash}
```

#### Sample Response

```

{
    "_links": {
        "curies": [
            {
                "href": "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-basepathmapping-{rel}.html",
                "name": "basepathmapping",
                "templated": true
            },
            {
                "href": "http://docs.aws.amazon.com/apigateway/latest/developerguide/restapi-domainname-{rel}.html",
                "name": "domainname",
                "templated": true
            }
        ],
        "self": {
            "href": "/domainnames?limit=10"
        },
        "domainname:by-name": {
            "href": "/domainnames/{domain_name}{?domain_name_id}",
            "templated": true
        },
        "domainname:create": {
            "href": "/domainnames"
        },
        "item": [
            {
                "href": "/domainnames/a.b.c.com"
            },
            {
                "href": "/domainnames/private.a.b.c.com?domain_name_id=a1b2c3"
            }
        ]
    },
    "_embedded": {
        "item": [
            {
                "_links": {
                    "self": {
                        "href": "/domainnames/a.b.c.com"
                    },
                    "basepathmapping:by-base-path": {
                        "href": "/domainnames/a.b.c.com/basepathmappings/{base_path}",
                        "templated": true
                    },
                    "basepathmapping:create": {
                        "href": "/domainnames/a.b.c.com/basepathmappings"
                    },
                    "domainname:basepathmappings": {
                        "href":  "/domainnames/a.b.c.com/basepathmappings{?limit}",,
                        "templated": true
                    },
                    "domainname:delete": {
                        "href": "/domainnames/a.b.c.com"
                    },
                    "domainname:update": {
                        "href": "/domainnames/a.b.c.com"
                    }
                },
                "certificateArn": "arn:aws:acm:us-east-1:012345678910:certificate/abcdef",
                "certificateUploadDate": "2016-04-18T22:16:04Z",
                "distributionDomainName": "d3ih7ecqtec0mt.cloudfront.net",
                "domainName": "a.b.c.com"
            },
            {
                "_links": {
                    "self": {
                        "href": "/domainnames/private.a.b.c.com?domain_name_id=a1b2c3"
                    },
                    "basepathmapping:by-base-path": {
                        "href": "/domainnames/prviate.a.b.c.com/basepathmappings/{base_path}?domain_name_id=a1b2c3",
                        "templated": true
                    },
                    "basepathmapping:create": {
                        "href": "/domainnames/private.a.b.c.com/basepathmappings?domain_name_id=a1b2c3"
                    },
                    "domainname:basepathmappings": {
                        "href": "/domainnames/private.a.b.c.com/basepathmappings?domain_name_id=a1b2c3{&limit}",
                        "templated": true
                    },
                    "domainname:delete": {
                        "href": "/domainnames/private.a.b.c.com?domain_name_id=a1b2c3"
                    },
                    "domainname:update": {
                        "href": "/domainnames/private.a.b.c.com?domain_name_id=a1b2c3"
                    }
                },
                "certificateArn": "arn:aws:acm:us-west-2:012345678910:certificate/aaaaaaa",
                "certificateUploadDate": "2024-09-19T21:51:28Z",
                "domainName": "private.a.b.c.com",
                "domainNameArn": "arn:aws:apigateway:us-west-2:012345678910:/domainnames/private.a.b.c.com+a1b2c3",
                "domainNameId": "a1b2c3",
                "domainNameStatus": "AVAILABLE",
                "endpointConfiguration": {
                    "vpcEndpointIds": null,
                    "types": "PRIVATE"
                },
                "securityPolicy": "TLS_1_2"
            }
        ]
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getdomainnames.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/getdomainnames.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDomainNameAccessAssociations

GetExport

All content copied from https://docs.aws.amazon.com/.
