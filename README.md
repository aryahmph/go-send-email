## Getting Started

1. Start with cloning this repo on your local machine :

```
$ git clone https://github.com/aryahmph/go-send-email
$ cd go-send-email
```

2. Create database named `go_send_email` or you can modify it in database.go
3. Create file named `app.config.json` and filled id like this :

```json
{
  "name": "Your Email Name",
  "email": "your.email@gmail.com",
  "password": "Your Password"
}
```

4. Config your Google account to enable Less secure apps feature for sending email from
   program. https://myaccount.google.com/lesssecureapps

5. Start program and connect to http://localhost:3000