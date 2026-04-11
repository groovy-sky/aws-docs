This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinitionVersion RunAs

The user and
group permissions used to run the Lambda function. This setting overrides the
default access identity that's specified for the group (by default, ggc\_user and
ggc\_group). You can override the user, group, or both. For more information, see [Run as](../../../greengrass/v1/developerguide/lambda-group-config.md#lambda-access-identity.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

###### Important

Running as the root user increases risks to your data and device. Do not run as root
(UID/GID=0) unless your business case requires it. For more information and
requirements, see [Running a Lambda Function as Root](../../../greengrass/v1/developerguide/lambda-group-config.md#lambda-running-as-root).

In an
CloudFormation template, `RunAs` is a property of the [`Execution`](../userguide/aws-properties-greengrass-functiondefinitionversion-execution.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Gid" : Integer,
  "Uid" : Integer
}

```

### YAML

```yaml

  Gid: Integer
  Uid: Integer

```

## Properties

`Gid`

The group ID whose permissions are used to run the Lambda function. You
can use the **getent group** command on your core device to look up the
group ID.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Uid`

The user ID whose permissions are used to run the Lambda function. You can
use the **getent passwd** command on your core device to look up the user
ID.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionRunAsConfig](../../../../reference/greengrass/v1/apireference/definitions-functionrunasconfig.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceAccessPolicy

AWS::Greengrass::Group

All content copied from https://docs.aws.amazon.com/.
