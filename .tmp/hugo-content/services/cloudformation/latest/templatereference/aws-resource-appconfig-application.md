This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Application

The `AWS::AppConfig::Application` resource creates an application. In AWS AppConfig, an application is simply an organizational construct like a folder. This
organizational construct has a relationship with some unit of executable code. For example,
you could create an application called MyMobileApp to organize and manage configuration data
for a mobile application installed by your users.

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
  "Type" : "AWS::AppConfig::Application",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::Application
Properties:
  Description: String
  Name: String
  Tags:
    - Tags

```

## Properties

`Description`

A description of the application.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the application.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata to assign to the application. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which
you define.

_Required_: No

_Type_: [Array](aws-properties-appconfig-application-tags.md) of [Tags](aws-properties-appconfig-application-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationId`

The application ID.

## Examples

### AWS AppConfig application example

The following example creates a simple AWS AppConfig application named
MyTestApplication. An application in AWS AppConfig is a logical unit of code that
provides capabilities for your customers. For example, an application can be a
microservice that runs on Amazon EC2 instances, a mobile application installed
by your users, a serverless application using Amazon API Gateway and AWS Lambda, or any system you run on behalf of others.

#### JSON

```json

BasicApplication": {
    "Type": "AWS::AppConfig::Application",
    "Properties": {
      "Name": "MyTestApplication",
      "Description": "A sample test application.",
      "Tags": [
        {
          "Key": "Env",
          "Value": "test"
        }
      ]
    }
  }
}

```

#### YAML

```yaml

BasicApplication:
    Type: AWS::AppConfig::Application
    Properties:
      Name: "MyTestApplication"
      Description: "A sample test application."
      Tags:
        - Key: Env
          Value: test

```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Creating an\
Application](../../../systems-manager/latest/userguide/appconfig-creating-application.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS AppConfig

Tags

All content copied from https://docs.aws.amazon.com/.
