This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection SparkEmrPropertiesInput

The Spark EMR properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputeArn" : String,
  "InstanceProfileArn" : String,
  "JavaVirtualEnv" : String,
  "LogUri" : String,
  "ManagedEndpointArn" : String,
  "PythonVirtualEnv" : String,
  "RuntimeRole" : String,
  "TrustedCertificatesS3Uri" : String
}

```

### YAML

```yaml

  ComputeArn: String
  InstanceProfileArn: String
  JavaVirtualEnv: String
  LogUri: String
  ManagedEndpointArn: String
  PythonVirtualEnv: String
  RuntimeRole: String
  TrustedCertificatesS3Uri: String

```

## Properties

`ComputeArn`

The compute ARN of Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso(-[bef])?))?:(elasticmapreduce|emr-serverless|emr-containers):.*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileArn`

The instance profile ARN of Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JavaVirtualEnv`

The java virtual env of the Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^[\S]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogUri`

The log URI of the Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^s3://.+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedEndpointArn`

The managed endpoint ARN of the EMR on EKS cluster.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PythonVirtualEnv`

The Python virtual env of the Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^[\S]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeRole`

The runtime role of the Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustedCertificatesS3Uri`

The certificates S3 URI of the Spark EMR.

_Required_: No

_Type_: String

_Pattern_: `^s3://.+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3PropertiesInput

SparkGlueArgs

All content copied from https://docs.aws.amazon.com/.
