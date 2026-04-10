This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ControlTower::EnabledControl

The resource represents an enabled control. It specifies an asynchronous operation
that creates AWS resources on the specified organizational unit and the
accounts it contains. The resources created will vary according to the control that you
specify.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ControlTower::EnabledControl",
  "Properties" : {
      "ControlIdentifier" : String,
      "Parameters" : [ EnabledControlParameter, ... ],
      "Tags" : [ Tag, ... ],
      "TargetIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::ControlTower::EnabledControl
Properties:
  ControlIdentifier: String
  Parameters:
    - EnabledControlParameter
  Tags:
    - Tag
  TargetIdentifier: String

```

## Properties

`ControlIdentifier`

The ARN of the control. Only **Strongly recommended** and
**Elective** controls are permitted, with the exception of the
**Region deny** control. For information on how to find the `controlIdentifier`, see [the overview page](../../../../reference/controltower/latest/apireference/welcome.md).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[0-9a-zA-Z_\-:\/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

Array of `EnabledControlParameter` objects.

_Required_: No

_Type_: Array of [EnabledControlParameter](aws-properties-controltower-enabledcontrol-enabledcontrolparameter.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-controltower-enabledcontrol-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIdentifier`

The ARN of the organizational unit. For information on how to find the `targetIdentifier`, see [the overview page](../../../../reference/controltower/latest/apireference/welcome.md).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[0-9a-zA-Z_\-:\/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the physical ID of the resource. The format is of the
form `target | control`. For example:

`arn:aws:organizations::123456789012:ou/o-myorg/ou-my-ouid |
                arn:aws:controltower:us-west-2::control/AWS-GR_AUTOSCALING_LAUNCH_CONFIG_PUBLIC_IP_DISABLED`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Example

The following example creates an enabled control named
_MyExampleControl_ for an OU named
_ou-zzxx-zzx0zzz2_.

#### JSON

```json

{
    "Resources": {
        "MyExampleControl": {
            "Properties": {
                "ControlIdentifier": "arn:aws:controltower:us-east-2::control/EXAMPLE_NAME",
                 "TargetIdentifier": "arn:aws:organizations::01234567890:ou/o-EXAMPLE/ou-zzxx-zzx0zzz2"
            },
            "Type": "AWS::ControlTower::EnabledControl"
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyExampleControl:
    Properties:
      ControlIdentifier: 'arn:aws:controltower:us-east-2::control/EXAMPLE_NAME'
      TargetIdentifier: 'arn:aws:organizations::01234567890:ou/o-EXAMPLE/ou-zzxx-zzx0zzz2'
    Type: 'AWS::ControlTower::EnabledControl'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EnabledControlParameter

All content copied from https://docs.aws.amazon.com/.
