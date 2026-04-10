This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition FirelensConfiguration

The FireLens configuration for the container. This is used to specify and configure a
log router for container logs. For more information, see [Custom log](../../../amazonecs/latest/developerguide/using-firelens.md) routing in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Options" : {Key: Value, ...},
  "Type" : String
}

```

### YAML

```yaml

  Options:
    Key: Value
  Type: String

```

## Properties

`Options`

The options to use when configuring the log router. This field is optional and can be
used to specify a custom configuration file or to add additional metadata, such as the
task, task definition, cluster, and container instance details to the log event. If
specified, the syntax to use is
`"options":{"enable-ecs-log-metadata":"true|false","config-file-type:"s3|file","config-file-value":"arn:aws:s3:::mybucket/fluent.conf|filepath"}`.
For more information, see [Creating a task definition that uses a FireLens configuration](../../../amazonecs/latest/developerguide/using-firelens.md#firelens-taskdef)
in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The log router to use. The valid values are `fluentd` or `fluentbit`.

_Required_: Yes

_Type_: String

_Allowed values_: `fluentd | fluentbit`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FargatePlatformConfiguration

Host

All content copied from https://docs.aws.amazon.com/.
