This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Application ApplicationVersionLifecycleConfig

The application version lifecycle settings for an application. Defines the rules that
Elastic Beanstalk applies to an application's versions in order to avoid hitting the
per-region limit for application versions.

When Elastic Beanstalk deletes an application version from its database, you can no
longer deploy that version to an environment. The source bundle remains in S3 unless you
configure the rule to delete it.

`ApplicationVersionLifecycleConfig` is a property of the [ApplicationResourceLifecycleConfig](../userguide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.md)
property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxAgeRule" : MaxAgeRule,
  "MaxCountRule" : MaxCountRule
}

```

### YAML

```yaml

  MaxAgeRule:
    MaxAgeRule
  MaxCountRule:
    MaxCountRule

```

## Properties

`MaxAgeRule`

Specify a max age rule to restrict the length of time that application versions are
retained for an application.

_Required_: No

_Type_: [MaxAgeRule](aws-properties-elasticbeanstalk-application-maxagerule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCountRule`

Specify a max count rule to restrict the number of application versions that are
retained for an application.

_Required_: No

_Type_: [MaxCountRule](aws-properties-elasticbeanstalk-application-maxcountrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationResourceLifecycleConfig

MaxAgeRule

All content copied from https://docs.aws.amazon.com/.
