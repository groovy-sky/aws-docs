This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::SourceLocation AccessConfiguration

Access configuration parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessType" : String,
  "SecretsManagerAccessTokenConfiguration" : SecretsManagerAccessTokenConfiguration
}

```

### YAML

```yaml

  AccessType: String
  SecretsManagerAccessTokenConfiguration:
    SecretsManagerAccessTokenConfiguration

```

## Properties

`AccessType`

The type of authentication used to access content from `HttpConfiguration::BaseUrl` on your source location. Accepted value: `S3_SIGV4`.

`S3_SIGV4` \- AWS Signature Version 4 authentication for Amazon S3 hosted virtual-style access. If your source location base URL is an Amazon S3 bucket, MediaTailor can use AWS Signature Version 4 (SigV4) authentication to access the bucket where your source content is stored. Your MediaTailor source location baseURL must follow the S3 virtual hosted-style request URL format. For example, https://bucket-name.s3.Region.amazonaws.com/key-name.

Before you can use `S3_SIGV4`, you must meet these requirements:

• You must allow MediaTailor to access your S3 bucket by granting mediatailor.amazonaws.com principal access in IAM. For information about configuring access in IAM, see Access management in the IAM User Guide.

• The mediatailor.amazonaws.com service principal must have permissions to read all top level manifests referenced by the VodSource packaging configurations.

• The caller of the API must have s3:GetObject IAM permissions to read all top level manifests referenced by your MediaTailor VodSource packaging configurations.

_Required_: No

_Type_: String

_Allowed values_: `S3_SIGV4 | SECRETS_MANAGER_ACCESS_TOKEN | AUTODETECT_SIGV4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessTokenConfiguration`

AWS Secrets Manager access token configuration parameters.

_Required_: No

_Type_: [SecretsManagerAccessTokenConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-sourcelocation-secretsmanageraccesstokenconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaTailor::SourceLocation

DefaultSegmentDeliveryConfiguration
