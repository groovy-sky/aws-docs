This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ApprovedOrigin

The approved origin for the instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::ApprovedOrigin",
  "Properties" : {
      "InstanceId" : String,
      "Origin" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::ApprovedOrigin
Properties:
  InstanceId: String
  Origin: String

```

## Properties

`InstanceId`

The Amazon Resource Name (ARN) of the instance.

_Minimum_: `1`

_Maximum_: `100`

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Origin`

Domain name to be added to the allow-list of the instance.

_Maximum_: `267`

_Required_: Yes

_Type_: String

_Maximum_: `267`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

## Examples

### Specify an Approved Origin for an Amazon Connect instance

The following example specifies an Approved Origin for an Amazon Connect
instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies an Approved Origin for an Amazon Connect instance
Resources:
  ApprovedOrigin:
    Type: AWS::Connect::ApprovedOrigin
    Properties:
      InstanceId: arn:aws:connect:region-name:aws-account-id:instance/instance-arn
      Origin: "https://aws.amazon.com"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Connect::ContactFlow

All content copied from https://docs.aws.amazon.com/.
