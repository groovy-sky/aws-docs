This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelExplainabilityJobDefinition MonitoringOutputConfig

The output configuration for monitoring jobs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "MonitoringOutputs" : [ MonitoringOutput, ... ]
}

```

### YAML

```yaml

  KmsKeyId: String
  MonitoringOutputs:
    - MonitoringOutput

```

## Properties

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker AI uses to
encrypt the model artifacts at rest using Amazon S3 server-side encryption.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitoringOutputs`

Monitoring outputs for monitoring jobs. This is where the output of the periodic
monitoring jobs is uploaded.

_Required_: Yes

_Type_: Array of [MonitoringOutput](aws-properties-sagemaker-modelexplainabilityjobdefinition-monitoringoutput.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringOutput

MonitoringResources

All content copied from https://docs.aws.amazon.com/.
