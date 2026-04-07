This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MWAAServerless::Workflow EncryptionConfiguration

Configuration for encrypting workflow data at rest and in transit. Amazon Managed Workflows for Apache Airflow Serverless provides comprehensive encryption capabilities to protect sensitive workflow data, parameters, and execution logs. When using customer-managed keys, the service integrates with AWSAWS KMS to provide fine-grained access control and audit capabilities. Encryption is applied consistently across the distributed execution environment including task containers, metadata storage, and log streams.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "Type" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  Type: String

```

## Properties

`KmsKeyId`

The ID or ARN of the AWS KMS key to use for encryption. Required when `Type` is `CUSTOMER_MANAGED_KEY`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of encryption to use. Values are `AWS_MANAGED_KEY` (AWS manages the encryption key) or `CUSTOMER_MANAGED_KEY` (you provide a KMS key).

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_MANAGED_KEY | CUSTOMER_MANAGED_KEY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MWAAServerless::Workflow

LoggingConfiguration
