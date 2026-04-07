This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::S3TableIntegration

Creates an integration between CloudWatch and S3 Tables for analytics. This integration enables
querying CloudWatch telemetry data using analytics engines like Amazon Athena, Amazon Redshift, and Apache
Spark.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ObservabilityAdmin::S3TableIntegration",
  "Properties" : {
      "Encryption" : EncryptionConfig,
      "LogSources" : [ LogSource, ... ],
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ObservabilityAdmin::S3TableIntegration
Properties:
  Encryption:
    EncryptionConfig
  LogSources:
    - LogSource
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`Encryption`

Defines the encryption configuration for S3 Table integrations, including the encryption
algorithm and KMS key settings.

_Required_: Yes

_Type_: [EncryptionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-s3tableintegration-encryptionconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogSources`

A data source with an S3 Table integration for query access in the `logs` namespace.

_Required_: No

_Type_: Array of [LogSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-s3tableintegration-logsource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that grants permissions for the S3 Table
integration to access necessary resources.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws([a-z0-9\-]+)?:([a-zA-Z0-9\-]+):([a-z0-9\-]+)?:([0-9]{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The key-value pairs to associate with the S3 Table integration resource for categorization
and management purposes.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-s3tableintegration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The Amazon Resource Name (ARN) of the created S3 Table integration.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the S3 Table integration.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WAFLoggingParameters

EncryptionConfig
