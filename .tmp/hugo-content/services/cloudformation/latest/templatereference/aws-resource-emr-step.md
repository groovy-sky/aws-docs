This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Step

Use `Step` to specify a cluster (job flow) step, which runs only on the master node. Steps are used to submit data processing jobs to a cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::Step",
  "Properties" : {
      "ActionOnFailure" : String,
      "EncryptionKeyArn" : String,
      "HadoopJarStep" : HadoopJarStepConfig,
      "JobFlowId" : String,
      "LogUri" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::EMR::Step
Properties:
  ActionOnFailure: String
  EncryptionKeyArn: String
  HadoopJarStep:
    HadoopJarStepConfig
  JobFlowId: String
  LogUri: String
  Name: String

```

## Properties

`ActionOnFailure`

This specifies what action to take when the cluster step fails. Possible values are `CANCEL_AND_WAIT` and `CONTINUE`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionKeyArn`

The KMS key ARN to encrypt the logs published to the given Amazon S3 destination. When omitted, EMR falls back to cluster-level logging behavior.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HadoopJarStep`

The `HadoopJarStepConfig` property type specifies a job flow step consisting of a JAR file whose main function will be executed. The main function submits a job for the cluster to execute as a step on the master node, and then waits for the job to finish or fail before executing subsequent steps.

_Required_: Yes

_Type_: [HadoopJarStepConfig](aws-properties-emr-step-hadoopjarstepconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobFlowId`

A string that uniquely identifies the cluster (job flow).

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogUri`

The Amazon S3 destination URI for log publishing. When omitted, EMR falls back to cluster-level logging behavior.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the cluster step.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the ID of the step.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Id`

The identifier of the cluster step.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::SecurityConfiguration

HadoopJarStepConfig

All content copied from https://docs.aws.amazon.com/.
