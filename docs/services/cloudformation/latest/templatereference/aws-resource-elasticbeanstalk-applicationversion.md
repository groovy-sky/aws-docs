This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::ApplicationVersion

The AWS::ElasticBeanstalk::ApplicationVersion resource is an AWS Elastic Beanstalk
resource type that specifies an application version, an iteration of deployable code, for an
Elastic Beanstalk application.

###### Note

After you create an application version with a specified Amazon S3 bucket and key
location, you can't change that Amazon S3 location. If you change the Amazon S3 location, an
attempt to launch an environment from the application version will fail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticBeanstalk::ApplicationVersion",
  "Properties" : {
      "ApplicationName" : String,
      "Description" : String,
      "SourceBundle" : SourceBundle
    }
}

```

### YAML

```yaml

Type: AWS::ElasticBeanstalk::ApplicationVersion
Properties:
  ApplicationName: String
  Description: String
  SourceBundle:
    SourceBundle

```

## Properties

`ApplicationName`

The name of the Elastic Beanstalk application that is associated with this application
version.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of this application version.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceBundle`

The Amazon S3 bucket and key that identify the location of the source bundle for this
version.

###### Note

The Amazon S3 bucket must be in the same region as the
environment.

_Required_: Yes

_Type_: [SourceBundle](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticbeanstalk-applicationversion-sourcebundle.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

#### JSON

```json

"myAppVersion": {
  "Type" : "AWS::ElasticBeanstalk::ApplicationVersion",
  "Properties" : {
    "ApplicationName" : {"Ref" : "myApp"},
    "Description" : "my sample version",
    "SourceBundle" : {
      "S3Bucket" : { "Fn::Join" :
        ["-", [ "elasticbeanstalk-samples", { "Ref" : "AWS::Region" } ] ] },
      "S3Key" : "php-newsample-app.zip"
    }
  }
}
```

#### YAML

```yaml

myAppVersion:
  Type: AWS::ElasticBeanstalk::ApplicationVersion
  Properties:
    ApplicationName:
      Ref: "myApp"
    Description: "my sample version"
    SourceBundle:
      S3Bucket:
        Fn::Join:
          - "-"
          -
            - "elasticbeanstalk-samples"
            - Ref: "AWS::Region"
      S3Key: "php-newsample-app.zip"
```

## See also

- For a complete Elastic Beanstalk sample template, see [Elastic\
Beanstalk Template Snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-elasticbeanstalk.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MaxCountRule

SourceBundle
