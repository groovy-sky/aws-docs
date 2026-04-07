This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpointPolicy CdnAuthConfiguration

The settings to enable CDN authorization headers in MediaPackage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CdnIdentifierSecretArns" : [ String, ... ],
  "SecretsRoleArn" : String
}

```

### YAML

```yaml

  CdnIdentifierSecretArns:
    - String
  SecretsRoleArn: String

```

## Properties

`CdnIdentifierSecretArns`

The ARN for the secret in Secrets Manager that your CDN uses for authorization to access the endpoint.

_Required_: Yes

_Type_: Array of String

_Minimum_: `20 | 1`

_Maximum_: `2048 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsRoleArn`

The ARN for the IAM role that gives MediaPackage read access to Secrets Manager and AWS KMS for CDN authorization.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaPackageV2::OriginEndpointPolicy

Next
