# Manage and monitor AMI usage

AWS provides several features to help you manage and monitor your AMI usage effectively.
You can track which accounts are using your shared AMIs, identify when your AMIs
were last used, and discover which resources in your AWS account are referencing
specific AMIs.

The following table provides an overview of the features for managing and monitoring AMI
usage:

FeatureUse caseKey benefits[AMI usage reports](your-ec2-ami-usage.md)Gain visibility into which AWS accounts are using your AMIs and how much each AMI is
being used.

- Identify the AWS accounts and resource types referencing your AMIs so that you can
safely deregister or disable AMIs.

- Identify unused AMIs for deregistration to reduce storage costs.

- Identify your most used AMIs.

[Last used tracking](ami-last-launched-time.md)Check when your AMI was last used.

- Identify unused AMIs so that you can safely
deregister AMIs.

- Identify unused AMIs for deregistration to reduce
storage costs.

[AMI reference check](ec2-ami-references.md)Ensure your AWS resources are using the latest compliant AMIs.

- Audit the use of AMIs in your account.

- Check where specific AMIs are being referenced.

- Maintain compliance by updating your resources to reference the latest
AMIs.

###### Topics

- [View your AMI usage](your-ec2-ami-usage.md)

- [Check when an Amazon EC2 AMI was last used](ami-last-launched-time.md)

- [Identify your resources referencing specified AMIs](ec2-ami-references.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AMI ancestry

View your AMI usage
