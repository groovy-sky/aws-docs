# Using widgets on CloudWatch dashboards

You can use the following types of widgets on CloudWatch dashboards:

- **Graph widget** – You can add graphs
that contain one or more metrics
to your CloudWatch dashboard.
The types
of graphs
that you can add
to your dashboard
include
**_Line_**,
**_Stacked area_**,
**_Number_**,
**_Gauge_**,
**_Bar_**, and
**_Pie_**.
You can remove graphs
from your dashboard
when you don't need them anymore.

- **Manual metrics widget** – If a metric hasn't published data in the past 14 days, you can't find it when
searching for metrics to add to a graph on a CloudWatch dashboard. You can add any metric manually to an existing graph.

- **Metrics explorer widget** – Metrics explorer widgets include graphs of multiple resources that have
the same tag, or share the same resource property such as an instance type. These
widgets stay up to date, as resources that match are created or deleted. Adding
metrics explorer widgets to your dashboard helps you to troubleshoot your
environment more efficiently.

For example, you can monitor your fleet of EC2 instances by assigning tags that
represent their environments, such as production or test. You can then use these
tags to filter and aggregate the operational metrics, such as
`CPUUtilization`, to understand the health and performance of the EC2
instances that are associated with each tag.

- **Line widget** – With the line widget, you can compare metrics over periods of time. You also can
use the widget's
mini-map
zoom feature to inspect sections of line graphs without changing between zoomed-in
and zoomed-out views. The procedures in this section describe how to add and remove
a line widget
on
a CloudWatch dashboard. For information about using the widget's
mini-map
zoom feature with line graphs, see [Zooming in on a line or stacked area graph](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/zoom-graph.html).

- **Number widget** – With the number widget, you can look at the latest metric values and trends
as
soon as they appear.
Because
the
number widget
includes
the sparkline feature,
you
can visualize the top and bottom halves of metric trends in a single graph. The
procedures in this section describe how to add and remove a number widget from a
CloudWatch dashboard.

- **Gauge widget** – With the gauge widget, you can visualize metric values that go between ranges.
For example, you can use the gauge widget to graph percentages and CPU utilization,
so that
you can observe and diagnose any performance issues that occur.
The procedures in this section describe how to add and remove a gauge widget from a
CloudWatch dashboard.

- **Text widget** – A text widget contains a block of text in [Markdown](https://docs.aws.amazon.com/general/latest/gr/aws-markdown.html) format.
You can add, edit, or remove text widgets from your CloudWatch dashboard.

- **Alarm widget** – To add an alarm widget to a dashboard,
choose
one of the
following
options:

- Add
a single alarm in a widget, which
displays
the graph of the alarm's metric and
also
displays the alarm status.

- Add
an _alarm status widget_, which displays the status of
multiple alarms in a grid. Only the alarm names and current status are displayed,
Graphs
are not displayed.
You
can include
up
to 100 alarms
in
one alarm status widget.

- **Table widget** – With the data table widget, you can see the raw datapoints of your metric and a quick summary of that
raw data. Because the data table widget is not a chart to abstract the actual data away from you, it
is easier to understand the datapoints being presented. The procedures in this section describe
how to add and remove a data table widget from a CloudWatch dashboard. The following image shows an
example of a table widget with columns for the Minimum, Maximum, Sum, and Average statistics for a set of
CloudWatch metrics.

- **Linked graphs** – You can link the graphs on your dashboard together, so that when you zoom in or zoom out
on one graph, the other graphs zoom in or zoom out at the same time. You can unlink graphs
to limit zoom to one graph.

###### Topics

- [Adding a graph widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_graph_dashboard.html)

- [Removing a graph widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/remove_graph_dashboard.html)

- [Graph metrics manually on a dashboard](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_old_metrics_to_graph.html)

- [Editing a graph](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/edit_graph_dashboard.html)

- [Renaming a graph](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/rename_graph_dashboard.html)

- [Moving a graph](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/move_graph_dashboard.html)

- [Changing the size of graphs](#resize_graph_dashboard)

- [Changing the size of a graph temporarily](#resize_graph_temp_dashboard)

- [Add a metrics explorer widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_metrics_explorer_dashboard.html)

- [Adding a line graph widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_line_dashboard.html)

- [Removing a line graph](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/remove_line_dashboard.html)

- [Adding a number widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_number_dashboard.html)

- [Removing a number widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/remove_number_dashboard.html)

- [Adding a gauge widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_gauge_dashboard.html)

- [Removing a gauge widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/remove_gauge_dashboard.html)

- [Using a custom widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_custom_widget_dashboard.html)

- [Adding a text widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_text_dashboard.html)

- [Editing a text widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/edit_text_dashboard.html)

- [Removing a text widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/remove_text_dashboard.html)

- [Adding an alarm](add-alarm-dashboard.md)

- [Adding an alarm status widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_alarm_widget_dashboard.html)

- [Removing an alert widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/remove_alarm_dashboard.html)

- [Using a data table widget](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/add_remove_table_dashboard.html)

- [Linking graphs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/link_graphs_dashboard.html)

- [Unlinking graphs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/unlink_graphs_dashboard.html)

## Changing the size of graphs on CloudWatch dashboards

###### To change the size of a graph

1. Open the CloudWatch console at
    [https://console.aws.amazon.com/cloudwatch/](https://console.aws.amazon.com/cloudwatch).

2. In the navigation pane, choose
    **Dashboards**,
    and
    then
    choose a dashboard.

3. To increase or decrease the size, hover over the graph and drag the lower
    right corner of the graph. Stop dragging the corner when you have the size that you want.

4. Choose **Save dashboard**.

## Changing the size of a graph temporarily on a CloudWatch dashboard

###### To enlarge a graph temporarily

1. Open the CloudWatch console at
    [https://console.aws.amazon.com/cloudwatch/](https://console.aws.amazon.com/cloudwatch).

2. In the navigation pane, choose
    **Dashboards**,
    and
    then
    choose a dashboard.

3. Select the graph. Alternatively, hover over the title of the graph and choose
    **Widget actions**, **Enlarge**.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tutorial: Creating dashboard with the Lambda function name as a variable

Adding a graph widget
