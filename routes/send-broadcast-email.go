package routes

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/template"
)

func SendBroadcastEmail(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST(
			"/api/send-broadcast",
			func(c echo.Context) error {
				data := struct {
					Id            string `form:"id"`
					Email         string `form:"email"`
					Cc            string `form:"cc"`
					Subject       string `form:"subject"`
					IncidentStart string `form:"incidentStart"`
					Description   string `form:"description"`
					Location      string `form:"location"`
					AssignedTeams string `form:"assignedTeams"`
					Update        string `form:"update"`
					TicketNumber  string `form:"ticketNumber"`
				}{}

				if err := c.Bind(&data); err != nil {
					return apis.NewBadRequestError("Failed to read request data.", err)
				}

				id := data.Id
				email := data.Email
				cc := data.Cc
				subject := data.Subject
				incidentStart := data.IncidentStart
				description := data.Description
				location := data.Location
				assignedTeams := data.AssignedTeams
				update := data.Update
				ticketNumber := data.TicketNumber

				var logoUrl = "https://www.telecom.com.fj/wp-content/themes/prototype/img/logo.png"

				registry := template.NewRegistry()

				html, _ := registry.LoadFiles(
					"templates/broadcast-email.html",
				).Render(map[string]any{
					"logoUrl":       logoUrl,
					"incidentStart": incidentStart,
					"description":   description,
					"location":      location,
					"assignedTeams": assignedTeams,
					"update":        update,
					"url":           fmt.Sprintf("%s/tickets/%s", app.Settings().Meta.AppUrl, id),
				})

				message := &mailer.Message{
					From: mail.Address{
						Address: app.Settings().Meta.SenderAddress,
						Name:    app.Settings().Meta.SenderName,
					},

					To:      []mail.Address{{Address: email}},
					Cc:      []mail.Address{{Address: cc}},
					Subject: fmt.Sprintf("%s - %s", subject, ticketNumber),
					HTML:    html,
				}

				err := app.NewMailClient().Send(message)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"details": "Failed to send email!",
					})
				}

				return c.JSON(http.StatusOK, "Resent")
			},
			apis.ActivityLogger(app),
			apis.RequireRecordAuth(),
		)

		return nil
	})
}
