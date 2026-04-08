# GetProfileAssociation

Retrieves a Route 53 Profile association for a VPC. A VPC can have only one Profile association, but a Profile can be associated with up to 5000 VPCs.

## Request Syntax

```nohighlight

GET /profileassociation/ProfileAssociationId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ProfileAssociationId](#API_route53profiles_GetProfileAssociation_RequestSyntax)**

The identifier of the association you want to get information about.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "ProfileAssociation": {
      "CreationTime": number,
      "Id": "string",
      "ModificationTime": number,
      "Name": "string",
      "OwnerId": "string",
      "ProfileId": "string",
      "ResourceId": "string",
      "Status": "string",
      "StatusMessage": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ProfileAssociation](#API_route53profiles_GetProfileAssociation_ResponseSyntax)**

Information about the Profile association that you specified in a `GetProfileAssociation` request.

Type: [ProfileAssociation](api-route53profiles-profileassociation.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified operation.

HTTP Status Code: 400

**ResourceNotFoundException**

The resource you are associating is not found.

**ResourceType**

The resource type that caused the resource not found exception.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

**ValidationException**

You have provided an invalid command.

HTTP Status Code: 400

## Examples

### GetProfileAssociation Example

This example illustrates one usage of GetProfileAssociation.

#### Sample Request

```

GET /profileassociation/rpassoc-489ce212fexample HTTP/1.1
host:route53profiles.us-east-1.amazonaws.com
Accept-Encoding: identity
X-Amz-Date:20240319T230136Z
User-Agent: aws-cli/1.32.63 botocore/1.34.63 Python/3.8.18
Authorization: AWS4-HMAC-SHA256
    Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-1/route53profiles/aws4_request,
    SignedHeaders=host;x-amz-date;x-amz-security-token
    Signature=[calculated-signature]
{} # RequestBody is empty
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Tue, 19 Mar 2024 23:01:36 GMT
Content-Type: application/json
Content-Length: 321
Connection: keep-alive
x-amzn-RequestId: dcd9d91e-1a5a-481f-82b7-bafe7dexample
Access-Control-Allow-Origin: *
x-amz-apigw-id: U5eX0FdmIexample=
X-Amzn-Trace-Id: Root=1-65fa10fe-6e5a93a56a325ab8example
{
    "ProfileAssociation": {
        "CreationTime": 1710886843.849,
        "Id": "rrpassoc-489ce212fexample",
        "ModificationTime": 1710886989.178,
        "Name": "test-association",
        "OwnerId": "123456789012",
        "ProfileId": "rp-4987774726example",
        "ResourceId": "vpc-0af3b96b3example",
        "Status": "COMPLETE",
        "StatusMessage": "Created Profile Association"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53profiles-2018-05-10/getprofileassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53profiles-2018-05-10/getprofileassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetProfile

GetProfileResourceAssociation
