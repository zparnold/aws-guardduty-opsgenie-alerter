# The AWS GuardDuty OpsGenie Alerter

AWS GuardDuty is super useful. But without some way of getting that data out of GuardDuty in a timely fashion, it makes
it no more important than any other fancy dashboard that has metrics on it. This tool was built to help solve that. Until
AWS builds more robust infrastructure around alerting on specific severity levels or types of GuardDuty messages, we built
a tool leveraging AWS CloudWatch Events, AWS Lambda, and our Alerting tool, OpsGenie.

*The AWS Official Docs on this are here: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html*

