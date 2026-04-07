This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::DomainConfiguration ServerCertificateConfig

The server certificate configuration.

For more information, see [Configurable\
endpoints](https://docs.aws.amazon.com/iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableOCSPCheck" : Boolean,
  "OcspAuthorizedResponderArn" : String,
  "OcspLambdaArn" : String
}

```

### YAML

```yaml

  EnableOCSPCheck: Boolean
  OcspAuthorizedResponderArn: String
  OcspLambdaArn: String

```

## Properties

`EnableOCSPCheck`

A Boolean value that indicates whether Online Certificate Status Protocol (OCSP) server
certificate check is enabled or not. For more information, see [Configurable\
endpoints](https://docs.aws.amazon.com/iot/latest/developerguide/iot-custom-endpoints-configurable.html) from the AWS IoT Core Developer Guide.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OcspAuthorizedResponderArn`

The Amazon Resource Name (ARN) for an X.509 certificate stored in ACM.
If provided, AWS IoT Core will use this certificate to validate the signature
of the received OCSP response. The OCSP responder must sign responses using either this
authorized responder certificate or the issuing certificate, depending on whether the ARN
is provided or not. The certificate must be in the same account and region as the
domain configuration.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-cn|-us-gov|-iso-b|-iso)?:acm:[a-z]{2}-(gov-|iso-|isob-)?[a-z]{4,9}-\d{1}:\d{12}:certificate/[a-zA-Z0-9/-]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OcspLambdaArn`

The Amazon Resource Name (ARN) for a Lambda function that acts as a Request for Comments
(RFC) 6960-compliant Online Certificate Status Protocol (OCSP) responder, supporting basic
OCSP responses. The Lambda function accepts a base64-encoding of the OCSP request in the
Distinguished Encoding Rules (DER) format. The Lambda function's response is also a
base64-encoded OCSP response in the DER format. The response size must not exceed 4
kilobytes (KiB). The Lambda function must be in the same account and region as
the domain configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClientCertificateConfig

ServerCertificateSummary
