# Updating the CloudWatch agent container image

###### Important

If you are upgrading or installing Container Insights on an Amazon EKS cluster, we
recommend that you use the Amazon CloudWatch Observability EKS add-on for the installation,
instead of using the instructions in this section. Additionally, to retrieve
accelerated computing metrics, you must use the Amazon CloudWatch Observability EKS
add-on or the CloudWatch agent operator. For more information and instructions, see [Quick start with the Amazon CloudWatch Observability EKS add-on](container-insights-setup-eks-addon.md).

If you need to update your container image to the latest version, use the steps in
this section.

###### To update your container image

1. Verify if the `amazoncloudwatchagent` Customer Resource Definition
    (CRD) already exists by entering the following command.

```nohighlight

kubectl get crds amazoncloudwatchagents.cloudwatch.aws.amazon.com -n amazon-cloudwatch
```

If this command returns an error that the CRD is missing, the cluster doesn't
    have Container Insights with enhanced observabilit for Amazon EKS configured with the
    CloudWatch agent operator. In this case, see [Upgrading to Container Insights with enhanced observability for Amazon EKS in CloudWatch](container-insights-upgrade-enhanced.md).

2. Apply the latest `cwagent-version.yaml` file by entering the
    following command.

```nohighlight

curl https://raw.githubusercontent.com/aws-samples/amazon-cloudwatch-container-insights/main/k8s-quickstart/cwagent-version.yaml | kubectl apply -f -
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Upgrading to Container Insights with enhanced observability for Amazon EKS in CloudWatch

Deleting the CloudWatch agent and Fluent Bit for Container Insights

All content copied from https://docs.aws.amazon.com/.
