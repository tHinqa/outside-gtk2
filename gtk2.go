//Package gtk2 is useful only to install all subcomponents

//TODO(t): Bring to attention synonymous use of gchar char and
// possible simgle byte output *char

package gtk2

import (
	_ "github.com/tHinqa/outside-gtk2/atk"
	_ "github.com/tHinqa/outside-gtk2/cairo"
	_ "github.com/tHinqa/outside-gtk2/expat"
	_ "github.com/tHinqa/outside-gtk2/fontconfig"
	_ "github.com/tHinqa/outside-gtk2/freetype2"
	_ "github.com/tHinqa/outside-gtk2/gail"
	_ "github.com/tHinqa/outside-gtk2/gdk"
	_ "github.com/tHinqa/outside-gtk2/gio"
	_ "github.com/tHinqa/outside-gtk2/glib"
	_ "github.com/tHinqa/outside-gtk2/gmodule"
	_ "github.com/tHinqa/outside-gtk2/gobject"
	_ "github.com/tHinqa/outside-gtk2/gtk"
	_ "github.com/tHinqa/outside-gtk2/gtksourceview"
	_ "github.com/tHinqa/outside-gtk2/intl"
	_ "github.com/tHinqa/outside-gtk2/pango"
	_ "github.com/tHinqa/outside-gtk2/pangowin32"
	_ "github.com/tHinqa/outside-gtk2/png"
	_ "github.com/tHinqa/outside-gtk2/zlib"
)
