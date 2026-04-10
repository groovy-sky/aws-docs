This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template Visual

A visual displayed on a sheet in an analysis, dashboard, or template.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BarChartVisual" : BarChartVisual,
  "BoxPlotVisual" : BoxPlotVisual,
  "ComboChartVisual" : ComboChartVisual,
  "CustomContentVisual" : CustomContentVisual,
  "EmptyVisual" : EmptyVisual,
  "FilledMapVisual" : FilledMapVisual,
  "FunnelChartVisual" : FunnelChartVisual,
  "GaugeChartVisual" : GaugeChartVisual,
  "GeospatialMapVisual" : GeospatialMapVisual,
  "HeatMapVisual" : HeatMapVisual,
  "HistogramVisual" : HistogramVisual,
  "InsightVisual" : InsightVisual,
  "KPIVisual" : KPIVisual,
  "LineChartVisual" : LineChartVisual,
  "PieChartVisual" : PieChartVisual,
  "PivotTableVisual" : PivotTableVisual,
  "PluginVisual" : PluginVisual,
  "RadarChartVisual" : RadarChartVisual,
  "SankeyDiagramVisual" : SankeyDiagramVisual,
  "ScatterPlotVisual" : ScatterPlotVisual,
  "TableVisual" : TableVisual,
  "TreeMapVisual" : TreeMapVisual,
  "WaterfallVisual" : WaterfallVisual,
  "WordCloudVisual" : WordCloudVisual
}

```

### YAML

```yaml

  BarChartVisual:
    BarChartVisual
  BoxPlotVisual:
    BoxPlotVisual
  ComboChartVisual:
    ComboChartVisual
  CustomContentVisual:
    CustomContentVisual
  EmptyVisual:
    EmptyVisual
  FilledMapVisual:
    FilledMapVisual
  FunnelChartVisual:
    FunnelChartVisual
  GaugeChartVisual:
    GaugeChartVisual
  GeospatialMapVisual:
    GeospatialMapVisual
  HeatMapVisual:
    HeatMapVisual
  HistogramVisual:
    HistogramVisual
  InsightVisual:
    InsightVisual
  KPIVisual:
    KPIVisual
  LineChartVisual:
    LineChartVisual
  PieChartVisual:
    PieChartVisual
  PivotTableVisual:
    PivotTableVisual
  PluginVisual:
    PluginVisual
  RadarChartVisual:
    RadarChartVisual
  SankeyDiagramVisual:
    SankeyDiagramVisual
  ScatterPlotVisual:
    ScatterPlotVisual
  TableVisual:
    TableVisual
  TreeMapVisual:
    TreeMapVisual
  WaterfallVisual:
    WaterfallVisual
  WordCloudVisual:
    WordCloudVisual

```

## Properties

`BarChartVisual`

A bar chart.

For more information, see [Using bar charts](../../../quicksight/latest/user/bar-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [BarChartVisual](aws-properties-quicksight-template-barchartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BoxPlotVisual`

A box plot.

For more information, see [Using box plots](../../../quicksight/latest/user/box-plots.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [BoxPlotVisual](aws-properties-quicksight-template-boxplotvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComboChartVisual`

A combo chart.

For more information, see [Using combo charts](../../../quicksight/latest/user/combo-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [ComboChartVisual](aws-properties-quicksight-template-combochartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomContentVisual`

A visual that contains custom content.

For more information, see [Using custom visual content](../../../quicksight/latest/user/custom-visual-content.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [CustomContentVisual](aws-properties-quicksight-template-customcontentvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmptyVisual`

An empty visual.

_Required_: No

_Type_: [EmptyVisual](aws-properties-quicksight-template-emptyvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilledMapVisual`

A filled map.

For more information, see [Creating filled maps](../../../quicksight/latest/user/filled-maps.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [FilledMapVisual](aws-properties-quicksight-template-filledmapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunnelChartVisual`

A funnel chart.

For more information, see [Using funnel charts](../../../quicksight/latest/user/funnel-visual-content.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [FunnelChartVisual](aws-properties-quicksight-template-funnelchartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GaugeChartVisual`

A gauge chart.

For more information, see [Using gauge charts](../../../quicksight/latest/user/gauge-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [GaugeChartVisual](aws-properties-quicksight-template-gaugechartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeospatialMapVisual`

A geospatial map or a points on map visual.

For more information, see [Creating point maps](../../../quicksight/latest/user/point-maps.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [GeospatialMapVisual](aws-properties-quicksight-template-geospatialmapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeatMapVisual`

A heat map.

For more information, see [Using heat maps](../../../quicksight/latest/user/heat-map.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [HeatMapVisual](aws-properties-quicksight-template-heatmapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HistogramVisual`

A histogram.

For more information, see [Using histograms](../../../quicksight/latest/user/histogram-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [HistogramVisual](aws-properties-quicksight-template-histogramvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightVisual`

An insight visual.

For more information, see [Working with insights](../../../quicksight/latest/user/computational-insights.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [InsightVisual](aws-properties-quicksight-template-insightvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KPIVisual`

A key performance indicator (KPI).

For more information, see [Using KPIs](../../../quicksight/latest/user/kpi.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [KPIVisual](aws-properties-quicksight-template-kpivisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineChartVisual`

A line chart.

For more information, see [Using line charts](../../../quicksight/latest/user/line-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [LineChartVisual](aws-properties-quicksight-template-linechartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PieChartVisual`

A pie or donut chart.

For more information, see [Using pie charts](../../../quicksight/latest/user/pie-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [PieChartVisual](aws-properties-quicksight-template-piechartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PivotTableVisual`

A pivot table.

For more information, see [Using pivot tables](../../../quicksight/latest/user/pivot-table.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [PivotTableVisual](aws-properties-quicksight-template-pivottablevisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PluginVisual`

The custom plugin visual type.

_Required_: No

_Type_: [PluginVisual](aws-properties-quicksight-template-pluginvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RadarChartVisual`

A radar chart visual.

For more information, see [Using radar charts](../../../quicksight/latest/user/radar-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [RadarChartVisual](aws-properties-quicksight-template-radarchartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SankeyDiagramVisual`

A sankey diagram.

For more information, see [Using Sankey diagrams](../../../quicksight/latest/user/sankey-diagram.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [SankeyDiagramVisual](aws-properties-quicksight-template-sankeydiagramvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScatterPlotVisual`

A scatter plot.

For more information, see [Using scatter plots](../../../quicksight/latest/user/scatter-plot.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [ScatterPlotVisual](aws-properties-quicksight-template-scatterplotvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableVisual`

A table visual.

For more information, see [Using tables as visuals](../../../quicksight/latest/user/tabular.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [TableVisual](aws-properties-quicksight-template-tablevisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreeMapVisual`

A tree map.

For more information, see [Using tree maps](../../../quicksight/latest/user/tree-map.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [TreeMapVisual](aws-properties-quicksight-template-treemapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaterfallVisual`

A waterfall chart.

For more information, see [Using waterfall charts](../../../quicksight/latest/user/waterfall-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [WaterfallVisual](aws-properties-quicksight-template-waterfallvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordCloudVisual`

A word cloud.

For more information, see [Using word clouds](../../../quicksight/latest/user/word-cloud.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [WordCloudVisual](aws-properties-quicksight-template-wordcloudvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisibleRangeOptions

VisualCustomAction

All content copied from https://docs.aws.amazon.com/.
