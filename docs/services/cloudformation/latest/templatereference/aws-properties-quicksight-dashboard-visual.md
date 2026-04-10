This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard Visual

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
  "LayerMapVisual" : LayerMapVisual,
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
  LayerMapVisual:
    LayerMapVisual
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

_Type_: [BarChartVisual](aws-properties-quicksight-dashboard-barchartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BoxPlotVisual`

A box plot.

For more information, see [Using box plots](../../../quicksight/latest/user/box-plots.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [BoxPlotVisual](aws-properties-quicksight-dashboard-boxplotvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComboChartVisual`

A combo chart.

For more information, see [Using combo charts](../../../quicksight/latest/user/combo-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [ComboChartVisual](aws-properties-quicksight-dashboard-combochartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomContentVisual`

A visual that contains custom content.

For more information, see [Using custom visual content](../../../quicksight/latest/user/custom-visual-content.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [CustomContentVisual](aws-properties-quicksight-dashboard-customcontentvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmptyVisual`

An empty visual.

_Required_: No

_Type_: [EmptyVisual](aws-properties-quicksight-dashboard-emptyvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilledMapVisual`

A filled map.

For more information, see [Creating filled maps](../../../quicksight/latest/user/filled-maps.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [FilledMapVisual](aws-properties-quicksight-dashboard-filledmapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunnelChartVisual`

A funnel chart.

For more information, see [Using funnel charts](../../../quicksight/latest/user/funnel-visual-content.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [FunnelChartVisual](aws-properties-quicksight-dashboard-funnelchartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GaugeChartVisual`

A gauge chart.

For more information, see [Using gauge charts](../../../quicksight/latest/user/gauge-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [GaugeChartVisual](aws-properties-quicksight-dashboard-gaugechartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeospatialMapVisual`

A geospatial map or a points on map visual.

For more information, see [Creating point maps](../../../quicksight/latest/user/point-maps.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [GeospatialMapVisual](aws-properties-quicksight-dashboard-geospatialmapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeatMapVisual`

A heat map.

For more information, see [Using heat maps](../../../quicksight/latest/user/heat-map.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [HeatMapVisual](aws-properties-quicksight-dashboard-heatmapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HistogramVisual`

A histogram.

For more information, see [Using histograms](../../../quicksight/latest/user/histogram-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [HistogramVisual](aws-properties-quicksight-dashboard-histogramvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightVisual`

An insight visual.

For more information, see [Working with insights](../../../quicksight/latest/user/computational-insights.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [InsightVisual](aws-properties-quicksight-dashboard-insightvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KPIVisual`

A key performance indicator (KPI).

For more information, see [Using KPIs](../../../quicksight/latest/user/kpi.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [KPIVisual](aws-properties-quicksight-dashboard-kpivisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LayerMapVisual`

The properties for a layer map visual

_Required_: No

_Type_: [LayerMapVisual](aws-properties-quicksight-dashboard-layermapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineChartVisual`

A line chart.

For more information, see [Using line charts](../../../quicksight/latest/user/line-charts.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [LineChartVisual](aws-properties-quicksight-dashboard-linechartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PieChartVisual`

A pie or donut chart.

For more information, see [Using pie charts](../../../quicksight/latest/user/pie-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [PieChartVisual](aws-properties-quicksight-dashboard-piechartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PivotTableVisual`

A pivot table.

For more information, see [Using pivot tables](../../../quicksight/latest/user/pivot-table.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [PivotTableVisual](aws-properties-quicksight-dashboard-pivottablevisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PluginVisual`

The custom plugin visual type.

_Required_: No

_Type_: [PluginVisual](aws-properties-quicksight-dashboard-pluginvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RadarChartVisual`

A radar chart visual.

For more information, see [Using radar charts](../../../quicksight/latest/user/radar-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [RadarChartVisual](aws-properties-quicksight-dashboard-radarchartvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SankeyDiagramVisual`

A sankey diagram.

For more information, see [Using Sankey diagrams](../../../quicksight/latest/user/sankey-diagram.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [SankeyDiagramVisual](aws-properties-quicksight-dashboard-sankeydiagramvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScatterPlotVisual`

A scatter plot.

For more information, see [Using scatter plots](../../../quicksight/latest/user/scatter-plot.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [ScatterPlotVisual](aws-properties-quicksight-dashboard-scatterplotvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableVisual`

A table visual.

For more information, see [Using tables as visuals](../../../quicksight/latest/user/tabular.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [TableVisual](aws-properties-quicksight-dashboard-tablevisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreeMapVisual`

A tree map.

For more information, see [Using tree maps](../../../quicksight/latest/user/tree-map.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [TreeMapVisual](aws-properties-quicksight-dashboard-treemapvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaterfallVisual`

A waterfall chart.

For more information, see [Using waterfall charts](../../../quicksight/latest/user/waterfall-chart.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [WaterfallVisual](aws-properties-quicksight-dashboard-waterfallvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordCloudVisual`

A word cloud.

For more information, see [Using word clouds](../../../quicksight/latest/user/word-cloud.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: [WordCloudVisual](aws-properties-quicksight-dashboard-wordcloudvisual.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VisibleRangeOptions

VisualAxisSortOption

All content copied from https://docs.aws.amazon.com/.
