This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::ServerCertificate

Uploads a server certificate entity for the AWS account. The server certificate
entity includes a public key certificate, a private key, and an optional certificate
chain, which should all be PEM-encoded.

We recommend that you use [AWS Certificate Manager](../../../acm/index.md) to
provision, manage, and deploy your server certificates. With ACM you can request a
certificate, deploy it to AWS resources, and let ACM handle certificate renewals for
you. Certificates provided by ACM are free. For more information about using ACM,
see the [AWS Certificate Manager User\
Guide](../../../acm/latest/userguide.md).

For more information about working with server certificates, see [Working\
with server certificates](../../../iam/latest/userguide/id-credentials-server-certs.md) in the _IAM User Guide_. This
topic includes a list of AWS services that can use the server certificates that you
manage with IAM.

For information about the number of server certificates you can upload, see [IAM and AWS STS\
quotas](../../../iam/latest/userguide/reference-iam-quotas.md) in the _IAM User Guide_.

###### Note

Because the body of the public key certificate, private key, and the certificate
chain can be large, you should use POST rather than GET when calling
`UploadServerCertificate`. For information about setting up
signatures and authorization through the API, see [Signing AWS API\
requests](../../../../general/latest/gr/signing-aws-api-requests.md) in the _AWS General Reference_. For general
information about using the Query API with IAM, see [Calling the API by making HTTP query\
requests](../../../iam/latest/userguide/programming.md) in the _IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::ServerCertificate",
  "Properties" : {
      "CertificateBody" : String,
      "CertificateChain" : String,
      "Path" : String,
      "PrivateKey" : String,
      "ServerCertificateName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::ServerCertificate
Properties:
  CertificateBody: String
  CertificateChain: String
  Path: String
  PrivateKey: String
  ServerCertificateName: String
  Tags:
    - Tag

```

## Properties

`CertificateBody`

The contents of the public key certificate.

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateChain`

The contents of the public key certificate chain.

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `2097152`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

The path for the server certificate. For more information about paths, see [IAM\
identifiers](../../../iam/latest/userguide/using-identifiers.md) in the _IAM User Guide_.

This parameter is optional. If it is not included, it defaults to a slash (/).
This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting
of either a forward slash (/) by itself or a string that must begin and end with forward slashes.
In addition, it can contain any ASCII character from the ! ( `\u0021`) through the DEL character ( `\u007F`), including
most punctuation characters, digits, and upper and lowercased letters.

###### Note

If you are uploading a server certificate specifically for use with Amazon
CloudFront distributions, you must specify a path using the `path`
parameter. The path must begin with `/cloudfront` and must include a
trailing slash (for example, `/cloudfront/test/`).

_Required_: No

_Type_: String

_Pattern_: `(\u002F)|(\u002F[\u0021-\u007F]+\u002F)`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKey`

The contents of the private key in PEM-encoded format.

The [regex pattern](http://wikipedia.org/wiki/regex)
used to validate this parameter is a string of characters consisting of the following:

- Any printable ASCII
character ranging from the space character ( `\u0020`) through the end of the ASCII character range

- The printable characters in the Basic Latin and Latin-1 Supplement character set
(through `\u00FF`)

- The special characters tab ( `\u0009`), line feed ( `\u000A`), and
carriage return ( `\u000D`)

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerCertificateName`

The name for the server certificate. Do not include the path in this value. The name
of the certificate cannot contain any spaces.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags that are attached to the server certificate. For more information about tagging, see [Tagging IAM resources](../../../iam/latest/userguide/id-tags.md) in the
_IAM User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-iam-servercertificate-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ServerCertificateName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the specified
`AWS::IAM::ServerCertificate` resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
