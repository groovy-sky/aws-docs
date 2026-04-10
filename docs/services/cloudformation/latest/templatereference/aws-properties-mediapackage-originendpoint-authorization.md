This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint Authorization

Parameters for enabling CDN authorization on the endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CdnIdentifierSecret" : String,
  "SecretsRoleArn" : String
}

```

### YAML

```yaml

  CdnIdentifierSecret: String
  SecretsRoleArn: String

```

## Properties

`CdnIdentifierSecret`

The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that your Content Delivery Network (CDN) uses for authorization to access your endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsRoleArn`

The Amazon Resource Name (ARN) for the IAM role that allows AWS Elemental MediaPackage to communicate with AWS Secrets Manager.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaPackage::OriginEndpoint

CmafEncryption

All content copied from https://docs.aws.amazon.com/.
