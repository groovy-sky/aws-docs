This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Application

The AWS::ElasticBeanstalk::Application resource is an AWS Elastic Beanstalk Beanstalk resource
type that specifies an Elastic Beanstalk application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticBeanstalk::Application",
  "Properties" : {
      "ApplicationName" : String,
      "Description" : String,
      "ResourceLifecycleConfig" : ApplicationResourceLifecycleConfig
    }
}

```

### YAML

```yaml

Type: AWS::ElasticBeanstalk::Application
Properties:
  ApplicationName: String
  Description: String
  ResourceLifecycleConfig:
    ApplicationResourceLifecycleConfig

```

## Properties

`ApplicationName`

A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. For
more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must replace
the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

Your description of the application.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLifecycleConfig`

Specifies an application resource lifecycle configuration to prevent your application from
accumulating too many versions.

_Required_: No

_Type_: [ApplicationResourceLifecycleConfig](aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

#### JSON

```json

{
   "Type" : "AWS::ElasticBeanstalk::Application",
   "Properties" : {
      "ApplicationName" : "SampleAWSElasticBeanstalkApplication",
      "Description" : "AWS Elastic Beanstalk PHP Sample Application"
   }
}
```

#### YAML

```yaml

Type: AWS::ElasticBeanstalk::Application
Properties:
  ApplicationName: "SampleAWSElasticBeanstalkApplication"
  Description: "AWS Elastic Beanstalk PHP Sample Application"
```

## See also

- For a complete Elastic Beanstalk sample template, see [Elastic\
Beanstalk Template Snippets](../userguide/quickref-elasticbeanstalk.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Elastic Beanstalk

ApplicationResourceLifecycleConfig
