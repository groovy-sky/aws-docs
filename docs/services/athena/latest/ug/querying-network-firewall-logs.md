# Query AWS Network Firewall logs

AWS Network Firewall is a managed service that you can use to deploy essential network protections
for your Amazon Virtual Private Cloud instances. AWS Network Firewall works together with AWS Firewall Manager so you can build
policies based on AWS Network Firewall rules and then centrally apply those policies across your VPCs
and accounts. For more information about AWS Network Firewall, see [AWS Network Firewall](https://aws.amazon.com/network-firewall).

You can configure AWS Network Firewall logging for traffic that you forward to your firewall's
stateful rules engine. Logging gives you detailed information about network traffic,
including the time that the stateful engine received a packet, detailed information about
the packet, and any stateful rule action taken against the packet. The logs are published to
the log destination that you've configured, where you can retrieve and view them. For more
information, see [Logging network\
traffic from AWS Network Firewall](https://docs.aws.amazon.com/network-firewall/latest/developerguide/firewall-logging.html) in the _AWS Network Firewall Developer Guide_.

###### Topics

- [Create and query a table for alert logs](https://docs.aws.amazon.com/athena/latest/ug/querying-network-firewall-logs-sample-alert-logs-table.html)

- [Create and query a table for netflow logs](https://docs.aws.amazon.com/athena/latest/ug/querying-network-firewall-logs-sample-netflow-logs-table.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GuardDuty

Alert logs
