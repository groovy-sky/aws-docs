This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::EmailIdentity DkimSigningAttributes

Used to configure or change the DKIM authentication settings for an email domain
identity. You can use this operation to do any of the following:

- Update the signing attributes for an identity that uses Bring Your Own DKIM
(BYODKIM).

- Update the key length that should be used for Easy DKIM.

- Change from using no DKIM authentication to using Easy DKIM.

- Change from using no DKIM authentication to using BYODKIM.

- Change from using Easy DKIM to using BYODKIM.

- Change from using BYODKIM to using Easy DKIM.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainSigningPrivateKey" : String,
  "DomainSigningSelector" : String,
  "NextSigningKeyLength" : String
}

```

### YAML

```yaml

  DomainSigningPrivateKey: String
  DomainSigningSelector: String
  NextSigningKeyLength: String

```

## Properties

`DomainSigningPrivateKey`

\[Bring Your Own DKIM\] A private key that's used to generate a DKIM signature.

The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using
base64 encoding.

###### Important

Rather than embedding sensitive information directly in your CFN templates, we
recommend you use dynamic parameters in the stack template to reference sensitive
information that is stored and managed outside of CFN, such as in the AWS Systems Manager Parameter Store or AWS Secrets
Manager.

For more information, see the [Do not embed\
credentials in your templates](../userguide/best-practices.md#creds) best practice.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainSigningSelector`

\[Bring Your Own DKIM\] A string that's used to identify a public key in the DNS
configuration for a domain.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NextSigningKeyLength`

\[Easy DKIM\] The key length of the future DKIM key pair to be generated. This can be
changed at most once per day.

Valid Values: `RSA_1024_BIT | RSA_2048_BIT`

_Required_: No

_Type_: String

_Pattern_: `RSA_1024_BIT|RSA_2048_BIT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DkimAttributes

FeedbackAttributes
