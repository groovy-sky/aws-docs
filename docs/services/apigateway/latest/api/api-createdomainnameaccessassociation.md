# CreateDomainNameAccessAssociation

Creates a domain name access association resource between an access association source and a private custom
domain name.

## Request Syntax

```nohighlight

POST /domainnameaccessassociations HTTP/1.1
Content-type: application/json

{
   "accessAssociationSource": "string",
   "accessAssociationSourceType": "string",
   "domainNameArn": "string",
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[accessAssociationSource](#API_CreateDomainNameAccessAssociation_RequestSyntax)**

The identifier of the domain name access association source. For a VPCE, the value is the VPC endpoint ID.

Type: String

Required: Yes

**[accessAssociationSourceType](#API_CreateDomainNameAccessAssociation_RequestSyntax)**

The type of the domain name access association source.

Type: String

Valid Values: `VPCE`

Required: Yes

**[domainNameArn](#API_CreateDomainNameAccessAssociation_RequestSyntax)**

The ARN of the domain name.

Type: String

Required: Yes

**[tags](#API_CreateDomainNameAccessAssociation_RequestSyntax)**

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

Type: String to string map

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "accessAssociationSource": "string",
   "accessAssociationSourceType": "string",
   "domainNameAccessAssociationArn": "string",
   "domainNameArn": "string",
   "tags": {
      "string" : "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[accessAssociationSource](#API_CreateDomainNameAccessAssociation_ResponseSyntax)**

The identifier of the domain name access association source. For a VPCE, the value is the VPC endpoint ID.

Type: String

**[accessAssociationSourceType](#API_CreateDomainNameAccessAssociation_ResponseSyntax)**

The type of the domain name access association source.

Type: String

Valid Values: `VPCE`

**[domainNameAccessAssociationArn](#API_CreateDomainNameAccessAssociation_ResponseSyntax)**

The ARN of the domain name access association resource.

Type: String

**[domainNameArn](#API_CreateDomainNameAccessAssociation_ResponseSyntax)**

The ARN of the domain name.

Type: String

**[tags](#API_CreateDomainNameAccessAssociation_ResponseSyntax)**

The collection of tags. Each tag element is associated with a given resource.

Type: String to string map

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**ConflictException**

The request configuration has conflicts. For details, see the accompanying error message.

HTTP Status Code: 409

**LimitExceededException**

The request exceeded the rate limit. Retry after the specified time period.

HTTP Status Code: 429

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## Examples

### Create a domain name access association

This example illustrates one usage of CreateDomainNameAccessAssociation.

#### Sample Request

```

POST /domainnameaccessassociation HTTP/1.1
Content-Type: application/json
Host: apigateway.us-east-1.amazonaws.com
X-Amz-Date: 20170223T200249Z
Authorization: AWS4-HMAC-SHA256 Credential={access_key_ID}/20170223/us-east-1/apigateway/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature={sig4_hash}

{
   "accessAssociationSource": "vpce-abcd1234",
   "accessAssociationSourceType": "VPCE",
   "domainNameArn": "arn:aws:apigateway:us-east-1:012345678910:/domainnames/private.a.b.c.com+abcd1234"
}
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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/createdomainnameaccessassociation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/createdomainnameaccessassociation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDomainName

CreateModel

All content copied from https://docs.aws.amazon.com/.
