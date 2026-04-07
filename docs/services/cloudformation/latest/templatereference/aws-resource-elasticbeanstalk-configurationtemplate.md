This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::ConfigurationTemplate

The AWS::ElasticBeanstalk::ConfigurationTemplate resource is an AWS Elastic Beanstalk
resource type that specifies an Elastic Beanstalk configuration template, associated with a
specific Elastic Beanstalk application. You define application configuration settings in a
configuration template. You can then use the configuration template to deploy different
versions of the application with the same configuration settings.

###### Note

The Elastic Beanstalk console and documentation often refer to configuration templates
as _saved configurations_. When you set configuration options in a saved
configuration (configuration template), Elastic Beanstalk applies them with a particular
precedence as part of applying options from multiple sources. For more information, see
[Configuration Options](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the _AWS Elastic Beanstalk Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticBeanstalk::ConfigurationTemplate",
  "Properties" : {
      "ApplicationName" : String,
      "Description" : String,
      "EnvironmentId" : String,
      "OptionSettings" : [ ConfigurationOptionSetting, ... ],
      "PlatformArn" : String,
      "SolutionStackName" : String,
      "SourceConfiguration" : SourceConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::ElasticBeanstalk::ConfigurationTemplate
Properties:
  ApplicationName: String
  Description: String
  EnvironmentId: String
  OptionSettings:
    - ConfigurationOptionSetting
  PlatformArn: String
  SolutionStackName: String
  SourceConfiguration:
    SourceConfiguration

```

## Properties

`ApplicationName`

The name of the Elastic Beanstalk application to associate with this configuration
template.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

An optional description for this configuration.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentId`

The ID of an environment whose settings you want to use to create the configuration
template. You must specify `EnvironmentId` if you don't specify
`PlatformArn`, `SolutionStackName`, or
`SourceConfiguration`.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OptionSettings`

Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these
values override the values obtained from the solution stack or the source configuration
template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the
_AWS Elastic Beanstalk Developer Guide_.

_Required_: No

_Type_: Array of [ConfigurationOptionSetting](aws-properties-elasticbeanstalk-configurationtemplate-configurationoptionsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlatformArn`

The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the _AWS Elastic Beanstalk Developer Guide_.

###### Note

If you specify `PlatformArn`, then don't specify `SolutionStackName`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SolutionStackName`

The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For
example, `64bit Amazon Linux 2013.09 running Tomcat 7 Java 7`. A solution stack
specifies the operating system, runtime, and application server for a configuration template.
It also determines the set of configuration options as well as the possible and default
values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the
_AWS Elastic Beanstalk Developer Guide_.

You must specify `SolutionStackName` if you don't specify
`PlatformArn`, `EnvironmentId`, or
`SourceConfiguration`.

Use the [`ListAvailableSolutionStacks`](https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ListAvailableSolutionStacks.html) API to obtain a list of available
solution stacks.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceConfiguration`

An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified
configuration template to create a new configuration.

Values specified in `OptionSettings` override any values obtained from the
`SourceConfiguration`.

You must specify `SourceConfiguration` if you don't specify
`PlatformArn`, `EnvironmentId`, or
`SolutionStackName`.

Constraint: If both solution stack name and source configuration are specified, the
solution stack of the source configuration template must match the specified solution stack
name.

_Required_: Conditional

_Type_: [SourceConfiguration](aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`TemplateName`

The name of the configuration template.

Constraint: This name must be unique per application.

## Examples

#### JSON

```json

"myConfigTemplate" : {
    "Type" : "AWS::ElasticBeanstalk::ConfigurationTemplate",
    "Properties" : {
      "ApplicationName" :{"Ref" : "myApp"},
      "Description" : "my sample configuration template",
      "EnvironmentId" : "",
      "SourceConfiguration" : {
        "ApplicationName" : {"Ref" : "mySecondApp"},
        "TemplateName" : {"Ref" : "mySourceTemplate"}
      },
      "SolutionStackName" : "64bit Amazon Linux running PHP 5.3",
      "OptionSettings" : [ {
        "Namespace" : "aws:autoscaling:launchconfiguration",
        "OptionName" : "EC2KeyName",
        "Value" : { "Ref" : "KeyName" }
      } ]
    }
  }
```

#### YAML

```yaml

myConfigTemplate:
  Type: AWS::ElasticBeanstalk::ConfigurationTemplate
  Properties:
    ApplicationName:
      Ref: "myApp"
    Description: "my sample configuration template"
    EnvironmentId: ""
    SourceConfiguration:
      ApplicationName:
        Ref: "mySecondApp"
      TemplateName:
        Ref: "mySourceTemplate"
    SolutionStackName: "64bit Amazon Linux running PHP 5.3"
    OptionSettings:
      -
        Namespace: "aws:autoscaling:launchconfiguration"
        OptionName: "EC2KeyName"
        Value:
          Ref: "KeyName"
```

## See also

- [AWS::ElasticBeanstalk::Application](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html)

- [Configuration Options](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the _AWS Elastic Beanstalk Developer_
_Guide_

- For a complete Elastic Beanstalk sample template, see [Elastic\
Beanstalk Template Snippets](../userguide/quickref-elasticbeanstalk.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SourceBundle

ConfigurationOptionSetting
