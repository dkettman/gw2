package main

import (
	"fmt"
	"os"

	"github.com/dkettman/gw2"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {

	cfg, cfgErr := gw2.LoadConfig("config.json")

	if cfgErr != nil {
		fmt.Printf("%v", cfgErr)
		os.Exit(1)
	}

	c := gw2.NewClient(cfg)

	var configDialog = func() {
		dialog := gtk.NewDialog()
		dialog.SetTitle("Configuration Options")

		configvbox := dialog.GetVBox()

		label := gtk.NewLabel("API Key:")
		configvbox.Add(label)

		input := gtk.NewEntry()
		input.SetEditable(true)
		input.SetText(c.Config.APIKey)
		input.SetWidthChars(73)
		configvbox.Add(input)

		button := gtk.NewButtonWithLabel("OK")
		button.Connect("clicked", func() {
			if input.GetText() != c.Config.APIKey {
				c.Config.APIKey = input.GetText()
			}
			dialog.Destroy()
		})
		configvbox.Add(button)
		dialog.ShowAll()
	}

	var refreshMenu = func() {
		err := c.GetAccountInfo()
		//		updateDisplay()
		if err != nil {
			dialog := gtk.NewDialog()
			dialog.SetTitle("ERROR!")

			errorvbox := dialog.GetVBox()

			label := gtk.NewLabel("Unable to update account info:")
			errorvbox.Add(label)

			label = gtk.NewLabel(err.Error())
			errorvbox.Add(label)

			button := gtk.NewButtonWithLabel("OK")
			button.Connect("clicked", func() {
				dialog.Destroy()
			})
			errorvbox.Add(button)
			dialog.ShowAll()
		}
	}

	if c.Config.APIKey != "" {
		err := c.GetAccountInfo()
		if err != nil {
			panic("AHHHHHHHHH")
		}
	}

	// err := c.GetAccountInfo()
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// 	os.Exit(1)
	// }

	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Guild Wars 2 Account Explorer")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		//fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "Byeeee")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.NewVBox(false, 1)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	var menuitem *gtk.MenuItem
	menuitem = gtk.NewMenuItemWithMnemonic("_Refresh")
	menuitem.Connect("activate", refreshMenu)
	submenu.Append(menuitem)

	menuitem = gtk.NewMenuItemWithMnemonic("_Config")
	menuitem.Connect("activate", configDialog)
	submenu.Append(menuitem)

	menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk.MainQuit()
	})
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	accountFrame := gtk.NewFrame("Account Information")
	accountFramebox1 := gtk.NewVBox(false, 1)
	accountFrame.Add(accountFramebox1)

	accountHBox := gtk.NewHBox(false, 1)
	lblAccountName := gtk.NewLabel("Account Name: " + c.Account.Name)
	lblWorldName := gtk.NewLabel("World Name: " + c.Account.World)
	lblGold := gtk.NewLabel("Gold: ")
	accountHBox.Add(lblAccountName)
	accountHBox.Add(lblWorldName)
	accountHBox.Add(lblGold)
	accountFramebox1.Add(accountHBox)
	vpaned.Pack2(accountFrame, false, false)

	vpane2 := gtk.NewVPaned()
	vbox.Add(vpane2)

	charInfo := gtk.NewLabel("")

	charCombo := gtk.NewComboBoxText()
	for _, v := range c.Account.Characters {
		charCombo.AppendText(v.Core.Name)
	}
	charCombo.Connect("changed", func() {
		char := c.Account.Characters[charCombo.GetActiveText()]
		char.GetDetails(&c)
		charInfo.SetText(fmt.Sprintf("%v", char))
	})

	vpane2.Add(charCombo)
	vpane2.Add(charInfo)
	window.Add(vbox)

	window.SetSizeRequest(600, 600)

	window.ShowAll()
	gtk.Main()

}

func uniq(strings []string) (ret []string) {
	return
}
