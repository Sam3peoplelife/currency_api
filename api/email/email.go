package email

import (
    "fmt"
    "log"
    "gopkg.in/gomail.v2"
    "api/models"
    "gorm.io/gorm"
    "github.com/robfig/cron/v3"
    "api/utils"
)

//<summary>
//Function to create email
//</summary>
func sendEmail(to, subject, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", "myemail@example.com")
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer("smtp.example.com", 587, "myemail@example.com", "mypassword")

    if err := d.DialAndSend(m); err != nil {
        return err
    }
    return nil
}

//<summary>
//Function to send daily mails
//</summary>
func sendDailyEmails(db *gorm.DB) {
    rate, err := utils.FetchExchangeRate()
    if err != nil {
        log.Printf("Error fetching exchange rate: %v", err)
        return
    }

    var users []models.User
    db.Find(&users)

    for _, user := range users {
        subject := "Daily USD to UAH Exchange Rate"
        body := fmt.Sprintf("The current exchange rate is: %f UAH for 1 USD.", rate)
        if err := sendEmail(user.Email, subject, body); err != nil {
            log.Printf("Error sending email to %s: %v", user.Email, err)
        }
    }
}

//<summary>
//Creating scheduler
//</summary>
func StartEmailScheduler(db *gorm.DB) {
    c := cron.New()
    c.AddFunc("@daily", func() { sendDailyEmails(db) })
    c.Start()
}

