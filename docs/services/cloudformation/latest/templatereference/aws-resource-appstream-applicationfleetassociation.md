This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::ApplicationFleetAssociation

This resource associates the specified application with the specified fleet. This is only supported for Elastic fleets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::ApplicationFleetAssociation",
  "Properties" : {
      "ApplicationArn" : String,
      "FleetName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::ApplicationFleetAssociation
Properties:
  ApplicationArn: String
  FleetName: String

```

## Properties

`ApplicationArn`

The ARN of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FleetName`

The name of the fleet.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns a combination of the `FleetName` and
`ApplicationArn`, such as
`aabcdefgFleet|arn:aws:appstream:us-west-2:123456789123:application/abcdefg`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppStream::ApplicationEntitlementAssociation

AWS::AppStream::DirectoryConfig
