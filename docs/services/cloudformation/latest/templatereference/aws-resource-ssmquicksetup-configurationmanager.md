This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMQuickSetup::ConfigurationManager

Creates a Quick Setup configuration manager resource. This object is a collection of
desired state configurations for multiple configuration definitions and summaries
describing the deployments of those definitions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSMQuickSetup::ConfigurationManager",
  "Properties" : {
      "ConfigurationDefinitions" : [ ConfigurationDefinition, ... ],
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SSMQuickSetup::ConfigurationManager
Properties:
  ConfigurationDefinitions:
    - ConfigurationDefinition
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`ConfigurationDefinitions`

The definition of the Quick Setup configuration that the configuration manager
deploys.

_Required_: Yes

_Type_: Array of [ConfigurationDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssmquicksetup-configurationmanager-configurationdefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the configuration.

_Required_: No

_Type_: String

_Pattern_: `^.{0,512}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the configuration

_Required_: No

_Type_: String

_Pattern_: `^[ A-Za-z0-9_-]{1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs of metadata to assign to the configuration manager.

_Required_: No

_Type_: Object of String

_Pattern_: `^[A-Za-z0-9 +=@_\/:.-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the configuration manager.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The datetime stamp when the configuration manager was created.

`LastModifiedAt`

The datetime stamp when the configuration manager was last updated.

`ManagerArn`

The ARN of the Quick Setup configuration.

`StatusSummaries`

Summaries of the state of the configuration manager. These summaries include an
aggregate of the statuses from the configuration definition associated with the
configuration manager. This includes deployment statuses, association statuses,
drift statuses, health checks, and more.

## Examples

### AWS Config Recording - local

Set up AWS Config Recording for the local account.

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
   SSMQuickSetupTestConfigurationManager:
      Type: "AWS::SSMQuickSetup::ConfigurationManager"
      Properties:
      Name: "Example"
      Description: "Example template"
      ConfigurationDefinitions:
         - Type: "AWSQuickSetupType-PatchPolicy"
            Parameters:
            TargetAccounts:
               Ref: AWS::AccountId
            TargetRegions:
               Ref: AWS::Region
            LocalDeploymentAdministrationRoleArn: !Sub "arn:aws:iam::${AWS::AccountId}:role/AWS-QuickSetup-StackSet-Example-AdministrationRole"
            LocalDeploymentExecutionRoleName: "AWS-QuickSetup-StackSet-Example-ExecutionRole"
      Tags:
         exampleTag: "exampleTagValue"
Outputs:
  Sample:
    Description: This is an example
    Value: !GetAtt SSMQuickSetupTestConfigurationManager.ManagerArn
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Systems Manager Quick Setup

ConfigurationDefinition
