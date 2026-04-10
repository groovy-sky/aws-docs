This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition ResourceAccessPolicy

A
list of the [resources](../userguide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.md) in the group that the function can access, with the corresponding
read-only or read-write permissions. The maximum is 10 resources.

###### Note

This property applies only to Lambda functions that run in a Greengrass
container.

In
an CloudFormation template, `ResourceAccessPolicy` is a property of the
[`Environment`](../userguide/aws-properties-greengrass-functiondefinition-environment.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Permission" : String,
  "ResourceId" : String
}

```

### YAML

```yaml

  Permission: String
  ResourceId: String

```

## Properties

`Permission`

The read-only or read-write access that the Lambda function has to the
resource. Valid values are `ro` or `rw`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceId`

The ID of the resource. This ID is assigned to the resource when you create the resource
definition.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [ResourceAccessPolicy](../../../../reference/greengrass/v1/apireference/definitions-resourceaccesspolicy.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunctionDefinitionVersion

RunAs

All content copied from https://docs.aws.amazon.com/.
