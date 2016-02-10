package cmd

import vultr "github.com/intrixius/vultr/lib"

func GetClient() *vultr.Client {
	return vultr.NewClient(*apiKey, nil)
}
