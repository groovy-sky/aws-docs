This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::CertificateProvider

Creates a certificate provider. AWS IoT Core certificate provider lets you
customize how to sign a certificate signing request (CSR) in fleet provisioning. For more
information, see [Self-managed certificate\
signing using AWS IoT Corecertificate provider](../../../iot/latest/developerguide/provisioning-cert-provider.md) from the
_AWS IoT Core Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::CertificateProvider",
  "Properties" : {
      "AccountDefaultForOperations" : [ String, ... ],
      "CertificateProviderName" : String,
      "LambdaFunctionArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::CertificateProvider
Properties:
  AccountDefaultForOperations:
    - String
  CertificateProviderName: String
  LambdaFunctionArn: String
  Tags:
    - Tag

```

## Properties

`AccountDefaultForOperations`

A list of the operations that the certificate provider will use to generate
certificates. Valid value: `CreateCertificateFromCsr`.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateProviderName`

The name of the certificate provider.

_Required_: No

_Type_: String

_Pattern_: `[\w=,@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LambdaFunctionArn`

The ARN of the Lambda function.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the certificate provider.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-certificateprovider-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the certificate provider. For example:

`{ "Ref": "MyCertificateProvider" }`

A value similar to the following is returned:

`a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the certificate. For example:

`{ "Fn::GetAtt": ["MyCertificateProvider", "Arn"] }`

A value similar to the following is returned:

`arn:aws:iot:ap-southeast-2:123456789012:certprovider/my-certificate-provider`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::Certificate

Tag

All content copied from https://docs.aws.amazon.com/.
