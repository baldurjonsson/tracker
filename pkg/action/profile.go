package action

import (
	"bufio"
	"fmt"
	"os"

	"github.com/baldurjonsson/tracker/pkg/store"
	"github.com/urfave/cli/v2"
)

func ShowProfile(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	fmt.Println("Email: ", s.Profile.Profile.Email)
	fmt.Println("Name: ", s.Profile.Profile.Name)
	return nil
}

func SetProfile(c *cli.Context) error {
	b := bufio.NewReader(os.Stdin)
	s := c.Context.Value("store").(*store.Store)
	email := prompt(b, "Email")
	name := prompt(b, "Name")
	if email == "" {
		return nil
	}
	s.Profile.Profile.Email = email
	s.Profile.Profile.Name = name
	return s.Save()
}
