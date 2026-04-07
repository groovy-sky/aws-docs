This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolDomain CustomDomainConfigType

The configuration for a hosted UI custom domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateArn" : String
}

```

### YAML

```yaml

  CertificateArn: String

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) of an AWS Certificate Manager SSL certificate. You use
this certificate for the subdomain of your custom domain.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Cognito::UserPoolDomain

AWS::Cognito::UserPoolGroup
