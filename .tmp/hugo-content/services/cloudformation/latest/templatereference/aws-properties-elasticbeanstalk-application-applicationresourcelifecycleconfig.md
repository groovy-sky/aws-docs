This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Application ApplicationResourceLifecycleConfig

The resource lifecycle configuration for an application. Defines lifecycle settings for
resources that belong to the application, and the service role that Elastic Beanstalk assumes
in order to apply lifecycle settings. The version lifecycle configuration defines lifecycle
settings for application versions.

`ApplicationResourceLifecycleConfig` is a property of the [AWS::ElasticBeanstalk::Application](../userguide/aws-properties-beanstalk.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServiceRole" : String,
  "VersionLifecycleConfig" : ApplicationVersionLifecycleConfig
}

```

### YAML

```yaml

  ServiceRole: String
  VersionLifecycleConfig:
    ApplicationVersionLifecycleConfig

```

## Properties

`ServiceRole`

The ARN of an IAM service role that Elastic Beanstalk has permission to
assume.

The `ServiceRole` property is required the first time that you provide a `ResourceLifecycleConfig` for the application.
After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again.
You can, however, specify it in subsequent updates to change the Service Role to another value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionLifecycleConfig`

Defines lifecycle settings for application versions.

_Required_: No

_Type_: [ApplicationVersionLifecycleConfig](aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElasticBeanstalk::Application

ApplicationVersionLifecycleConfig

All content copied from https://docs.aws.amazon.com/.
