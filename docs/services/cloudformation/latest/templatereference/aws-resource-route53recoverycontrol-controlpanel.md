This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::ControlPanel

Creates a new control panel in Amazon Route 53 Application Recovery Controller. A control panel represents a group of routing controls that can be changed together
in a single transaction. You can use a control panel to centrally view the operational status of applications
across your organization, and trigger multi-app failovers in a single transaction, for example, to fail over from one
AWS Region (cell) to another.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryControl::ControlPanel",
  "Properties" : {
      "ClusterArn" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryControl::ControlPanel
Properties:
  ClusterArn: String
  Name: String
  Tags:
    - Tag

```

## Properties

`ClusterArn`

The Amazon Resource Name (ARN) of the cluster for the control panel.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9:\/_-]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the control panel. You can use any non-white space character in the name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the control panel.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53recoverycontrol-controlpanel-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ControlPanelArn` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ControlPanelArn`

The Amazon Resource Name (ARN) of the control panel.

`DefaultControlPanel`

The boolean flag that is set to true for the default control panel in the cluster.

`RoutingControlCount`

The number of routing controls in the control panel.

`Status`

The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING\_DELETION.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
