This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition ResourceDefinitionVersion

A resource definition version contains a list of resources. (In CloudFormation,
resources are named _resource instances_.)

###### Note

After you create a resource definition version that contains the resources you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In an CloudFormation template, `ResourceDefinitionVersion` is the property
type of the `InitialVersion` property in the [`AWS::Greengrass::ResourceDefinition`](../userguide/aws-resource-greengrass-resourcedefinition.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Resources" : [ ResourceInstance, ... ]
}

```

### YAML

```yaml

  Resources:
    - ResourceInstance

```

## Properties

`Resources`

The resources in this version.

_Required_: Yes

_Type_: Array of [ResourceInstance](aws-properties-greengrass-resourcedefinition-resourceinstance.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [ResourceDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-resourcedefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceDataContainer

ResourceDownloadOwnerSetting

All content copied from https://docs.aws.amazon.com/.
