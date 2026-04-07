This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ControlTower::LandingZone

Creates a new landing zone. This API call starts an asynchronous operation that creates and configures a landing zone,
based on the parameters specified in the manifest JSON file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ControlTower::LandingZone",
  "Properties" : {
      "Manifest" : ,
      "RemediationTypes" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::ControlTower::LandingZone
Properties:
  Manifest:

  RemediationTypes:
    - String
  Tags:
    - Tag
  Version: String

```

## Properties

`Manifest`

The landing zone manifest JSON text file that specifies the landing zone configurations.

_Required_: Yes

_Type_:

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemediationTypes`

The types of remediation actions configured for the landing zone, such as automatic drift correction or compliance enforcement.

_Required_: No

_Type_: Array of String

_Allowed values_: `INHERITANCE_DRIFT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to be applied to the landing zone.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-controltower-landingzone-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The landing zone's current deployed version.

_Required_: Yes

_Type_: String

_Pattern_: `\d+.\d+`

_Minimum_: `3`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the **LandingZoneIdentifier**. For example:

`"LandingZoneIdentifier": "arn:aws:controltower:us-west-2:123456789012:landingzone/1A2B3C4D5E6F7G8H`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The ARN of the landing zone.

`DriftStatus`

The drift status of the landing zone.

`LandingZoneIdentifier`

The unique identifier of the landing zone.

`LatestAvailableVersion`

The latest available version of the landing zone.

`Status`

The landing zone deployment status. One of `ACTIVE`, `PROCESSING`, `FAILED`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
