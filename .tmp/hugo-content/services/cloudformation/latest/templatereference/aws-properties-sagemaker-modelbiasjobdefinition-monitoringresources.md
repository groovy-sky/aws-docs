This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelBiasJobDefinition MonitoringResources

Identifies the resources to deploy for a monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterConfig" : ClusterConfig
}

```

### YAML

```yaml

  ClusterConfig:
    ClusterConfig

```

## Properties

`ClusterConfig`

The configuration for the cluster resources used to run the processing job.

_Required_: Yes

_Type_: [ClusterConfig](aws-properties-sagemaker-modelbiasjobdefinition-clusterconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringOutputConfig

NetworkConfig

All content copied from https://docs.aws.amazon.com/.
