This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::ConfigurationTemplate ConfigurationOptionSetting

The `ConfigurationOptionSetting` property type specifies an option for an AWS Elastic Beanstalk configuration template.

The `OptionSettings` property of the [AWS::ElasticBeanstalk::ConfigurationTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-beanstalk-configurationtemplate.html)
resource contains a list of `ConfigurationOptionSetting` property types.

For a list of possible namespaces and option values, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the
_AWS Elastic Beanstalk Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Namespace" : String,
  "OptionName" : String,
  "ResourceName" : String,
  "Value" : String
}

```

### YAML

```yaml

  Namespace: String
  OptionName: String
  ResourceName: String
  Value: String

```

## Properties

`Namespace`

A unique namespace that identifies the option's associated AWS resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionName`

The name of the configuration option.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceName`

A unique resource name for the option setting. Use it for a time–based scaling configuration option.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The current value for the configuration option.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConfigurationOptionSetting](https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionSetting.html) in the _AWS Elastic Beanstalk API Reference_

- [Configuration Options](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the _AWS Elastic Beanstalk Developer_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElasticBeanstalk::ConfigurationTemplate

SourceConfiguration
