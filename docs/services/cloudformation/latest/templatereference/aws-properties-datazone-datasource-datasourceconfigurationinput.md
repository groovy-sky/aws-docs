This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource DataSourceConfigurationInput

The configuration of the data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlueRunConfiguration" : GlueRunConfigurationInput,
  "RedshiftRunConfiguration" : RedshiftRunConfigurationInput,
  "SageMakerRunConfiguration" : SageMakerRunConfigurationInput
}

```

### YAML

```yaml

  GlueRunConfiguration:
    GlueRunConfigurationInput
  RedshiftRunConfiguration:
    RedshiftRunConfigurationInput
  SageMakerRunConfiguration:
    SageMakerRunConfigurationInput

```

## Properties

`GlueRunConfiguration`

The configuration of the AWS Glue data source.

_Required_: No

_Type_: [GlueRunConfigurationInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-gluerunconfigurationinput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftRunConfiguration`

The configuration of the Amazon Redshift data source.

_Required_: No

_Type_: [RedshiftRunConfigurationInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SageMakerRunConfiguration`

The Amazon SageMaker run configuration.

_Required_: No

_Type_: [SageMakerRunConfigurationInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-sagemakerrunconfigurationinput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::DataSource

FilterExpression
