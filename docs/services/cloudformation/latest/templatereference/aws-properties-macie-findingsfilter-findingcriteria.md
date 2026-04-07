This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::FindingsFilter FindingCriteria

Specifies, as a map, one or more property-based conditions for a findings filter. A _findings filter_, also referred
to as a _filter rule_, is a set of custom criteria that specifies which findings to include or exclude
from the results of a query for findings. You can also configure a findings filter to suppress (automatically archive) findings that
match the filter's criteria. For more information,
see [Filtering Macie findings](https://docs.aws.amazon.com/macie/latest/user/findings-filter-overview.html) in
the _Amazon Macie User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Criterion" : {Key: Value, ...}
}

```

### YAML

```yaml

  Criterion:
    Key: Value

```

## Properties

`Criterion`

Specifies a condition that defines the property, operator, and one or more values to
use to filter the results.

_Required_: No

_Type_: Object of [CriterionAdditionalProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-macie-findingsfilter-criterionadditionalproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CriterionAdditionalProperties

Tag
