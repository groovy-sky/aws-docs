This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster StepConfig

`StepConfig` is a property of the `AWS::EMR::Cluster` resource. The `StepConfig` property type specifies a cluster (job flow) step, which runs only on the master node. Steps are used to submit data processing jobs to the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionOnFailure" : String,
  "HadoopJarStep" : HadoopJarStepConfig,
  "Name" : String
}

```

### YAML

```yaml

  ActionOnFailure: String
  HadoopJarStep:
    HadoopJarStepConfig
  Name: String

```

## Properties

`ActionOnFailure`

The action to take when the cluster step fails. Possible values are `CANCEL_AND_WAIT` and `CONTINUE`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HadoopJarStep`

The JAR file used for the step.

_Required_: Yes

_Type_: [HadoopJarStepConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emr-cluster-hadoopjarstepconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the step.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SpotResizingSpecification

Tag
