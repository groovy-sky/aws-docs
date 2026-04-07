# Supported CloudWatch metrics

Amazon CloudWatch is a monitoring service for AWS resources. You can use CloudWatch to collect and
track metrics, set alarms, and automatically react to changes in your AWS resources.
ACM publishes metrics twice per day for every certificate in an account until
expiration.

The `AWS/CertificateManager` namespace includes the following metric.

MetricDescriptionUnitDimensions`DaysToExpiry`Number of days until a certificate expires. ACM stops publishing
this metric after a certificate expires.IntegerCertificateArn

- Value: ARN of the certificate

For more information about CloudWatch metrics, see the following topics:

- [Using Amazon CloudWatch\
Metrics](../../../amazoncloudwatch/latest/monitoring/working-with-metrics.md)

- [Creating Amazon CloudWatch\
Alarms](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

API calls for integrated services

Use AWS Certificate Manager with the SDK for Java
