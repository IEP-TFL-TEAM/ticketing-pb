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

func SendIncidentCreationNotification(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST(
			"/api/send-incident-creation-notification",
			func(c echo.Context) error {
				data := struct {
					Id            string `form:"id"`
					Email         string `form:"email"`
					Cc            string `form:"cc"`
					Subject       string `form:"subject"`
					IncidentStart string `form:"incidentStart"`
					Description   string `form:"description"`
					TicketNumber  string `form:"ticketNumber"`
					ActionType    string `form:"actionType"`
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
				ticketNumber := data.TicketNumber
				actionType := data.ActionType

				var action string
				var greeting string
				var whoToGreet string

				switch actionType {
				case "create":
					greeting = "Bula"
					whoToGreet = "Team"
					action = "A new Incident has been Created!"

				case "reassign":
					greeting = "Bula"
					whoToGreet = ""
					action = "Your team has been re-assigned this ticket."

				default:
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"details": "Invalid action type!",
					})
				}

				var logoUrl = "https://www.telecom.com.fj/wp-content/themes/prototype/img/logo.png"

				registry := template.NewRegistry()

				html, _ := registry.LoadFiles(
					"templates/incident-creation-email.html",
				).Render(map[string]any{
					"logoUrl":       logoUrl,
					"incidentStart": incidentStart,
					"subject":       subject,
					"description":   description,
					"url":           fmt.Sprintf("%s/tickets/%s", app.Settings().Meta.AppUrl, id),
					"action":        action,
					"greeting":      greeting,
					"whoToGreet":    whoToGreet,
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

				return c.JSON(http.StatusOK, "Email Sent")
			},
			apis.ActivityLogger(app),
			apis.RequireRecordAuth(),
		)

		return nil
	})
}
