This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::SecurityControl

The `AWS::SecurityHub::SecurityControl` resource specifies custom parameter values for
an AWS Security Hub CSPM control. For a list of controls that support custom
parameters, see [Security Hub CSPM controls reference](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-controls-reference.html). You can also use this resource to specify the use of default parameter values for
a control. For more information about custom parameters,
see [Custom control\
parameters](https://docs.aws.amazon.com/securityhub/latest/userguide/custom-control-parameters.html) in the _AWS Security Hub CSPM User Guide_.

Tags aren't supported for this resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::SecurityControl",
  "Properties" : {
      "LastUpdateReason" : String,
      "Parameters" : {Key: Value, ...},
      "SecurityControlArn" : String,
      "SecurityControlId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::SecurityControl
Properties:
  LastUpdateReason: String
  Parameters:
    Key: Value
  SecurityControlArn: String
  SecurityControlId: String

```

## Properties

`LastUpdateReason`

The most recent reason for updating the customizable properties of a security control. This differs from the
`UpdateReason` field of the [`BatchUpdateStandardsControlAssociations`](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_BatchUpdateStandardsControlAssociations.html) API, which tracks the
reason for updating the enablement status of a control. This field accepts alphanumeric
characters in addition to white spaces, dashes, and underscores.

_Required_: No

_Type_: String

_Pattern_: `^([^-]|[-_ a-zA-Z0-9])+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

An object that identifies the name of a control parameter, its current value, and whether it has been customized.

_Required_: Yes

_Type_: Object of [ParameterConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-securitycontrol-parameterconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityControlArn`

The Amazon Resource Name (ARN) for a security control across standards, such as
`arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This
parameter doesn't mention a specific standard.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityControlId`

The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a
number, such as APIGateway.3.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the identifier of the security control. For example, `Config.1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Configuring control parameters

This example configures a parameter for the control ACM.1.

#### JSON

```json

{
  "Description": "Example template to configure control parameters",
  "Resources": {
    "ExampleSecurityControl": {
      "Type": "AWS::SecurityHub::SecurityControl",
      "Properties": {
        "SecurityControlId": "ACM.1",
        "Parameters": {
          "daysToExpiration": {
            "ValueType": "CUSTOM",
            "Value": {
              "Integer": 15
            }
          }
        },
        "LastUpdateReason": "Internal compliance requirement"
      }
    }
  }
}
```

#### YAML

```yaml

Description: Example template to configure control parameters
Resources:
  ExampleSecurityControl:
    Type: 'AWS::SecurityHub::SecurityControl'
    Properties:
      SecurityControlId: 'ACM.1'
      Parameters:
        daysToExpiration:
          ValueType: 'CUSTOM'
          Value:
            Integer: 15
      LastUpdateReason: 'Internal compliance requirement'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityHub::ProductSubscription

ParameterConfiguration
