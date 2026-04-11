# DisassociateResourceFromProfile

Dissoaciated a specified resource, from the Route 53 Profile.

## Request Syntax

```nohighlight

DELETE /profileresourceassociation/profileid/ProfileId/resourcearn/ResourceArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ProfileId](#API_route53profiles_DisassociateResourceFromProfile_RequestSyntax)**

The ID of the Profile.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[ResourceArn](#API_route53profiles_DisassociateResourceFromProfile_RequestSyntax)**

The Amazon Resource Name (ARN) of the resource.

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## Request Body

The request does not have a request body.

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

**[ProfileResourceAssociation](#API_route53profiles_DisassociateResourceFromProfile_ResponseSyntax)**

Information about the `DisassociateResourceFromProfile` request, including the status of the request.

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

### DissociateProfile Example

This example illustrates one usage of DisassociateResourceFromProfile.

#### Sample Request

```

DELETE /profileresourceassociation/profileid/rp-4987774726example/resourcearn/arn%3Aaws%3Aroute53resolver%3Aus-east-1%123456789012%3Afirewall-rule-group%2Frslvr-frg-cfe7f72example HTTP/1.1
host:route53profiles.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Length: 0
X-Amz-Date:20240319T222606Z
User-Agent: aws-cli/1.32.63 botocore/1.34.63 Python/3.8.18
Authorization: AWS4-HMAC-SHA256
    Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-1/route53profiles/aws4_request,
    SignedHeaders=host;x-amz-date;x-amz-security-token
    Signature=[calculated-signature]
# RequestBody is empty
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Tue, 19 Mar 2024 22:26:06 GMT
Content-Type: application/json
Content-Length: 508
Connection: keep-alive
x-amzn-RequestId: dcd9d91e-1a5a-481f-82b7-bafe7dexample
Access-Control-Allow-Origin: *
x-amz-apigw-id: U5eX0FdmIexample=
X-Amzn-Trace-Id: Root=1-65fa10fe-6e5a93a56a325ab8example

{
    "ProfileResourceAssociation": {
        "CreationTime": 1710887120.062,
        "Id": "rpr-001913120a7example",
        "ModificationTime": 1710887166.643,
        "Name": "test-resource-association",
        "OwnerId": "123456789012",
        "ProfileId": "rp-4987774726example",
        "ResourceArn": "arn:aws:route53resolver:us-east-1:123456789012:firewall-rule-group/rslvr-frg-cfe7f72example",
        "ResourceProperties": "{\"priority\":105}",
        "ResourceType": "FIREWALL_RULE_GROUP",
        "Status": "DELETING",
        "StatusMessage": "Deleting the Profile to DNS Firewall rule group association"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53profiles-2018-05-10/disassociateresourcefromprofile.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateProfile

GetProfile
