package main

import (
	"strconv"

	"github.com/conformal/gotk3/gtk"
)

var win *gtk.Window

func main() {
	gtk.Init(nil)

	win, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	win.SetTitle("PhysicsCalculator")
	win.SetDefaultSize(300, 300)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.Add(PCMenuNew())
	win.ShowAll()
	gtk.Main()
}

func PCMenuNew() *gtk.Grid {
	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	gbtn, _ := gtk.ButtonNewWithLabel("Weight")
	gbtn.SetVExpand(true)
	gbtn.SetHExpand(true)
	gbtn.Connect("clicked", func() {
		win.Resize(300, 150)
		grid.Destroy()
		win.Add(PCWinNew("Fg: ", "m: ", "g: ", MCalc))
		win.ShowAll()
	})
	rbtn, _ := gtk.ButtonNewWithLabel("Friction")
	rbtn.SetVExpand(true)
	rbtn.SetHExpand(true)
	rbtn.Connect("clicked", func() {
		win.Resize(300, 150)
		grid.Destroy()
		win.Add(PCWinNew("Fr: ", "Fg: ", "µ: ", MCalc))
		win.ShowAll()
	})
	wbtn, _ := gtk.ButtonNewWithLabel("Work")
	wbtn.SetVExpand(true)
	wbtn.SetHExpand(true)
	wbtn.Connect("clicked", func() {
		win.Resize(300, 150)
		grid.Destroy()
		win.Add(PCWinNew("W: ", "F: ", "s: ", MCalc))
		win.ShowAll()
	})
	lbtn, _ := gtk.ButtonNewWithLabel("Power")
	lbtn.SetVExpand(true)
	lbtn.SetHExpand(true)
	lbtn.Connect("clicked", func() {
		win.Resize(300, 150)
		grid.Destroy()
		win.Add(PCWinNew("P: ", "W: ", "t: ", DCalc))
		win.ShowAll()
	})

	grid.Attach(gbtn, 0, 0, 1, 1)
	grid.Attach(rbtn, 1, 0, 1, 1)
	grid.Attach(wbtn, 0, 1, 1, 1)
	grid.Attach(lbtn, 1, 1, 1, 1)

	return grid
}

func PCWinNew(fls, sls, tls string, fcalc func(*gtk.Entry, *gtk.Entry, *gtk.Entry)) *gtk.Grid {
	grid, _ := gtk.GridNew()
	grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)

	fl, _ := gtk.LabelNew(fls)
	fl.SetHExpand(true)
	fl.SetVExpand(true)

	fli, _ := gtk.EntryNew()
	fli.SetHExpand(true)
	fli.SetVExpand(true)

	sl, _ := gtk.LabelNew(sls)
	sl.SetHExpand(true)
	sl.SetVExpand(true)

	sli, _ := gtk.EntryNew()
	sli.SetHExpand(true)
	sli.SetVExpand(true)

	tl, _ := gtk.LabelNew(tls)
	tl.SetHExpand(true)
	tl.SetVExpand(true)

	tli, _ := gtk.EntryNew()
	tli.SetHExpand(true)
	tli.SetVExpand(true)

	cbtn, _ := gtk.ButtonNewWithLabel("Calc")
	cbtn.SetVExpand(true)
	cbtn.SetHExpand(true)
	cbtn.Connect("clicked", func() {
		fcalc(fli, sli, tli)
	})
	bbtn, _ := gtk.ButtonNewWithLabel("Back")
	bbtn.SetHExpand(true)
	bbtn.Connect("clicked", func() {
		grid.Destroy()
		win.Add(PCMenuNew())
		win.Resize(300, 300)
		win.ShowAll()
	})

	grid.Add(fl)
	grid.Add(fli)
	grid.Add(sl)
	grid.Add(sli)
	grid.Add(tl)
	grid.Add(tli)

	grid.Attach(cbtn, 0, 1, 6, 2)
	grid.Attach(bbtn, 0, 3, 6, 1)

	return grid
}

func MCalc(fli, sli, tli *gtk.Entry) {
	flt, _ := fli.GetText()
	slt, _ := sli.GetText()
	tlt, _ := tli.GetText()

	f, ef := strconv.ParseFloat(flt, 64)
	s, es := strconv.ParseFloat(slt, 64)
	t, et := strconv.ParseFloat(tlt, 64)

	if ef != nil && es == nil && et == nil {
		f = s * t
		fli.SetText(strconv.FormatFloat(f, 'f', 2, 64))
	} else if ef == nil && es != nil && et == nil {
		s = f / t
		sli.SetText(strconv.FormatFloat(s, 'f', 2, 64))
	} else if ef == nil && es == nil && et != nil {
		t = f / s
		tli.SetText(strconv.FormatFloat(t, 'f', 2, 64))
	}
}
func DCalc(fli, sli, tli *gtk.Entry) {
	flt, _ := fli.GetText()
	slt, _ := sli.GetText()
	tlt, _ := tli.GetText()

	f, ef := strconv.ParseFloat(flt, 64)
	s, es := strconv.ParseFloat(slt, 64)
	t, et := strconv.ParseFloat(tlt, 64)

	if ef != nil && es == nil && et == nil {
		f = s / t
		fli.SetText(strconv.FormatFloat(f, 'f', 2, 64))
	} else if ef == nil && es != nil && et == nil {
		s = f * t
		sli.SetText(strconv.FormatFloat(s, 'f', 2, 64))
	} else if ef == nil && es == nil && et != nil {
		t = s / f
		tli.SetText(strconv.FormatFloat(t, 'f', 2, 64))
	}

}
