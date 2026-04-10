This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Environment OptionSetting

The `OptionSetting` property type specifies an option for an AWS Elastic Beanstalk environment.

The `OptionSettings` property of the [AWS::ElasticBeanstalk::Environment](../userguide/aws-properties-beanstalk-environment.md) resource contains a list of
`OptionSetting` property types.

For a list of possible namespaces and option values, see [Option Values](../../../elasticbeanstalk/latest/dg/command-options.md) in the
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

- [ConfigurationOptionSetting](../../../elasticbeanstalk/latest/api/api-configurationoptionsetting.md) in the _AWS Elastic Beanstalk API Reference_

- [Configuration Options](../../../elasticbeanstalk/latest/dg/command-options.md) in the _AWS Elastic Beanstalk Developer_
_Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElasticBeanstalk::Environment

Tag

All content copied from https://docs.aws.amazon.com/.
