# Edit a monitor in Internet Monitor

Using the **Action** menu, you can edit a monitor in Amazon CloudWatch Internet Monitor after you create it. For example, you can
edit a monitor to do the following:

- Change the percentage of application traffic to monitor

- Set or update the city-networks maximum limit

- Change health event thresholds for availability or performance scores

- Add or remove resources

- Enable or update publishing events to Amazon S3

Note that you can't change the name of a monitor after you create it.

To make changes to a monitor, use the following procedure.

###### To edit a monitor

1. Open the CloudWatch console at
    [https://console.aws.amazon.com/cloudwatch/](https://console.aws.amazon.com/cloudwatch).

2. In the left navigation pane, under **Network Monitoring**, choose **Internet monitors**.

3. Choose your monitor, and then choose the **Action** menu.

4. Choose **Update monitor**.

5. Make the desired updates. For example, to change the percentage of traffic to monitor, under **Application traffic**
**to monitor**, select or enter a percentage.

6. Choose **Update**.

For more information about the options that you can update, see the following:

- To learn more about resources that you add in Internet Monitor, see [Add resources to your monitor](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMMonitorResources.html).

- To learn more about the application traffic percentage, see [Choose a percentage of traffic to monitor for your application](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMTrafficPercentage.html).

- To learn more about changing the threshold for health events, see [Change health event\
thresholds](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-get-started.change-threshold.html#IMUpdateThresholdFromOverview).

- To learn more about the city-networks maximum limit, see [Choose a city-networks maximum limit](imcitynetworksmaximum.md).

- To learn more about opting to publish events to S3, see [Publish internet measurements to Amazon S3 in Internet Monitor](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-get-started.Publish-to-S3.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use a monitor

Delete a monitor
