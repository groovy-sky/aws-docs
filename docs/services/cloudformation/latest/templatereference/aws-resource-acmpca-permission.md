This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::Permission

Grants permissions to the AWS Certificate Manager (ACM) service
principal ( `acm.amazonaws.com`) to perform [IssueCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html), [GetCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificate.html), and [ListPermissions](https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListPermissions.html)
actions on a CA. These actions are needed for the ACM principal to renew private PKI
certificates requested through ACM and residing in the same AWS account
as the CA.

###### About permissions

- If the private CA and the certificates it issues reside in the same account,
you can use `AWS::ACMPCA::Permission` to grant permissions for ACM to
carry out automatic certificate renewals.

- For automatic certificate renewal to succeed, the ACM service principal needs
permissions to create, retrieve, and list permissions.

- If the private CA and the ACM certificates reside in different accounts, then
permissions cannot be used to enable automatic renewals. Instead, the ACM
certificate owner must set up a resource-based policy to enable cross-account
issuance and renewals. For more information, see [Using a Resource Based Policy\
with AWS Private CA](https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html).

###### Note

To update an `AWS::ACMPCA::Permission` resource, you must first delete
the existing permission resource from the CloudFormation stack and then create a new
permission resource with updated properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ACMPCA::Permission",
  "Properties" : {
      "Actions" : [ String, ... ],
      "CertificateAuthorityArn" : String,
      "Principal" : String,
      "SourceAccount" : String
    }
}

```

### YAML

```yaml

Type: AWS::ACMPCA::Permission
Properties:
  Actions:
    - String
  CertificateAuthorityArn: String
  Principal: String
  SourceAccount: String

```

## Properties

`Actions`

The private CA actions that can be performed by the designated AWS
service. Supported actions are `IssueCertificate`,
`GetCertificate`, and `ListPermissions`.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CertificateAuthorityArn`

The Amazon Resource Number (ARN) of the private CA from which the permission was
issued.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:acm-pca:[\w+=/,.@-]*:[0-9]*:[\w+=,.@-]+(/[\w+=,.@-]+)*`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

The AWS service or entity that holds the permission. At this time, the only valid
principal is `acm.amazonaws.com`.

_Required_: Yes

_Type_: String

_Pattern_: `[^*]+`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceAccount`

The ID of the account that assigned the permission.

_Required_: No

_Type_: String

_Pattern_: `[0-9]+`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ACMPCA::CertificateAuthorityActivation

Next
