This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachine

Provisions a state machine. A state machine consists of a collection of states that can
do work ( `Task` states), determine to which states to transition next
( `Choice` states), stop an execution with an error ( `Fail`
states), and so on. State machines are specified using a JSON-based, structured
language.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::StepFunctions::StateMachine",
  "Properties" : {
      "Definition" : Json,
      "DefinitionS3Location" : S3Location,
      "DefinitionString" : String,
      "DefinitionSubstitutions" : {Key: Value, ...},
      "EncryptionConfiguration" : EncryptionConfiguration,
      "LoggingConfiguration" : LoggingConfiguration,
      "RoleArn" : String,
      "StateMachineName" : String,
      "StateMachineType" : String,
      "Tags" : [ TagsEntry, ... ],
      "TracingConfiguration" : TracingConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::StepFunctions::StateMachine
Properties:
  Definition: Json
  DefinitionS3Location:
    S3Location
  DefinitionString:
    String
  DefinitionSubstitutions:
    Key: Value
  EncryptionConfiguration:
    EncryptionConfiguration
  LoggingConfiguration:
    LoggingConfiguration
  RoleArn: String
  StateMachineName: String
  StateMachineType: String
  Tags:
    - TagsEntry
  TracingConfiguration:
    TracingConfiguration

```

## Properties

`Definition`

The Amazon States Language definition of the state machine. The state machine definition must be in JSON or YAML, and the format of the object must
match the format of your CloudFormationtemplate file. See [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionS3Location`

The name of the S3 bucket where the state machine definition is stored. The state machine definition must be a JSON or YAML file.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachine-s3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionString`

The Amazon States Language definition of the state machine. The state machine definition must be in JSON. See [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1048576`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionSubstitutions`

A map (string to string) that specifies the mappings for placeholder variables in the
state machine definition. This enables the customer to inject values obtained at runtime,
for example from intrinsic functions, in the state machine definition. Variables can be
template parameter names, resource logical IDs, resource attributes, or a variable in a
key-value map.

Substitutions must follow the syntax: `${key_name}` or `${variable_1,variable_2,...}`.

_Required_: No

_Type_: Object

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

Encryption configuration for the state machine.

_Required_: No

_Type_: [EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachine-encryptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingConfiguration`

Defines what execution history events are logged and where they are logged.

###### Note

By default, the `level` is set to `OFF`. For more information
see [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.

_Required_: No

_Type_: [LoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachine-loggingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role to use for this state machine.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateMachineName`

The name of the state machine.

A name must _not_ contain:

- white space

- brackets `< > { } [ ]`

- wildcard characters `? *`

- special characters ``" # % \ ^ | ~ ` $ & , ; : /``

- control characters ( `U+0000-001F`, `U+007F-009F`)

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StateMachineType`

Determines whether a `STANDARD` or `EXPRESS` state machine is
created. The default is `STANDARD`. You cannot update the `type` of a
state machine once it has been created. For more information on `STANDARD` and
`EXPRESS` workflows, see [Standard Versus Express\
Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-standard-vs-express.html) in the AWS Step Functions Developer Guide.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | EXPRESS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The list of tags to add to a resource.

Tags may only contain Unicode letters, digits, white space, or these symbols: `_ . : / = + - @`.

_Required_: No

_Type_: Array of [TagsEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachine-tagsentry.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TracingConfiguration`

Selects whether or not the state machine's AWS X-Ray tracing is enabled.

_Required_: No

_Type_: [TracingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-stepfunctions-statemachine-tracingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the Ref intrinsic function, Ref
returns the ARN of the created state machine. For example:

`{ "Ref": "MyStateMachine" }`

Returns a value similar to the following:

`arn:aws:states:us-east-1:111122223333:stateMachine:HelloWorld-StateMachine`

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

`Arn`

Returns the ARN of the resource.

`Name`

Returns the name of the state machine. For example:

`{ "Fn::GetAtt": ["MyStateMachine", "Name"] }`

Returns the name of your state machine:

`HelloWorld-StateMachine`

If you did not specify the name it will be similar to the following:

`MyStateMachine-1234abcdefgh`

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`StateMachineRevisionId`

Identifier for a state machine revision, which is an immutable, read-only snapshot of a state machine’s definition and configuration.

## Examples

The following examples create a Step Functions state machine.

- [Using a Single-Line Property](#aws-resource-stepfunctions-statemachine--examples--Using_a_Single-Line_Property)

- [Using the Fn::Join Intrinsic Function](#aws-resource-stepfunctions-statemachine--examples--Using_the_Fn::Join_Intrinsic_Function)

- [Including Tags](#aws-resource-stepfunctions-statemachine--examples--Including_Tags)

- [Using DefinitionSubstitutions](#aws-resource-stepfunctions-statemachine--examples--Using_DefinitionSubstitutions)

- [hello\_world.json](#aws-resource-stepfunctions-statemachine--examples--hello_world.json)

### Using a Single-Line Property

#### JSON

```json

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Description":"An example template for a Step Functions state machine.",
   "Resources":{
      "MyStateMachine":{
         "Type":"AWS::StepFunctions::StateMachine",
         "Properties":{
            "StateMachineName":"HelloWorld-StateMachine",
            "StateMachineType":"STANDARD",
            "DefinitionString":"{\"StartAt\": \"HelloWorld\",
            \"States\": {\"HelloWorld\": {\"Type\": \"Task\", \"Resource\":
            \"arn:aws:lambda:us-east-1:111122223333;:function:HelloFunction\", \"End\": true}}}",
            "RoleArn":"arn:aws:iam::111122223333:role/service-role/StatesExecutionRole-us-east-1;"
         }
      }
   }
}
```

### Using the Fn::Join Intrinsic Function

#### JSON

```json

{
    "AWSTemplateFormatVersion" : "2010-09-09",
    "Description" : "An example template for a Step Functions state machine.",
    "Resources": {
       "MyStateMachine": {
          "Type": "AWS::StepFunctions::StateMachine",
             "Properties": {
                "StateMachineName" : "HelloWorld-StateMachine",
                "StateMachineType":"STANDARD",
                "DefinitionString" : {
                   "Fn::Join": [
                      "\n",
                      [
                         "{",
                         "    \"StartAt\": \"HelloWorld\",",
                         "    \"States\" : {",
                         "        \"HelloWorld\" : {",
                         "            \"Type\" : \"Task\", ",
                         "            \"Resource\" : \"arn:aws:lambda:us-east-1:111122223333:function:HelloFunction\",",
                         "            \"End\" : true",
                         "        }",
                         "    }",
                         "}"
                      ]
                   ]
                },
   	      "RoleArn" : "arn:aws:iam::111122223333:role/service-role/StatesExecutionRole-us-east-1",
            "Tags": [
                    {
                        "Key": "keyname1",
                        "Value": "value1"
                    },
                    {
                        "Key": "keyname2",
                        "Value": "value2"
                    }
                ]
            }
        }
    }
}
```

### Including Tags

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: An example template for a Step Functions state machine.
Resources:
  MyStateMachine:
    Type: AWS::StepFunctions::StateMachine
    Properties:
      StateMachineName: HelloWorld-StateMachine
      DefinitionString: |-
        {
          "StartAt": "HelloWorld",
          "States": {
            "HelloWorld": {
              "Type": "Task",
              "Resource": "arn:aws:lambda:us-east-1:111122223333:function:HelloFunction",
              "End": true
            }
          }
        }
      RoleArn: arn:aws:iam::111122223333:role/service-role/StatesExecutionRole-us-east-1
      Tags:
        -
          Key: "keyname1"
          Value: "value1"
        -
          Key: "keyname2"
          Value: "value2"
```

### Using DefinitionSubstitutions

In this example template, `HelloFunction:` is defined for the
`DefinitionSubstitutions` property. In the
`hello_world.json` definition file, that
follows `${HelloFunction}` will be replaced by
`arn:aws:lambda:us-east-1:111122223333:function:HelloFunction`.

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: An example template for a Step Functions state machine.
Resources:
  MyStateMachine:
    Type: AWS::StepFunctions::StateMachine
    Properties:
      StateMachineName: HelloWorld-StateMachine
      DefinitionS3Location:
        Bucket: example_bucket
        Key: hello_world.json
      DefinitionSubstitutions:
        HelloFunction: arn:aws:lambda:us-east-1:111122223333:function:HelloFunction
      RoleArn: arn:aws:iam::111122223333:role/service-role/StatesExecutionRole-us-east-1

```

### hello\_world.json

A definition file where `${HelloFunction}` will be replaced by
`arn:aws:lambda:us-east-1:111122223333:function:HelloFunction`. from
the preceding example template.

#### JSON

```json

{
  "StartAt": "HelloWorld",
  "States": {
    "HelloWorld": {
      "Type": "Task",
      "Resource": "${HelloFunction}",
      "End": true
    }
  }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagsEntry

CloudWatchLogsLogGroup
