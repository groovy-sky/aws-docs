This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Environment

The `AWS::AppConfig::Environment` resource creates an environment, which is a
logical deployment group of AWS AppConfig targets, such as applications in a
`Beta` or `Production` environment. You define one or more
environments for each AWS AppConfig application. You can also define environments for
application subcomponents such as the `Web`, `Mobile` and
`Back-end` components for your application. You can configure Amazon CloudWatch alarms for each environment. The system monitors alarms during a
configuration deployment. If an alarm is triggered, the system rolls back the
configuration.

AWS AppConfig requires that you create resources and deploy a configuration in the
following order:

1. Create an application

2. Create an environment

3. Create a configuration profile

4. Choose a pre-defined deployment strategy or create your own

5. Deploy the configuration

For more information, see [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::Environment",
  "Properties" : {
      "ApplicationId" : String,
      "DeletionProtectionCheck" : String,
      "Description" : String,
      "Monitors" : [ Monitor, ... ],
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::Environment
Properties:
  ApplicationId: String
  DeletionProtectionCheck: String
  Description: String
  Monitors:
    - Monitor
  Name: String
  Tags:
    - Tag

```

## Properties

`ApplicationId`

The application ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeletionProtectionCheck`

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting an environment if your application called either [GetLatestConfiguration](../../../../reference/appconfig/2019-10-09/apireference/api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](../../../../reference/appconfig/2019-10-09/apireference/api-getconfiguration.md) in the
environment during the specified interval.

This parameter supports the following values:

- `BYPASS`: Instructs AWS AppConfig to bypass the deletion
protection check and delete a configuration profile even if deletion protection would
have otherwise prevented it.

- `APPLY`: Instructs the deletion protection check to run, even if
deletion protection is disabled at the account level. `APPLY` also forces
the deletion protection check to run against resources created in the past hour,
which are normally excluded from deletion protection checks.

- `ACCOUNT_DEFAULT`: The default setting, which instructs AWS AppConfig to implement the deletion protection value specified in the
`UpdateAccountSettings` API.

_Required_: No

_Type_: String

_Allowed values_: `ACCOUNT_DEFAULT | APPLY | BYPASS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the environment.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Monitors`

Amazon CloudWatch alarms to monitor during the deployment process.

_Required_: No

_Type_: Array of [Monitor](aws-properties-appconfig-environment-monitor.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the environment.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which
you define.

_Required_: No

_Type_: Array of [Tag](aws-properties-appconfig-environment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the environment ID.

### Fn::GetAtt

`EnvironmentId`

The environment ID.

## Examples

### AWS AppConfig environment example

The following example creates an AWS AppConfig environment named
MyTestEnvironment. An environment is a logical deployment group of AWS AppConfig
targets, such as applications in a Beta or Production environment. You can also define
environments for application subcomponents such as the Web, Mobile, and Back-end
components for your application.

#### JSON

```json

Resources": {
    "BasicEnvironment": {
      "Type": "AWS::AppConfig::Environment",
      "DependsOn": "DependentApplication",
      "Properties": {
        "ApplicationId": null,
        "Name": "MyTestEnvironment",
        "Description": "My test environment",
        "Tags": [
          {
            "Key": "Env",
            "Value": "test"
          }
        ]
      }
    }
  }
}

```

#### YAML

```yaml

Resources:
  BasicEnvironment:
    Type: AWS::AppConfig::Environment
    Properties:
      ApplicationId: !Ref DependentApplication
      Name: "MyTestEnvironment"
      Description: "My test environment"
      Tags:
        - Key: Env
          Value: test

```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Creating an\
environment](../../../systems-manager/latest/userguide/appconfig-creating-environment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Monitor

All content copied from https://docs.aws.amazon.com/.
