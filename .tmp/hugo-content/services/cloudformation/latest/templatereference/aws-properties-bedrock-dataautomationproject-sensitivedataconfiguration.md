This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject SensitiveDataConfiguration

Configuration for detecting and redacting sensitive data in content.
Use this to control whether sensitive data is detected only or both detected and redacted,
specify the scope of detection (standard output, custom output, or both),
and configure specific PII entity types to detect along with how they should be masked when redacted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DetectionMode" : String,
  "DetectionScope" : [ String, ... ],
  "PiiEntitiesConfiguration" : PIIEntitiesConfiguration
}

```

### YAML

```yaml

  DetectionMode: String
  DetectionScope:
    - String
  PiiEntitiesConfiguration:
    PIIEntitiesConfiguration

```

## Properties

`DetectionMode`

Specifies the mode for handling sensitive data detection.
Set to DETECTION to only identify sensitive data without modifying content - this produces one output file per detection scope containing detection information with original unredacted content.
Set to DETECTION\_AND\_REDACTION to both identify and mask sensitive data - this produces two output files per detection scope: one unredacted file with detection information and one redacted file with masking applied to sensitive content.
For example, if detectionScope includes both STANDARD and CUSTOM with DETECTION\_AND\_REDACTION mode, four output files will be generated (two for standard output and two for custom output).

_Required_: No

_Type_: String

_Allowed values_: `DETECTION | DETECTION_AND_REDACTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectionScope`

Defines which BDA output types to apply sensitive data detection to.
Specify STANDARD to detect sensitive data in standard output, CUSTOM to detect in custom output (blueprint-based extraction), or both to apply detection to both output types. If not specified, defaults to both STANDARD and CUSTOM.
The number of output files generated depends on both the detection mode and the scopes selected - each scope specified will produce its own set of output files according to the detection mode configured.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PiiEntitiesConfiguration`

Configuration for detecting and redacting Personally Identifiable Information (PII) entities.

_Required_: No

_Type_: [PIIEntitiesConfiguration](aws-properties-bedrock-dataautomationproject-piientitiesconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PIIEntitiesConfiguration

SpeakerLabelingConfiguration

All content copied from https://docs.aws.amazon.com/.
