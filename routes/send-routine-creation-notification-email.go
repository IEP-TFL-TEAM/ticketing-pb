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

func SendRoutineCreationNotification(app *pocketbase.PocketBase) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST(
			"/api/send-routine-creation-notification",
			func(c echo.Context) error {
				data := struct {
					Id           string `form:"id"`
					Email        string `form:"email"`
					Subject      string `form:"subject"`
					StartDate    string `form:"startDate"`
					Objective    string `form:"objective"`
					TicketNumber string `form:"ticketNumber"`
				}{}

				if err := c.Bind(&data); err != nil {
					return apis.NewBadRequestError("Failed to read request data.", err)
				}

				id := data.Id
				email := data.Email
				subject := data.Subject
				startDate := data.StartDate
				objective := data.Objective
				ticketNumber := data.TicketNumber

				var logoUrl = "https://www.telecom.com.fj/wp-content/themes/prototype/img/logo.png"

				registry := template.NewRegistry()

				html, _ := registry.LoadFiles(
					"templates/routine-creation-email.html",
				).Render(map[string]any{
					"logoUrl":   logoUrl,
					"startDate": startDate,
					"subject":   subject,
					"objective": objective,
					"url":       fmt.Sprintf("%s/tickets/%s", app.Settings().Meta.AppUrl, id),
					"action":    "A new Routine Maintenance has been Created!",
				})

				message := &mailer.Message{
					From: mail.Address{
						Address: app.Settings().Meta.SenderAddress,
						Name:    app.Settings().Meta.SenderName,
					},

					To:      []mail.Address{{Address: email}},
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
