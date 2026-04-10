This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary ArtifactConfig

A structure that contains the configuration for canary artifacts,
including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Encryption" : S3Encryption
}

```

### YAML

```yaml

  S3Encryption:
    S3Encryption

```

## Properties

`S3Encryption`

A structure that contains the configuration
of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3
or later. For more information, see
[Encrypting canary artifacts](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-artifact-encryption.md).

_Required_: No

_Type_: [S3Encryption](aws-properties-synthetics-canary-s3encryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Synthetics::Canary

BaseScreenshot

All content copied from https://docs.aws.amazon.com/.
