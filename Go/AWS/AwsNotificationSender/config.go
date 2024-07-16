package main

// InitConfig initialize the configurations
func InitConfig() *Configuration {
	// Config is the instance of configuration
	var config Configuration
	config.SenderEmailID = DefaultFromEmailID
	config.SenderEmailPassword = DefaultSenderEmailPassword
	config.MailServerPort = DefaultMailServerPort
	config.MailServer = DefaultMailServer
	return &config
}

// SenderEmail - returns senderEmailId
func (c *Configuration) SenderEmail() string {
	return c.SenderEmailID
}

// SenderPassword - returns senderEmailPassword
func (c *Configuration) SenderPassword() string {
	return c.SenderEmailPassword
}

// MailSrvPort - returns mailServerPort
func (c *Configuration) MailSrvPort() string {
	return c.MailServerPort
}

// MailSrv - returns mailServer
func (c *Configuration) MailSrv() string {
	return c.MailServer
}
