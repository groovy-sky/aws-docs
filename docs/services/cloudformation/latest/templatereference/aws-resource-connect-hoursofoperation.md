This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation

Specifies hours of operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::HoursOfOperation",
  "Properties" : {
      "ChildHoursOfOperations" : [ HoursOfOperationsIdentifier, ... ],
      "Config" : [ HoursOfOperationConfig, ... ],
      "Description" : String,
      "HoursOfOperationOverrides" : [ HoursOfOperationOverride, ... ],
      "InstanceArn" : String,
      "Name" : String,
      "ParentHoursOfOperations" : [ HoursOfOperationsIdentifier, ... ],
      "Tags" : [ Tag, ... ],
      "TimeZone" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::HoursOfOperation
Properties:
  ChildHoursOfOperations:
    - HoursOfOperationsIdentifier
  Config:
    - HoursOfOperationConfig
  Description: String
  HoursOfOperationOverrides:
    - HoursOfOperationOverride
  InstanceArn: String
  Name: String
  ParentHoursOfOperations:
    - HoursOfOperationsIdentifier
  Tags:
    - Tag
  TimeZone: String

```

## Properties

`ChildHoursOfOperations`

Information about child hours of operations.

_Required_: No

_Type_: Array of [HoursOfOperationsIdentifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationsidentifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Config`

Configuration information for the hours of operation.

_Required_: Yes

_Type_: Array of [HoursOfOperationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationconfig.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the hours of operation.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HoursOfOperationOverrides`

Property description not available.

_Required_: No

_Type_: Array of [HoursOfOperationOverride](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationoverride.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the hours of operation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentHoursOfOperations`

Information about parent hours of operations.

_Required_: No

_Type_: Array of [HoursOfOperationsIdentifier](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-hoursofoperationsidentifier.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, { "Tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-hoursofoperation-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The time zone for the hours of operation.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the hours of operation. For example:

`{ "Ref": "myHoursOfOperation" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`HoursOfOperationArn`

The Amazon Resource Name (ARN) of the hours of operation.

## Examples

### Specify an hours of operation resource

The following example specifies an hours of operation resource for an Amazon Connect instance. In the following example, the hours of operation
claimed operates in Sunday 10:01 to 11:59 AM Pacific Standard Time.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies an hours of operation for an Amazon Connect instance
Resources:
  HoursOfOperation:
    Type: 'AWS::Connect::HoursOfOperation'
    Properties:
      Name: 'ExampleHoursOfOperation'
      Description: 'hours of operation created using cfn'
      InstanceArn: 'arn:aws:connect:region-name:aws-account-id:instance/instance-arn'
      TimeZone: 'Pacific/Midway'
      Config:
        - Day: 'SUNDAY'
          EndTime:
            Hours: 11
            Minutes: 59
          StartTime:
            Hours: 10
            Minutes: 1
      Tags:
        - Key: 'tagKey'
          Value: 'tagValue'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

HoursOfOperationConfig
