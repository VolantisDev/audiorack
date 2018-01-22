package main

import (
	"flag"
	"time"
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/volantisdev/vst-guide/pkg/bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/juju/errors"
)

const htmlAbout = `Welcome on <b>VST Guide</b>!<br>
Please visit <a href="https://vstguide.com">vstguide.com</a> for more information.`

var (
	AppName string
	BuiltAt string
	debug	= flag.Bool("d", false, "Enables debug mode")
	w		*astilectron.Window
)

func main() {
	flag.Parse()
	astilog.FlagInit()

	astilog.Debugf("Running app built at %s", BuiltAt)

	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName: AppName,
			AppIconDarwinPath: "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: *debug,
		Homepage: "../../../../web/dist/index.html",
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr("About"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", htmlAbout, func(m *bootstrap.MessageIn) {
							// Unmarshal payload
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								astilog.Error(errors.Annotate(err, "unmarshaling payload failed"))
								return
							}

							astilog.Infof("About modal has been displayed and payload is %s!", s)
						}); err != nil {
							astilog.Error(errors.Annotate(err, "sending about event failed"))
						}
						return
					},
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, iw *astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = iw
			go func() {
				time.Sleep(5 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					astilog.Error(errors.Annotate(err, "sending check.out.menu event failed"))
				}
			}()
			return nil
		},
		MessageHandler: handleMessages,
		RestoreAssets: RestoreAssets,
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center: astilectron.PtrBool(true),
			Height: astilectron.PtrInt(900),
			Width: astilectron.PtrInt(700),
		},
	}); err != nil {
		astilog.Fatal(errors.Annotate(err, "running bootstrap failed"))
	}
}
