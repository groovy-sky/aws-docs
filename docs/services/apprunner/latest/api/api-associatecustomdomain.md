# AssociateCustomDomain

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Associate your own domain name with the AWS App Runner subdomain URL of your App Runner service.

After you call `AssociateCustomDomain` and receive a successful response, use the information in the [CustomDomain](api-customdomain.md) record
that's returned to add CNAME records to your Domain Name System (DNS). For each mapped domain name, add a mapping to the target App Runner subdomain and one or
more certificate validation records. App Runner then performs DNS validation to verify that you own or control the domain name that you associated. App Runner tracks
domain validity in a certificate stored in [AWS Certificate Manager (ACM)](../../../acm/latest/userguide.md).

## Request Syntax

```nohighlight

{
   "DomainName": "string",
   "EnableWWWSubdomain": boolean,
   "ServiceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DomainName](#API_AssociateCustomDomain_RequestSyntax)**

A custom domain endpoint to associate. Specify a root domain (for example, `example.com`), a subdomain (for example,
`login.example.com` or `admin.login.example.com`), or a wildcard (for example, `*.example.com`).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z0-9*.-]{1,255}`

Required: Yes

**[EnableWWWSubdomain](#API_AssociateCustomDomain_RequestSyntax)**

Set to `true` to associate the subdomain `www.DomainName
                  ` with the App Runner service in addition to the base
domain.

Default: `true`

Type: Boolean

Required: No

**[ServiceArn](#API_AssociateCustomDomain_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner service that you want to associate a custom domain name with.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

## Response Syntax

```nohighlight

{
   "CustomDomain": {
      "CertificateValidationRecords": [
         {
            "Name": "string",
            "Status": "string",
            "Type": "string",
            "Value": "string"
         }
      ],
      "DomainName": "string",
      "EnableWWWSubdomain": boolean,
      "Status": "string"
   },
   "DNSTarget": "string",
   "ServiceArn": "string",
   "VpcDNSTargets": [
      {
         "DomainName": "string",
         "VpcId": "string",
         "VpcIngressConnectionArn": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[CustomDomain](#API_AssociateCustomDomain_ResponseSyntax)**

A description of the domain name that's being associated.

Type: [CustomDomain](api-customdomain.md) object

**[DNSTarget](#API_AssociateCustomDomain_ResponseSyntax)**

The App Runner subdomain of the App Runner service. The custom domain name is mapped to this target name.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

**[ServiceArn](#API_AssociateCustomDomain_ResponseSyntax)**

The Amazon Resource Name (ARN) of the App Runner service with which a custom domain name is associated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

**[VpcDNSTargets](#API_AssociateCustomDomain_ResponseSyntax)**

DNS Target records for the custom domains of this Amazon VPC.

Type: Array of [VpcDNSTarget](api-vpcdnstarget.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**InvalidStateException**

You can't perform this action when the resource is in its current state.

HTTP Status Code: 400

## Examples

### Associate a domain name and the www subdomain with a service

This example illustrates how to associate a custom domain name that you control with an App Runner service. The domain name is the root domain
`example.com`, including the special-case subdomain `www.example.com`.

###### Note

`CertificateValidationRecords` is an optional field and returns an empty response for `AssociateCustomDomain` APIs.

#### Sample Request

```json

$ aws apprunner associate-custom-domain --cli-input-json "`cat`"
{
  "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa",
  "DomainName": "example.com",
  "EnableWWWSubdomain": true
}
```

#### Sample Response

```json

{
  "DNSTarget": "zgz2t7wmhi.us-east-1.awsapprunner.com",
  "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa",
  "CustomDomain": {
    "DomainName": "example.com",
    "EnableWWWSubdomain": true,
    "Status": "creating"
  },
  "VpcDNSTargets": [
    {
      "DomainName": "psbqam834h.us-east-1.awsapprunner.com",
      "VpcId": "vpc-4a5b6c7d",
      "VpcIngressConnectionArn": "arn:aws:apprunner:us-east-1:123456789012:vpcingressconnection/my-ingress-connection-name/3f2eb10e2c494674952026f646844e3d"
    }
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/associatecustomdomain.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/associatecustomdomain.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

CreateAutoScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
