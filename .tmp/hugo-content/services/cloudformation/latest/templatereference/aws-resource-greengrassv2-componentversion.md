This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion

Creates a component. Components are software that run on AWS IoT Greengrass core devices.
After you develop and test a component on your core device, you can use this operation to
upload your component to AWS IoT Greengrass. Then, you can deploy the component to other
core devices.

You can use this operation to do the following:

- **Create components from recipes**

Create a component from a recipe, which is a file that defines the component's
metadata, parameters, dependencies, lifecycle, artifacts, and platform capability. For
more information, see [AWS IoT Greengrass component recipe reference](../../../greengrass/v2/developerguide/component-recipe-reference.md) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

To create a component from a recipe, specify `inlineRecipe` when you call
this operation.

- **Create components from Lambda**
**functions**

Create a component from an AWS Lambda function that runs on AWS IoT Greengrass. This creates a recipe and artifacts from the Lambda
function's deployment package. You can use this operation to migrate Lambda
functions from AWS IoT Greengrass V1 to AWS IoT Greengrass V2.

This function accepts Lambda functions in all supported versions of
Python, Node.js, and Java runtimes. AWS IoT Greengrass doesn't apply any additional
restrictions on deprecated Lambda runtime versions.

To create a component from a Lambda function, specify
`lambdaFunction` when you call this operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GreengrassV2::ComponentVersion",
  "Properties" : {
      "InlineRecipe" : String,
      "LambdaFunction" : LambdaFunctionRecipeSource,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::GreengrassV2::ComponentVersion
Properties:
  InlineRecipe: String
  LambdaFunction:
    LambdaFunctionRecipeSource
  Tags:
    Key: Value

```

## Properties

`InlineRecipe`

The recipe to use to create the component. The recipe defines the component's metadata,
parameters, dependencies, lifecycle, artifacts, and platform compatibility.

You must specify either `InlineRecipe` or `LambdaFunction`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LambdaFunction`

The parameters to create a component from a Lambda function.

You must specify either `InlineRecipe` or `LambdaFunction`.

_Required_: No

_Type_: [LambdaFunctionRecipeSource](aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Application-specific metadata to attach to the component version. You can use tags in
IAM policies to control access to AWS IoT Greengrass resources. You can
also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](../../../greengrass/v2/developerguide/tag-resources.md) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses the
following format, which is different from most `Tags` implementations in CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]{1,128}$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Arn`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the component version.

`ComponentName`

The name of the component.

`ComponentVersion`

The version of the component.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IoT Greengrass Version 2

ComponentDependencyRequirement

All content copied from https://docs.aws.amazon.com/.
