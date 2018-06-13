package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
	"os"
)

/*  GuardDuty severity levels defined here:
	https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings.html#guardduty_findings-severity
*/
const (
	LOW_RISK_THRESHOLD = 3.9
	MEDIUM_RISK_THRESHOLD = 6.9
	HIGH_RISK_THRESHOLD = 8.9
)

func sendGenieAlert(team string, teamID string, riskLevel float32, message string) {

	/* initialize the opsgenie client */
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(os.Getenv("OPSGENIE_API_KEY"))
	alertCli, _ := cli.AlertV2()

	/* YOU WILL NEED TO CONFIGURE THIS BASED ON YOUR APPLICATIONS SECURITY NEEDS */
	var priority alerts.Priority
	if riskLevel > LOW_RISK_THRESHOLD {
		priority = alerts.P1
	}

	teams := []alerts.TeamRecipient{
		&alerts.Team{Name: team},
		&alerts.Team{ID: teamID},
	}

	request := alerts.CreateAlertRequest{
		Message:  message,
		Teams:    teams,
		Priority: priority,
	}

	//Supress errors that have no Priority level
	if priority != "" {
		response, err := alertCli.Create(request)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Create request ID: " + response.RequestID)
		}
	}

}
