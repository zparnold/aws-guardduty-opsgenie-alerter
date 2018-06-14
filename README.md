# The AWS GuardDuty OpsGenie Alerter

AWS GuardDuty is super useful. But without some way of getting that data out of GuardDuty in a timely fashion, it makes
it no more important than any other fancy dashboard that has metrics on it. This tool was built to help solve that. Until
AWS builds more robust infrastructure around alerting on specific severity levels or types of GuardDuty messages, we built
a tool leveraging AWS CloudWatch Events, AWS Lambda, and our Alerting tool, OpsGenie.

*The AWS Official Docs on this are here: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html*

## How To:

1) Fork this repo into your own environment
2) Add your OpsGenie Team Name, Team ID, and API Key as environment variables in the `serverless.yml` file
3) Run 
```bash
$ serverless-deploy
```
4) Using the AWS CLI
```bash
$ aws events put-rule --name GuardDuty --event-pattern "{\"source\":[\"aws.guardduty\"],\"detail-type\":[\"GuardDuty Finding\"],\"detail\":{\"severity\":[5.0,8.0]}}"
```
You will likely have to modify the above to have all of the increments you are interested in getting alarms for.
Example:

Getting an alert for all medium/high risk situations will look like this: 
`aws events put-rule --name GuardDuty --event-pattern "{\"source\":[\"aws.guardduty\"],\"detail-type\":[\"GuardDuty Finding\"],\"detail\":{\"severity\":[4.0,4.1,4.2,4.3,4.4,4.5,4.6,4.7,4.8,4.9,5.0,5.1,5.2,5.3,5.4,5.5,5.6,5.7,5.8,5.9,6.0,6.1,6.2,6.3,6.4,6.5,6.6,6.7,6.8,6.9,7.0,7.1,7.2,7.3,7.4,7.5,7.6,7.7,7.8,7.9,8.0,8.1,8.2,8.3,8.4,8.5,8.6,8.7,8.8,8.9]}}"`

5) Then in the AWS Lambda Console, you will need to make an association between your function (just deployed) and the event
rule you just created in Step 4 (https://console.aws.amazon.com/lambda/home)
6) In the "Add Triggers" section of the Lambda console, select the "CloudWatch Events" tab
7) In the "Configure Triggers" box underneath "Rule" select our newly created rule (Above called "GuardDuty")
8) Make sure the "Enable Trigger" box is selected and then click the "Add" button


## GuardDuty Severity Levels
*https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings.html#guardduty_findings-severity*

The severity levels roughly go as follows:

* [0.1 - 3.9] = Low Risk
* [4.0 - 6.9] = Medium Risk
* [7.0 - 8.9] = High Risk


## Contributions:
We would love your ideas and contributions, simply open an issue or PR and I'll be sure to get back to you!