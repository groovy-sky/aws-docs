# View Internet Monitor metrics or set alarms in CloudWatch Metrics

You can view or set alarms on Internet Monitor metrics by using CloudWatch alarms and CloudWatch Metrics in the CloudWatch console. Internet Monitor publishes
metrics to your account, including metrics for performance, availability, round-trip time, and throughput (bytes per second).
To find all metrics for your monitor, in the CloudWatch Metrics dashboard, see
the custom namespace `AWS/InternetMonitor`.

To see examples for using several of these metrics to help determine values to choose for a city-networks maximum limit
for your monitor, see
[Choosing a city-network maximum value](imcitynetworksmaximum.md). To learn more about setting alarms
for Internet Monitor, see [Create alarms with Internet Monitor](cloudwatch-im-create-alarm.md).

Metrics are aggregated across all internet traffic to your VPCs, Network Load Balancers, CloudFront distributions,
or WorkSpaces directories in the monitor, and to all traffic to each AWS Region and internet edge location that is monitored.
Regions are defined by the service location, which can either be all locations or a specific Region, such as `us-east-1`.

Note: _city-networks_ are pairs of client locations and the ASNs the clients use (typically internet
service providers or ISPs).

Internet Monitor provides the following metrics.

MetricDescriptionPerformanceScoreA performance score represents the estimated percentage of traffic that is not seeing a performance drop.AvailabilityScoreAn availability score represents the estimated percentage of traffic that is not seeing an availability drop. BytesInBytes transferred in for your application internet traffic at all application city-networks.BytesOutBytes transferred out for your application internet traffic at all application city-networks.BytesInMonitoredBytes transferred in for your application internet traffic at monitored city-networks.BytesOutMonitoredBytes transferred out for your application internet traffic at monitored city-networks.Round-trip time (RTT)Round-trip time between the AWS Regions, ASNs (typically internet service providers or ISPs),
and locations (such as cities) specific to your VPCs, Network Load Balancers, CloudFront distributions, or WorkSpaces directories.CityNetworksMonitoredThe number of city-networks Internet Monitor monitored for your application internet traffic. This is never
more than the upper limit that you set as the maximum city-networks for the monitor.TrafficMonitoredPercentThe percentage of total application internet traffic for this monitor that is represented (included)
by the city-networks that Internet Monitor is monitoring. This is less than 100 (that is, less than 100%) if clients access your application
in more city-networks than the maximum city-networks limit that you have set for the monitor.CityNetworksFor100PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 100% of your
application internet traffic in Internet Monitor.CityNetworksFor99PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 99% of your
application internet traffic in Internet Monitor.CityNetworksFor95PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 95% of your
application internet traffic in Internet Monitor.CityNetworksFor90PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 90% of your
application internet traffic in Internet Monitor.CityNetworksFor75PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 75% of your
application internet traffic in Internet Monitor.CityNetworksFor50PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 50% of your
application internet traffic in Internet Monitor.CityNetworksFor25PercentTrafficThe number that you should set your city-networks maximum limit to if you want to monitor 25% of your
application internet traffic in Internet Monitor.

For more information, see [Metrics in Amazon CloudWatch](working-with-metrics.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatch Contributor Insights

Athena with S3 logs

All content copied from https://docs.aws.amazon.com/.
