This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DashboardPublishOptions

Dashboard publish options:

- `AvailabilityStatus` for `AdHocFilteringOption` \- This
status can be either `ENABLED` or `DISABLED`. When this is
set to `DISABLED`, Amazon Quick Sight disables the left filter pane on the
published dashboard, which can be used for ad hoc (one-time) filtering. This
option is `ENABLED` by default.

- `AvailabilityStatus` for `ExportToCSVOption` \- This
status can be either `ENABLED` or `DISABLED`. The visual
option to export data to .CSV format isn't enabled when this is set to
`DISABLED`. This option is `ENABLED` by default.

- `VisibilityState` for `SheetControlsOption` \- This
visibility state can be either `COLLAPSED` or `EXPANDED`.
This option is `COLLAPSED` by default.

- `AvailabilityStatus` for `QuickSuiteActionsOption` \- This status
can be either `ENABLED` or `DISABLED`. Features related to Actions in Amazon
Quick Suite on dashboards are disabled when this is set to `DISABLED`. This option is
`DISABLED` by default.

- `AvailabilityStatus` for `ExecutiveSummaryOption` \- This status
can be either `ENABLED` or `DISABLED`. The option to build an executive
summary is disabled when this is set to `DISABLED`. This option is `ENABLED`
by default.

- `AvailabilityStatus` for `DataStoriesSharingOption` \- This status
can be either `ENABLED` or `DISABLED`. The option to share a data story is
disabled when this is set to `DISABLED`. This option is `ENABLED`
by default.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdHocFilteringOption" : AdHocFilteringOption,
  "DataPointDrillUpDownOption" : DataPointDrillUpDownOption,
  "DataPointMenuLabelOption" : DataPointMenuLabelOption,
  "DataPointTooltipOption" : DataPointTooltipOption,
  "DataQAEnabledOption" : DataQAEnabledOption,
  "DataStoriesSharingOption" : DataStoriesSharingOption,
  "ExecutiveSummaryOption" : ExecutiveSummaryOption,
  "ExportToCSVOption" : ExportToCSVOption,
  "ExportWithHiddenFieldsOption" : ExportWithHiddenFieldsOption,
  "QuickSuiteActionsOption" : QuickSuiteActionsOption,
  "SheetControlsOption" : SheetControlsOption,
  "SheetLayoutElementMaximizationOption" : SheetLayoutElementMaximizationOption,
  "VisualAxisSortOption" : VisualAxisSortOption,
  "VisualMenuOption" : VisualMenuOption,
  "VisualPublishOptions" : DashboardVisualPublishOptions
}

```

### YAML

```yaml

  AdHocFilteringOption:
    AdHocFilteringOption
  DataPointDrillUpDownOption:
    DataPointDrillUpDownOption
  DataPointMenuLabelOption:
    DataPointMenuLabelOption
  DataPointTooltipOption:
    DataPointTooltipOption
  DataQAEnabledOption:
    DataQAEnabledOption
  DataStoriesSharingOption:
    DataStoriesSharingOption
  ExecutiveSummaryOption:
    ExecutiveSummaryOption
  ExportToCSVOption:
    ExportToCSVOption
  ExportWithHiddenFieldsOption:
    ExportWithHiddenFieldsOption
  QuickSuiteActionsOption:
    QuickSuiteActionsOption
  SheetControlsOption:
    SheetControlsOption
  SheetLayoutElementMaximizationOption:
    SheetLayoutElementMaximizationOption
  VisualAxisSortOption:
    VisualAxisSortOption
  VisualMenuOption:
    VisualMenuOption
  VisualPublishOptions:
    DashboardVisualPublishOptions

```

## Properties

`AdHocFilteringOption`

Ad hoc (one-time) filtering option.

_Required_: No

_Type_: [AdHocFilteringOption](aws-properties-quicksight-dashboard-adhocfilteringoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPointDrillUpDownOption`

The drill-down options of data points in a dashboard.

_Required_: No

_Type_: [DataPointDrillUpDownOption](aws-properties-quicksight-dashboard-datapointdrillupdownoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPointMenuLabelOption`

The data point menu label options of a dashboard.

_Required_: No

_Type_: [DataPointMenuLabelOption](aws-properties-quicksight-dashboard-datapointmenulabeloption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPointTooltipOption`

The data point tool tip options of a dashboard.

_Required_: No

_Type_: [DataPointTooltipOption](aws-properties-quicksight-dashboard-datapointtooltipoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataQAEnabledOption`

Adds Q&A capabilities to an Quick Sight dashboard. If no topic is linked, Dashboard Q&A uses the data values that are rendered on the dashboard. End users can use Dashboard Q&A to ask for different slices of the data that they see on the dashboard. If a topic is linked, Topic Q&A is used.

_Required_: No

_Type_: [DataQAEnabledOption](aws-properties-quicksight-dashboard-dataqaenabledoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataStoriesSharingOption`

Data stories sharing option.

_Required_: No

_Type_: [DataStoriesSharingOption](aws-properties-quicksight-dashboard-datastoriessharingoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutiveSummaryOption`

Executive summary option.

_Required_: No

_Type_: [ExecutiveSummaryOption](aws-properties-quicksight-dashboard-executivesummaryoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportToCSVOption`

Export to .csv option.

_Required_: No

_Type_: [ExportToCSVOption](aws-properties-quicksight-dashboard-exporttocsvoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportWithHiddenFieldsOption`

Determines if hidden fields are exported with a dashboard.

_Required_: No

_Type_: [ExportWithHiddenFieldsOption](aws-properties-quicksight-dashboard-exportwithhiddenfieldsoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuickSuiteActionsOption`

Determines if Actions in Amazon Quick Suite are enabled in a dashboard.

_Required_: No

_Type_: [QuickSuiteActionsOption](aws-properties-quicksight-dashboard-quicksuiteactionsoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetControlsOption`

Sheet controls option.

_Required_: No

_Type_: [SheetControlsOption](aws-properties-quicksight-dashboard-sheetcontrolsoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetLayoutElementMaximizationOption`

The sheet layout maximization options of a dashbaord.

_Required_: No

_Type_: [SheetLayoutElementMaximizationOption](aws-properties-quicksight-dashboard-sheetlayoutelementmaximizationoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualAxisSortOption`

The axis sort options of a dashboard.

_Required_: No

_Type_: [VisualAxisSortOption](aws-properties-quicksight-dashboard-visualaxissortoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualMenuOption`

The menu options of a visual in a dashboard.

_Required_: No

_Type_: [VisualMenuOption](aws-properties-quicksight-dashboard-visualmenuoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualPublishOptions`

The visual publish options of a visual in a dashboard.

_Required_: No

_Type_: [DashboardVisualPublishOptions](aws-properties-quicksight-dashboard-dashboardvisualpublishoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashboardError

DashboardSourceEntity

All content copied from https://docs.aws.amazon.com/.
