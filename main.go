package main

import (
	"time"

	"github.com/micro/cli"
	"github.com/pquerna/otp/totp"

	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
)

var (
	name     string
	version  string
	passcode = cli.Command{
		Name:    "passcode",
		Usage:   "generate passcode",
		Aliases: []string{"pc"},
		Action:  GeneratePasscode,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "secret,s",
				Usage: "secret code",
			},
		},
	}
	qrcode = cli.Command{
		Name:    "qrcode",
		Usage:   "generate qrcode",
		Action:  GenerateQRcode,
		Aliases: []string{"qr"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "issuer",
				Usage: "default syhlion",
				Value: "syhlion",
			},
			cli.StringFlag{
				Name:  "account",
				Usage: "default syhlion[@]gmail dot com",
				Value: "syhlion[@]gmail dot com",
			},
		},
	}
	check = cli.Command{
		Name:    "check",
		Usage:   "check passcode",
		Action:  CheckPasscode,
		Aliases: []string{"ch"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "secret,s",
				Usage: "secret code",
			},
		},
	}
)

func GenerateQRcode(c *cli.Context) {
	issuer := c.String("issuer")
	account := c.String("account")
	kkey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf bytes.Buffer
	img, err := kkey.Image(200, 200)
	if err != nil {
		fmt.Println(err)
		return
	}
	png.Encode(&buf, img)
	fmt.Printf("Issuer:       %s\n", kkey.Issuer())
	fmt.Printf("Account Name: %s\n", kkey.AccountName())
	fmt.Printf("Secret:       %s\n", kkey.Secret())
	fmt.Println("Writing PNG to qr-code.png....")
	ioutil.WriteFile("qr-code.png", buf.Bytes(), 0644)
	fmt.Println("finish!!!")

}
func GeneratePasscode(c *cli.Context) {
	secret := c.String("secret")
	if secret == "" {
		fmt.Println("secret empty")
		return
	}
	passcode, _ := totp.GenerateCode(secret, time.Now())
	fmt.Println("passcode:")
	fmt.Println(passcode)
}
func CheckPasscode(c *cli.Context) {
	secret := c.String("secret")
	if secret == "" {
		fmt.Println("secret empty")
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Passcode: ")
		text, _ := reader.ReadString('\n')
		valid := totp.Validate(text, secret)
		if valid {
			fmt.Println("Valid passcode!")
		} else {
			fmt.Println("Invalid passocde!")
		}

	}
}

func main() {
	gusher := cli.NewApp()
	gusher.Name = name
	gusher.Author = "Scott (syhlion)"
	gusher.Usage = "very simple to know totp"
	gusher.UsageText = "totp-tool [qrcode|check|passcode] -h"
	gusher.Version = version
	gusher.Compiled = time.Now()
	gusher.Commands = []cli.Command{
		qrcode,
		check,
		passcode,
	}
	gusher.Run(os.Args)
}
