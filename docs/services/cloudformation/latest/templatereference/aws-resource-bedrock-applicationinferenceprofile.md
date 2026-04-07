This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::ApplicationInferenceProfile

Specifies an inference profile as a resource in a top-level template. Use the
`ModelSource` field to specify the inference profile to copy into the
resource. For more information about using inference profiles in Amazon Bedrock, see
[Improve resilience with\
cross-region inference](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::ApplicationInferenceProfile",
  "Properties" : {
      "Description" : String,
      "InferenceProfileName" : String,
      "ModelSource" : InferenceProfileModelSource,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::ApplicationInferenceProfile
Properties:
  Description: String
  InferenceProfileName: String
  ModelSource:
    InferenceProfileModelSource
  Tags:
    - Tag

```

## Properties

`Description`

The description of the inference profile.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z:.][ _-]?)+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceProfileName`

The name of the inference profile.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][ _-]?)+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelSource`

Contains configurations for the inference profile to copy as the resource.

_Required_: No

_Type_: [InferenceProfileModelSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-applicationinferenceprofile-inferenceprofilemodelsource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags associated with the inference profile.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-applicationinferenceprofile-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the inference profile ID.

For example, `{ "Ref": "myInferenceProfile" }` could return the value
`"IP12345678"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time at which the inference profile was created.

`InferenceProfileArn`

The Amazon Resource Name (ARN) of the inference profile.

`InferenceProfileId`

The unique identifier of the inference profile.

`InferenceProfileIdentifier`

The ID or Amazon Resource Name (ARN) of the inference profile.

`Models`

A list of information about each model in the inference profile.

`Status`

The status of the inference profile. `ACTIVE` means that the inference profile is ready to be used.

`Type`

The type of the inference profile. The following types are possible:

- `SYSTEM_DEFINED` – The inference profile is defined by Amazon Bedrock. You can route inference requests across regions with these inference profiles.

- `APPLICATION` – The inference profile was created by a user. This type of inference profile can track metrics and costs when invoking the model in it. The inference profile may route requests to one or multiple regions.

`UpdatedAt`

The time at which the inference profile was last updated.

## Examples

### Create an inference profile

The following example creates an inference profile for the Anthropic Claude
3.5 Sonnet model in the US West (Oregon) region:

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Example Application Inference Profile
Resources:
  ExampleApplicationInferenceProfile:
    Type: AWS::Bedrock::ApplicationInferenceProfile
    Properties:
      ModelSource:
        CopyFrom: "arn:aws:bedrock:us-west-2::foundation-model/anthropic.claude-3-5-sonnet-20240620-v1:0"
      InferenceProfileName: Example Application Inference Profile Name
      Description: Description of ExampleApplicationInferenceProfile
```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Example Application Inference Profile",
    "Resources": {
        "ExampleApplicationInferenceProfile": {
            "Type": "AWS::Bedrock::ApplicationInferenceProfile",
            "Properties": {
                "ModelSource": {
                    "CopyFrom": "arn:aws:bedrock:us-west-2::foundation-model/anthropic.claude-3-5-sonnet-20240620-v1:0"
                },
                "InferenceProfileName": "Example Application Inference Profile Name",
                "Description": "Description of ExampleApplicationInferenceProfile"
            }
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AgentAliasRoutingConfigurationListItem

InferenceProfileModel
