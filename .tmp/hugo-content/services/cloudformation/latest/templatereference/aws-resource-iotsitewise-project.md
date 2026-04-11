This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Project

###### Important

The AWS IoT SiteWise Monitor feature will no longer be open to new customers starting November 7, 2025 . If you would like to use the AWS IoT SiteWise Monitor feature,
sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS IoT SiteWise Monitor availability change](../../../iot-sitewise/latest/appguide/iotsitewise-monitor-availability-change.md).

Creates a project in the specified portal.

###### Note

Make sure that the project name and description don't contain confidential
information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::Project",
  "Properties" : {
      "AssetIds" : [ String, ... ],
      "PortalId" : String,
      "ProjectDescription" : String,
      "ProjectName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::Project
Properties:
  AssetIds:
    - String
  PortalId: String
  ProjectDescription: String
  ProjectName: String
  Tags:
    - Tag

```

## Properties

`AssetIds`

A list that contains the IDs of each asset associated with the project.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalId`

The ID of the portal in which to create the project.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectDescription`

A description for the project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectName`

A friendly name for the project.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the project. For more information, see
[Tagging your AWS IoT SiteWise\
resources](../../../iot-sitewise/latest/userguide/tag-resources.md) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotsitewise-project-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ProjectId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ProjectArn`

The [ARN](../../../../general/latest/gr/aws-arns-and-namespaces.md) of the project, which has the following format.

`arn:${Partition}:iotsitewise:${Region}:${Account}:project/${ProjectId}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`ProjectId`

The ID of the project.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
