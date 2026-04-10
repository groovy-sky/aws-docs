This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Application MaxCountRule

A lifecycle rule that deletes the oldest application version when the maximum count is
exceeded.

`MaxCountRule` is a property of the [ApplicationVersionLifecycleConfig](../userguide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.md)
property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteSourceFromS3" : Boolean,
  "Enabled" : Boolean,
  "MaxCount" : Integer
}

```

### YAML

```yaml

  DeleteSourceFromS3: Boolean
  Enabled: Boolean
  MaxCount: Integer

```

## Properties

`DeleteSourceFromS3`

Set to `true` to delete a version's source bundle from Amazon S3 when
Elastic Beanstalk deletes the application version.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specify `true` to apply the rule, or `false` to disable
it.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCount`

Specify the maximum number of application versions to retain.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaxAgeRule

AWS::ElasticBeanstalk::ApplicationVersion

All content copied from https://docs.aws.amazon.com/.
