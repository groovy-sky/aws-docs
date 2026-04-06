# UpdateCertificateOptions

Updates a certificate. You can use this function to specify whether to opt in to or
out of recording your certificate in a certificate transparency log and exporting. For
more information, see [Opting Out of\
Certificate Transparency Logging](../../../../services/acm/latest/userguide/acm-bestpractices.md#best-practices-transparency) and [AWS Certificate Manager Exportable\
Managed Certificates](../../../../services/acm/latest/userguide/acm-exportable-certificates.md).

## Request Syntax

```nohighlight

{
   "CertificateArn": "string",
   "Options": {
      "CertificateTransparencyLoggingPreference": "string",
      "Export": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[CertificateArn](#API_UpdateCertificateOptions_RequestSyntax)**

ARN of the requested certificate to update. This must be of the form:

`arn:aws:acm:us-east-1:account:certificate/12345678-1234-1234-1234-123456789012
                  `

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: Yes

**[Options](#API_UpdateCertificateOptions_RequestSyntax)**

Use to update the options for your certificate. Currently, you can specify whether to
add your certificate to a transparency log or export your certificate. Certificate
transparency makes it possible to detect SSL/TLS certificates that have been mistakenly
or maliciously issued. Certificates that have not been logged typically produce an error
message in a browser.

Type: [CertificateOptions](api-certificateoptions.md) object

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidArnException**

The requested Amazon Resource Name (ARN) does not refer to an existing
resource.

HTTP Status Code: 400

**InvalidStateException**

Processing has reached an invalid state.

HTTP Status Code: 400

**LimitExceededException**

An ACM quota has been exceeded.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified certificate cannot be found in the caller's account or the caller's
account cannot be found.

HTTP Status Code: 400

## Examples

### UpdateCertificateOptions

This example illustrates one usage of UpdateCertificateOptions.

#### Sample Request

```

POST / HTTP/1.1
acm.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Length: 185
X-Amz-Target: CertificateManager.UpdateCertificateOptions
X-Amz-Date: 20180326T222032Z
User-Agent: aws-cli/1.14.28 Python/2.7.9 Windows/8 botocore/1.8.32
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256 Credential=key_ID/20151222/us-east-1/acm/aws4_request,
SignedHeaders=content-type;host;user-agent;x-amz-date;x-amz-target,
Signature=7ec7e70cd614724945545b22bc28296f77803d0c2524573d41c994668f07f435

{
  "CertificateArn": "arn:aws:acm:region:account:certificate/12345678-1234-1234-1234-123456789012",
  "CertificateOptions": {
    "CertificateTransparencyLoggingPreference": "DISABLED"
  }
}
```

### Example

This example illustrates one usage of UpdateCertificateOptions.

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: e6f55ecb-3143-11e8-af72-0bd5049841d5
Content-Type: application/x-amz-json-1.1
Content-Length: 0
Date: Tue, 22 Dec 2015 17:07:18 GMT

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/acm-2015-12-08/UpdateCertificateOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/UpdateCertificateOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RevokeCertificate

Data Types
