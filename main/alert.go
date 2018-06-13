package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"os"
)

func SendGenieAlert(team string, teamID string) (string, error){

	/* initialize the opsgenie client */
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(os.Getenv("OPSGENIE_API_KEY"))
	alertCli, _ := cli.AlertV2()
	message := "AWS GUARDDUTY ALERT: VULNERABILITY FOUND THAT NEEDS YOUR ATTENTION"

	teams := []alerts.TeamRecipient{
		&alerts.Team{Name: team},
		&alerts.Team{ID: teamID},
	}

	request := alerts.CreateAlertRequest{
		Message:  message,
		Teams:    teams,
		Priority: alerts.P1,
	}

	response, err := alertCli.Create(request)
	if err != nil {
		return "-1", err
	} else {
	}
	return response.RequestID, nil
}
