This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FraudDetector::Detector

Manages a detector and associated detector versions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FraudDetector::Detector",
  "Properties" : {
      "AssociatedModels" : [ Model, ... ],
      "Description" : String,
      "DetectorId" : String,
      "DetectorVersionStatus" : String,
      "EventType" : EventType,
      "RuleExecutionMode" : String,
      "Rules" : [ Rule, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::FraudDetector::Detector
Properties:
  AssociatedModels:
    - Model
  Description: String
  DetectorId: String
  DetectorVersionStatus: String
  EventType:
    EventType
  RuleExecutionMode: String
  Rules:
    - Rule
  Tags:
    - Tag

```

## Properties

`AssociatedModels`

The models to associate with this detector. You must provide the ARNs of all the models you want to associate.

_Required_: No

_Type_: Array of [Model](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-model.html)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The detector description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorId`

The name of the detector.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DetectorVersionStatus`

The status of the detector version. If a value is not provided for this property, AWS CloudFormation assumes `DRAFT` status.

Valid values: `ACTIVE | DRAFT`

_Required_: No

_Type_: String

_Allowed values_: `DRAFT | ACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventType`

The event type associated with this detector.

_Required_: Yes

_Type_: [EventType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-eventtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleExecutionMode`

The rule execution mode for the rules included in the detector version.

Valid values: `FIRST_MATCHED | ALL_MATCHED` Default value: `FIRST_MATCHED`

You can define and edit the rule mode at the detector version level, when it is in draft status.

If you specify `FIRST_MATCHED`, Amazon Fraud Detector
evaluates rules sequentially, first to last, stopping at the first matched rule. Amazon Fraud dectector then provides the outcomes for that single rule.

If you specifiy `ALL_MATCHED`, Amazon Fraud Detector evaluates all rules and returns the outcomes for all matched rules.

_Required_: No

_Type_: String

_Allowed values_: `FIRST_MATCHED | ALL_MATCHED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

The rules to include in the detector version.

_Required_: Yes

_Type_: Array of [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-rule.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the primary identifier for the resource, which is the ARN.

Example: `{"Ref": "arn:aws:frauddetector:us-west-2:123123123123:outcome/outcome_name"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The detector ARN.

`CreatedTime`

Timestamp of when detector was created.

`DetectorVersionId`

The name of the detector.

`EventType.Arn`

The detector ARN.

`EventType.CreatedTime`

Timestamp of when the detector was created.

`EventType.LastUpdatedTime`

Timestamp of when the detector was last updated.

`LastUpdatedTime`

Timestamp of when detector was last updated.

## See also

- [CreateDetectorVersion](https://docs.aws.amazon.com/frauddetector/latest/api/API_CreateDetectorVersion.html) in the _Amazon Fraud Detector API Reference_.

- [Create a detector version](https://docs.aws.amazon.com/frauddetector/latest/ug/create-a-detector-version.html) in the _Amazon Fraud Detector User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Fraud Detector

EntityType
