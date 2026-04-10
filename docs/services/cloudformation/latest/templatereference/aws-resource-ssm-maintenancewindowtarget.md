This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTarget

The `AWS::SSM::MaintenanceWindowTarget` resource registers a target with a
maintenance window for AWS Systems Manager. For more information, see [RegisterTargetWithMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-registertargetwithmaintenancewindow.md) in the _AWS Systems Manager API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::MaintenanceWindowTarget",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "OwnerInformation" : String,
      "ResourceType" : String,
      "Targets" : [ Targets, ... ],
      "WindowId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSM::MaintenanceWindowTarget
Properties:
  Description: String
  Name: String
  OwnerInformation: String
  ResourceType: String
  Targets:
    - Targets
  WindowId: String

```

## Properties

`Description`

A description for the target.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the maintenance window target.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,128}$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerInformation`

A user-provided value that will be included in any Amazon CloudWatch Events events that are
raised while running tasks for these targets in this maintenance window.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of target that is being registered with the maintenance window.

_Required_: Yes

_Type_: String

_Allowed values_: `INSTANCE | RESOURCE_GROUP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The targets to register with the maintenance window. In other words, the instances to
run commands on when the maintenance window runs.

You must specify targets by using the `WindowTargetIds` parameter.

_Required_: Yes

_Type_: [Array](aws-properties-ssm-maintenancewindowtarget-targets.md) of [Targets](aws-properties-ssm-maintenancewindowtarget-targets.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WindowId`

The ID of the maintenance window to register the target with.

_Required_: Yes

_Type_: String

_Pattern_: `^mw-[0-9a-f]{17}$`

_Minimum_: `20`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the maintenance window target ID, such as
`12a345b6-bbb7-4bb6-90b0-8c9577a2d2b9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`WindowTargetId`

The ID of the target.

## Examples

### Create a maintenance window that targets instances by using tags

The following example creates a Systems Manager maintenance window target
that targets managed instances with the tag key `ENV` and the tag
value `DEV`.

#### JSON

```json

{
    "Resources": {
        "MaintenanceWindowTarget": {
            "Type": "AWS::SSM::MaintenanceWindowTarget",
            "Properties": {
                "WindowId": "MaintenanceWindow",
                "ResourceType": "INSTANCE",
                "Targets": [
                    {
                        "Key": "tag:ENV",
                        "Values": [
                            "DEV"
                        ]
                    }
                ],
                "OwnerInformation": "SSM Step Function Demo",
                "Name": "SSMStepFunctionDemo",
                "Description": "A target for demonstrating maintenance windows and step functions"
            },
            "DependsOn": "MaintenanceWindow"
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  MaintenanceWindowTarget:
    Type: AWS::SSM::MaintenanceWindowTarget
    Properties:
      WindowId: MaintenanceWindow
      ResourceType: INSTANCE
      Targets:
      - Key: tag:ENV
        Values:
        - DEV
      OwnerInformation: SSM Step Function Demo
      Name: SSMStepFunctionDemo
      Description: A target for demonstrating maintenance windows and step functions
    DependsOn: MaintenanceWindow
```

## See also

- [AWS::SSM::MaintenanceWindow](../userguide/aws-resource-ssm-maintenancewindow.md)

- [AWS::SSM::MaintenanceWindowTask](../userguide/aws-resource-ssm-maintenancewindowtask.md)

- [RegisterTaskWithMaintenanceWindow](../../../../reference/systems-manager/latest/apireference/api-registertaskwithmaintenancewindow.md) in the _AWS Systems Manager API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Targets

All content copied from https://docs.aws.amazon.com/.
