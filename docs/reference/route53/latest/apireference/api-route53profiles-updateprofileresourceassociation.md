# UpdateProfileResourceAssociation

Updates the specified Route 53 Profile resourse association.

## Request Syntax

```nohighlight

PATCH /profileresourceassociation/ProfileResourceAssociationId HTTP/1.1
Content-type: application/json

{
   "Name": "string",
   "ResourceProperties": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[ProfileResourceAssociationId](#API_route53profiles_UpdateProfileResourceAssociation_RequestSyntax)**

ID of the resource association.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[Name](#API_route53profiles_UpdateProfileResourceAssociation_RequestSyntax)**

Name of the resource association.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**[ResourceProperties](#API_route53profiles_UpdateProfileResourceAssociation_RequestSyntax)**

If you are adding a DNS Firewall rule group, include also a priority. The priority indicates the processing order for the rule groups, starting with the priority assinged the lowest value.

The allowed values for priority are between 100 and 9900.

Type: String

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "ProfileResourceAssociation": {
      "CreationTime": number,
      "Id": "string",
      "ModificationTime": number,
      "Name": "string",
      "OwnerId": "string",
      "ProfileId": "string",
      "ResourceArn": "string",
      "ResourceProperties": "string",
      "ResourceType": "string",
      "Status": "string",
      "StatusMessage": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ProfileResourceAssociation](#API_route53profiles_UpdateProfileResourceAssociation_ResponseSyntax)**

Information about the `UpdateProfileResourceAssociation` request, including a status message.

Type: [ProfileResourceAssociation](api-route53profiles-profileresourceassociation.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified operation.

HTTP Status Code: 400

**ConflictException**

The request you submitted conflicts with an existing request.

HTTP Status Code: 400

**InternalServiceErrorException**

An internal server error occured. Retry your request.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters in this request are not valid.

**FieldName**

The parameter field name for the invalid parameter exception.

HTTP Status Code: 400

**LimitExceededException**

The request caused one or more limits to be exceeded.

**ResourceType**

The resource type that caused the limits to be exceeded.

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

### UpdateProfileResourceAssociation Example

This example illustrates one usage of UpdateProfileResourceAssociation.

#### Sample Request

```

Request:
PATCH /profileresourceassociation/rpr-001913120a7example HTTP/1.1
host:route53profiles.us-east-1.amazonaws.com
Accept-Encoding: identity
X-Amz-Date:20240319T233609Z
User-Agent: aws-cli/1.32.63 botocore/1.34.63 Python/3.8.18
Content-Type: application/json
Content-Length: 45
Authorization: AWS4-HMAC-SHA256
    Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-1/route53profiles/aws4_request,
    SignedHeaders=content-type;host;x-amz-date;x-amz-security-token,
    Signature=[calculated-signature]
{
    "ResourceProperties": "{\\"priority\\": 105}"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Tue, 19 Mar 2024 23:36:09 GMT
Content-Type: application/json
Content-Length: 508
Connection: keep-alive
x-amzn-RequestId: dcd9d91e-1a5a-481f-82b7-bafe7dexample
Access-Control-Allow-Origin: *
x-amz-apigw-id: U5eX0FdmIexample=
X-Amzn-Trace-Id: Root=1-65fa10fe-6e5a93a56a325ab8example
{
    "ProfileResourceAssociation": {
        "CreationTime": 1710887901.139,
        "Id": "rpr-001913120a7example",
        "ModificationTime": 1710891369.579,
        "Name": "test-resource-association",
        "OwnerId": "123456789012",
        "ProfileId": "rp-4987774726example",
        "ResourceArn": "arn:aws:route53resolver:us-east-1:123456789012:firewall-rule-group/rslvr-frg-cfe7f72example",
        "ResourceProperties": "{\"priority\":105}",
        "ResourceType": "FIREWALL_RULE_GROUP",
        "Status": "UPDATING",
        "StatusMessage": "Updating the Profile to DNS Firewall rule group association"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53profiles-2018-05-10/UpdateProfileResourceAssociation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UntagResource

Amazon Route 53 Resolver
