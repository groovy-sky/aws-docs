This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdNamespace

Creates an ID namespace object which will help customers provide metadata explaining
their dataset and how to use it. Each ID namespace must have a unique name. To modify an
existing ID namespace, use the UpdateIdNamespace API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EntityResolution::IdNamespace",
  "Properties" : {
      "Description" : String,
      "IdMappingWorkflowProperties" : [ IdNamespaceIdMappingWorkflowProperties, ... ],
      "IdNamespaceName" : String,
      "InputSourceConfig" : [ IdNamespaceInputSource, ... ],
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EntityResolution::IdNamespace
Properties:
  Description: String
  IdMappingWorkflowProperties:
    - IdNamespaceIdMappingWorkflowProperties
  IdNamespaceName: String
  InputSourceConfig:
    - IdNamespaceInputSource
  RoleArn: String
  Tags:
    - Tag
  Type: String

```

## Properties

`Description`

The description of the ID namespace.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdMappingWorkflowProperties`

Determines the properties of `IdMappingWorflow` where this
`IdNamespace` can be used as a `Source` or a
`Target`.

_Required_: No

_Type_: Array of [IdNamespaceIdMappingWorkflowProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdNamespaceName`

The name of the ID namespace.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputSourceConfig`

A list of `InputSource` objects, which have the fields
`InputSourceARN` and `SchemaName`.

_Required_: No

_Type_: Array of [IdNamespaceInputSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idnamespace-idnamespaceinputsource.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role. AWS Entity Resolution assumes
this role to access the resources defined in this `IdNamespace` on your behalf
as part of the workflow run.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `32`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-entityresolution-idnamespace-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of ID namespace. There are two types: `SOURCE` and
`TARGET`.

The `SOURCE` contains configurations for `sourceId` data that will
be processed in an ID mapping workflow.

The `TARGET` contains a configuration of `targetId` which all
`sourceIds` will resolve to.

_Required_: Yes

_Type_: String

_Allowed values_: `SOURCE | TARGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

IdNamespaceIdMappingWorkflowProperties
