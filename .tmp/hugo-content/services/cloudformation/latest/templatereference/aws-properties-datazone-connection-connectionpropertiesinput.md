This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection ConnectionPropertiesInput

The properties of a connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmazonQProperties" : AmazonQPropertiesInput,
  "AthenaProperties" : AthenaPropertiesInput,
  "GlueProperties" : GluePropertiesInput,
  "HyperPodProperties" : HyperPodPropertiesInput,
  "IamProperties" : IamPropertiesInput,
  "MlflowProperties" : MlflowPropertiesInput,
  "RedshiftProperties" : RedshiftPropertiesInput,
  "S3Properties" : S3PropertiesInput,
  "SparkEmrProperties" : SparkEmrPropertiesInput,
  "SparkGlueProperties" : SparkGluePropertiesInput,
  "WorkflowsMwaaProperties" : WorkflowsMwaaPropertiesInput,
  "WorkflowsServerlessProperties" : Json
}

```

### YAML

```yaml

  AmazonQProperties:
    AmazonQPropertiesInput
  AthenaProperties:
    AthenaPropertiesInput
  GlueProperties:
    GluePropertiesInput
  HyperPodProperties:
    HyperPodPropertiesInput
  IamProperties:
    IamPropertiesInput
  MlflowProperties:
    MlflowPropertiesInput
  RedshiftProperties:
    RedshiftPropertiesInput
  S3Properties:
    S3PropertiesInput
  SparkEmrProperties:
    SparkEmrPropertiesInput
  SparkGlueProperties:
    SparkGluePropertiesInput
  WorkflowsMwaaProperties:
    WorkflowsMwaaPropertiesInput
  WorkflowsServerlessProperties: Json

```

## Properties

`AmazonQProperties`

The Amazon Q properties of the connection.

_Required_: No

_Type_: [AmazonQPropertiesInput](aws-properties-datazone-connection-amazonqpropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AthenaProperties`

The Amazon Athena properties of a connection.

_Required_: No

_Type_: [AthenaPropertiesInput](aws-properties-datazone-connection-athenapropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueProperties`

The AWS Glue properties of a connection.

_Required_: No

_Type_: [GluePropertiesInput](aws-properties-datazone-connection-gluepropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HyperPodProperties`

The hyper pod properties of a connection.

_Required_: No

_Type_: [HyperPodPropertiesInput](aws-properties-datazone-connection-hyperpodpropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamProperties`

The IAM properties of a connection.

_Required_: No

_Type_: [IamPropertiesInput](aws-properties-datazone-connection-iampropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MlflowProperties`

The MLflow properties of a connection.

_Required_: No

_Type_: [MlflowPropertiesInput](aws-properties-datazone-connection-mlflowpropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftProperties`

The Amazon Redshift properties of a connection.

_Required_: No

_Type_: [RedshiftPropertiesInput](aws-properties-datazone-connection-redshiftpropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Properties`

The Amazon S3 properties of a connection.

_Required_: No

_Type_: [S3PropertiesInput](aws-properties-datazone-connection-s3propertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkEmrProperties`

The Spark EMR properties of a connection.

_Required_: No

_Type_: [SparkEmrPropertiesInput](aws-properties-datazone-connection-sparkemrpropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkGlueProperties`

The Spark AWS Glue properties of a connection.

_Required_: No

_Type_: [SparkGluePropertiesInput](aws-properties-datazone-connection-sparkgluepropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowsMwaaProperties`

The Amazon MWAA properties of a connection.

_Required_: No

_Type_: [WorkflowsMwaaPropertiesInput](aws-properties-datazone-connection-workflowsmwaapropertiesinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowsServerlessProperties`

The MWAA serverless properties of a connection.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthenticationCredentials

GlueConnectionInput

All content copied from https://docs.aws.amazon.com/.
