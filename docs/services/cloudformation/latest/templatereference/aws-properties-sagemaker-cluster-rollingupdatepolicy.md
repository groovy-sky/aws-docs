This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster RollingUpdatePolicy

Specifies a rolling deployment strategy for updating a SageMaker endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumBatchSize" : CapacitySizeConfig,
  "RollbackMaximumBatchSize" : CapacitySizeConfig
}

```

### YAML

```yaml

  MaximumBatchSize:
    CapacitySizeConfig
  RollbackMaximumBatchSize:
    CapacitySizeConfig

```

## Properties

`MaximumBatchSize`

Batch size for each rolling step to provision capacity and turn on traffic on the new
endpoint fleet, and terminate capacity on the old endpoint fleet. Value must be between
5% to 50% of the variant's total instance count.

_Required_: Yes

_Type_: [CapacitySizeConfig](aws-properties-sagemaker-cluster-capacitysizeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollbackMaximumBatchSize`

Batch size for rollback to the old endpoint fleet. Each rolling step to provision
capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the
new endpoint fleet. If this field is absent, the default value will be set to 100% of
total capacity which means to bring up the whole capacity of the old fleet at once
during rollback.

_Required_: No

_Type_: [CapacitySizeConfig](aws-properties-sagemaker-cluster-capacitysizeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Orchestrator

ScheduledUpdateConfig

All content copied from https://docs.aws.amazon.com/.
