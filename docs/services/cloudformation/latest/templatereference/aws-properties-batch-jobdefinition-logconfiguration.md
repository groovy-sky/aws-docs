This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition LogConfiguration

Log configuration options to send to a custom log driver for the container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogDriver" : String,
  "Options" : {Key: Value, ...},
  "SecretOptions" : [ Secret, ... ]
}

```

### YAML

```yaml

  LogDriver: String
  Options:
    Key: Value
  SecretOptions:
    - Secret

```

## Properties

`LogDriver`

The log driver to use for the container. The valid values that are listed for this parameter
are log drivers that the Amazon ECS container agent can communicate with by default.

The supported log drivers are `awsfirelens`, `awslogs`, `fluentd`, `gelf`,
`json-file`, `journald`, `logentries`, `syslog`, and
`splunk`.

###### Note

Jobs that are running on Fargate resources are restricted to the `awslogs` and
`splunk` log drivers.

awsfirelens

Specifies the firelens logging driver. For more information on configuring Firelens, see
[Send\
Amazon ECS logs to an AWS service or AWS Partner](../../../amazonecs/latest/developerguide/using-firelens.md) in the
_Amazon Elastic Container Service Developer Guide_.

awslogs

Specifies the Amazon CloudWatch Logs logging driver. For more information, see [Using the awslogs log driver](../../../batch/latest/userguide/using-awslogs.md)
in the _AWS Batch User Guide_ and [Amazon CloudWatch Logs logging\
driver](https://docs.docker.com/config/containers/logging/awslogs) in the Docker documentation.

fluentd

Specifies the Fluentd logging driver. For more information including usage and options,
see [Fluentd logging\
driver](https://docs.docker.com/config/containers/logging/fluentd) in the _Docker documentation_.

gelf

Specifies the Graylog Extended Format (GELF) logging driver. For more information
including usage and options, see [Graylog Extended Format logging\
driver](https://docs.docker.com/config/containers/logging/gelf) in the _Docker documentation_.

journald

Specifies the journald logging driver. For more information including usage and options,
see [Journald logging\
driver](https://docs.docker.com/config/containers/logging/journald) in the _Docker documentation_.

json-file

Specifies the JSON file logging driver. For more information including usage and options,
see [JSON File\
logging driver](https://docs.docker.com/config/containers/logging/json-file) in the _Docker documentation_.

splunk

Specifies the Splunk logging driver. For more information including usage and options,
see [Splunk logging\
driver](https://docs.docker.com/config/containers/logging/splunk) in the _Docker documentation_.

syslog

Specifies the syslog logging driver. For more information including usage and options,
see [Syslog logging\
driver](https://docs.docker.com/config/containers/logging/syslog) in the _Docker documentation_.

###### Note

If you have a custom driver that's not listed earlier that you want to work with the Amazon ECS
container agent, you can fork the Amazon ECS container agent project that's [available on GitHub](https://github.com/aws/amazon-ecs-agent) and customize it to
work with that driver. We encourage you to submit pull requests for changes that you want to
have included. However, Amazon Web Services doesn't currently support running modified copies of this
software.

This parameter requires version 1.18 of the Docker Remote API or greater on your
container instance. To check the Docker Remote API version on your container instance, log in to your
container instance and run the following command: `sudo docker version | grep "Server API version"`

_Required_: Yes

_Type_: String

_Allowed values_: `json-file | syslog | journald | gelf | fluentd | awslogs | splunk | awsfirelens`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

The configuration options to send to the log driver. This parameter requires version 1.19 of the Docker Remote API or greater on your
container instance. To check the Docker Remote API version on your container instance, log in to your
container instance and run the following command: `sudo docker version | grep "Server API version"`

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretOptions`

The secrets to pass to the log configuration. For more information, see [Specifying sensitive\
data](../../../batch/latest/userguide/specifying-sensitive-data.md) in the _AWS Batch User Guide_.

_Required_: No

_Type_: Array of [Secret](aws-properties-batch-jobdefinition-secret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LinuxParameters

MountPoint

All content copied from https://docs.aws.amazon.com/.
